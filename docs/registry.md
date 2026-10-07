# ProcessRegistry

How a voting process lives in the registry: its id, lifecycle and grace window, the
parameters it is created with, what the organizer can change and what can be read back.
Settlement of transitions and results is in [settlement.md](settlement.md), every revert in
[errors.md](errors.md).

## Process ids

A process id is a `bytes31`: the organizer address (20 bytes), a 4-byte registry prefix and a
7-byte per-organizer nonce. The prefix is the low 4 bytes of
`keccak256(abi.encodePacked(uint32 chainID, registry))`, so ids from another registry or chain
revert with `UnknownProcessIdPrefix`. `getNextProcessId(organizer)` returns the id the
organizer's next `newProcess` will get. The organizer is the `newProcess` caller, stored as
`organizationId`.

## Lifecycle

`ProcessStatus` is `READY` (0), `ENDED` (1), `CANCELED` (2), `PAUSED` (3), `RESULTS` (4). The
organizer moves a process with `setProcessStatus`:

| From | Organizer may set | Other ways out |
|---|---|---|
| `READY` | `PAUSED` (before the end time), `CANCELED`, `ENDED` (from `startTime` on) | `RESULTS` or `ENDED` from the results calls, once the grace window has closed |
| `PAUSED` | `READY`, `CANCELED`, `ENDED` (from `startTime` on) | same as `READY` |
| `ENDED` | none | `RESULTS` from the results calls |
| `CANCELED` | none | none |
| `RESULTS` | none | none |

A process ends at `startTime + duration` (`getProcessEndTime`). Setting `ENDED` by hand before
that sets `duration` to the time elapsed since `startTime` and emits `ProcessDurationChanged`;
past the end it leaves `duration` alone. `ENDED` is refused before `startTime` and `PAUSED` from
the end time on, both with `InvalidTimeBounds`: a process that never opened is voided with
`CANCELED`, and a pause cannot outlast voting.

Transitions settle while `startTime <= block.timestamp < getProcessGraceEnd(processId)` and the
process is `READY`, `ENDED`, or `PAUSED` past its end time. A pause during voting stops
settlement but not the clock.

