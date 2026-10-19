# Changelog

## 2026-10-19

v0.0.52: the protocol v5 registry.

- Blob evaluation points. The registry derives each blob's evaluation point from the cell digest
  the batch proof outputs (registers 46 to 53) and opens the blob there, in blob and L1-blob mode;
  it needs a davinci-zkvm v0.4.0 batch program.
- `setProcessStatus` refuses CANCELED from the grace end on (`InvalidTimeBounds`): a results
  transaction may then be pending. Before the grace end the organizer still cancels.
- The constructor refuses equal batch and results program keys, and `newProcess` refuses a zero
  duration.
- `DeployAll` deploys with the chain's own id and refuses a different `CHAIN_ID`.
- Test vectors for protocol v5: the davinci-circom v1.1.0 development ballot key (71 public
  inputs) and a Keccak census root above the BN254 modulus carried through creation and
  register 20. A test pins the vendored `ZiskVerifier`'s runtime code hash
  (`0xadab570d…dcc1`): it covers the solc metadata, so `remappings.txt` keeps the
  `poseidon-solidity/` entry although the library is gone.
- Adversarial test suites for settlement, the grace window, results and L1 blobs.
- Pins for davinci-zkvm v0.4.0: batch `0x1e4a74a1…d2a8`, results `0xde55d751…e4e9`,
  `rootCVadcopFinal` unchanged.

## 2026-10-15

v0.0.51: L1-blob mode, the ZisK 1.3.2-alpha OP Mainnet stack and the OP registry with the DKG and
Council adapters.

- L1-blob mode. `submitStateTransitionL1(processId, publicValues, proofBytes, commitments, ys,
  L1BlobProof)` settles a transition on OP while its blobs live in Ethereum blob transactions. The
  registry's immutable `L1BlobVerifier` walks a header chain anchored at `L1Block.hash()` (OP
  predeploy `0x4200…0015`) or at a header it `record()`ed, proves each blob transaction against its
  header's `transactionsRoot` (vendored OP `RLPReader`/`MerkleTrie`/`Bytes` from
  ethereum-optimism/optimism 04108aa, `src/libraries/op/`) and returns the versioned hashes the
  registry opens with the point-evaluation precompile. `MAX_AGE` is 7200 L1 blocks. New event
  `ProcessStateTransitionedL1(processId, l1BlockNumbers, versionedHashes)`, `OriginRecorded(number,
  hash)` and the `l1BlobVerifier()` getter.
- The mode is immutable. The constructor takes `address _l1BlobVerifier` as its 14th argument; zero
  is blob mode. `submitStateTransition` on an L1 registry and `submitStateTransitionL1` on a blob
  registry revert `WrongDAMode`. New `L1BlobVerifier` errors: `L1OriginMismatch`,
  `BrokenHeaderChain(index)`, `BadHeader(index)`, `TxAboveOrigin(tx)`, `NotBlobTx(tx)`,
  `BadVersionedHash(tx)`, `TxTooOld(tx)`. `MerkleTrie` reverts with `require` strings
  (docs/errors.md, docs/settlement.md).
- ZisK 1.3.2-alpha verifier. `src/verifiers/{PlonkVerifier,ZiskVerifier}.sol` and
  `src/interfaces/IZiskVerifier.sol` are vendored from the 1.3.2 snark setup (`VERSION()`
  `v1.3.2-alpha`); the procedure and expected hashes are in docs/development.md. Runtime code hash
  `0xadab570d3e3a4c172d0c75e5aef5d06098b94d70a9c8ff76af239a90fdcadcc1`. New pins: batch
  `0xe39e4600…988b`, results `0x39469455…d167`, `rootCVadcopFinal` `0x6b2490fb…b265`. The 1.3.0
  verifier hard-codes the old root, so a 1.3.2 registry needs its own verifier.
  `test/vectors/recorded_batch_snark.json` and `publics.json` are re-recorded from a real 1.3.2 job.
- Deploy tooling. `DeployAll` reads `L1_BLOCK` (deploys an `L1BlobVerifier`) or `L1_BLOB_VERIFIER`
  (reuses one; refused without code or on another `L1Block`). `deploy_all.sh` requires
  `EXPECT_DA_MODE=l1|blob` off the local chains and takes `EXPECT_GRACE`. `verify_deployment.py`
  gains `--l1-block`, `--l1-blob-verifier ADDR|0x0` and `--grace`, fails closed on an unreadable
  getter and scrubs RPC URLs from its errors.
- Go bindings: network `optimism` (10) with `L1BlobVerifierContract`,
  `ProcessRegistryOptimismAddress`, `ZiskVerifierOptimismAddress`, `L1BlobVerifierOptimismAddress`
  and `GetOptimismAddresses()`. Hardhat artifacts and TypeChain types regenerated, including
  `L1BlobVerifier`.
- OP Mainnet (chain id 10) is the production chain. `ProcessRegistry`
  `0x64d41aD49009204E8dAcc153Fc98a5Ee14cB6916` (block 158245426, 2026-10-15, commit `338ac77`,
  deployer `0x5eb38f2D16CEaabFF6a63C17505a14b45624f537`, L1-blob mode) with the DKG adapter
  `0x0c1E7120FaBf95fDe294839A78780334d27cfCff` (DKGManager `0x9e799a1d…2d41` from davinci-dkg,
  block 158245366; DKGAppManager `0x1e7f0bec…c39a`) and the Council adapter
  `0xa5A1e4B204976709EB545Af0189211c41e20193d` (Council manager `0x281b84de…396e`
  from davinci-dkg-council, circuits-v1 development setup, block 158245395). It reuses the
  `ZiskVerifier` `0x35Ad4cBc4C77a2Bc69B07e68f2535Da5C6d07AB7`, the `L1BlobVerifier`
  `0x3B95BC14F43ed8Addb83F84D10A949De74eE7529`, the 4 pins and grace 300/240/600/1800/60 of the
  previous registry. Record `deployments/10.json`; `verify_deployment.py` passed 32 checks.
- The previous OP registry `0x7bfa68879f118C98B864FB52d2aD3a112eF1558F` (block 158164175) is
  retired: no DKG or Council adapter, record `deployments/10-0x7bfa.json`.
- Retired, still on chain: `0xC5f88160a1Df51A3Ee56b970b12B8308799c05A7` (v0.2.0 vks on ZisK 1.3.0,
  `deployments/10-0xC5f8.json`), `0x8b63fAE96948D69CF92Eb40e4EE0fCE40962267d` (v0.1 vks, grace floor
  150 s, `deployments/10-0x8b63.json`) and the Gnosis registries (`0x20b96e…eA21` retired on
  2026-10-13). Gnosis Chain is retired.
- None of the OP contracts is source-verified on an explorer yet; Optimistic Etherscan
  verification of the new registry is pending.

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
