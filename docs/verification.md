# Checking a deployment

[`script/verify_deployment.py`](../script/verify_deployment.py) compares a live deployment with
a local build. It checks that:

- the runtime code of `ProcessRegistry` and `ZiskVerifier` matches the local build, with
  immutables masked;
- the registry's `batchProgramVK`, `resultsProgramVK`, `rootCVadcopFinal` and `ballotVKHash`
  equal the given pins, and so does the verifier's `getRootCVadcopFinal()`;
- the registry's `chainID` equals the RPC's chain id (and `--chain-id`, when given);
- when `dkgAdapter()` is set, the adapter's code matches the local build and
  `adapter.registry()` is the registry.

It also prints the grace and notice settings (`defaultGrace`, `graceFloor`, `graceCeil`,
`graceMaxTotal`, `noticeMin`) without judging them. Each check prints `OK` or `FAIL`, and the
exit status is 1 if any fails.

## Running it

It needs Python 3, `cast` (at `~/.foundry/bin/cast`, or pass `--cast <path>`) and a
`forge build` of the source the deployment was built from, since it reads `out/`. The runtime
code ends in a hash of the compiler metadata, which covers every source file of the contract,
comments included, so build from the exact commit. Forge records it as `commit` in
`broadcast/DeployAll.s.sol/<chain id>/run-latest.json`; the
[deployments table](../README.md#deployments) lists it too.

For the Gnosis deployment, with `RPC` a Gnosis Chain JSON-RPC endpoint:

```bash
git checkout 9b03f18
forge build
python3 script/verify_deployment.py --rpc "$RPC" --chain-id 100 \
    --registry 0x6702e0141B6b72bCF8C1bdff20A82A35C5502E7D \
    --batch-vk 0x6cfc89d562d0b22f04478a5c15b390433eb52f1b03147030b183076260da7a10 \
    --results-vk 0x7bc8c5e9235548386a44b1885732a2a7ffb1badddc8c7fba599d07ece47be794 \
    --root-c 0x05006517b6ccde5da4d890587ba62845b5af8a307c00e87d4b9d05099b16dc80 \
    --ballot-vk-hash 0xbf1e6590bb1ba883d601c4d7d1c6fa2722a78590716874019db6d68fc776bb0e
```

## Where the pins come from

The script only compares the registry with the values it is given, so those should come from
an independent source:

- `--batch-vk` and `--results-vk` are the program vks of the davinci-zkvm vote-batch and
  results guests: what `cargo-zisk setup` prints as `Root hash` for each ELF, and what
  davinci-zkvm pins in `rust-sdk/src/release.rs`.
- `--root-c` is the root of the ZisK vadcop-final setup the PLONK proofs are wrapped under. The
  vendored `ZiskVerifier` hardcodes it, and davinci-zkvm pins it next to the vks.
- `--ballot-vk-hash` is the `sha256` of the ballot proof verification key, genesis leaf `0x07`
  (`davinci.BallotVKLeaf` in the davinci-zkvm Go SDK).

davinci-zkvm also pins the `keccak256` of the `ZiskVerifier` runtime code
(`ZISK_VERIFIER_CODEHASH`), and sequencers refuse a registry whose verifier has other code. To
compare:

```bash
cast keccak $(cast code 0x150547716bD6f15D872508b66b2ae7ce17677C9C --rpc-url "$RPC")
```
