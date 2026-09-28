# DAVINCI contracts (zkVM)

[![License: AGPL v3](https://img.shields.io/badge/License-AGPL%20v3-blue.svg)](https://www.gnu.org/licenses/agpl-3.0)

Solidity contracts for the DAVINCI voting protocol. The `ProcessRegistry` stores voting
processes and settles the proofs produced by
[davinci-zkvm](https://github.com/vocdoni/davinci-zkvm): every state transition is a ZisK PLONK
proof of the vote-batch guest, bound to the EIP-4844 blobs that publish it, and the tally is
either a PLONK proof of the results guest or a threshold decryption by a davinci-dkg committee.
Settlement and results calls are permissionless: the proofs, the root chain and the blob
openings authenticate them.

This code is a work in progress and not meant for production use yet. The protocol is
described in the [whitepaper](https://whitepaper.vocdoni.io).

## Contents

- [Contracts](#contracts)
- [Deployments](#deployments)
- [Processes](#processes)
- [Creating a process](#creating-a-process)
- [Organizer controls](#organizer-controls)
- [State transitions](#state-transitions)
- [Results](#results)
- [Reading state](#reading-state)
- [Events](#events)
- [Errors](#errors)
- [Deploying](#deploying)
- [Checking a deployment](#checking-a-deployment)
- [Development](#development)
- [License](#license)

## Contracts

| Path | Role |
|---|---|
| `src/ProcessRegistry.sol` | Process storage and lifecycle, transition settlement, results. |
| `src/verifiers/ZiskVerifier.sol`, `PlonkVerifier.sol` | ZisK PLONK verifier, vendored from the ZisK snark setup. |
| `src/DavinciDKGAdapter.sol` | The registry's link to davinci-dkg. Created by the registry constructor when a DKG manager is configured. |
| `src/libraries/GenesisLib.sol` | Genesis state root of a process and the encryption key check. |
| `src/libraries/Sha256SmtLib.sol` | SHA-256 sparse Merkle tree (vocdoni/arbo layout, 64 levels, 8-byte keys): roots and inclusion proofs. |
| `src/libraries/PublicsLib.sol` | Decoding of the 512-byte PLONK public values. |
| `src/libraries/BlobsLib.sol` | `BLOBHASH` and point-evaluation helpers. |
| `src/libraries/BjjFormLib.sol` | BabyJubJub conversion between the circomlib form and the reduced form davinci-dkg uses. |
| `src/libraries/ProcessIdLib.sol` | Process id layout. |
| `src/libraries/DAVINCITypes.sol` | Shared enums and structs. |
| `src/interfaces/` | `IProcessRegistry`, `IZiskVerifier`, `ICensusValidator` (on-chain census contracts) and, under `dkg/`, the subset of the davinci-dkg interfaces the adapter calls. |

## Deployments

Gnosis Chain (chain id 100):

| Contract | Address |
|---|---|
| `ZiskVerifier` | `0xAe7b632A72cf474039E4128354576770beA33f34` |
| `ProcessRegistry` | `0x3CDE68c39E26ecf94bD029b6ED3b9F945441daf3` |
| `DavinciDKGAdapter` | `0x21FDE45181d31CcefAA722CE648b4BB37dd7645c` |
| davinci-dkg `DKGManager` | `0x6fA82Ffe5dfAdCe7f9d538FDab648bd01d2E15E6` |

The registry is pinned to:

| Immutable | Value |
|---|---|
| `batchProgramVK` | `0x6cfc89d562d0b22f04478a5c15b390433eb52f1b03147030b183076260da7a10` |
| `resultsProgramVK` | `0x7bc8c5e9235548386a44b1885732a2a7ffb1badddc8c7fba599d07ece47be794` |
| `rootCVadcopFinal` | `0x05006517b6ccde5da4d890587ba62845b5af8a307c00e87d4b9d05099b16dc80` |
| `ballotVKHash` | `0xbf1e6590bb1ba883d601c4d7d1c6fa2722a78590716874019db6d68fc776bb0e` |

[Checking a deployment](#checking-a-deployment) shows how to compare these against a local build.

## Processes

### Process ids

A process id is a `bytes31`: the organizer address (20 bytes), a 4-byte registry prefix and a
7-byte per-organizer nonce. The prefix is the low 4 bytes of
`keccak256(abi.encodePacked(uint32 chainID, registry))`, so ids from another registry or chain
revert with `UnknownProcessIdPrefix`. `getNextProcessId(organizer)` returns the id the
organizer's next `newProcess` will get. The organizer is the `newProcess` caller, stored as
`organizationId`.

### Statuses

`ProcessStatus` is `READY` (0), `ENDED` (1), `CANCELED` (2), `PAUSED` (3), `RESULTS` (4). The
organizer moves a process with `setProcessStatus`:

| From | Organizer may set | Other ways out |
|---|---|---|
| `READY` | `PAUSED`, `CANCELED`, `ENDED` | `RESULTS` or `ENDED` from the results calls, once the end time has passed |
| `PAUSED` | `READY`, `CANCELED`, `ENDED` | same as `READY` |
| `ENDED` | none | `RESULTS` from the results calls |
| `CANCELED` | none | none |
| `RESULTS` | none | none |

A process ends at `startTime + duration` (`getProcessEndTime`). Setting `ENDED` by hand also
sets `duration` to the time elapsed since `startTime` (0 if it had not started) and emits
`ProcessDurationChanged`. Transitions settle only while the process is `READY` and
`startTime <= block.timestamp < end`; `PAUSED` stops settlement but not the clock.

Results are accepted once the process is `ENDED`, or `READY`/`PAUSED` with its end time passed.
In sequencer mode `setProcessResults` moves it straight to `RESULTS`. In the DKG modes
`requestResultsDecryption` first moves it to `ENDED`, which takes it out of the organizer's
hands, and `finalizeResultsFromDKG` moves it to `RESULTS`.

## Creating a process

```solidity
function newProcess(
    DAVINCITypes.ProcessStatus status,
    uint256 startTime,
    uint256 duration,
    uint256 maxVoters,
    DAVINCITypes.BallotMode calldata ballotMode,
    DAVINCITypes.Census calldata census,
    string calldata metadata,
    DAVINCITypes.EncryptionKey calldata encryptionKey,
    DAVINCITypes.DKGParams calldata dkg
) external returns (bytes31 processId);
```

| Parameter | Rule |
|---|---|
| `status` | `READY` or `PAUSED`. |
| `startTime` | 0 means now; otherwise not in the past. |
| `duration` | `startTime + duration` must be in the future. |
| `maxVoters` | Non-zero, and `ballotMode.maxValue <= 10^12 / maxVoters` so every possible tally stays inside the sequencer's bounded decryption search. |
| `ballotMode` | See below. |
| `census` | See [Census origins](#census-origins). |
| `metadata` | Metadata URI, stored as is. |
| `encryptionKey` | The ElGamal public key in sequencer mode; `(0, 0)` in the DKG modes. |
| `dkg` | Key mode and DKG registration arguments; all zero in sequencer mode. |

The call emits `ProcessCreated(processId, organizer)`.

### Ballot mode

```solidity
struct BallotMode {
    bool uniqueValues;
    uint8 numFields;
    uint8 groupSize;
    uint8 costExponent;
    uint256 maxValue;
    uint256 minValue;
    uint256 maxValueSum;  // 0: each voter's weight is the limit
    uint256 minValueSum;  // 0: no lower limit
}
```

`numFields` is 1 to 16, the ballot capacity of the zkVM guest. `groupSize` may not exceed
`numFields`, `minValue <= maxValue` and `minValueSum <= maxValueSum`. The ballot mode is packed
into one state leaf, so `maxValue` and `minValue` must fit in 48 bits and the two sums in 63
bits.

### Census origins

```solidity
struct Census {
    CensusOrigin censusOrigin;
    bytes32 censusRoot;
    address contractAddress;
    string censusURI;
    bool onchainAllowAnyValidRoot;
}
```

| Value | `CensusOrigin` | `censusRoot` | Changes |
|---|---|---|---|
| 1 | `MERKLE_TREE_OFFCHAIN_STATIC_V1` | lean-IMT root, big-endian, non-zero | Fixed. |
| 2 | `MERKLE_TREE_OFFCHAIN_DYNAMIC_V1` | lean-IMT root, big-endian, non-zero | The organizer replaces it with `setProcessCensus`. |
| 3 | `MERKLE_TREE_ONCHAIN_DYNAMIC_V1` | Ignored. The registry stores the census contract's `getCensusRoot()` for information. | A batch may use any root the census contract held from the process creation block on. |
| 4 | `CSP_EDDSA_BABYJUBJUB_V1` | CSP signer address as a `uint160` | Fixed. |

`contractAddress` is required for origin 3, where it must be an `ICensusValidator` contract with
code that answers `getCensusRoot()`, and must be zero for every other origin. `censusURI` must
not be empty: it tells the sequencer where to fetch a Merkle census, or voters where to get CSP
signatures. `onchainAllowAnyValidRoot` must be false. `CENSUS_UNKNOWN` (0) is rejected. The
origin 4 name is historical: the davinci-zkvm guest verifies an ECDSA (secp256k1) CSP
signature.

For origin 3 the registry calls `getRootBlockNumber(root)` on every settlement. The census
contract has to answer `block.number` for its current root, the block a replaced root was
replaced in, and 0 for a root it never held. A batch proven against an evicted root stops
settling, and the census has to be append-only with fixed weights;
`src/interfaces/ICensusValidator.sol` explains why.

### Key modes

```solidity
enum KeyMode { SEQUENCER, DKG_AUTOMATIC, DKG_LOCKED }

struct DKGParams {
    KeyMode mode;
    bytes12 epochId; // DKG_LOCKED only: the epoch the proof of possession binds
    uint256 orgPKx;  // DKG_LOCKED only: organizer key, reduced form
    uint256 orgPKy;
    uint256 popAx;   // DKG_LOCKED only: Schnorr proof of possession
    uint256 popAy;
    uint256 popZ;
}
```

- `SEQUENCER`: the caller supplies `encryptionKey` and every `DKGParams` field is zero. Results
  come from the results guest.
- `DKG_AUTOMATIC`: `encryptionKey` is `(0, 0)` and only `mode` is read. The adapter registers a
  davinci-dkg application on `registrationEpoch()`, the newest Live epoch with a free pool key
  (it scans back at most 8 epochs), and the process key is that pool key. Only the committee
  decrypts the tally.
- `DKG_LOCKED`: as above, but on the caller's `epochId` and with an organizer key: the process
  key is the pool key plus `PK_org`. The organizer key and proof of possession are in the DKG's
  reduced (a = -1) form, exactly as `DKGAppManager.registerApplication` takes them. The proof
  binds the epoch and the application id, so a client reads the epoch from the adapter's
  `registrationEpoch()` and the id from `aidFor(getNextProcessId(organizer))` before building
  it. The committee cannot finish decrypting until the organizer secret is published with
  `revealProcessKey`.

The DKG modes need a registry deployed with a DKG manager (`dkgAdapter()` non-zero), otherwise
they revert with `DKGDisabled`. The application id is
`keccak256(abi.encode(chainid, registry, processId)) mod Q`, never 0 (`aidFor`). Each
application allows only the adapter to submit ciphertexts, at most 16 of them. The registry
converts the DKG key to circomlib form and stores `keyMode`, `dkgEpochId` and `dkgAid` on the
process.

### Encryption key and genesis root

Whatever its source, the key must be a canonical point (coordinates below the BN254 scalar
field) on the circomlib BabyJubJub curve with `x != 0`, which rules out the identity, the
order-2 point and reduced-form coordinates. Subgroup membership is left to the batch guest.

The registry computes the genesis state root itself. It is the root of a SHA-256 sparse Merkle
tree with six leaves:

| Key | Value |
|---|---|
| `0x00` | process id |
| `0x02` | packed ballot mode |
| `0x03` | `sha256(x ‖ y)` of the encryption key |
| `0x04` | results accumulator: `sha256` of 16 identity ciphertexts |
| `0x06` | census origin |
| `0x07` | `ballotVKHash`, the hash of the ballot proof verification key |

`genesisRoot(processId, ballotMode, encryptionKey, censusOrigin)` returns the same value.
`latestStateRoot` holds the raw SHA-256 digest of the tree root.

## Organizer controls

Only the organizer can call these; anyone else gets `Unauthorized`.

- `setProcessStatus(processId, status)`: see [Statuses](#statuses).
- `setProcessDuration(processId, duration)`: while `READY` or `PAUSED` and before the current end.
  The duration can only grow, and the new end must be in the future. Once the end has passed the
  tally may already be public, so the process cannot be reopened.
- `setProcessMaxVoters(processId, maxVoters)`: while `READY` or `PAUSED`. Non-zero, at least the
  current `votersCount`, and within the same result cap as `newProcess`.
- `setProcessCensus(processId, census)`: origin 2 processes only, while `READY` or `PAUSED` and
  before the end. The new census keeps origin 2, has a non-zero root, a non-empty URI and no
  contract address. Batches proven against the previous root no longer settle.

## State transitions

```solidity
function submitStateTransition(
    bytes31 processId,
    bytes calldata publicValues,   // 512 bytes
    bytes calldata proofBytes,     // abi-encoded uint256[24]
    bytes[] calldata commitments,  // 48-byte KZG commitment per blob
    bytes32[] calldata ys,         // evaluation of each blob at its point, big-endian
    bytes[] calldata kzgProofs     // 48-byte opening proof per blob
) external;
```

It must be sent as a blob transaction carrying exactly the transition's blobs, in order. The
target chain needs EIP-4844: the `BLOBHASH` opcode and the point-evaluation precompile at `0x0a`.

`publicValues` are the guest's 64 `u32` output registers, each written as an 8-byte
little-endian word. A 256-bit value spans 8 registers as little-endian limbs, so a state root
reads back as its raw SHA-256 digest and the census root is compared as the byte-reversed,
big-endian integer. The registry reads these registers of the vote-batch guest (the full layout
is in davinci-zkvm's `circuit/CIRCUIT.md`):

| Registers | Value |
|---|---|
| 0 | `ok` |
| 1 | `fail_mask` |
| 2..9 | state root before |
| 10..17 | state root after |
| 18 | votes in the batch |
| 19 | overwrites among them |
| 20..27 | census root |
| 28..35 | blobs digest |
| 36 | `n_blobs` |
| 42 | `occupied_before` |

The checks, in order:

1. The process exists, is `READY` and within its voting window.
2. `publicValues` is 512 bytes, `ok == 1` and `fail_mask == 0`.
3. The state root before equals `latestStateRoot`.
4. The census root matches: for origins 1, 2 and 4 it equals the stored root; for origin 3 the
   census contract's `getRootBlockNumber(root)` must be non-zero, at most `block.number` and at
   least the process creation block. The census call gets 100k gas and must return a full word.
5. `occupied_before` equals `votersCount`, the number of distinct ballot slots written so far.
   The guest cannot see the tree, so the registry pins it.
6. `votersCount + votes - overwrites <= maxVoters`.
7. `n_blobs > 0`, the three blob arrays have `n_blobs` entries, the transaction carries no blob
   past `n_blobs`, commitments and proofs are 48 bytes, and
   `sha256(commitment_0 ‖ y_0 ‖ commitment_1 ‖ y_1 ‖ ...)` equals the blobs digest.
8. `ZiskVerifier.verifySnarkProof(batchProgramVK, rootCVadcopFinal, publicValues, proofBytes)`
   accepts the PLONK proof. The verifier hashes the program vk, the public values and the setup
   root together, so a proof of another program or setup fails.
9. Every blob `i` is present and opens to `y_i` at
   `z_i = sha256(pid ‖ rootBefore ‖ commitment_i) mod r_BLS`, with the process id and the state
   root before as 32-byte big-endian integers. This is the point at which the guest evaluated
   the blob it laid out itself, so the published blobs are the ones the proof covers.

On success `latestStateRoot` becomes the root after, `votersCount` grows by
`votes - overwrites`, `overwrittenVotesCount` by `overwrites`, `batchNumber` by one, and the
registry emits `ProcessStateTransitioned`.

## Results

### Sequencer mode

```solidity
function setProcessResults(bytes31 processId, bytes calldata publicValues, bytes calldata proofBytes) external;
```

The results guest proves the tally of a state root. The registry requires a sequencer-mode
process that is not `CANCELED` or `RESULTS` and has ended, `ok == 1` and `fail_mask == 0`, a
state root (registers 2..9) equal to `latestStateRoot`, and a PLONK proof verified against
`resultsProgramVK`. Result `i` is read from registers `10 + 2i` (low 32 bits) and `11 + 2i`
(high 32 bits) for each of the `numFields` fields. The process moves to `RESULTS` and the call
emits `ProcessStatusChanged` and `ProcessResultsSet`.

### DKG modes

DKG processes do not use `setProcessResults` (it reverts with `InvalidKeyMode`). Both steps are
permissionless.

```solidity
function requestResultsDecryption(bytes31 processId, uint256[64] calldata accumulator, bytes32[] calldata siblings) external;
function finalizeResultsFromDKG(bytes31 processId) external;
function revealProcessKey(bytes31 processId, uint256 sk) external; // DKG_LOCKED only
```

`requestResultsDecryption` runs once per process, after it has ended. `accumulator` holds the
16 ElGamal ciphertexts of the results leaf as `(C1x, C1y, C2x, C2y)` each, circomlib form,
big-endian. `siblings` is the SMT inclusion proof of leaf `0x04`, root to leaf, zero-padded,
at most 64 entries and ending in a zero entry. The registry:

1. requires every coordinate to be below the BN254 scalar field, so `(0, 1)` is the only
   identity encoding;
2. checks that `sha256` of the 64 words is the value of leaf `0x04` under `latestStateRoot`;
3. moves the process to `ENDED` if it is not there already, since the plaintexts become public
   on the DKG before finalization and the organizer must not be able to cancel after reading
   them;
4. skips fields that are identity in both halves (they count as 0), rejects a field with only
   one identity half, and submits the rest to the DKG, converted to reduced form, at
   contiguous ciphertext indices;
5. emits `ResultsDecryptionRequested(processId, epochId, aid, firstIndex, count)`.

If every declared field is identity nothing is submitted and the results are finalized to
zero at once.

`finalizeResultsFromDKG` reads the combined plaintexts once the committee has decrypted every
submitted ciphertext, stores `numFields` results in field order (0 for skipped fields), moves
the process to `RESULTS` and emits `ProcessStatusChanged` and `ProcessResultsSet`. Until then it
reverts with `ResultsNotReady`.

`revealProcessKey` forwards the organizer secret of a `DKG_LOCKED` process to
`DKGAppManager.revealOrganizerSecret`, which checks `sk·G == PK_org` and accepts it once. The
registry does not restrict when it is called.

DKG liveness is election liveness. Once the ciphertexts are submitted there is no fallback: if
more than `n - t` members of the registration epoch's committee are gone, the combines never
complete and the process never reaches `RESULTS`.

## Reading state

- `getProcess(processId)` returns the full `DAVINCITypes.Process`: status, organizer, key,
  `latestStateRoot`, `result`, times, `maxVoters`, `votersCount` (distinct ballot slots
  written), `overwrittenVotesCount`, `creationBlock`, `batchNumber`, metadata, ballot mode,
  census, and the DKG fields `keyMode`, `dkgEpochId`, `dkgAid`, `dkgFirstIndex`, `dkgCount`,
  `dkgZeroSkipped` (bit `i` set when field `i` was skipped as identity) and
  `dkgResultsRequested`.
- `getNextProcessId(organizer)`, `getProcessEndTime(processId)`, `genesisRoot(...)`,
  `aidFor(processId)`.
- The immutables `ziskVerifier`, `batchProgramVK`, `resultsProgramVK`, `rootCVadcopFinal`,
  `ballotVKHash` and `dkgAdapter`; `getSTVerifierVKeyHash()` and `getRVerifierVKeyHash()` return
  the batch and results vks.
- `chainID`, `pidPrefix`, `processCount`, `processNonce(organizer)`.
- On the adapter: `registrationEpoch()`, `aidFor(processId)`, `registry`, `manager`,
  `appManager`.

## Events

| Event | Emitted by |
|---|---|
| `ProcessCreated(bytes31 indexed processId, address indexed creator)` | `newProcess` |
| `ProcessStatusChanged(bytes31 indexed processId, ProcessStatus oldStatus, ProcessStatus newStatus)` | `setProcessStatus`, `setProcessResults`, `requestResultsDecryption` (to `ENDED`, and to `RESULTS` when every field is identity), `finalizeResultsFromDKG` |
| `ProcessDurationChanged(bytes31 indexed processId, uint256 duration)` | `setProcessDuration`, `setProcessStatus` to `ENDED` |
| `ProcessMaxVotersChanged(bytes31 indexed processId, uint256 maxVoters)` | `setProcessMaxVoters` |
| `CensusUpdated(bytes31 indexed processId, bytes32 censusRoot, string censusURI)` | `setProcessCensus` |
| `ProcessStateTransitioned(bytes31 indexed processId, address indexed sender, bytes32 oldStateRoot, bytes32 newStateRoot, uint256 newVotersCount, uint256 newOverwrittenVotesCount, uint256 nBlobs)` | `submitStateTransition` |
| `ProcessResultsSet(bytes31 indexed processId, address indexed sender, uint256[] result)` | `setProcessResults`, `finalizeResultsFromDKG`, and `requestResultsDecryption` when every field is identity |
| `ResultsDecryptionRequested(bytes31 indexed processId, bytes12 epochId, bytes32 aid, uint16 firstIndex, uint8 count)` | `requestResultsDecryption` |

## Errors

Errors of `IProcessRegistry` unless noted.

| Error | Raised by | Condition |
|---|---|---|
| `InvalidProcessId` | calls taking a process id | zero id |
| `UnknownProcessIdPrefix` | calls taking a process id | id from another registry or chain |
| `ProcessNotFound` | calls taking a process id | no such process |
| `Unauthorized` | organizer controls | caller is not the organizer |
| `ProcessAlreadyExists` | `newProcess` | the id is taken |
| `InvalidStatus` | several | initial status not `READY`/`PAUSED`; transition not allowed; process not `READY` (settlement) or not `READY`/`PAUSED` (organizer controls); process `CANCELED` or already `RESULTS` (results calls) |
| `InvalidStartTime` | `newProcess` | start time in the past |
| `InvalidDuration` | `newProcess`, `setProcessDuration` | end not in the future; new duration zero or not longer |
| `InvalidTimeBounds` | several | past the end (`setProcessDuration`, `setProcessCensus`); outside the voting window (settlement); not ended yet (results calls) |
| `InvalidMaxVoters` | `newProcess`, `setProcessMaxVoters` | zero, or below `votersCount` |
| `MaxPossibleResultCapExceeded` | `newProcess`, `setProcessMaxVoters` | `maxValue > 10^12 / maxVoters` |
| `InvalidMaxCount` | `newProcess` | `numFields` is 0 or above 16 |
| `InvalidGroupSize` | `newProcess` | `groupSize > numFields` |
| `InvalidMaxMinValueBounds` | `newProcess` | `minValue > maxValue` |
| `InvalidValueSumBounds` | `newProcess` | `minValueSum > maxValueSum` |
| `BallotModeMaxValueTooLarge`, `BallotModeMinValueTooLarge` | `newProcess` | value above 2^48 - 1 |
| `BallotModeMaxValueSumTooLarge`, `BallotModeMinValueSumTooLarge` | `newProcess` | sum above 2^63 - 1 |
| `InvalidCensusOrigin` | `newProcess`, `setProcessCensus` | `CENSUS_UNKNOWN`; a different origin on update |
| `InvalidCensusConfig` | `newProcess` | `onchainAllowAnyValidRoot` set |
| `InvalidCensusAddress` | `newProcess`, `setProcessCensus` | origin 3 contract without code or not answering `getCensusRoot()`; a non-zero address for another origin |
| `InvalidCensusRoot` | `newProcess`, `setProcessCensus`, `submitStateTransition` | zero root; CSP root wider than 160 bits; batch census root not accepted |
| `InvalidCensusURI` | `newProcess`, `setProcessCensus` | empty URI |
| `CensusNotUpdatable` | `setProcessCensus` | process census is not origin 2 |
| `InvalidEncryptionKey` | `newProcess` | key not canonical, not on the curve or `x == 0`; a non-zero key in a DKG mode |
| `InvalidDKGParams` | `newProcess` | non-zero DKG fields in sequencer mode |
| `DKGDisabled` | `newProcess`, `aidFor` | DKG mode on a registry without a DKG manager |
| `InvalidPublicValues` | settlement, `setProcessResults` | `publicValues` not 512 bytes |
| `CircuitFailed` | settlement, `setProcessResults` | `ok != 1` or `fail_mask != 0` |
| `InvalidStateRoot` | settlement, `setProcessResults` | root does not match `latestStateRoot` |
| `InvalidOccupiedBefore` | `submitStateTransition` | `occupied_before != votersCount` |
| `MaxVotersReached` | `submitStateTransition` | the batch would exceed `maxVoters` |
| `NoBlobs` | `submitStateTransition` | `n_blobs == 0` |
| `BlobCountMismatch` | `submitStateTransition` | array lengths differ from `n_blobs`, or the transaction carries an extra blob |
| `InvalidBlobCommitmentLength`, `InvalidKZGProofLength` | `submitStateTransition` | not 48 bytes |
| `InvalidBlobsDigest` | `submitStateTransition` | commitment/evaluation digest differs from the proof's |
| `MissingBlob(index)` | `submitStateTransition` | the transaction has no blob at `index` |
| `InvalidBlobOpening(index)` | `submitStateTransition` | the point-evaluation precompile rejects blob `index` |
| `InvalidProof` (`ZiskVerifier`) | settlement, `setProcessResults` | PLONK proof does not verify |
| `InvalidKeyMode` | results calls, `revealProcessKey` | call not valid for the process's key mode |
| `ResultsAlreadyRequested` | `requestResultsDecryption` | second request |
| `InvalidAccumulator` | `requestResultsDecryption` | a coordinate not below the field, or a field with only one identity half |
| `InvalidInclusionProof` | `requestResultsDecryption` | accumulator not included under `latestStateRoot` |
| `ResultsNotReady` | `finalizeResultsFromDKG` | not requested yet, or a combine is still missing |
| `NoLiveEpoch` (adapter) | `newProcess` | `DKG_AUTOMATIC` found no Live epoch with a free pool key |
| `NonContiguousIndex` (adapter) | `requestResultsDecryption` | the DKG assigned non-consecutive ciphertext indices |
| `InvalidVerifierConfig` | constructor | zero verifier address or pin |

Errors from davinci-dkg pass through unchanged. On `newProcess` these include `InvalidEpoch`
or `InvalidPhase` for an unknown or not yet Live `DKG_LOCKED` epoch, `InvalidSchnorrProof`,
and `NotRegistrar` when the DKG's registrar is not this adapter; on `revealProcessKey`,
`InvalidOrganizerSecret` and `AlreadyRevealed`. `IProcessRegistry` also
declares `InvalidBlockNumber`, `InvalidMaxValue`, `InvalidMinValue`, `InvalidMinTotalCost`,
`InvalidUniqueValues`, `CannotAcceptResult`, `ProcessNotEnded` and `ProofInvalid`, which this
registry never raises.

## Deploying

`script/DeployAll.s.sol` deploys `ZiskVerifier` and then `ProcessRegistry`:

```solidity
constructor(
    uint32 chainID,
    address ziskVerifier,
    bytes32 batchProgramVK,
    bytes32 resultsProgramVK,
    bytes32 rootCVadcopFinal,
    bytes32 ballotVKHash,
    address dkgManager
)
```

It reads:

| Variable | Value |
|---|---|
| `PRIVATE_KEY` | deployer key |
| `CHAIN_ID` | chain id, stored in the registry and folded into process ids; must fit in `uint32` |
| `BATCH_PROGRAM_VK` | program vk of the vote-batch guest (`bytes32`) |
| `RESULTS_PROGRAM_VK` | program vk of the results guest |
| `ROOT_C_VADCOP_FINAL` | root of the ZisK vadcop-final setup; the script aborts unless it equals `ZiskVerifier.getRootCVadcopFinal()` |
| `BALLOT_VK_HASH` | `sha256` of the ballot proof VK wire bytes, genesis leaf `0x07` (`davinci.BallotVKLeaf` in the davinci-zkvm Go SDK) |
| `DKG_MANAGER` | optional davinci-dkg `DKGManager`; unset or zero disables the DKG modes |

The program vks are what `cargo-zisk setup` prints as `Root hash` for each guest ELF;
davinci-zkvm pins the released ones in `rust-sdk/src/release.rs`. A zero verifier address or
pin reverts with `InvalidVerifierConfig`.

Local node, with the variables above exported (`CHAIN_ID=31337`):

```bash
anvil
forge script script/DeployAll.s.sol --rpc-url http://localhost:8545 --broadcast
```

Single chain, with the variables above plus `RPC_URL` in `.env` (see `.env.example`):

```bash
./deploy_all.sh
```

`deploy_all.sh` runs the script with `--broadcast --slow` and the optimizer settings of
`foundry.toml`, verifies the sources according to `VERIFY_MODE` (`auto` skips chain ids 31337
and 1337, `true`, `false`; `ETHERSCAN_API_KEY` and optionally `ETHERSCAN_API_URL`), then
regenerates `golang-types/addresses.go` from the broadcast logs.

Several chains: put shared values and `DEPLOY_CHAINS=base,sepolia,...` in `.env`, and
`CHAIN_ID`, `RPC_URL` and anything chain-specific in `.env.<chain>` (or `.env-<chain>`), then:

```bash
./deploy_all_contracts_to_all_chains.sh
```

It reloads `.env`, clears the chain-scoped variables, loads each chain's file and calls
`deploy_all.sh`. `DKG_MANAGER` is not among the variables it clears, so when it is set for one
chain, set it (to zero if needed) in every chain file.

### With DKG support

1. Deploy the davinci-dkg contracts.
2. Deploy with `DKG_MANAGER` set. The registry constructor creates the `DavinciDKGAdapter`,
   which reads `appManager()` from the manager; the script logs its address and
   `dkgAdapter()` returns it.
3. From the `DKGAppManager` admin, call `setRegistrar(adapter)`. Until then anyone can register
   applications on that DKG, which lets them take an application id ahead of a process or drain
   the pool keys, so do it before the first epoch goes Live.

## Checking a deployment

`script/verify_deployment.py` compares a live deployment with this checkout. It needs Python 3,
`cast`, and a prior `forge build` with the same compiler settings, since it reads `out/`.

```bash
forge build
python3 script/verify_deployment.py --rpc "$GNOSIS_RPC_URL" --chain-id 100 \
    --registry 0x3CDE68c39E26ecf94bD029b6ED3b9F945441daf3 \
    --batch-vk 0x6cfc89d562d0b22f04478a5c15b390433eb52f1b03147030b183076260da7a10 \
    --results-vk 0x7bc8c5e9235548386a44b1885732a2a7ffb1badddc8c7fba599d07ece47be794 \
    --root-c 0x05006517b6ccde5da4d890587ba62845b5af8a307c00e87d4b9d05099b16dc80 \
    --ballot-vk-hash 0xbf1e6590bb1ba883d601c4d7d1c6fa2722a78590716874019db6d68fc776bb0e
```

It checks that:

- the runtime code of `ProcessRegistry` and `ZiskVerifier` matches the local build, with
  immutables masked;
- the registry's `batchProgramVK`, `resultsProgramVK`, `rootCVadcopFinal` and `ballotVKHash`
  equal the given pins, and so does the verifier's `getRootCVadcopFinal()`;
- the registry's `chainID` equals the RPC's chain id (and `--chain-id`, when given);
- when `dkgAdapter()` is set, the adapter's code matches the local build, `adapter.registry()`
  is the registry and `DKGAppManager.registrar()` is the adapter. An unset registrar fails.

Each check prints `OK` or `FAIL`; the exit status is 1 if any fails. Use `--cast` when `cast`
is not at `~/.foundry/bin/cast`.

## Development

Prerequisites: [Foundry](https://getfoundry.sh/); Node.js, [abigen](https://geth.ethereum.org/docs/tools/abigen)
and [jq](https://jqlang.org/) for the bindings; Go (version in `test/vectors/go.mod`) to
regenerate the test vectors.

```bash
git clone --recurse-submodules https://github.com/vocdoni/davinci-contracts.git
cd davinci-contracts
npm install   # only needed for the bindings and linters
```

### Build and test

```bash
forge build
forge build --sizes                          # contract sizes, as CI runs it
forge test
forge test -vvv --match-path test/DKG.t.sol
forge test --gas-report
```

| Test | Covers |
|---|---|
| `ProcessRegistry.t.sol` | status machine, duration, max voters, ballot mode and census validation |
| `NewProcess.t.sol` | census origins and encryption key checks at creation |
| `Genesis.t.sol` | genesis roots against the Go reference |
| `Transition.t.sol` | settlement with real KZG openings (PLONK verifier mocked) |
| `DynamicCensus.t.sol` | census origins 2 and 3 |
| `Results.t.sol` | `setProcessResults` |
| `DKG.t.sol` | DKG key modes, against a mock DKG with real BabyJubJub arithmetic |
| `Publics.t.sol`, `ZiskVerifier.t.sol` | public values decoding and the vendored verifier, against a recorded batch PLONK |
| `BlobsLib.t.sol`, `ProcessIdLib.t.sol` | library helpers |

The fixtures in `test/vectors/` come from the Go program in the same directory, which builds
them with the davinci-zkvm Go SDK. The tests deploy the registry at the address the fixtures
were built for (anvil account 0, nonce 0, chain id 1337). To regenerate them, check out
davinci-zkvm and davinci-circom next to this repository and run:

```bash
cd test/vectors && go run .
```

CI does the same against davinci-zkvm `main` and a pinned davinci-circom commit, and fails if
the output differs from the committed files.

### Bindings

```bash
./build_all.sh
```

This cleans, runs `forge build`, compiles with Hardhat to produce `artifacts/` and the
TypeChain types in `typechain-types/`, and runs `./go_bind.sh`, which writes the Go bindings
`golang-types/ProcessRegistry.go` and `golang-types/ICensusValidator.go` (package `contracts`)
plus `golang-types/addresses.go` from the broadcast logs.

The npm package `@vocdoni/davinci-contracts` ships the compiled TypeChain types (`npm run
prepare` builds `dist/`). In CI, pushes that touch `src/**/*.sol` regenerate and commit the
bindings, pushes to `release` bump the patch version and tag it, and `v*` tags publish to npm.

### Linters

```bash
npm run lint:sol
npm run prettier
npm run slither
npm run mythril
```

### Docker

`docker-compose.yml` runs Foundry (`FOUNDRY_VERSION`, default v1.8.3) in an image built
from this checkout, submodules included, so check them out first.

```bash
docker compose --profile test run --rm test   # forge build --sizes, forge test
docker compose --profile local up -d          # anvil with the contracts deployed
docker compose --profile local down           # the chain is gone after this
```

`local` starts anvil on port `ANVIL_PORT` (default 8545): chain id 31337, Osaka, blocks every
`ANVIL_BLOCK_TIME` seconds (default 1). It deploys from anvil's account 0, pinned to the
davinci-zkvm release and without DKG, so the addresses are always `ZiskVerifier`
`0x5FbDB2315678afecb367f032d93F642f64180aa3` and `ProcessRegistry`
`0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512`. A davinci-sequencer runs against it with
`--blob-source anvil`.

`deploy` runs `deploy_all.sh` with the variables of [Deploying](#deploying) from `.env`; the key
never enters the image. The broadcast record stays in the container, so copy it out before
removing it:

```bash
docker compose --profile deploy up deploy
docker compose --profile deploy cp deploy:/app/broadcast/DeployAll.s.sol/. broadcast/DeployAll.s.sol/
docker compose --profile deploy rm -f deploy
```

The `golang-types/addresses.go` it regenerates stays in the container too; run
`helpers/write_contract_addresses.sh` on the host after copying the record.

## License

GNU Affero General Public License v3.0, see [LICENSE.md](LICENSE.md). Source files carry their
own SPDX identifiers.
