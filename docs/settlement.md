# Settlement and results

What the registry checks when a sequencer settles a batch of votes and when the tally is
published. Every call here is permissionless: the proofs, the chain of state roots and the blob
openings authenticate it. When each call is allowed is described in
[registry.md](registry.md#lifecycle).

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
is in davinci-zkvm's
[`circuit/CIRCUIT.md`](https://github.com/vocdoni/davinci-zkvm/blob/main/circuit/CIRCUIT.md)):

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

1. The process exists, is `READY`, `ENDED`, or `PAUSED` past its end time, and
   `startTime <= block.timestamp < getProcessGraceEnd(processId)`.
2. `publicValues` is 512 bytes, `ok == 1` and `fail_mask == 0`.
3. The state root before equals `latestStateRoot`.
4. The census root matches: for origins 1, 2 and 4 it equals the stored root; for origin 3 the
   census contract's `getRootBlockNumber(root)` must be non-zero, at most `block.number` and at
   least the process creation block. The census call gets 100k gas and must return a full word.
5. `occupied_before` equals `votersCount`, the number of distinct ballot slots written so far.
   The guest cannot see the tree, so the registry pins it.
6. `votes > 0` (`EmptyTransition` otherwise) and `votersCount + votes - overwrites <= maxVoters`.
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
`votes - overwrites`, `overwrittenVotesCount` by `overwrites`, `batchNumber` by one,
`lastVoteAt` becomes `block.timestamp`, and the registry emits `ProcessStateTransitioned`.

## Results in sequencer mode

```solidity
function setProcessResults(bytes31 processId, bytes calldata publicValues, bytes calldata proofBytes) external;
```

The results guest proves the tally of a state root. The registry requires a sequencer-mode
process that is not `CANCELED` or `RESULTS`, has ended and has a closed grace window
(`GraceOpen` otherwise), `ok == 1` and `fail_mask == 0`, a state root (registers 2..9) equal to
`latestStateRoot`, and a PLONK proof verified against `resultsProgramVK`. Result `i` is read
from registers `10 + 2i` (low 32 bits) and `11 + 2i` (high 32 bits) for each of the `numFields`
fields. The process moves to `RESULTS` and the call emits `ProcessStatusChanged` and
`ProcessResultsSet`.

## Results in the DKG modes

DKG and `COUNCIL` processes do not use `setProcessResults` (it reverts with `InvalidKeyMode`).
Both go through the calls below; the registry picks the adapter by the process's key mode, the
`CouncilAdapter` for `COUNCIL` and the `DavinciDKGAdapter` otherwise.

```solidity
function requestResultsDecryption(bytes31 processId, uint256[64] calldata accumulator, bytes32[] calldata siblings) external;
function finalizeResultsFromDKG(bytes31 processId) external;
function revealProcessKey(bytes31 processId, uint256 sk) external; // DKG_LOCKED only
```

`requestResultsDecryption` runs once per process, after it has ended and its grace window has
closed. `accumulator` holds the 16 ElGamal ciphertexts of the results leaf as
`(C1x, C1y, C2x, C2y)` each, circomlib form, big-endian. `siblings` is the SMT inclusion proof
of leaf `0x04`, root to leaf, zero-padded, 1 to 64 entries and ending in a zero entry. The
registry:

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

For `COUNCIL` the active ciphertexts go to the Council manager as one request
(`submitRequest(ceremonyId, processId, cts)`), in circomlib form, with `firstIndex` 0;
`epochId` and `aid` in the event are the ceremony id and the request id. The manager checks
each half is canonical, in the prime subgroup and not the identity.

If every declared field is identity nothing is submitted and the results are finalized to
zero at once.

`finalizeResultsFromDKG` reads the combined plaintexts once the committee has decrypted every
submitted ciphertext, stores `numFields` results in field order (0 for skipped fields), moves
the process to `RESULTS` and emits `ProcessStatusChanged` and `ProcessResultsSet`. Until then it
reverts with `ResultsNotReady` (`GraceOpen` while the grace window is open).

For `COUNCIL`, finalization reads the whole request at once and is ready only when every
field is combined. `revealProcessKey` reverts with `InvalidKeyMode`: there is no organizer key.

`revealProcessKey` forwards the organizer secret of a `DKG_LOCKED` process to
`DKGAppManager.revealOrganizerSecret`, which checks `sk·G == PK_org` and accepts it once. The
registry does not restrict when it is called.

DKG liveness is election liveness. Once the ciphertexts are submitted there is no fallback: if
more than `n - t` members of the registration epoch's committee are gone, the combines never
complete and the process never reaches `RESULTS`.
