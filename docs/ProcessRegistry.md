# ProcessRegistry

## Overview

The `ProcessRegistry` contract is a core component of the Vocdoni voting system that manages the complete lifecycle of voting processes. It serves as the central registry for all voting processes, handling their creation, state management, vote aggregation, and results finalization.

### Key Responsibilities

- **Process Creation**: Creates new voting processes with configurable parameters
- **State Management**: Manages process status transitions throughout the voting lifecycle
- **Census Management**: Handles voter eligibility data and updates
- **Vote Aggregation**: Processes state transitions with zero-knowledge proof verification
- **Results Finalization**: Verifies and stores final voting results

## Architecture

### Contract Inheritance

```
ProcessRegistry ──implements──> IProcessRegistry
```

### Dependencies

- **IZKVerifier**: Verifies zero-knowledge proofs for state transitions and results
- **ProcessIdLib**: Computes unique process identifiers

## State Variables

### Constants

| Name                | Type    | Value | Description                           |
| ------------------- | ------- | ----- | ------------------------------------- |
| `MAX_CENSUS_ORIGIN` | `uint8` | 2     | Maximum value for census origin enum  |
| `MAX_STATUS`        | `uint8` | 4     | Maximum value for process status enum |

### Storage Variables

| Name           | Type                          | Description                                |
| -------------- | ----------------------------- | ------------------------------------------ |
| `processes`    | `mapping(bytes32 => Process)` | Maps process IDs to process data           |
| `processNonce` | `mapping(address => uint64)`  | Tracks process creation nonce per address  |
| `processCount` | `uint32`                      | Total number of processes created          |
| `chainID`      | `uint32`                      | Chain identifier for process ID generation |
| `stVerifier`   | `address`                     | State transition ZK verifier contract      |
| `rVerifier`    | `address`                     | Results ZK verifier contract               |

## Data Structures

### Enums

#### ProcessStatus

Defines the possible states of a voting process:

```solidity
enum ProcessStatus {
    READY, // Process is active and accepting votes
    ENDED, // Process has ended, awaiting results
    CANCELED, // Process was canceled
    PAUSED, // Process is temporarily paused
    RESULTS // Results have been finalized
}
```

#### CensusOrigin

Specifies the type of census used for voter eligibility:

```solidity
enum CensusOrigin {
    CENSUS_UNKNOWN,                 // Unknown census type
    MERKLE_TREE_OFFCHAIN_STATIC_V1, // Off-chain static Merkle tree
    CSP_EDDSA_BN254_V1              // CSP over BN254
}
```

### Structs

#### Process

Complete process data structure:

```solidity
struct Process {
    ProcessStatus status; // Current process status
    address organizationId; // Process creator and controller
    EncryptionKey encryptionKey; // Public key for vote encryption
    uint256 latestStateRoot; // Current state root
    uint256[] result; // Final voting results
    uint256 startTime; // Process start timestamp
    uint256 duration; // Process duration in seconds
    uint256 votersCount; // Total voters that participated
    uint256 overwrittenVotesCount; // Total vote overwrites
    string metadataURI; // URI to process metadata
    BallotMode ballotMode; // Voting configuration
    Census census; // Voter eligibility data
}
```

#### BallotMode

