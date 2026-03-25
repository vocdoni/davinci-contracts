# Changelog

## 2026-03-10 (`76d2742dd5f15e63f1a57e476daf53b0d674303b`)

### TypeScript consumer

- `DAVINCITypes.CensusOrigin` was simplified. `CENSUS_UNKNOWN` was removed, `MERKLE_TREE_OFFCHAIN_STATIC_V1` and `MERKLE_TREE_OFFCHAIN_DYNAMIC_V1` were collapsed into `MERKLE_TREE_OFFCHAIN`, and `MERKLE_TREE_ONCHAIN_DYNAMIC_V1` was renamed to `MERKLE_TREE_ONCHAIN`. Numeric enum values changed.
- `DAVINCITypes.ParamsMod` was added. `ProcessRegistry.newProcess` now requires a new final `paramsMod` argument after `encryptionKey`.
- `DAVINCITypes.Process` now includes `paramsMod`.
- `ProcessRegistry` now exposes `setProcessMetadata(bytes31 processId, string metadata)`.
- Event names changed: `ProcessDurationChanged` -> `ProcessDurationUpdated`, `ProcessStatusChanged` -> `ProcessStatusUpdated`, and `ProcessMaxVotersChanged` -> `ProcessMaxVotersUpdated`.
- `ProcessMetadataUpdated(bytes31 processId, string metadata)` was added.
- `DAVINCITypes.StateTransitionBatchProofInputs` was added and is now the exported proof input struct used by `ProcessRegistry`.
- `DAVINCITypes.BallotMode` field order changed so `uniqueValues` now comes before `costFromWeight`.

### Behavior

- Process mutability is now controlled per process through `paramsMod`. Status, census, duration, metadata, and max-voter updates can be rejected with `Unauthorized()` when the corresponding flag is disabled.
- Metadata updates are now supported only while the process is `READY` or `PAUSED`, and are blocked for canceled, ended, or results processes.
- Dynamic census behavior is now encoded through `CensusOrigin` plus `ParamsMod.census` instead of dedicated dynamic census enum variants.
- State transitions using on-chain census validation now reject unknown census roots (`getRootBlockNumber(...) == 0`).
- `setProcessStatus(..., ENDED)` normalizes `duration` to `max(endTimestamp - startTime, 0)`. This requires `paramsMod.duration = true` whenever the stored duration changes, including manual endings before or after the scheduled end time.
- Census validation consolidated for `newProcess` and `setProcessCensus`. Fixed on-chain censuses must now include both a committed `censusRoot` and a validator `contractAddress`. Off-chain and CSP censuses must include a non-zero `censusRoot` and now reject `onchainAllowAnyValidRoot`.
- `setProcessCensus` now replaces the full census struct, not just `censusRoot` and `censusURI`, while still requiring the census origin to remain unchanged.
- State transitions for on-chain censuses now always validate the submitted root through `ICensusValidator`. Fixed on-chain censuses must match the committed root, and mutable on-chain censuses reject unknown or future roots and, unless `onchainAllowAnyValidRoot` is set, roots older than the process creation block.

## 2026-02-19 (`72f4b9724db099b62a7f06f9041c25a9e948a995`)

- `ProcessId` uses `bytes31`.
- New deployments.
- CI updated, generates bindings automatically.

## 2026-02-11 (`13c0515f996cf3010ac88bb02feb057a18a41c88`)

### TypeScript consumer

- `ProcessRegistry.newProcess` TypeChain signature removed the `initStateRoot` argument. Calls must drop the previous final parameter.
- `BallotMode` TypeChain struct now includes required `groupSize: BigNumberish` (Solidity `uint8`) between `numFields` and `costExponent`.
- Type namespaces for process-related structs changed from `IProcessRegistry.*` to `DAVINCITypes.*` in generated types.
- `ProcessRegistry__factory` deployment typing now requires linked library addresses via `ProcessRegistryLibraryAddresses`.
- `ProcessRegistry__factory` now requires the `StateRootLib` link key: `"src/libraries/StateRootLib.sol:StateRootLib"`.
- New generated TypeChain exports were added for `PoseidonT3`, `PoseidonT4`, and `StateRootLib` (types + factories).
