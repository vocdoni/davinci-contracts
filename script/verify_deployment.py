#!/usr/bin/env python3
"""Checks a live ProcessRegistry deployment against this checkout's build.

Compares the on-chain runtime code of ZiskVerifier, ProcessRegistry and both adapters
with `forge build` output (immutable slots masked), then reads every immutable back and
compares it with the expected pins, verifier and managers. The grace and notice settings
are compared with --pins-from and --grace, and printed otherwise. Exit status 1 on any
mismatch. --record writes what was verified as JSON, only when every check passes.

    python3 script/verify_deployment.py --rpc URL --registry 0x... \
        --batch-vk 0x... --results-vk 0x... --root-c 0x... --ballot-vk-hash 0x... \
        [--grace 300,240,600,1800,60] [--l1-blob-verifier 0x...|0x0]
    python3 script/verify_deployment.py --rpc URL --registry 0x... --pins-from 0x... \
        --dkg-manager 0x... --council-manager 0x... [--l1-block 0x...] [--broadcast run-latest.json --record out.json]
"""
import argparse
import json
import re
from urllib.parse import urlsplit
import subprocess
import sys
import urllib.request
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
UA = {"Content-Type": "application/json", "User-Agent": "davinci-verify"}


def scrub(text, url):
    """`text` without the parts of the RPC URL (path, host, userinfo, query) a server may echo."""
    text = str(text)
    try:
        u = urlsplit(url)
        parts = [u.path.strip("/"), u.query, u.netloc, u.netloc.rpartition("@")[0]]
    except ValueError:
        parts = []
    for part in sorted({p for p in parts if len(p) > 1}, key=len, reverse=True):
        text = text.replace(part, "<rpc>")
    return text


def failure(e, url):
    """One line for a transport or decoding failure: type and reason or code, scrubbed."""
    return f"{type(e).__name__}: {scrub(getattr(e, 'reason', getattr(e, 'code', '')), url)}"


def rpc(url, method, params):
    body = json.dumps({"jsonrpc": "2.0", "id": 1, "method": method, "params": params})
    req = urllib.request.Request(url, data=body.encode(), headers=UA)
    try:
        out = json.load(urllib.request.urlopen(req, timeout=30))
    except Exception as e:  # noqa: BLE001 - one line; the URL may carry a key and is never echoed
        raise SystemExit(f"{method}: {failure(e, url)}")
    if not isinstance(out, dict):
        raise SystemExit(f"{method}: malformed response")
    if "error" in out:
        raise SystemExit(f"{method}: {scrub(out['error'], url)}")
    if "result" not in out:
        raise SystemExit(f"{method}: response without a result")
    return out["result"]


ADDRESS_WORD = re.compile(r"0x0{24}[0-9a-fA-F]{40}")


def optional_address(url, registry, calldata, what="councilAdapter()"):
    """The registry's address getter `what`, or None when the registry has no such function.

    Absent means only what a registry from before the COUNCIL mode answers: an execution
    revert without revert data, or an empty return. A transport failure, any other JSON-RPC
    error, a revert carrying data or a result that is not one address word raises SystemExit,
    so an unreachable or misbehaving RPC fails the verification instead of passing it.
    """
    body = json.dumps(
        {"jsonrpc": "2.0", "id": 1, "method": "eth_call", "params": [{"to": registry, "data": calldata}, "latest"]}
    )
    req = urllib.request.Request(url, data=body.encode(), headers=UA)
    try:
        out = json.load(urllib.request.urlopen(req, timeout=30))
    except Exception as e:  # noqa: BLE001 - every transport or decoding failure fails closed
        raise SystemExit(f"{what}: {failure(e, url)}")
    if not isinstance(out, dict):
        raise SystemExit(f"{what}: malformed response")
    if "error" in out:
        err = out["error"]
        if (
            isinstance(err, dict)
            and "execution reverted" in str(err.get("message", "")).lower()
            and err.get("data") in (None, "", "0x")
        ):
            return None
        raise SystemExit(f"{what}: {scrub(err, url)}")
    result = out.get("result")
    if result == "0x":
        return None
    if not isinstance(result, str) or not ADDRESS_WORD.fullmatch(result):
        raise SystemExit(f"{what}: malformed result {result!r}")
    return "0x" + result[-40:]


def selector(sig, cast="cast"):
    out = subprocess.run([cast, "sig", sig], capture_output=True, text=True, check=True)
    return out.stdout.strip()


def artifact(name):
    p = ROOT / "out" / f"{name}.sol" / f"{name}.json"
    return json.loads(p.read_text())


