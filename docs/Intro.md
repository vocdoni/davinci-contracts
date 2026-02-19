# Introduction

This documentation is up to date with the latest version of the code available.

## Overview

The Vocdoni davinci-contracts repository implements the EVM smart contracts that power the DAVINCI (Decentralized Autonomous Vote Integrity Network with Cryptographic Inference) protocol - a cutting-edge voting system that leverages applied cryptography and blockchain technology to enable secure, verifiable, coercion-resistant and anonymous digital voting.

This contract suite forms the foundation of a specialized zkRollup system that inherits security from Ethereum. The system relies exclusively on cryptographic proofs to ensure integrity and security, eliminating the need for centralized authorities.

## Key Protocol Features

### Zero-Knowledge Proofs

The system uses zkSNARKs at multiple levels:

- **Vote Proofs**: Voters prove their ballot is valid without revealing choices
- **State Transition Proofs**: Prove correct vote aggregation and state updates
- **Results Proofs**: Final tally is proven correct while maintaining vote privacy

### Homomorphic Encryption

Using ElGamal encryption, the system can:

- Aggregate encrypted votes without decryption
- Maintain complete vote privacy
- Support vote overwrites for receipt-freeness

### Decentralized Infrastructure

- Decentralized vote processing through cryptographic proofs
- No single point of failure
- Trustless verification of all state transitions

## Core Components

The protocol consists of smart contracts that work together to orchestrate the entire voting lifecycle.

### 1. **ProcessRegistry**

The heart of the voting system that handles:

- Creation and lifecycle management of voting processes
- State transitions with zero-knowledge proof verification
- Results finalization and verification
- Process status management (READY, PAUSED, ENDED, CANCELED, RESULTS)
- Census updates and duration management

## How the Contracts Interact

### 1. **Voting Process Creation Flow**

- [User]
    1. ProcessRegistry.newProcess()
    2. Provide process parameters (ballot mode, census, duration, etc.)
    3. Process is created with initial state root

### 2. **Vote Processing Flow**

- [Voter]

    1. Generates encrypted vote + zkSNARK proof
    2. Submits vote (off-chain)

- [Sequencer]
    1. Collects and validates votes
    2. ProcessRegistry.submitStateTransition()
        - Verifies zkSNARK proof on-chain
        - Updates state root
        - Increments vote count and overwrite count

### 4. **Results Finalization Flow**

- Process ends by time
    - [Results Provider]
        1. Generates results proof
        2. ProcessRegistry.setProcessResults() verifies results
        3. Process status changes to RESULTS
- Process ends manually
    - [Process creator]
        1. Set process status to ENDED
    - [Results Provider]
        1. Generates results proof
        2. ProcessRegistry.setProcessResults() verifies results
        3. Process status changes to RESULTS

## Process Lifecycle

1. **Initialization**

    - Process initialized with encryption key and initial state root
    - Status set to READY or PAUSED

2. **Voting Phase**

    - Voters encrypt ballots using public key
    - Generate zkSNARK proofs of validity
    - Submit votes off-chain

3. **Aggregation**

    - Votes are verified and accumulated
    - State transitions submitted with proofs
    - Each transition updates the state root and vote counts

4. **Finalization**
    - Process ends by duration or manually (status set to ENDED)
    - Results proof generated and verified
    - Final results published on-chain
    - Process status changes to RESULTS

## Process Management Features

The ProcessRegistry provides several management capabilities:

- **Status Management**: Processes can be READY, PAUSED, ENDED, CANCELED, or RESULTS
- **Census Updates**: It is possible to update census data for ongoing processes
- **Duration Extension**: Process duration can be extended before it ends
- **Manual Termination**: Processes can be ended manually

## Integration Example

```solidity

// 1. Create a voting process
ProcessRegistry.BallotMode memory ballotMode = ProcessRegistry.BallotMode({
    maxCount: 3,
    maxValue: 5,
    uniqueValues: true,
    // ... other parameters
});

ProcessRegistry.Census memory census = ProcessRegistry.Census({
    censusOrigin: ProcessRegistry.CensusOrigin.OFF_CHAIN_TREE,
    censusRoot: bytes32(0xabc...),
    censusURI: "ipfs://census-data"
});

ProcessRegistry.EncryptionKey memory encryptionKey = ProcessRegistry.EncryptionKey({
    publicKey: publicKeyBytes,
    publicKeyHash: keccak256(publicKeyBytes)
});

bytes32 processId = processRegistry.newProcess(
    ProcessRegistry.ProcessStatus.READY,
    block.timestamp,
    7 days,
    ballotMode,
    census,
    "ipfs://process-metadata",
    encryptionKey,
    initialStateRoot
);

// 2. Process continues with off-chain voting and on-chain verification
processRegistry.submitStateTransition(
    processId,
    proof,
    input
);

// 3. Process ends, submit results
processRegistry.setProcessResults(
    processId,
    proof,
    input
);
```
