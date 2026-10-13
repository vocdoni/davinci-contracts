# Deploying

[`script/DeployAll.s.sol`](../script/DeployAll.s.sol) deploys `ZiskVerifier` (or reuses a
deployed one) and then `ProcessRegistry`, which creates the `DavinciDKGAdapter` itself when a
DKG manager is configured and the `CouncilAdapter` when a Council manager is. The target chain needs EIP-4844: the `BLOBHASH` opcode and the point-evaluation
precompile at `0x0a`.

```solidity
constructor(
    uint32 chainID,
    address ziskVerifier,
    bytes32 batchProgramVK,
    bytes32 resultsProgramVK,
    bytes32 rootCVadcopFinal,
    bytes32 ballotVKHash,
    address dkgManager,
    address councilManager,
    uint32 defaultGrace,
    uint32 graceFloor,
    uint32 graceCeil,
    uint32 graceMaxTotal,
    uint32 noticeMin,
    address l1BlobVerifier
)
```

## Configuration

`DeployAll.s.sol` reads:

| Variable | Value |
|---|---|
| `PRIVATE_KEY` | deployer key |
| `CHAIN_ID` | chain id, stored in the registry and folded into process ids; must fit in `uint32` |
| `BATCH_PROGRAM_VK` | program vk of the davinci-zkvm vote-batch guest (`bytes32`) |
| `RESULTS_PROGRAM_VK` | program vk of the davinci-zkvm results guest |
| `ROOT_C_VADCOP_FINAL` | root of the ZisK vadcop-final setup; the script aborts unless it equals `ZiskVerifier.getRootCVadcopFinal()` |
| `BALLOT_VK_HASH` | `sha256` of the ballot proof VK wire bytes, genesis leaf `0x07` |
| `ZISK_VERIFIER` | optional deployed `ZiskVerifier` to reuse instead of deploying one; it must report `ROOT_C_VADCOP_FINAL`, and the script logs its code hash (the sequencers' `ZISK_VERIFIER_CODEHASH` pin) |
| `PINS_FROM_REGISTRY` | optional live `ProcessRegistry` on the same chain to inherit from: its `ZiskVerifier` is reused and its four pins and five grace settings copied, so none of those variables is needed. Any of them that is also set must equal the inherited value, otherwise the script aborts before broadcasting |
| `DKG_MANAGER` | optional davinci-dkg `DKGManager`; unset or zero disables the DKG modes. It must have code and an `appManager()`. Use a davinci-dkg deployment that binds application ids to their registrant ([vocdoni/davinci-dkg#14](https://github.com/vocdoni/davinci-dkg/issues/14)): an older one accepts the adapter's ids too, but lets anyone take a process's id first and block its creation |
| `COUNCIL_MANAGER` | optional Council manager; unset or zero disables the `COUNCIL` mode. It must have code and differ from `DKG_MANAGER` |
| `L1_BLOB_VERIFIER` | optional deployed `L1BlobVerifier`; non-zero makes the registry an L1-blob-mode one (see [settlement.md](settlement.md#l1-blob-mode)). Unset or zero is blob mode |
| `L1_BLOCK` | optional OP Stack `L1Block` address (`0x4200000000000000000000000000000000000015` on OP chains): when `L1_BLOB_VERIFIER` is unset, the script deploys `new L1BlobVerifier(L1_BLOCK)` before the registry and uses it. It must have code. A mock on a local chain |
| `GRACE_DEFAULT` | optional, seconds: the grace window of a new process (`defaultGrace`, default 180) |
| `GRACE_FLOOR` | optional, seconds: the minimum for `setProcessGrace` (`graceFloor`, default 150) |
| `GRACE_CEIL` | optional, seconds: the maximum for `setProcessGrace` (`graceCeil`, default 600) |
| `GRACE_MAX_TOTAL` | optional, seconds: the cap on the window past the end (`graceMaxTotal`, default 1800) |
| `NOTICE_MIN` | optional, seconds: the minimum notice for shortening a process (`noticeMin`, default 60) |

`deploy_all.sh` also reads:

| Variable | Value |
|---|---|
| `RPC_URL` | JSON-RPC endpoint of the target chain |
| `VERIFY_MODE` | `auto` (default: verify sources except on chain ids 31337 and 1337), `true` or `false` |
| `ETHERSCAN_API_KEY` | explorer key for source verification |
| `ETHERSCAN_API_URL` | optional verifier endpoint for explorers other than Etherscan |
| `DEBUG_DEPLOY` | `true` traces the script |
| `EXPECT_DA_MODE` | `l1` or `blob`: the mode the run must produce, checked before `forge script` against `L1_BLOB_VERIFIER` / `L1_BLOCK` (a misspelled name otherwise silently deploys a blob-mode registry). Optional, but required on chains 10 and 11155420 |
| `EXPECT_GRACE` | optional `DEFAULT,FLOOR,CEIL,MAX_TOTAL,NOTICE` in seconds the registry must read back (e.g. `300,240,600,1800,60`); a misspelled `GRACE_*` silently keeps its default |

[`.env.example`](../.env.example) lists them all.

The pins must match the davinci-zkvm release the sequencers run, or they will refuse the
registry. The program vks are what `cargo-zisk setup` prints as `Root hash` for each guest ELF;
davinci-zkvm pins the released ones in `rust-sdk/src/release.rs`. `BALLOT_VK_HASH` is
`davinci.BallotVKLeaf` in the davinci-zkvm Go SDK. The values of the current release are the
ones in the [deployments table](../README.md#deployments).

A zero verifier address or pin reverts with `InvalidVerifierConfig`, and grace bounds that are
not `0 < GRACE_FLOOR <= GRACE_DEFAULT <= GRACE_CEIL <= GRACE_MAX_TOTAL`, or `NOTICE_MIN=0`,
with `InvalidGrace`. The defaults are production values; a local or test chain can deploy with
short ones, for example
`GRACE_DEFAULT=10 GRACE_FLOOR=2 GRACE_CEIL=60 GRACE_MAX_TOTAL=60 NOTICE_MIN=5`. The script
prints the values it deployed with.

## Local chain

With Docker, the `local` profile of `docker-compose.yml` starts anvil (chain id 31337, Osaka
hardfork) and deploys from anvil's account 0, pinned to the davinci-zkvm release and without
DKG:

```bash
docker compose --profile local up -d
docker compose --profile local down   # the chain is gone after this
```

The addresses are always `ZiskVerifier` `0x5FbDB2315678afecb367f032d93F642f64180aa3` and
`ProcessRegistry` `0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512`. anvil listens on `ANVIL_PORT`
(default 8545) and mines a block every `ANVIL_BLOCK_TIME` seconds (default 1); the image uses
Foundry `FOUNDRY_VERSION` (default v1.8.3). The grace and notice variables pass through from the
host environment (unset keeps the defaults), so exporting the short values above before `up`
gives a chain with short windows. A davinci-sequencer runs against it with
`--blob-source anvil`. The images are built from the checkout, submodules included, so check
them out first.

Without Docker:

```bash
anvil --chain-id 31337 --hardfork osaka
```

and in another shell:

```bash
export PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 CHAIN_ID=31337 \
  BATCH_PROGRAM_VK=0xa0ad5c4c7e39d4a7007933158181418a752e6c1563a211af84c579640db435fa \
  RESULTS_PROGRAM_VK=0xe693589453ecaee63eba697c48d4c5321a714bb18c2d84a97437b3f368578f68 \
  ROOT_C_VADCOP_FINAL=0x05006517b6ccde5da4d890587ba62845b5af8a307c00e87d4b9d05099b16dc80 \
  BALLOT_VK_HASH=0xbf1e6590bb1ba883d601c4d7d1c6fa2722a78590716874019db6d68fc776bb0e
forge script script/DeployAll.s.sol --rpc-url http://localhost:8545 --broadcast
```

## Public chains

Single chain, with the variables above in `.env`:

```bash
./deploy_all.sh
```

It runs the script with `--broadcast --slow` and the optimizer settings of `foundry.toml` and
verifies the sources according to `VERIFY_MODE`. It then runs
[`script/verify_deployment.py`](verification.md) against the new registry, with the pins (or
`PINS_FROM_REGISTRY`) and both managers as the expected values (unset expects that mode
disabled), and writes what it verified to `deployments/<chain id>.json`: addresses of the
registry, verifier, both adapters and their managers, the `ZiskVerifier` code hash, pins, grace
settings, commit, deployer, block and transaction hashes. A failed check stops the script before
it regenerates `golang-types/addresses.go` from the broadcast logs and leaves no record. It needs
`jq` and `python3` besides Foundry. `DeployAll` keeps its default for a grace variable that is
unset or misspelled, so set `EXPECT_GRACE` (`DEFAULT,FLOOR,CEIL,MAX_TOTAL,NOTICE`, for example
`300,240,600,1800,60`) as well: the script then also checks the deployed grace settings against
it, and writes no record on a difference.

Several chains: put shared values and `DEPLOY_CHAINS=base,sepolia,...` in `.env`, and
`CHAIN_ID`, `RPC_URL` and anything chain-specific in `.env.<chain>` (or `.env-<chain>`). Like
`GRACE_*` and `NOTICE_MIN`, `EXPECT_GRACE` is not reset between chains, so a chain that overrides
`GRACE_*` must override `EXPECT_GRACE` too. Then:

```bash
./deploy_all_contracts_to_all_chains.sh
```

It reloads `.env`, clears the chain-scoped variables (`DKG_MANAGER` and `COUNCIL_MANAGER`
among them), loads each chain's file and calls `deploy_all.sh`. The grace and notice variables
are not chain-scoped: a value set in one chain file carries over to the chains after it, so set
them in `.env` or in every chain file.

With Docker, the `deploy` profile runs `deploy_all.sh` with the variables from `.env`; the key
never enters the image. The broadcast record stays in the container, so copy it out before
removing it, then regenerate the Go addresses on the host:

```bash
docker compose --profile deploy up deploy
docker compose --profile deploy cp deploy:/app/broadcast/DeployAll.s.sol/. broadcast/DeployAll.s.sol/
docker compose --profile deploy cp deploy:/app/deployments/. deployments/
docker compose --profile deploy rm -f deploy
helpers/write_contract_addresses.sh
```

## DKG support

1. Deploy the [davinci-dkg](https://github.com/vocdoni/davinci-dkg) contracts.
2. Deploy with `DKG_MANAGER` set. The registry constructor creates the `DavinciDKGAdapter`,
   which reads `appManager()` from the manager; the script logs its address and
   `dkgAdapter()` returns it.

## L1-blob mode

For OP Stack chains. Deploy with `L1_BLOCK` set to the chain's `L1Block` predeploy, or with
`L1_BLOB_VERIFIER` set to a verifier deployed earlier (it is stateless and ownerless, so one can
serve several registries on the same chain). The verifier is created before the registry, which
does not change the other contracts' deployment order, but the registry's address then comes from a later nonce than the fixed local one above; the script logs `L1BlobVerifier deployed
at` and the verifier and `L1Block` addresses. `verify_deployment.py` reads `l1BlobVerifier()`,
compares its code with the local build and its `l1Block()` with `--l1-block` (default the OP
predeploy; `deploy_all.sh` passes `L1_BLOCK` when set) and, with `--l1-blob-verifier`, requires
the registry to hold the expected verifier (`deploy_all.sh` passes `L1_BLOB_VERIFIER`, else the
one it deployed, else `0x0` for blob mode). The mode is immutable: switching needs a
new registry. The target chain needs the point-evaluation precompile at `0x0a`, but not
`BLOBHASH` of its own blobs. A local dry run needs a stand-in `L1Block` (`test/mocks/MockL1Block.sol`).
See [settlement.md](settlement.md#l1-blob-mode) for the proof and its limits.

## OP Mainnet

Deployed on 2026-10-13 reusing the `ZiskVerifier` and the `L1BlobVerifier` of the first OP registry
(the verifier is stateless apart from its permissionless origin record), on the davinci-zkvm v0.2.0
pins and a grace window that fits the Ethereum leg, from the owner's account, with no `.env`
(variables set in a subshell; the RPC and the key never written to disk):

```bash
CHAIN_ID=10 VERIFY_MODE=false \
ZISK_VERIFIER=0x6781FFF50D615CB3F331CCD3Bdad3416d4ce0d5c \
L1_BLOB_VERIFIER=0x3B95BC14F43ed8Addb83F84D10A949De74eE7529 EXPECT_DA_MODE=l1 \
BATCH_PROGRAM_VK=0xa0ad5c4c7e39d4a7007933158181418a752e6c1563a211af84c579640db435fa \
RESULTS_PROGRAM_VK=0xe693589453ecaee63eba697c48d4c5321a714bb18c2d84a97437b3f368578f68 \
ROOT_C_VADCOP_FINAL=0x05006517b6ccde5da4d890587ba62845b5af8a307c00e87d4b9d05099b16dc80 \
BALLOT_VK_HASH=0xbf1e6590bb1ba883d601c4d7d1c6fa2722a78590716874019db6d68fc776bb0e \
GRACE_DEFAULT=300 GRACE_FLOOR=240 GRACE_CEIL=600 GRACE_MAX_TOTAL=1800 NOTICE_MIN=60 \
EXPECT_GRACE=300,240,600,1800,60 \
RPC_URL=... PRIVATE_KEY=... ./deploy_all.sh
```

The script ran `forge script` once without `--broadcast` first. The first OP registry
(2026-10-12, `L1_BLOCK` instead of `L1_BLOB_VERIFIER`, the v0.1 pins, the default grace) is
retired; its record is `deployments/10-0x8b63.json`. Addresses are in the
[deployments table](../README.md#deployments).

## Council support

Deploy with `COUNCIL_MANAGER` set to a Council manager (vocdoni/davinci-dkg-council). The registry
constructor creates the `CouncilAdapter` after the DKG adapter, so the DKG adapter's address
does not move; the script logs it and `councilAdapter()` returns it. Each ceremony organizer
then allows that adapter and authorizes the process creators on the manager.

## Gnosis production beta

The production beta registry on Gnosis takes the new davinci-dkg manager (with the
application id binding of vocdoni/davinci-dkg#14) and the Council manager, and inherits the
verifier, pins and grace settings of the previous production registry. It was deployed on
2026-10-07 with this `.env` (addresses in the [deployments table](../README.md#deployments)):

```bash
CHAIN_ID=100
RPC_URL=https://rpc.gnosischain.com
PRIVATE_KEY=...                     # the deployer key, never committed
ETHERSCAN_API_KEY=...               # Gnosisscan, for source verification
PINS_FROM_REGISTRY=0x6702e0141B6b72bCF8C1bdff20A82A35C5502E7D
DKG_MANAGER=0xC6Fb38c42ed3FB35D363a702218746d5C7Da36BF  # davinci-dkg v0.10.0, with #14
COUNCIL_MANAGER=0x2f5b110864cbad4017fe8ac59111812278f5f71f
```

```bash
./deploy_all.sh
```

The new registry reuses `ZiskVerifier` `0x150547716bD6f15D872508b66b2ae7ce17677C9C`. Commit
`deployments/100.json`, `broadcast/DeployAll.s.sol/100/run-latest.json` (which replaces the
previous registry's record) and `golang-types/addresses.go`, and fill in the
[deployments table](../README.md#deployments). Before broadcasting, a rehearsal on a fork runs
the same script against an Anvil fork of Gnosis with any funded key and `VERIFY_MODE=false`;
do it in a scratch checkout, since it rewrites the chain 100 broadcast record:

```bash
anvil --fork-url https://rpc.gnosischain.com --port 8546
# .env as above with RPC_URL=http://127.0.0.1:8546, VERIFY_MODE=false, an anvil key and, for
# DKG_MANAGER, a davinci-dkg suite deployed on the fork (davinci-dkg solidity/script/DeployAll.s.sol)
./deploy_all.sh
```

### Council test registry on Gnosis (retired)

A TEST registry for the Council round trip, deployed 2026-10-07 from this branch at `f4abc5d`
(block 48,627,101, tx `0xc3c4b4e6ff079457888f867f54404840c5ea71ed21009dcc982e1a3ef2013a04`),
reusing the production `ZiskVerifier` (`ZISK_VERIFIER`) and the production pins and grace
settings, so released sequencers and provers accept it:

| Contract | Address |
|---|---|
| `ProcessRegistry` | `0x847a16CC56E0Ef57FEc28735105941a0299cDC62` |
| `CouncilAdapter` | `0x4817493b792db101dcc75306754242ceFd928E40` (Council manager `0x2f5b110864cbad4017fe8ac59111812278f5f71f`, on a development trusted setup; the production beta uses it too) |
| `DavinciDKGAdapter` | `0x9d356d42eC5a04ABeaDA1AE31D83Eb431f120958` (davinci-dkg `DKGManager` `0x9999F38Ff8Bf959E98Ddd5D4551f82775219c01B`) |
| `ZiskVerifier` | `0x150547716bD6f15D872508b66b2ae7ce17677C9C` (reused) |

**Retired 2026-10-07.** Superseded by the production beta registry R2
(`0x20b96e465CA7C3536B9C733571ec1eCf42b2eA21`). No sequencer follows this registry.

The three new contracts are source-verified on Gnosisscan and Blockscout, and
`script/verify_deployment.py` passes against a build of `f4abc5d`. Its broadcast record is not
committed: `broadcast/DeployAll.s.sol/100/run-latest.json` stays the production one.

## After a deployment

Forge writes the record to `broadcast/DeployAll.s.sol/<chain id>/run-latest.json`, including the
commit it was built from. The repository tracks that file for every public chain, and
`golang-types/addresses.go` is generated from it; a run that reused a `ZiskVerifier` takes its
address from the registry's constructor arguments. `deploy_all.sh` also writes the verified
record `deployments/<chain id>.json`, tracked as well (not for local chain ids). Check the result
again at any time with [verification.md](verification.md).