def masked_match(onchain_hex, art):
    want = bytearray.fromhex(art["deployedBytecode"]["object"][2:])
    got = bytearray.fromhex(onchain_hex[2:])
    if len(got) != len(want):
        return False, f"length {len(got)} != {len(want)}"
    for refs in art["deployedBytecode"].get("immutableReferences", {}).values():
        for r in refs:
            got[r["start"]:r["start"] + r["length"]] = bytes(r["length"])
    if got != want:
        diff = next(i for i in range(len(got)) if got[i] != want[i])
        return False, f"first differing byte at {diff}"
    return True, "ok"


GRACE = ("defaultGrace", "graceFloor", "graceCeil", "graceMaxTotal", "noticeMin")
PINS = (
    ("batchProgramVK", "batch_vk"),
    ("resultsProgramVK", "results_vk"),
    ("rootCVadcopFinal", "root_c"),
    ("ballotVKHash", "ballot_vk_hash"),
)


def broadcast_meta(path, registry):
    """Commit, deployer, block and transaction hashes of the run that created registry."""
    run = json.loads(Path(path).read_text())
    txs = [t for t in run.get("transactions", []) if t.get("transactionType") == "CREATE"]
    reg = [t for t in txs if t.get("contractName") == "ProcessRegistry"]
    if len(reg) != 1 or str(reg[0].get("contractAddress", "")).lower() != registry.lower():
        raise SystemExit(f"{path}: no ProcessRegistry CREATE at {registry}")
    block = {r["transactionHash"]: int(r["blockNumber"], 16) for r in run.get("receipts", [])}
    return {
        "commit": run.get("commit"),
        "deployer": reg[0]["transaction"]["from"],
        "block": block.get(reg[0]["hash"]),
        "transactions": {t["contractName"]: t["hash"] for t in txs},
    }


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--rpc", required=True)
    ap.add_argument("--registry", required=True)
    ap.add_argument("--chain-id", type=int)
    ap.add_argument("--batch-vk")
    ap.add_argument("--results-vk")
    ap.add_argument("--root-c")
    ap.add_argument("--ballot-vk-hash")
    ap.add_argument(
        "--pins-from",
        metavar="REGISTRY",
        help="expect the verifier, pins and grace settings of this live registry (a deployment made "
        "with PINS_FROM_REGISTRY); explicit pin flags must agree with it",
    )
    ap.add_argument("--verifier", help="expected ziskVerifier()")
    ap.add_argument("--dkg-manager", help="expected DavinciDKGAdapter manager; zero expects DKG disabled")
    ap.add_argument("--council-manager", help="expected CouncilAdapter manager; zero expects COUNCIL disabled")
    ap.add_argument(
        "--l1-block",
        default="0x4200000000000000000000000000000000000015",
        help="expected L1BlobVerifier.l1Block() when the registry is in L1-blob mode (OP Stack predeploy by default)",
    )
    ap.add_argument(
        "--grace",
        metavar="DEFAULT,FLOOR,CEIL,MAX_TOTAL,NOTICE",
        help="expected defaultGrace, graceFloor, graceCeil, graceMaxTotal and noticeMin in seconds "
        "(DeployAll keeps its default for a GRACE_* variable that is unset or misspelled)",
    )
    ap.add_argument(
        "--l1-blob-verifier",
        metavar="ADDR|0x0",
        help="expected registry.l1BlobVerifier(); 0x0 expects blob mode. Without it the mode is only printed",
    )
    ap.add_argument("--broadcast", help="forge run-latest.json of the deployment, for the record")
    ap.add_argument("--record", help="write the verified deployment as JSON here (only when every check passes)")
    ap.add_argument("--cast", default=str(Path.home() / ".foundry/bin/cast"))
    a = ap.parse_args()
    try:
        u = urlsplit(a.rpc)
        bad = u.scheme not in ("http", "https") or not u.netloc
    except ValueError:
        bad = True
    if bad or re.search(r"[\s\x00-\x1f\x7f]", a.rpc):
        ap.error("--rpc must be an http(s) URL without whitespace (not echoed: it may carry a key)")
    if a.l1_blob_verifier is not None and not re.fullmatch(r"0x[0-9a-fA-F]{1,40}", a.l1_blob_verifier):
        ap.error("--l1-blob-verifier takes a hex address or 0x0")
    expect_grace = None
    if a.grace is not None:
        fields = a.grace.split(",")
        if len(fields) != len(GRACE) or not all(re.fullmatch(r"[0-9]+", x) for x in fields):
            ap.error("--grace takes five non-negative integers: DEFAULT,FLOOR,CEIL,MAX_TOTAL,NOTICE")
        values = [int(x) for x in fields]
        if max(values) > 0xFFFFFFFF:
            ap.error("--grace values must fit in uint32 (at most 4294967295)")
        expect_grace = dict(zip(GRACE, values))

    ok = True

    def check(label, cond, detail=""):
        nonlocal ok
        print(f"{'OK  ' if cond else 'FAIL'} {label} {detail}".rstrip())
        ok &= bool(cond)

    def get(to, sig, *args):
        data = selector(sig, a.cast) + "".join(x[2:].lower().rjust(64, "0") for x in args)
        return rpc(a.rpc, "eth_call", [{"to": to, "data": data}, "latest"])

    word = lambda h: "0x" + h[2:].rjust(64, "0")[-64:]
    addr = lambda h: "0x" + h[-40:]
    same = lambda x, y: x.lower() == y.lower()

    chain = int(rpc(a.rpc, "eth_chainId", []), 16)
    if a.chain_id is not None:
        check("rpc chain id", chain == a.chain_id, str(chain))

    # Expected pins: explicit flags, else (and checked against) the --pins-from registry.
    want = {name: getattr(a, flag) for name, flag in PINS}
    want_grace = None
    if a.pins_from:
        for name, flag in PINS:
            inherited = word(get(a.pins_from, name + "()"))
            if want[name] is not None:
                check(f"--{flag.replace('_', '-')} == {a.pins_from}.{name}", same(want[name], inherited))
            want[name] = inherited
        source_verifier = addr(get(a.pins_from, "ziskVerifier()"))
        if a.verifier is not None:
            check(f"--verifier == {a.pins_from}.ziskVerifier", same(a.verifier, source_verifier))
        a.verifier = source_verifier
        want_grace = {n: int(get(a.pins_from, n + "()"), 16) for n in GRACE}
    missing = [f"--{flag.replace('_', '-')}" for name, flag in PINS if want[name] is None]
    if missing:
        ap.error(f"{', '.join(missing)} required without --pins-from")

    verifier = addr(get(a.registry, "ziskVerifier()"))
    if a.verifier is not None:
        check("registry.ziskVerifier", same(verifier, a.verifier), verifier)

    reg_code = rpc(a.rpc, "eth_getCode", [a.registry, "latest"])
    ver_code = rpc(a.rpc, "eth_getCode", [verifier, "latest"])
    m, why = masked_match(reg_code, artifact("ProcessRegistry"))
    check("ProcessRegistry runtime code == local build (immutables masked)", m, why)
    m, why = masked_match(ver_code, artifact("ZiskVerifier"))
    check("ZiskVerifier runtime code == local build", m, why)

    pins = {}
    for name, _ in PINS:
        pins[name] = word(get(a.registry, name + "()"))
        check(f"registry.{name}", same(pins[name], want[name]))
    check("verifier.getRootCVadcopFinal", same(word(get(verifier, "getRootCVadcopFinal()")), want["rootCVadcopFinal"]))
    reg_chain = int(get(a.registry, "chainID()"), 16)
    check("registry.chainID == rpc chain id", reg_chain == chain, str(reg_chain))

    record = {
        "chainId": chain,
        "pinsFrom": a.pins_from,
        "processRegistry": a.registry,
        "ziskVerifier": verifier,
        "ziskVerifierCodeHash": None,
        "pins": pins,
        "dkgAdapter": None,
        "dkgManager": None,
        "dkgAppManager": None,
        "councilAdapter": None,
        "councilManager": None,
        "l1BlobVerifier": None,
        "l1Block": None,
    }
    if a.record:
        out = subprocess.run([a.cast, "keccak", ver_code], capture_output=True, text=True, check=True)
        record["ziskVerifierCodeHash"] = out.stdout.strip()

    adapter = addr(get(a.registry, "dkgAdapter()"))
    if int(adapter, 16) != 0:
        adp_code = rpc(a.rpc, "eth_getCode", [adapter, "latest"])
        m, why = masked_match(adp_code, artifact("DavinciDKGAdapter"))
        check("DavinciDKGAdapter runtime code == local build (immutables masked)", m, why)
        adp_reg = addr(get(adapter, "registry()"))
        check("adapter.registry == registry", same(adp_reg, a.registry), adp_reg)
        manager = addr(get(adapter, "manager()"))
        if a.dkg_manager is not None:
            check("adapter.manager == --dkg-manager", same(manager, a.dkg_manager), manager)
        app_manager = addr(get(adapter, "appManager()"))
        check("adapter.appManager == manager.appManager", same(app_manager, addr(get(manager, "appManager()"))))
        # vocdoni/davinci-dkg#14: application ids live in the adapter's namespace.
        aid = word(get(a.registry, "aidFor(bytes31)", "0x" + "00" * 30 + "0100"))  # bytes31(uint248(1))
        check("registry.aidFor in the adapter's namespace", same(addr(aid), adapter), aid)
        record.update(dkgAdapter=adapter, dkgManager=manager, dkgAppManager=app_manager)
        print(f"adapter  {adapter} (manager {manager})")
    else:
        if a.dkg_manager is not None:
            check("DKG disabled as --dkg-manager expects", int(a.dkg_manager, 16) == 0)
        print("adapter  none (DKG disabled)")
    try:
        council = optional_address(a.rpc, a.registry, selector("councilAdapter()", a.cast))
    except SystemExit as e:  # unreadable is a failure, never "no adapter"
        check("registry.councilAdapter() readable", False, str(e))
        council = "unreadable"
    if council == "unreadable":
        pass
    elif council is not None and int(council, 16) != 0:
        cad_code = rpc(a.rpc, "eth_getCode", [council, "latest"])
        m, why = masked_match(cad_code, artifact("CouncilAdapter"))
        check("CouncilAdapter runtime code == local build (immutables masked)", m, why)
        cad_reg = addr(get(council, "registry()"))
        check("councilAdapter.registry == registry", same(cad_reg, a.registry), cad_reg)
        cmanager = addr(get(council, "manager()"))
        if a.council_manager is not None:
            check("councilAdapter.manager == --council-manager", same(cmanager, a.council_manager), cmanager)
        check("council manager has code", rpc(a.rpc, "eth_getCode", [cmanager, "latest"]) not in ("0x", ""))
        record.update(councilAdapter=council, councilManager=cmanager)
        print(f"council {council} (manager {cmanager})")
    else:
        if a.council_manager is not None:
            check("COUNCIL disabled as --council-manager expects", int(a.council_manager, 16) == 0)
        if council is None:
            print("council none (no councilAdapter(): a registry from before the COUNCIL mode)")
        else:
            print("council none (COUNCIL disabled)")
    try:
        l1v = optional_address(a.rpc, a.registry, selector("l1BlobVerifier()", a.cast), "l1BlobVerifier()")
    except SystemExit as e:  # unreadable is a failure, never "blob mode" or "pre-L1 registry"
        check("registry.l1BlobVerifier() readable", False, str(e))
        l1v = "unreadable"
    if a.l1_blob_verifier is not None:
        want_l1v = int(a.l1_blob_verifier, 16)
        if want_l1v == 0:
            check("registry is in blob mode as --l1-blob-verifier 0x0 expects", l1v is None or (l1v != "unreadable" and int(l1v, 16) == 0), str(l1v))
        else:
            check("registry.l1BlobVerifier == --l1-blob-verifier", l1v not in (None, "unreadable") and same(l1v, a.l1_blob_verifier), str(l1v))
    if l1v == "unreadable":
        pass
    elif l1v is not None and int(l1v, 16) != 0:
        v_code = rpc(a.rpc, "eth_getCode", [l1v, "latest"])
        m, why = masked_match(v_code, artifact("L1BlobVerifier"))
        check("L1BlobVerifier runtime code == local build (immutables masked)", m, why)
        l1b = addr(get(l1v, "l1Block()"))
        check("l1BlobVerifier.l1Block == --l1-block", same(l1b, a.l1_block), l1b)
        record.update(l1BlobVerifier=l1v, l1Block=l1b)
        print(f"l1 blob  {l1v} (L1Block {l1b})")
    elif l1v is None:
        print("l1 blob  n/a (a registry from before the L1-blob mode)")
    else:
        print("l1 blob  none (blob mode)")
    print(f"verifier {verifier}")
    grace = {}
    for name in GRACE:
        try:
            grace[name] = int(get(a.registry, name + "()"), 16)
        except SystemExit:  # a registry from before the grace window
            grace[name] = None
        if want_grace is not None:
            check(f"registry.{name} == {a.pins_from}.{name}", grace[name] == want_grace[name], f"{grace[name]} s")
        if expect_grace is not None:
            ok_g = grace[name] == expect_grace[name]
            check(f"registry.{name} == --grace", ok_g, f"{grace[name]} s" + ("" if ok_g else f" (expected {expect_grace[name]})"))
        if want_grace is None and expect_grace is None:
            print(f"{name:13} {'n/a' if grace[name] is None else str(grace[name]) + ' s'}")
    record["grace"] = grace

    if a.broadcast:
        try:
            record.update(broadcast_meta(a.broadcast, a.registry))
        except (OSError, KeyError, ValueError, SystemExit) as e:
            check("broadcast record matches the registry", False, str(e))
    if a.record:
        if ok:
            Path(a.record).parent.mkdir(parents=True, exist_ok=True)
            Path(a.record).write_text(json.dumps(record, indent=2) + "\n")
            print(f"record   {a.record}")
        else:
            print("record   not written: a check failed")
    sys.exit(0 if ok else 1)


if __name__ == "__main__":
    main()
