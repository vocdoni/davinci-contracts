# Changelog

## 2026-10-07

- Gnosis production beta registry `0x20b96e465CA7C3536B9C733571ec1eCf42b2eA21` (block
  48,633,301): DKG modes on davinci-dkg `0xC6Fb38c42ed3FB35D363a702218746d5C7Da36BF`, `COUNCIL`
  on the Council manager `0x2f5b110864cbad4017fe8ac59111812278f5f71f`, verifier, pins and grace
  settings inherited from the previous registry `0x6702e014…E7D`, now retired.
- `DavinciDKGAdapter.aidFor` (and so `registry.aidFor` and the stored `dkgAid`) returns
  `salt << 160 | adapter`, with `salt` the top 92 bits of
  `keccak256(abi.encode(chainid, registry, processId))`, instead of that hash mod Q. Paired with
  davinci-dkg's rule that an application id's low 160 bits are its registrant, nobody can
  register a process's id ahead of the adapter and block its creation (vocdoni/davinci-dkg#14).
  Clients that read the id from `aidFor` need no change; ones that computed it must follow.
- Deployment: `DeployAll` takes `PINS_FROM_REGISTRY` to reuse a live registry's verifier and copy
  its pins and grace settings, and rejects managers without code. `deploy_all.sh` checks every
  deployment with `verify_deployment.py` (now also the Council adapter, both managers and the
  application id namespace) and writes the verified record to `deployments/<chain id>.json`.
- `KeyMode.COUNCIL` (3): the process key is the key of a Live Council ceremony (invite-only
  threshold DKG), named by `DKGParams.epochId`; the organizer and PoP fields must be zero. The
  constructor takes `address _councilManager` after `_dkgManager` (zero disables the mode,
  `CouncilDisabled`) and creates a `CouncilAdapter`, exposed as `councilAdapter()`, after the
  DKG adapter. The adapter binds the process with the creator (the `newProcess` caller), stores
  the request id as `dkgAid`, submits the accumulator as one request and reads back the whole
  vector. `requestResultsDecryption` and `finalizeResultsFromDKG` pick the adapter by key mode;
  `revealProcessKey` rejects `COUNCIL` with `InvalidKeyMode`. `DeployAll` reads
  `COUNCIL_MANAGER`.
- Council decryption gate (Council protocol v2): `CouncilAdapter.isDecryptionOpen(cid)` reads
  the manager's gate. For `COUNCIL` only, `requestResultsDecryption` no longer finalizes an
  all-identity tally before the ceremony opens decryption (the process stays `ENDED`), and
  `finalizeResultsFromDKG` reverts with the new `DecryptionNotOpen` until it opens, on both
  paths. The vendored `ICouncilManager` is the v2 adapter surface: `isDecryptionOpen` added,
  `getRequest` replaced by `getRequestMeta`; `ICouncilManagerErrors` gains `DecryptionNotOpen`.
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
