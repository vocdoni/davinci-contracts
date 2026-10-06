# Deploying

[`script/DeployAll.s.sol`](../script/DeployAll.s.sol) deploys `ZiskVerifier` and then
`ProcessRegistry`, which creates the `DavinciDKGAdapter` itself when a DKG manager is
configured. The target chain needs EIP-4844: the `BLOBHASH` opcode and the point-evaluation
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
    uint32 defaultGrace,
    uint32 graceFloor,
    uint32 graceCeil,
    uint32 graceMaxTotal,
    uint32 noticeMin
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
| `DKG_MANAGER` | optional davinci-dkg `DKGManager`; unset or zero disables the DKG modes |
| `COUNCIL_MANAGER` | optional Council manager; unset or zero disables the `COUNCIL` mode |
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
  BATCH_PROGRAM_VK=0x6cfc89d562d0b22f04478a5c15b390433eb52f1b03147030b183076260da7a10 \
  RESULTS_PROGRAM_VK=0x7bc8c5e9235548386a44b1885732a2a7ffb1badddc8c7fba599d07ece47be794 \
  ROOT_C_VADCOP_FINAL=0x05006517b6ccde5da4d890587ba62845b5af8a307c00e87d4b9d05099b16dc80 \
  BALLOT_VK_HASH=0xbf1e6590bb1ba883d601c4d7d1c6fa2722a78590716874019db6d68fc776bb0e
forge script script/DeployAll.s.sol --rpc-url http://localhost:8545 --broadcast
```

## Public chains

Single chain, with the variables above in `.env`:

```bash
./deploy_all.sh
```

It runs the script with `--broadcast --slow` and the optimizer settings of `foundry.toml`,
verifies the sources according to `VERIFY_MODE`, then regenerates `golang-types/addresses.go`
from the broadcast logs.

Several chains: put shared values and `DEPLOY_CHAINS=base,sepolia,...` in `.env`, and
`CHAIN_ID`, `RPC_URL` and anything chain-specific in `.env.<chain>` (or `.env-<chain>`), then:

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
docker compose --profile deploy rm -f deploy
helpers/write_contract_addresses.sh
```

## DKG support

1. Deploy the [davinci-dkg](https://github.com/vocdoni/davinci-dkg) contracts.
2. Deploy with `DKG_MANAGER` set. The registry constructor creates the `DavinciDKGAdapter`,
   which reads `appManager()` from the manager; the script logs its address and
   `dkgAdapter()` returns it.

## Council support

Deploy with `COUNCIL_MANAGER` set to a Council manager (vocdoni/davinci-dkg-council). The registry
constructor creates the `CouncilAdapter` after the DKG adapter, so the DKG adapter's address
does not move; the script logs it and `councilAdapter()` returns it. Each ceremony organizer
then allows that adapter and authorizes the process creators on the manager.

## After a deployment

Forge writes the record to `broadcast/DeployAll.s.sol/<chain id>/run-latest.json`, including the
commit it was built from. The repository tracks that file for every public chain, and
`golang-types/addresses.go` is generated from it. Check the result with
[verification.md](verification.md).