Defines how votes are cast and counted, see [blog post](https://blog.vocdoni.io/vocdoni-ballot-protocol) explaining the ballot protocol.

```solidity
struct BallotMode {
    bool costFromWeight; // Use voter weight as maxTotalCost
    bool forceUniqueness; // Require unique choices
    uint8 maxCount; // Max choices per ballot
    uint8 costExponent; // Exponent for cost calculation
    uint256 maxValue; // Maximum value per choice
    uint256 minValue; // Minimum value per choice
    uint256 maxTotalCost; // Max sum of all choices
    uint256 minTotalCost; // Min sum of all choices
}
```

#### Census

Voter eligibility configuration:

```solidity
struct Census {
    CensusOrigin censusOrigin; // Type of census
    bytes32 censusRoot; // Merkle root of census
    string censusURI; // URI to census data
}
```

#### EncryptionKey

Public key for vote encryption:

```solidity
struct EncryptionKey {
    uint256 x; // X coordinate on the curve
    uint256 y; // Y coordinate on the curve
}
```

## Functions

#### newProcess

```solidity
function newProcess(
    ProcessStatus status,
    uint256 startTime,
    uint256 duration,
    BallotMode calldata ballotMode,
    Census calldata census,
    string calldata metadata,
    EncryptionKey calldata encryptionKey,
    uint256 initStateRoot
) external returns (bytes32)
```

Creates a new voting process with specified parameters.

**Parameters:**

- `status`: Initial process status (READY or PAUSED)
- `startTime`: Process start time (0 for current block timestamp)
- `duration`: Process duration in seconds
- `ballotMode`: Voting configuration parameters
- `census`: Voter eligibility configuration
- `metadata`: URI to process metadata
- `encryptionKey`: Public key for vote encryption
- `initStateRoot`: Initial state root

**Returns:** Unique process ID

**Validation:**

- Process ID must not already exist
- Ballot mode parameters must be valid
- Census parameters must be valid
- Status must be READY or PAUSED
- Start time must be current or future
- End time must be in the future

**Effects:**

- Stores process data
- Increments process count and nonce
- Emits `ProcessCreated` event

#### setProcessStatus

```solidity
function setProcessStatus(bytes32 processId, ProcessStatus newStatus) external
```

Updates the status of an existing process.

**Parameters:**

- `processId`: Process identifier
- `newStatus`: New status to set

**Access Control:** Only process owner

**Valid Transitions:**

- READY → PAUSED, CANCELED, ENDED
- PAUSED → READY, CANCELED, ENDED
- ENDED → RESULTS (via `setProcessResults`)
- CANCELED → No transitions
- RESULTS → No transitions

**Special Behavior:**

- When setting to ENDED, duration is updated to actual elapsed time

**Events:** `ProcessStatusChanged`, `ProcessDurationChanged` (if ENDED)

#### setProcessCensus

```solidity
function setProcessCensus(bytes32 processId, Census calldata census) external
```

Updates the census of an ongoing process.

**Parameters:**

- `processId`: Process identifier
- `census`: New census configuration

**Access Control:** Only process owner

**Validation:**

- Census origin must match original
- Max votes can only increase
- Census root and URI must be valid
- Process must be READY or PAUSED
- Process must not have ended

**Events:** `CensusUpdated`

#### setProcessDuration

```solidity
function setProcessDuration(bytes32 processId, uint256 _duration) external
```

Extends the duration of an ongoing process.

**Parameters:**

- `processId`: Process identifier
- `_duration`: New duration in seconds

**Access Control:** Only process owner

**Validation:**

- Process must be READY or PAUSED
- New duration must be non-zero
- New end time must be in the future
- New end time must be later than current end time

**Events:** `ProcessDurationChanged`

#### submitStateTransition

```solidity
function submitStateTransition(
    bytes32 processId,
    bytes calldata proof,
    bytes calldata input
) external
```

Submits a state transition with zero-knowledge proof.

**Parameters:**

- `processId`: Process identifier
- `proof`: Zero-knowledge proof
- `input`: Public inputs (encoded as uint256[4])

**Input Format:**

- `input[0]`: Previous state root
- `input[1]`: New state root
- `input[2]`: Vote count
- `input[3]`: Vote overwrite count

**Validation:**

- Process must exist and be READY
- Process must not have ended
- ZK proof must be valid
- Previous state root must match

**Effects:**

- Updates latest state root
- Increments vote counts
- Emits `ProcessStateTransitioned`

#### setProcessResults

```solidity
function setProcessResults(
    bytes32 processId,
    bytes calldata proof,
    bytes calldata input
) external
```

Finalizes process results with zero-knowledge proof.

**Parameters:**

- `processId`: Process identifier
- `proof`: Zero-knowledge proof
- `input`: Public inputs (encoded as uint256[9])

**Input Format:**

- `input[0]`: Final state root
- `input[1-8]`: Voting results

**Validation:**

- Process must exist and be ENDED
- Process must have ended by time
- ZK proof must be valid
- State root must match

**Effects:**

- Sets status to RESULTS
- Stores voting results
- Emits `ProcessStatusChanged` and `ProcessResultsSet`

### View Functions

#### getProcess

```solidity
function getProcess(bytes32 processId) external view returns (Process memory)
```

Returns complete process data.

#### getProcessEndTime

```solidity
function getProcessEndTime(bytes32 processId) external view returns (uint256)
```

Returns the calculated end time (startTime + duration).

#### getSTVerifierVKeyHash

```solidity
function getSTVerifierVKeyHash() external view returns (bytes32)
```

Returns the state transition verifier proving key hash.

#### getRVerifierVKeyHash

```solidity
function getRVerifierVKeyHash() external view returns (bytes32)
```

Returns the results verifier proving key hash.

#### getNextProcessId

```solidity
function getNextProcessId() external view returns (bytes32)
```

Returns the next process ID for the calling address.

## Events

### ProcessCreated

```solidity
event ProcessCreated(bytes32 indexed processId, address indexed creator)
```

Emitted when a new process is created.

### CensusUpdated

```solidity
event CensusUpdated(
    bytes32 indexed processId,
    bytes32 censusRoot,
    string censusURI,
)
```

Emitted when process census is updated.

### ProcessDurationChanged

```solidity
event ProcessDurationChanged(bytes32 indexed processId, uint256 duration)
```

Emitted when process duration is modified.

### ProcessStateTransitioned

```solidity
event ProcessStateTransitioned(
    bytes32 indexed processId,
    address indexed sender,
    uint256 oldStateRoot,
    uint256 newStateRoot,
    uint256 newVotersCount,
    uint256 newOverwrittenVotesCount
)
```

Emitted when state root is updated via state transition.

### ProcessResultsSet

```solidity
event ProcessResultsSet(
    bytes32 indexed processId,
    address indexed sender,
    uint256[] result
)
```

Emitted when process results are finalized.

### ProcessStatusChanged

```solidity
event ProcessStatusChanged(
    bytes32 indexed processId,
    ProcessStatus oldStatus,
    ProcessStatus newStatus
)
```

Emitted when process status changes.

## Errors

### Process Validation Errors

- `ProcessAlreadyExists()`: Process ID already exists
- `ProcessNotFound()`: Process does not exist
- `InvalidProcessId()`: Process ID is zero

### Status and Timing Errors

- `InvalidStatus()`: Invalid status or transition
- `InvalidStartTime()`: Start time is in the past
- `InvalidDuration()`: Invalid duration value
- `InvalidTimeBounds()`: Time constraints violated
- `ProcessNotEnded()`: Process hasn't ended yet

### Ballot Mode Errors

- `InvalidMaxCount()`: Max count is zero
- `InvalidMaxValue()`: Max value less than min value
- `InvalidMinValue()`: Invalid minimum value
- `InvalidMaxTotalCost()`: Invalid max total cost
- `InvalidMinTotalCost()`: Invalid min total cost
- `InvalidTotalCostBounds()`: Total cost bounds invalid

### Census Errors

- `InvalidCensusRoot()`: Census root is zero
- `InvalidCensusURI()`: Census URI is empty
- `InvalidCensusOrigin()`: Invalid census origin

### Verification Errors

- `InvalidStateRoot()`: State root mismatch
- `ProofInvalid()`: ZK proof verification failed
- `CannotAcceptResult()`: Cannot set results

### Access Control Errors

- `Unauthorized()`: Caller is not authorized

## Security Considerations and others

### Access Control

- Only process owners (organizationId) can modify their processes

### Zero-Knowledge Proofs

- State transitions require valid ZK proofs
- Results require valid ZK proofs
- Proofs ensure integrity of vote aggregation

### State Integrity

- Process IDs are deterministically computed
- State roots chain previous states
- Vote counts are monotonically increasing

## Integration Guidelines

### Creating a Process

1. Prepare process parameters
2. Call `newProcess` with parameters
3. Store returned process ID

### Managing Process Lifecycle

1. Start with READY or PAUSED status
2. Use `setProcessStatus` for transitions
3. Update census if needed with `setProcessCensus`
4. Extend duration if needed with `setProcessDuration`

### Processing Votes

1. Submit state transitions with ZK proof
2. Each transition includes ZK proof
3. State root chains from previous
4. Vote counts accumulate

### Finalizing Results

1. Wait for process to end by time
2. Ensure status is ENDED
3. Submit results with ZK proof
4. Status automatically changes to RESULTS
5. Results become immutable
