#!/usr/bin/env python3
"""Checks a live ProcessRegistry deployment against this checkout's build.

Compares the on-chain runtime code of ZiskVerifier, ProcessRegistry and both adapters
with `forge build` output (immutable slots masked), then reads every immutable back and
compares it with the expected pins, verifier and managers. The grace and notice settings
are compared with --pins-from and printed otherwise. Exit status 1 on any mismatch.
--record writes what was verified as JSON, only when every check passes.

    python3 script/verify_deployment.py --rpc URL --registry 0x... \
        --batch-vk 0x... --results-vk 0x... --root-c 0x... --ballot-vk-hash 0x...
    python3 script/verify_deployment.py --rpc URL --registry 0x... --pins-from 0x... \
        --dkg-manager 0x... --council-manager 0x... [--broadcast run-latest.json --record out.json]
"""
import argparse
import json
import re
import subprocess
import sys
import urllib.request
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
UA = {"Content-Type": "application/json", "User-Agent": "davinci-verify"}


def rpc(url, method, params):
    body = json.dumps({"jsonrpc": "2.0", "id": 1, "method": method, "params": params})
    req = urllib.request.Request(url, data=body.encode(), headers=UA)
    out = json.load(urllib.request.urlopen(req, timeout=30))
    if "error" in out:
        raise SystemExit(f"{method}: {out['error']}")
    return out["result"]


ADDRESS_WORD = re.compile(r"0x0{24}[0-9a-fA-F]{40}")


def council_adapter(url, registry, calldata):
    """The registry's councilAdapter(), or None when the registry has no such function.

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
        raise SystemExit(f"councilAdapter(): {e}")
    if not isinstance(out, dict):
        raise SystemExit(f"councilAdapter(): malformed response {out!r}")
    if "error" in out:
        err = out["error"]
        if (
            isinstance(err, dict)
            and "execution reverted" in str(err.get("message", "")).lower()
            and err.get("data") in (None, "", "0x")
        ):
            return None
        raise SystemExit(f"councilAdapter(): {err}")
    result = out.get("result")
    if result == "0x":
        return None
    if not isinstance(result, str) or not ADDRESS_WORD.fullmatch(result):
        raise SystemExit(f"councilAdapter(): malformed result {result!r}")
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
    ap.add_argument("--broadcast", help="forge run-latest.json of the deployment, for the record")
    ap.add_argument("--record", help="write the verified deployment as JSON here (only when every check passes)")
    ap.add_argument("--cast", default=str(Path.home() / ".foundry/bin/cast"))
    a = ap.parse_args()

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
        council = council_adapter(a.rpc, a.registry, selector("councilAdapter()", a.cast))
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
    print(f"verifier {verifier}")
    grace = {}
    for name in GRACE:
        try:
            grace[name] = int(get(a.registry, name + "()"), 16)
        except SystemExit:  # a registry from before the grace window
            grace[name] = None
        if want_grace is not None:
            check(f"registry.{name} == {a.pins_from}.{name}", grace[name] == want_grace[name], f"{grace[name]} s")
        else:
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
