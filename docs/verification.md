# Checking a deployment

[`script/verify_deployment.py`](../script/verify_deployment.py) compares a live deployment with
a local build. It checks that:

- the runtime code of `ProcessRegistry` and `ZiskVerifier` matches the local build, with
  immutables masked;
- the registry's `batchProgramVK`, `resultsProgramVK`, `rootCVadcopFinal` and `ballotVKHash`
  equal the given pins, and so does the verifier's `getRootCVadcopFinal()`;
- the registry's `chainID` equals the RPC's chain id (and `--chain-id`, when given);
- with `--verifier`, the registry's `ziskVerifier()` is that address;
- when `dkgAdapter()` is set, the adapter's code matches the local build, `adapter.registry()`
  is the registry, `adapter.appManager()` is its manager's `appManager()`, `aidFor` returns ids
  whose low 160 bits are the adapter (the namespace of vocdoni/davinci-dkg#14) and, with
  `--dkg-manager`, `adapter.manager()` is that address;
- when `councilAdapter()` is set, its code matches the local build, it points back at the
  registry, its manager has code and, with `--council-manager`, is that address. A
  `councilAdapter()` that cannot be read fails the check;
- `--dkg-manager 0x0` or `--council-manager 0x0` expect that mode disabled.

It prints the grace and notice settings (`defaultGrace`, `graceFloor`, `graceCeil`,
`graceMaxTotal`, `noticeMin`) without judging them, except with `--grace
DEFAULT,FLOOR,CEIL,MAX_TOTAL,NOTICE` (seconds), which expects those values, or with
`--pins-from <registry>`: then the pins, the verifier and the grace settings are expected to
equal that registry's (a deployment made with `PINS_FROM_REGISTRY`), and any pin flag also given
must agree with it. In L1-blob mode it compares the `L1BlobVerifier`'s code with the local build
and its `l1Block()` with `--l1-block` (default the OP predeploy). `--l1-blob-verifier ADDR|0x0`
expects the registry's `l1BlobVerifier()` to be that address (`0x0`: blob mode); an unreadable
getter fails the run. Each
check prints `OK` or `FAIL`, and the exit status is 1 if any fails. `--record <file>` writes
what was verified as JSON when every check passes, with `--broadcast <run-latest.json>` adding
the commit, deployer, block and transaction hashes; `deploy_all.sh` does this after every
deployment.

## Running it

It needs Python 3, `cast` (at `~/.foundry/bin/cast`, or pass `--cast <path>`) and a
`forge build` of the source the deployment was built from, since it reads `out/`. The runtime
code ends in a hash of the compiler metadata, which covers every source file of the contract,
comments included, so build from the exact commit. Forge records it as `commit` in
`broadcast/DeployAll.s.sol/<chain id>/run-latest.json`; the
[deployments table](../README.md#deployments) lists it too.

For the Gnosis production beta registry (the values of `deployments/100.json`), with `RPC` a
Gnosis Chain JSON-RPC endpoint:

```bash
git checkout 4e4f6c8
forge build
python3 script/verify_deployment.py --rpc "$RPC" --chain-id 100 \
    --registry 0x20b96e465CA7C3536B9C733571ec1eCf42b2eA21 \
    --pins-from 0x6702e0141B6b72bCF8C1bdff20A82A35C5502E7D \
    --batch-vk 0x6cfc89d562d0b22f04478a5c15b390433eb52f1b03147030b183076260da7a10 \
    --results-vk 0x7bc8c5e9235548386a44b1885732a2a7ffb1badddc8c7fba599d07ece47be794 \
    --root-c 0x05006517b6ccde5da4d890587ba62845b5af8a307c00e87d4b9d05099b16dc80 \
    --ballot-vk-hash 0xbf1e6590bb1ba883d601c4d7d1c6fa2722a78590716874019db6d68fc776bb0e \
    --dkg-manager 0xC6Fb38c42ed3FB35D363a702218746d5C7Da36BF \
    --council-manager 0x2f5b110864cbad4017fe8ac59111812278f5f71f
```

For the OP Mainnet ZisK 1.3.2 registry, with `RPC` an OP Mainnet JSON-RPC endpoint (the registry and verifier
addresses are TBD until the deploy step fills them, and its record `deployments/10-<registry>.json` is not written yet):

```bash
python3 script/verify_deployment.py --rpc "$RPC" --chain-id 10 \
    --registry <TBD registry> \
    --verifier <TBD verifier> \
    --batch-vk 0xe39e460085ecddecf365649a61183a71c0fe778aa9c06fd2aa1a6c43cddb988b \
    --results-vk 0x394694558f39d0b2da0bbb09322333a0f567ee8d16bea2f354543a086587d167 \
    --root-c 0x6b2490fba336fb00f0c8719a961de309a8db6f43a9bfb7575ca5f014d609b265 \
    --ballot-vk-hash 0xbf1e6590bb1ba883d601c4d7d1c6fa2722a78590716874019db6d68fc776bb0e \
    --dkg-manager 0x0 --council-manager 0x0 \
    --l1-block 0x4200000000000000000000000000000000000015 --grace 300,240,600,1800,60 \
    --l1-blob-verifier 0x3B95BC14F43ed8Addb83F84D10A949De74eE7529
```

For the retired OP Mainnet registry (`deployments/10.json`, v0.2.0 vks, ZisK 1.3.0):

```bash
git checkout 44fdca1
forge build
python3 script/verify_deployment.py --rpc "$RPC" --chain-id 10 \
    --registry 0xC5f88160a1Df51A3Ee56b970b12B8308799c05A7 \
    --verifier 0x6781FFF50D615CB3F331CCD3Bdad3416d4ce0d5c \
    --batch-vk 0xa0ad5c4c7e39d4a7007933158181418a752e6c1563a211af84c579640db435fa \
    --results-vk 0xe693589453ecaee63eba697c48d4c5321a714bb18c2d84a97437b3f368578f68 \
    --root-c 0x05006517b6ccde5da4d890587ba62845b5af8a307c00e87d4b9d05099b16dc80 \
    --ballot-vk-hash 0xbf1e6590bb1ba883d601c4d7d1c6fa2722a78590716874019db6d68fc776bb0e \
    --dkg-manager 0x0 --council-manager 0x0 \
    --l1-block 0x4200000000000000000000000000000000000015 --grace 300,240,600,1800,60 \
    --l1-blob-verifier 0x3B95BC14F43ed8Addb83F84D10A949De74eE7529
```

For the previous Gnosis registry:

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
(`ZISK_VERIFIER_CODEHASH`, `0xadab570d3e3a4c172d0c75e5aef5d06098b94d70a9c8ff76af239a90fdcadcc1` for ZisK 1.3.2-alpha; the Gnosis verifier below is the older 1.3.0 one), and sequencers refuse a registry whose verifier has other code. To
compare:

```bash
cast keccak $(cast code 0x150547716bD6f15D872508b66b2ae7ce17677C9C --rpc-url "$RPC")
```
