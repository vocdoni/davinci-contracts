#!/usr/bin/env python3
"""Checks a live ProcessRegistry deployment against this checkout's build.

Compares the on-chain runtime code of ZiskVerifier and ProcessRegistry with
`forge build` output (immutable slots masked), then reads every immutable
back and compares it with the expected pins. Exit status 1 on any mismatch.

    python3 script/verify_deployment.py --rpc URL --registry 0x... \
        --batch-vk 0x... --results-vk 0x... --root-c 0x... --ballot-vk-hash 0x...
"""
import argparse
import json
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
        # The registrar gate on the DKG side blocks aid front-running and pool
        # exhaustion; it is a manual owner call, so verify it was made.
        app_mgr = "0x" + get(adapter, "appManager()")[-40:]
        registrar = "0x" + get(app_mgr, "registrar()")[-40:]
        if int(registrar, 16) == 0:
            check("appManager.registrar set (gate enabled)", False, "registrar is zero")
        else:
            check("appManager.registrar == adapter", registrar.lower() == adapter.lower(), registrar)
        print(f"adapter  {adapter}")
    else:
        print("adapter  none (DKG disabled)")
    print(f"verifier {verifier}")
    sys.exit(0 if ok else 1)


if __name__ == "__main__":
    main()
