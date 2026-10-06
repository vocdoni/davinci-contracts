#!/usr/bin/env python3
"""Checks a live ProcessRegistry deployment against this checkout's build.

Compares the on-chain runtime code of ZiskVerifier and ProcessRegistry with
`forge build` output (immutable slots masked), then reads every immutable
back and compares it with the expected pins; the grace and notice settings are
printed, not compared. Exit status 1 on any mismatch.

    python3 script/verify_deployment.py --rpc URL --registry 0x... \
        --batch-vk 0x... --results-vk 0x... --root-c 0x... --ballot-vk-hash 0x...
"""
import argparse
import json
import re
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
    import subprocess

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


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--rpc", required=True)
    ap.add_argument("--registry", required=True)
    ap.add_argument("--chain-id", type=int)
    ap.add_argument("--batch-vk", required=True)
    ap.add_argument("--results-vk", required=True)
    ap.add_argument("--root-c", required=True)
    ap.add_argument("--ballot-vk-hash", required=True)
    ap.add_argument("--cast", default=str(Path.home() / ".foundry/bin/cast"))
    a = ap.parse_args()

    ok = True

    def check(label, cond, detail=""):
        nonlocal ok
        print(f"{'OK  ' if cond else 'FAIL'} {label} {detail}".rstrip())
        ok &= bool(cond)

    chain = int(rpc(a.rpc, "eth_chainId", []), 16)
    if a.chain_id is not None:
        check("rpc chain id", chain == a.chain_id, str(chain))

    def get(to, sig):
        return rpc(a.rpc, "eth_call", [{"to": to, "data": selector(sig, a.cast)}, "latest"])

    word = lambda h: "0x" + h[2:].rjust(64, "0")[-64:]
    verifier = "0x" + get(a.registry, "ziskVerifier()")[-40:]

    reg_code = rpc(a.rpc, "eth_getCode", [a.registry, "latest"])
    ver_code = rpc(a.rpc, "eth_getCode", [verifier, "latest"])
    m, why = masked_match(reg_code, artifact("ProcessRegistry"))
    check("ProcessRegistry runtime code == local build (immutables masked)", m, why)
    m, why = masked_match(ver_code, artifact("ZiskVerifier"))
    check("ZiskVerifier runtime code == local build", m, why)

    check("registry.batchProgramVK", word(get(a.registry, "batchProgramVK()")) == a.batch_vk.lower())
    check("registry.resultsProgramVK", word(get(a.registry, "resultsProgramVK()")) == a.results_vk.lower())
    check("registry.rootCVadcopFinal", word(get(a.registry, "rootCVadcopFinal()")) == a.root_c.lower())
    check("verifier.getRootCVadcopFinal", word(get(verifier, "getRootCVadcopFinal()")) == a.root_c.lower())
    check("registry.ballotVKHash", word(get(a.registry, "ballotVKHash()")) == a.ballot_vk_hash.lower())
    reg_chain = int(get(a.registry, "chainID()"), 16)
    check("registry.chainID == rpc chain id", reg_chain == chain, str(reg_chain))

    adapter = "0x" + get(a.registry, "dkgAdapter()")[-40:]
    if int(adapter, 16) != 0:
        adp_code = rpc(a.rpc, "eth_getCode", [adapter, "latest"])
        m, why = masked_match(adp_code, artifact("DavinciDKGAdapter"))
        check("DavinciDKGAdapter runtime code == local build (immutables masked)", m, why)
        adp_reg = "0x" + get(adapter, "registry()")[-40:]
        check("adapter.registry == registry", adp_reg.lower() == a.registry.lower(), adp_reg)
        print(f"adapter  {adapter}")
    else:
        print("adapter  none (DKG disabled)")
    try:
        council = council_adapter(a.rpc, a.registry, selector("councilAdapter()", a.cast))
    except SystemExit as e:  # unreadable is a failure, never "no adapter"
        check("registry.councilAdapter() readable", False, str(e))
        council = "unreadable"
    if council == "unreadable":
        pass
    elif council is None:
        print("council none (no councilAdapter(): a registry from before the COUNCIL mode)")
    elif int(council, 16) != 0:
        cad_code = rpc(a.rpc, "eth_getCode", [council, "latest"])
        m, why = masked_match(cad_code, artifact("CouncilAdapter"))
        check("CouncilAdapter runtime code == local build (immutables masked)", m, why)
        cad_reg = "0x" + get(council, "registry()")[-40:]
        check("councilAdapter.registry == registry", cad_reg.lower() == a.registry.lower(), cad_reg)
        print(f"council {council} (manager 0x{get(council, 'manager()')[-40:]})")
    else:
        print("council none (COUNCIL disabled)")
    print(f"verifier {verifier}")
    for name in ("defaultGrace", "graceFloor", "graceCeil", "graceMaxTotal", "noticeMin"):
        try:
            print(f"{name:13} {int(get(a.registry, name + '()'), 16)} s")
        except SystemExit:  # a registry from before the grace window
            print(f"{name:13} n/a")
    sys.exit(0 if ok else 1)


if __name__ == "__main__":
    main()
