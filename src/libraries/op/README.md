# Vendored Optimism libraries

`RLPReader`, `RLPErrors`, `Bytes` and `MerkleTrie` from ethereum-optimism/optimism
(`packages/contracts-bedrock/src/libraries`), commit 04108aa450c85599b732e8f28473617382523b93, MIT licensed. Only the import paths
changed. Used by `L1BlobVerifier` to read L1 headers and transaction-trie proofs.
