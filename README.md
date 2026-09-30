# DAVINCI contracts

Solidity contracts of the [DAVINCI](https://davinci.vote) voting protocol. Organizers create
voting processes in the `ProcessRegistry`, and sequencers settle each batch of votes and the
final tally on it with zero-knowledge proofs.

[![test](https://github.com/vocdoni/davinci-contracts/actions/workflows/test.yml/badge.svg?branch=zkvm)](https://github.com/vocdoni/davinci-contracts/actions/workflows/test.yml)
[![License: AGPL v3](https://img.shields.io/badge/License-AGPL%20v3-blue.svg)](https://www.gnu.org/licenses/agpl-3.0)

This code is a work in progress and not meant for production use yet. The protocol is
described in the [whitepaper](https://whitepaper.vocdoni.io).

## Overview

An organizer registers a process with its census, ballot mode, timing, metadata and the public
key ballots are encrypted to. Voters send encrypted ballots to a sequencer
([davinci-sequencer](https://github.com/vocdoni/davinci-sequencer)), which groups them into
batches and proves each state transition with
[davinci-zkvm](https://github.com/vocdoni/davinci-zkvm) as a ZisK PLONK. The sequencer settles it
with `submitStateTransition` in a blob transaction: the EIP-4844 blobs publish the new state,
and the registry checks the proof, the chain of state roots, the census root and that the blobs
are the ones the proof covers. Settlement is permissionless, so any sequencer can settle any
process.

Voting ends at `startTime + duration`. Batches still in flight keep settling during a short
grace window, and once it closes the tally is published over the final state root. With a
sequencer key it comes with a second PLONK from the zkVM results guest (`setProcessResults`).
With a key from a [davinci-dkg](https://github.com/vocdoni/davinci-dkg) committee, the committee
threshold-decrypts it (`requestResultsDecryption`, then `finalizeResultsFromDKG`).

| Contract | Role |
|---|---|
| [`ProcessRegistry`](src/ProcessRegistry.sol) | Processes and their lifecycle, transition settlement, results. |
| [`ZiskVerifier`](src/verifiers/ZiskVerifier.sol) | ZisK PLONK verifier, vendored from the ZisK snark setup. |
| [`DavinciDKGAdapter`](src/DavinciDKGAdapter.sol) | The registry's link to davinci-dkg, created by the registry when a DKG manager is configured. |
| [`src/libraries/`](src/libraries) | Genesis root, SHA-256 sparse Merkle tree, public values, blobs, process ids, BabyJubJub forms. |
| [`src/interfaces/`](src/interfaces) | `IProcessRegistry`, `IZiskVerifier`, `ICensusValidator` (on-chain censuses) and the davinci-dkg subset the adapter calls. |

## Deployments

Gnosis Chain (chain id 100), built from commit `9b03f18`:

| Contract | Address |
|---|---|
| `ZiskVerifier` | `0x150547716bD6f15D872508b66b2ae7ce17677C9C` |
| `ProcessRegistry` | `0x6702e0141B6b72bCF8C1bdff20A82A35C5502E7D` |
| `DavinciDKGAdapter` | `0xE9559c78E7ff8c19937A0657a092A221E90CCBC3` |
| davinci-dkg `DKGManager` | `0x9999F38Ff8Bf959E98Ddd5D4551f82775219c01B` |

The registry is pinned to:

| Immutable | Value |
|---|---|
| `batchProgramVK` | `0x6cfc89d562d0b22f04478a5c15b390433eb52f1b03147030b183076260da7a10` |
| `resultsProgramVK` | `0x7bc8c5e9235548386a44b1885732a2a7ffb1badddc8c7fba599d07ece47be794` |
| `rootCVadcopFinal` | `0x05006517b6ccde5da4d890587ba62845b5af8a307c00e87d4b9d05099b16dc80` |
| `ballotVKHash` | `0xbf1e6590bb1ba883d601c4d7d1c6fa2722a78590716874019db6d68fc776bb0e` |
| `defaultGrace` / `graceFloor` / `graceCeil` | 180 / 150 / 600 s |
| `graceMaxTotal` | 1800 s |
| `noticeMin` | 60 s |

[docs/verification.md](docs/verification.md) shows how to check the deployment against the
source.

## Quick start

A local chain with the contracts deployed and pinned to the davinci-zkvm release. It needs
Docker, and [Foundry](https://getfoundry.sh/) for `cast`:

```bash
git clone --recurse-submodules --branch zkvm https://github.com/vocdoni/davinci-contracts.git
cd davinci-contracts
docker compose --profile local up -d
docker compose --profile local logs -f deploy-local   # waits for the deployment
cast call 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512 "batchProgramVK()(bytes32)" --rpc-url http://localhost:8545
```

anvil listens on port 8545 with chain id 31337. `ZiskVerifier` is at
`0x5FbDB2315678afecb367f032d93F642f64180aa3` and `ProcessRegistry` at
`0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512`. `docker compose --profile local down` removes the
chain.

## Usage

### Creating a process

The organizer asks the registry for its next process id, gets the election key for that id from
a sequencer, and calls `newProcess`. On the Quick start chain, with a davinci-sequencer on its
default port:

```bash
RPC=http://localhost:8545
REGISTRY=0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512
ORGANIZER_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80  # anvil account 0
SEQUENCER=http://localhost:9090

PID=$(cast call $REGISTRY "getNextProcessId(address)(bytes31)" \
  $(cast wallet address $ORGANIZER_KEY) --rpc-url $RPC)
KEY=$(curl -s -X POST $SEQUENCER/processes/keys -H 'Content-Type: application/json' \
  -d "{\"processId\":\"$PID\"}")

CENSUS_ROOT=0x2a3f0c1d9e8b7a6f5e4d3c2b1a09f8e7d6c5b4a3928170f6e5d4c3b2a1908f7e  # lean-IMT root
printf '{"title":"Example"}' > metadata.json  # serve it at the metadata URI
METADATA_HASH=0x$(sha256sum metadata.json | cut -d' ' -f1)

cast send $REGISTRY "newProcess(uint8,uint256,uint256,uint256,\
(bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,bytes32,address,string,bool),\
string,bytes32,(uint256,uint256),(uint8,bytes12,uint256,uint256,uint256,uint256,uint256))" \
  0 0 3600 1000 \
  "(false,3,1,1,1,0,1,0)" \
  "(1,$CENSUS_ROOT,0x0000000000000000000000000000000000000000,https://example.org/census,false)" \
  https://example.org/metadata.json $METADATA_HASH \
  "($(echo $KEY | jq -r .x),$(echo $KEY | jq -r .y))" \
  "(0,0x000000000000000000000000,0,0,0,0,0)" \
  --private-key $ORGANIZER_KEY --rpc-url $RPC
```

The arguments, in order: status `READY` (0), start now (0), one hour, at most 1000 voters; a
ballot of three fields worth 0 or 1 with at most one set; an off-chain Merkle census (origin 1)
with its root and URI; the metadata URI and the SHA-256 of the exact bytes served there; the
election key; and all-zero DKG parameters, which select sequencer mode. For the other census
origins, the DKG key modes and every rule the registry enforces, see
[docs/registry.md](docs/registry.md#creating-a-process).

### Managing a process

The organizer can pause, resume, end or cancel a process with `setProcessStatus`, and until the
end time change its duration (`setProcessDuration`, shortening only with `noticeMin` seconds of
notice), grace window (`setProcessGrace`), voter cap (`setProcessMaxVoters`), census root
(`setProcessCensus`, dynamic off-chain censuses only) and metadata (`setProcessMetadata`).
Batches keep settling after the end until `getProcessGraceEnd(processId)`, and results are
accepted from then on.

### Settling transitions and results

These calls are permissionless and made by sequencers:

- `submitStateTransition(processId, publicValues, proofBytes, commitments, ys, kzgProofs)`, sent
  as a blob transaction carrying the transition's blobs;
- `setProcessResults(processId, publicValues, proofBytes)` in sequencer mode;
- `requestResultsDecryption(processId, accumulator, siblings)` and then
  `finalizeResultsFromDKG(processId)` in the DKG modes, plus `revealProcessKey` for
  `DKG_LOCKED`.

[docs/settlement.md](docs/settlement.md) lists what each one checks.

### Reading state

`getProcess(processId)` returns the whole process: status, times, census, ballot mode, key,
`latestStateRoot`, voter counts and, once published, `result`. The registry emits an event for
every change (`ProcessCreated`, `ProcessStatusChanged`, `ProcessStateTransitioned`,
`ProcessResultsSet` and others), so a process can be followed from the log alone:

```bash
cast call $REGISTRY "getProcessGraceEnd(bytes31)(uint256)" $PID --rpc-url $RPC
cast logs --address $REGISTRY "ProcessStateTransitioned(bytes31 indexed,address indexed,bytes32,bytes32,uint256,uint256,uint256)" $PID --from-block 0 --rpc-url $RPC
```

Contracts can call the registry through [`IProcessRegistry`](src/interfaces/IProcessRegistry.sol)
and the types in [`DAVINCITypes`](src/libraries/DAVINCITypes.sol). Applications can use the
generated bindings: TypeChain types in `typechain-types/` and Go bindings in `golang-types/`
(package `contracts`). The npm package `@vocdoni/davinci-contracts` is cut from the `release`
branch and can lag the contracts described here; `npm install && npm run prepare` builds the
types of this checkout into `dist/`.

### Deploying

`./deploy_all.sh` deploys `ZiskVerifier` and `ProcessRegistry` with the settings in `.env`
(see [`.env.example`](.env.example)):

| Variable | Value |
|---|---|
| `PRIVATE_KEY`, `RPC_URL`, `CHAIN_ID` | deployer key and target chain |
| `BATCH_PROGRAM_VK`, `RESULTS_PROGRAM_VK` | program vks of the davinci-zkvm vote-batch and results guests |
| `ROOT_C_VADCOP_FINAL` | root of the ZisK setup; must match the vendored verifier |
| `BALLOT_VK_HASH` | hash of the ballot proof verification key |
| `DKG_MANAGER` | optional davinci-dkg manager; enables the DKG key modes |
| `GRACE_DEFAULT`, `GRACE_FLOOR`, `GRACE_CEIL`, `GRACE_MAX_TOTAL`, `NOTICE_MIN` | optional grace window and notice settings, in seconds (180, 150, 600, 1800, 60) |

Sequencers refuse a registry pinned to other values than the davinci-zkvm release they run; the
current ones are in the [deployments table](#deployments).
[docs/deployment.md](docs/deployment.md) covers source verification, several chains, Docker and
DKG support.

## Documentation

- [docs/registry.md](docs/registry.md): process ids, lifecycle, grace window, creation
  parameters, census origins, key modes, organizer controls, reading state and events.
- [docs/settlement.md](docs/settlement.md): what `submitStateTransition` and the results calls
  check.
- [docs/errors.md](docs/errors.md): every custom error and when it is raised.
- [docs/deployment.md](docs/deployment.md): constructor, configuration, local and public
  deployments.
- [docs/verification.md](docs/verification.md): checking a live deployment against the source.
- [docs/development.md](docs/development.md): tests, fixtures, bindings, linters.
- [CHANGELOG.md](CHANGELOG.md): interface changes.

## Development

```bash
git clone --recurse-submodules --branch zkvm https://github.com/vocdoni/davinci-contracts.git
cd davinci-contracts
forge build
forge test
```

[docs/development.md](docs/development.md) describes the test suites, how the Go-generated
fixtures are regenerated, the bindings and the linters.

## License

GNU Affero General Public License v3.0, see [LICENSE.md](LICENSE.md). Source files carry their
own SPDX identifiers.
