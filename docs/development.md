# Development

Prerequisites: [Foundry](https://getfoundry.sh/); Node.js, [abigen](https://geth.ethereum.org/docs/tools/abigen)
and [jq](https://jqlang.org/) for the bindings; Go (version in `test/vectors/go.mod`) to
regenerate the test vectors.

```bash
git clone --recurse-submodules https://github.com/vocdoni/davinci-contracts.git
cd davinci-contracts
npm install   # only needed for the bindings and linters
```

## Build and test

```bash
forge build
forge build --sizes                          # contract sizes, as CI runs it
forge test
forge test -vvv --match-path test/DKG.t.sol
forge test --gas-report
forge fmt --check
```

| Test | Covers |
|---|---|
| `ProcessRegistry.t.sol` | status machine, duration, max voters, ballot mode and census validation |
| `NewProcess.t.sol` | census origins and encryption key checks at creation |
| `ProcessMetadata.t.sol` | metadata URI and hash at creation and on update |
| `Genesis.t.sol` | genesis roots against the Go reference |
| `Transition.t.sol` | settlement with real KZG openings (PLONK verifier mocked) |
| `DynamicCensus.t.sol` | census origins 2 and 3 |
| `Results.t.sol` | `setProcessResults`, including the grace gate |
| `Grace.t.sol` | the grace window, `setProcessGrace`, `setProcessMaxVoters` past the end, shortening with notice |
| `DKG.t.sol` | DKG key modes, against a mock DKG with real BabyJubJub arithmetic |
| `Publics.t.sol`, `ZiskVerifier.t.sol` | public values decoding and the vendored verifier, against a recorded batch PLONK |
| `BlobsLib.t.sol`, `ProcessIdLib.t.sol` | library helpers |

The `test` profile of `docker-compose.yml` runs the build and tests of the CI forge job in the
Foundry image:

```bash
docker compose --profile test run --rm test   # forge build --sizes, forge test
```

## Test vectors

The fixtures in `test/vectors/` come from the Go program in the same directory, which builds
them with the davinci-zkvm Go SDK. The tests deploy the registry at the address the fixtures
were built for (anvil account 0, nonce 0, chain id 1337). To regenerate them, check out
davinci-zkvm and davinci-circom next to this repository and run:

```bash
cd test/vectors && go run .
```

CI does the same against davinci-zkvm `main` and a pinned davinci-circom commit, and fails if
the output differs from the committed files.

## Bindings

```bash
./build_all.sh
```

This cleans, runs `forge build`, compiles with Hardhat to produce `artifacts/` and the
TypeChain types in `typechain-types/`, and runs `./go_bind.sh`, which writes the Go bindings
`golang-types/ProcessRegistry.go` and `golang-types/ICensusValidator.go` (package `contracts`)
plus `golang-types/addresses.go` from the broadcast logs. `npm run prepare` compiles the
TypeChain types into `dist/`, the contents of the npm package `@vocdoni/davinci-contracts`.

In CI, pushes that touch `src/**/*.sol` regenerate and commit the bindings, pushes to `release`
bump the patch version and tag it, and `v*` tags publish to npm.

## Linters

```bash
npm run lint:sol
npm run prettier
npm run slither
npm run mythril
```

## Vendored code

`src/verifiers/PlonkVerifier.sol` and `ZiskVerifier.sol` are copied from the ZisK snark setup
and kept byte-identical: `forge fmt` skips them, and reformatting would change the deployed
code that [verification.md](verification.md) compares. `src/interfaces/dkg/` is the subset of
the davinci-dkg interfaces the adapter calls, ABI-identical to upstream.
