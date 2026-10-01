# Changelog

## Unreleased

- `newProcess` takes `bytes32 metadataHash` right after the metadata URI: the SHA-256 of the
  exact bytes served there. An empty URI or a zero hash reverts `InvalidMetadata`.
  `DAVINCITypes.Process` gains `metadataHash` after `metadataURI`.
- `setProcessMetadata(processId, metadataURI, metadataHash)` lets the organizer replace both
  while the process is `READY` or `PAUSED` and before its end.
- `ProcessMetadataUpdated(processId, metadataURI, metadataHash)` is emitted by `newProcess` and
  `setProcessMetadata`.

## 2026-09-28

- `submitStateTransition` verifies the davinci-zkvm batch PLONK through `ZiskVerifier`
  against pinned program vks, binds EIP-4844 blobs with point-evaluation openings and
  checks root continuity, the census root and `occupied_before`.
- `setProcessResults` verifies the davinci-zkvm results PLONK.
- Dynamic census origins with a settlement-time root check.
- Key modes: `newProcess` takes a trailing `DKGParams`. `DKG_AUTOMATIC` and `DKG_LOCKED`
  take the process key from davinci-dkg through `DavinciDKGAdapter`, and results come from
  `requestResultsDecryption` + `finalizeResultsFromDKG`; `revealProcessKey` unlocks a
  locked process. The constructor takes `dkgManager` as a 7th argument.
- `script/verify_deployment.py` checks a live deployment against the local build.
- Gnosis chain (100) deployment.

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
