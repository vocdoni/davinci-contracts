# Errors

Custom errors of `IProcessRegistry` unless noted. "Settlement" is `submitStateTransition`;
"results calls" are `setProcessResults`, `requestResultsDecryption` and
`finalizeResultsFromDKG`.

| Error | Raised by | Condition |
|---|---|---|
| `InvalidProcessId` | calls taking a process id | zero id |
| `UnknownProcessIdPrefix` | calls taking a process id | id from another registry or chain |
| `ProcessNotFound` | calls taking a process id | no such process |
| `Unauthorized` | organizer controls | caller is not the organizer |
| `ProcessAlreadyExists` | `newProcess` | the id is taken |
| `InvalidStatus` | several | initial status not `READY`/`PAUSED`; transition not allowed; process not `READY`/`ENDED`, or `PAUSED` before its end (settlement) or not `READY`/`PAUSED` (organizer controls); process `CANCELED` or already `RESULTS` (results calls) |
| `InvalidStartTime` | `newProcess` | start time in the past |
| `InvalidDuration` | `newProcess`, `setProcessDuration` | end not in the future; new duration zero or unchanged; an earlier end less than `noticeMin` away |
| `InvalidTimeBounds` | several | past the end (`setProcessDuration`, `setProcessMaxVoters`, `setProcessCensus`, `setProcessMetadata`, `setProcessGrace`); `ENDED` before `startTime` or `PAUSED` from the end on (`setProcessStatus`); before `startTime` or at or past the grace end (settlement); not ended yet (`setProcessResults`, `requestResultsDecryption`) |
| `InvalidGrace` | constructor, `setProcessGrace` | grace outside `[graceFloor, graceCeil]`; constructor bounds not `0 < graceFloor <= defaultGrace <= graceCeil <= graceMaxTotal` with `noticeMin > 0` |
| `GraceOpen` | results calls | the grace window has not closed |
| `EmptyTransition` | settlement | the batch carries no vote |
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
| `InvalidCensusRoot` | `newProcess`, `setProcessCensus`, settlement | zero root; CSP root wider than 160 bits; batch census root not accepted |
| `InvalidCensusURI` | `newProcess`, `setProcessCensus` | empty URI |
| `InvalidMetadata` | `newProcess`, `setProcessMetadata` | empty metadata URI or zero hash |
| `CensusNotUpdatable` | `setProcessCensus` | process census is not origin 2 |
| `InvalidEncryptionKey` | `newProcess` | key not canonical, not on the curve or `x == 0`; a non-zero key in a DKG or Council mode |
| `InvalidDKGParams` | `newProcess` | non-zero DKG fields in sequencer mode; in `COUNCIL`, a zero ceremony id or a non-zero organizer/PoP field (raised by the Council adapter, same selector) |
| `DKGDisabled` | `newProcess`, `aidFor` | DKG mode on a registry without a DKG manager |
| `CouncilDisabled` | `newProcess` | `COUNCIL` on a registry without a Council manager |
| `InvalidPublicValues` | settlement, `setProcessResults` | `publicValues` not 512 bytes |
| `CircuitFailed` | settlement, `setProcessResults` | `ok != 1` or `fail_mask != 0` |
| `InvalidStateRoot` | settlement, `setProcessResults` | root does not match `latestStateRoot` |
| `InvalidOccupiedBefore` | settlement | `occupied_before != votersCount` |
| `MaxVotersReached` | settlement | the batch would exceed `maxVoters` |
| `NoBlobs` | settlement | `n_blobs == 0` |
| `BlobCountMismatch` | settlement | array lengths differ from `n_blobs`, or the transaction carries an extra blob |
| `InvalidBlobCommitmentLength`, `InvalidKZGProofLength` | settlement | not 48 bytes |
| `InvalidBlobsDigest` | settlement | commitment/evaluation digest differs from the proof's |
| `MissingBlob(index)` | settlement | the transaction has no blob at `index` |
| `InvalidBlobOpening(index)` | settlement | the point-evaluation precompile rejects blob `index` |
| `InvalidProof` (`ZiskVerifier`) | settlement, `setProcessResults` | PLONK proof does not verify |
| `InvalidKeyMode` | results calls, `revealProcessKey` | call not valid for the process's key mode |
| `ResultsAlreadyRequested` | `requestResultsDecryption` | second request |
| `InvalidAccumulator` | `requestResultsDecryption` | a coordinate not below the field, or a field with only one identity half |
| `InvalidInclusionProof` | `requestResultsDecryption` | accumulator not included under `latestStateRoot` |
| `ResultsNotReady` | `finalizeResultsFromDKG` | not requested yet, or a combine is still missing |
| `NoLiveEpoch` (adapter) | `newProcess` | `DKG_AUTOMATIC` found no Live epoch with a free pool key |
| `NonContiguousIndex` (adapter) | `requestResultsDecryption` | the DKG assigned non-consecutive ciphertext indices |
| `NotRegistry` (adapter) | adapter `register`, `submit`, `reveal` | caller is not the registry that created the adapter |
| `UnknownRequest` (Council adapter) | `plaintexts` | no request submitted under that request id and ceremony |
| `InvalidFieldRange` (Council adapter) | `plaintexts` | `first != 0` or `count` differs from the request's field count |
| `RequestMismatch` (Council adapter) | `requestResultsDecryption`, `plaintexts` | the manager answered for another request or with the wrong number of values |
| `UnsupportedKeyMode` (Council adapter) | adapter `reveal` | Council has no organizer key |
| `InvalidVerifierConfig` | constructor | zero verifier address or pin |

Errors from davinci-dkg pass through unchanged. On `newProcess` these include `InvalidEpoch`
or `InvalidPhase` for an unknown or not yet Live `DKG_LOCKED` epoch, and `InvalidSchnorrProof`;
on `revealProcessKey`, `InvalidOrganizerSecret` and `AlreadyRevealed`. Council manager errors
pass through too: on `newProcess`, `UnknownCeremony`, `WrongPhase`, `NotAllowedAdapter` and
`NotAuthorizedCreator`; on `requestResultsDecryption`, the request admission errors
(`NotInSubgroup`, `InvalidPoint`, `NonCanonical`, `BadFieldCount`, ...). `IProcessRegistry` also
declares `InvalidBlockNumber`, `InvalidMaxValue`, `InvalidMinValue`, `InvalidMinTotalCost`,
`InvalidUniqueValues`, `CannotAcceptResult`, `ProcessNotEnded` and `ProofInvalid`, which this
registry never raises.