Results are accepted once the process is `ENDED`, or `READY`/`PAUSED` with its end time passed,
and its grace window has closed. In sequencer mode `setProcessResults` moves it straight to
`RESULTS`. In the DKG modes `requestResultsDecryption` first moves it to `ENDED`, which takes it
out of the organizer's hands, and `finalizeResultsFromDKG` moves it to `RESULTS`; for
`COUNCIL` not before the ceremony opens decryption ([settlement](settlement.md#the-council-decryption-gate)).

## Grace window

Votes cast just before the end can still be queued or proving when it passes. The grace window
lets them settle: transitions keep landing after the end until

```
graceEnd = min(end + graceMaxTotal, max(end, lastVoteAt) + grace)
```

with `end = startTime + duration`, `lastVoteAt` the `block.timestamp` of the last settled
transition (0 before the first) and `grace` the process's idle window in seconds.
`getProcessGraceEnd(processId)` returns it. Every settled transition moves `lastVoteAt`, so the
window stays open while batches keep landing at most `grace` apart, and never past
`end + graceMaxTotal`. Once `block.timestamp` reaches `graceEnd` nothing settles, `lastVoteAt`
stops moving and the window never reopens: past the end the duration, `grace` and `maxVoters`
are fixed, and `ENDED` no longer moves the end.

A process ended by hand gets the same window, counted from the moment it was ended. A process
paused during voting stays blocked until the end time and then settles through the window like
`READY`, so a pause cannot hold the window shut until it expires. A transition must carry at
least one vote (`EmptyTransition` otherwise), so a batch that only re-encrypts cannot extend
the window; a batch of overwrites alone is a vote and does.

`setProcessResults`, `requestResultsDecryption` and `finalizeResultsFromDKG` revert with
`GraceOpen` while `block.timestamp < graceEnd`, so the tally is always taken over the final
root. Nothing but the window decides when they unlock: there is no organizer or node signal, so
any sequencer, known to the organizer or not, gets `grace` seconds after the latest landed batch
to land its own.

`grace` starts at the registry's `defaultGrace`. The organizer can set it within
`[graceFloor, graceCeil]` with `setProcessGrace` while the process is `READY` or `PAUSED` and
before its end.

Trust: the registry cannot tell when a vote was cast, so a sequencer can settle votes cast after
the end while the window is open. That is bounded, since the window closes at most
`graceMaxTotal` after the end and each extension costs a batch carrying a vote. It is also
visible: every transition's block timestamp is public, so anyone can count the batches a
process settled after its end.

For census origin 3 the window also accepts census roots that appeared after the end: the
registry checks that the census contract held the root at some block since the process was
created, and does not know the block the election ended at. A census contract used with this
registry should therefore stop changing at the election end. Honest nodes refuse votes after
the end anyway, so this only matters together with a sequencer that admits late votes.

### Shortening with notice

`setProcessDuration` can also move the end earlier, provided the new end is at least `noticeMin`
seconds after the call (and after `startTime`). That makes an announcement such as "voting
closes in one minute" a single transaction: nodes see `ProcessDurationChanged` before the new
end arrives, so they can settle their queues while voting is still open and stop admitting
votes at the announced moment. No vote admitted under the old end turns late, and the grace
window follows the new end. `setProcessStatus(ENDED)` remains the way to close at once.

## Creating a process

```solidity
function newProcess(
    DAVINCITypes.ProcessStatus status,
    uint256 startTime,
    uint256 duration,
    uint256 maxVoters,
    DAVINCITypes.BallotMode calldata ballotMode,
    DAVINCITypes.Census calldata census,
    string calldata metadataURI,
    bytes32 metadataHash,
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
| `ballotMode` | See [Ballot mode](#ballot-mode). |
| `census` | See [Census origins](#census-origins). |
| `metadataURI` | Where the metadata document is served. Non-empty. |
| `metadataHash` | SHA-256 of that document; see [Metadata](#metadata). Non-zero. |
| `encryptionKey` | The ElGamal public key in sequencer mode; `(0, 0)` in the DKG modes. |
| `dkg` | Key mode and DKG registration arguments; all zero in sequencer mode. See [Key modes](#key-modes). |

The call emits `ProcessCreated(processId, organizer)`, then
`ProcessMetadataUpdated(processId, metadataURI, metadataHash)`. The process starts with
`grace = defaultGrace`.

### Metadata

The metadata document carries what the registry does not: the title, the question and which
option each ballot field stands for. `metadataHash` binds the process to that content, as the
census root binds the census, so a host that serves other bytes no longer matches. The hash is
SHA-256 over the exact bytes served at `metadataURI`, with no JSON canonicalisation. A client
fetches the document, hashes the raw body and compares the digest with `metadataHash` from
`getProcess`; reformatting the document changes its hash.

The organizer can replace both with `setProcessMetadata` until voting closes. `newProcess` and
every update emit `ProcessMetadataUpdated`, so the event log alone is the complete metadata
history of a process.

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
[`ICensusValidator.sol`](../src/interfaces/ICensusValidator.sol) explains why.
[davinci-onchain-census-contract](https://github.com/vocdoni/davinci-onchain-census-contract)
provides such a census.

### Key modes

```solidity
enum KeyMode { SEQUENCER, DKG_AUTOMATIC, DKG_LOCKED, COUNCIL }

struct DKGParams {
    KeyMode mode;
    bytes12 epochId; // DKG_LOCKED: the epoch the proof of possession binds; COUNCIL: the ceremony id
    uint256 orgPKx;  // DKG_LOCKED only: organizer key, reduced form
    uint256 orgPKy;
    uint256 popAx;   // DKG_LOCKED only: Schnorr proof of possession
    uint256 popAy;
    uint256 popZ;
}
```

- `SEQUENCER`: the caller supplies `encryptionKey` and every `DKGParams` field is zero. The
  sequencer that generated the key decrypts the tally and proves it with the results guest.
- `DKG_AUTOMATIC`: `encryptionKey` is `(0, 0)` and `epochId` is ignored: the adapter registers a
  davinci-dkg application on `registrationEpoch()`, the newest Live epoch with a free pool key
  (it scans back at most 8 epochs), and the process key is that pool key. The organizer key and
  proof fields are forwarded, and the DKG ignores them in this mode. Only the committee
  decrypts the tally.
- `DKG_LOCKED`: as above, but on the caller's `epochId` and with an organizer key: the process
  key is the pool key plus `PK_org`. The organizer key and proof of possession are in the DKG's
  reduced (a = -1) form, exactly as `DKGAppManager.registerApplication` takes them. The proof
  binds the epoch and the application id, so a client reads the epoch from the adapter's
  `registrationEpoch()` and the id from `aidFor(getNextProcessId(organizer))` before building
  it. The committee cannot finish decrypting until the organizer secret is published with
  `revealProcessKey`.

- `COUNCIL`: `encryptionKey` is `(0, 0)`, `epochId` names a Council ceremony and every other
  `DKGParams` field is zero (`InvalidDKGParams` otherwise, as for a zero ceremony id). The
  `CouncilAdapter` binds the process to the ceremony with `bindProcess(ceremonyId, processId,
  creator)`, where `creator` is the `newProcess` caller; the manager requires the ceremony to be
  Live, the adapter allowed and the creator authorized by the ceremony's organizer, and its
  errors (`NotAllowedAdapter`, `NotAuthorizedCreator`, `WrongPhase`, `UnknownCeremony`) pass
  through. The process key is the ceremony key, already in circomlib form; every process bound to
  a ceremony shares it. The registry stores the ceremony id as `dkgEpochId` and the Council
  request id as `dkgAid`; the adapter maps the request id back to the process id
  (`bindings(requestId)`). Only the ceremony's committee decrypts the tally, and no result,
  not even an all-zero one, is published before the ceremony opens decryption.

The DKG modes need a registry deployed with a DKG manager (`dkgAdapter()` non-zero), otherwise
they revert with `DKGDisabled`; `COUNCIL` needs a Council manager (`councilAdapter()`
non-zero), otherwise it reverts with `CouncilDisabled`. The application id (`aidFor`) is
`salt << 160 | adapter`, where `salt` is the top 92 bits of
`keccak256(abi.encode(chainid, registry, processId))` and `adapter` the `DavinciDKGAdapter`
address. davinci-dkg only lets an account register ids whose low 160 bits are its own address,
so no one but the adapter can register a process's id, even though anyone can compute it in
advance (vocdoni/davinci-dkg#14); the id is below `2^252`, so a valid field element. Each
application allows only the adapter to submit ciphertexts, at most 16 of them. The registry
converts the DKG key to circomlib form and stores `keyMode`, `dkgEpochId` and `dkgAid` on the
process.

### Encryption key and genesis root

Whatever its source, the key must be a canonical point (coordinates below the BN254 scalar
field) on the circomlib BabyJubJub curve with `x != 0`, which rules out the identity, the
order-2 point and reduced-form coordinates. Subgroup membership is left to the batch guest.

The registry computes the genesis state root itself. It is the root of a SHA-256 sparse Merkle
tree (vocdoni/arbo layout, 64 levels, 8-byte keys) with six leaves:

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

Only the organizer can call these; anyone else gets `Unauthorized`. All of them except
`setProcessStatus` work only while the process is `READY` or `PAUSED` and before its end, so the
parameters of an election are fixed once voting closes.

| Function | Rule |
|---|---|
| `setProcessStatus(processId, status)` | See [Lifecycle](#lifecycle). |
| `setProcessDuration(processId, duration)` | Non-zero. The new end can be any later time, or an earlier one at least `noticeMin` seconds away (see [Shortening with notice](#shortening-with-notice)). |
| `setProcessGrace(processId, grace)` | `graceFloor <= grace <= graceCeil`. Emits `ProcessGraceChanged`. |
| `setProcessMaxVoters(processId, maxVoters)` | Non-zero, at least the current `votersCount`, and within the same result cap as `newProcess`. |
| `setProcessCensus(processId, census)` | Origin 2 only. The new census keeps origin 2, has a non-zero root, a non-empty URI and no contract address. Batches proven against the previous root no longer settle. |
| `setProcessMetadata(processId, metadataURI, metadataHash)` | Non-empty URI and non-zero hash (see [Metadata](#metadata)). |

## Reading state

- `getProcess(processId)` returns the full `DAVINCITypes.Process`: status, organizer, key,
  `latestStateRoot`, `result`, times, `maxVoters`, `votersCount` (distinct ballot slots
  written), `overwrittenVotesCount`, `creationBlock`, `batchNumber`, `metadataURI`,
  `metadataHash`, ballot mode, census, and the DKG fields `keyMode`, `dkgEpochId`,
  `dkgFirstIndex`, `dkgCount`, `dkgZeroSkipped` (bit `i` set when field `i` was skipped as
  identity), `dkgResultsRequested` and `dkgAid`, then `grace` and `lastVoteAt`, in that order.
- `getNextProcessId(organizer)`, `getProcessEndTime(processId)`, `getProcessGraceEnd(processId)`,
  `genesisRoot(...)`, `aidFor(processId)`.
- The immutables `ziskVerifier`, `batchProgramVK`, `resultsProgramVK`, `rootCVadcopFinal`,
  `ballotVKHash`, `dkgAdapter` and `councilAdapter`; `getSTVerifierVKeyHash()` and `getRVerifierVKeyHash()` return
  the batch and results vks. The window settings `defaultGrace`, `graceFloor`, `graceCeil`,
  `graceMaxTotal` and `noticeMin` (seconds, `uint32`) are immutables too, for nodes to read at
  boot.
- `chainID`, `pidPrefix`, `processCount`, `processNonce(organizer)`.
- On the adapter: `registrationEpoch()`, `aidFor(processId)`, `registry`, `manager`,
  `appManager`. On the Council adapter: `bindings(requestId)` (ceremony id, submitted field
  count, process id), `isDecryptionOpen(ceremonyId)` (the manager's decryption gate),
  `registry`, `manager`.

## Events

| Event | Emitted by |
|---|---|
| `ProcessCreated(bytes31 indexed processId, address indexed creator)` | `newProcess` |
| `ProcessStatusChanged(bytes31 indexed processId, ProcessStatus oldStatus, ProcessStatus newStatus)` | `setProcessStatus`, `setProcessResults`, `requestResultsDecryption` (to `ENDED`, and to `RESULTS` when every field is identity, for `COUNCIL` only once decryption is open), `finalizeResultsFromDKG` |
| `ProcessDurationChanged(bytes31 indexed processId, uint256 duration)` | `setProcessDuration`, `setProcessStatus` to `ENDED` before the end time |
| `ProcessMaxVotersChanged(bytes31 indexed processId, uint256 maxVoters)` | `setProcessMaxVoters` |
| `ProcessGraceChanged(bytes31 indexed processId, uint32 grace)` | `setProcessGrace` |
| `CensusUpdated(bytes31 indexed processId, bytes32 censusRoot, string censusURI)` | `setProcessCensus` |
| `ProcessMetadataUpdated(bytes31 indexed processId, string metadataURI, bytes32 metadataHash)` | `newProcess` (the initial values), `setProcessMetadata` |
| `ProcessStateTransitioned(bytes31 indexed processId, address indexed sender, bytes32 oldStateRoot, bytes32 newStateRoot, uint256 newVotersCount, uint256 newOverwrittenVotesCount, uint256 nBlobs)` | `submitStateTransition` |
| `ProcessResultsSet(bytes31 indexed processId, address indexed sender, uint256[] result)` | `setProcessResults`, `finalizeResultsFromDKG`, and `requestResultsDecryption` when every field is identity (for `COUNCIL` only once decryption is open) |
| `ResultsDecryptionRequested(bytes31 indexed processId, bytes12 epochId, bytes32 aid, uint16 firstIndex, uint8 count)` | `requestResultsDecryption` |
