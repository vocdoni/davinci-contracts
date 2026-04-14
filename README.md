# Vocdoni DAVINCI Contracts

**DISCLAIMER**: The **code** in this repository is a **work-in-progress** and it is not meant to be used in production environments.

[![License: AGPL v3](https://img.shields.io/badge/License-AGPL%20v3-blue.svg)](https://www.gnu.org/licenses/agpl-3.0)
[![Version](https://img.shields.io/badge/version-0.0.37-brightgreen.svg)](https://github.com/vocdoni/davinci-contracts/releases)

Smart contracts powering DAVINCI's (Decentralized Autonomous Vote Integrity Network with Cryptographic Inference) digital voting protocol - a cutting-edge voting system that leverages zero-knowledge proofs and blockchain technology to enable secure, verifiable, coercion-resistant, and anonymous digital voting.

## 📋 Table of Contents

- [Overview](#overview)
- [Architecture](#architecture)
- [Installation](#installation)
- [Development](#development)
- [Testing](#testing)
- [Deployment](#deployment)
- [Documentation](#documentation)
- [Contributing](#contributing)
- [License](#license)

## 🔍 Overview

The Vocdoni DAVINCI contracts work together with a set of sequencers that implement a specialized zkRollup system that enables secure digital voting with complete privacy guarantees. The system uses multiple layers of cryptographic proofs:

- **Identity Proofs**: Voters prove their right to participate via identity proofs
- **Vote Proofs**: Voters prove their ballot is valid without revealing choices
- **State Transition Proofs**: Prove correct vote aggregation and state updates
- **Results Proofs**: Final tally is proven correct while maintaining vote privacy

## 🏗️ Architecture

### Core Components

1. **ProcessRegistry**: Handles voting process lifecycle, state transitions, and results
2. **ZK Verifiers**: On-chain verification of zkSNARK proofs for state transitions and results
3. **Process ID Library**: Utilities for generating unique process identifiers

### Deployed libraries

- **Sepolia**
    - PoseidonT3: `0x1464bD48D1635E9B9F65cFd629d8E9f507A952dD`
    - PoseidonT4: `0xd747896B912C1585b04007c103D10A04e71bfb25`
- **Arbitrum**
    - PoseidonT3: `0xE17D701EA8f34022F97fC2Ec68c73D42bF99D0BD`
    - PoseidonT4: `0x2DaF913D423128258b2F378E320F9D9D3Be5eCf5`
    - BlobsLib: `0xca08064a83C7Bea0DB7514eE37F6f4a193b2B53c`
    - ProcessIdLib: `0xDda6c75d32c375946C8ae9be41B2F3539dB1118A`
    - StateRootLib: `0x808276962217AD1ED3af7D51bFc791903CAd9389`

## 📦 Installation

### Prerequisites

- [Node.js](https://nodejs.org/) >= 16.0.0
- [Foundry](https://getfoundry.sh/)
- [Git](https://git-scm.com/)
- [Abigen](https://geth.ethereum.org/docs/tools/abigen)
- [jq](https://jqlang.org/)

### Setup

1. Clone the repository:

```bash
git clone https://github.com/vocdoni/davinci-contracts.git
cd davinci-contracts
```

2. Install dependencies:

```bash
npm install
forge install
```

3. Set up environment variables:

```bash
cp .env.example .env
# Edit .env with your configuration
```

4. Build the project:

```bash
./build_all.sh
```

## 🛠️ Development

### Building

```bash
# Clean and build everything
./build_all.sh

# Or build individually
forge build
npx hardhat compile
```

### Code Quality

```bash
# Linting
npm run lint:sol
npm run prettier

# Security analysis
npm run slither
npm run mythril
```

### TypeScript Support

The project includes TypeScript bindings:

```bash
npm run typechain
```

### Go Bindings

Generate Go bindings for contract integration:

```bash
./go_bind.sh
```

## 🧪 Testing

Run the comprehensive test suite:

```bash
# Run all tests
forge test

# Run with verbosity
forge test -vvv

# Run specific test file
forge test --match-path test/ProcessRegistry.t.sol

# Gas reporting
forge test --gas-report
```

## 🚢 Deployment

### Local Development

1. Start a local node:

```bash
anvil
```

2. Deploy contracts:

```bash
forge script script/DeployAll.s.sol --rpc-url http://localhost:8545 --broadcast
```

### Testnet/Mainnet Deployment

1. Configure shared values in `.env` and chain-specific values in `.env.<chain>`.

Shared `.env`:

```bash
PRIVATE_KEY=your_deployment_key
ETHERSCAN_API_KEY=your_explorer_key
VERIFY_MODE=auto
DEPLOY_CHAINS=base,sepolia
```

Chain file, for example `.env.base`:

```bash
RPC_URL=your_rpc_endpoint
CHAIN_ID=your_chain_id
ACTIVATE_BLOBS=True

# Optional: reuse already deployed libraries.
# If any of these are unset or point to an address without bytecode,
# deploy_all.sh will deploy that library and print export lines you can reuse.
POSEIDON_T3_ADDRESS=
POSEIDON_T4_ADDRESS=
STATE_ROOT_LIB_ADDRESS=
PROCESS_ID_LIB_ADDRESS=
BLOBS_LIB_ADDRESS=
```

2. Deploy:

```bash
./deploy_all.sh
```

`deploy_all.sh` loads `.env` first and then uses the currently exported
single-chain variables. For direct single-chain use, source one chain file into
your shell first or export the variables manually.

### Multi-Chain Deployment

Use a shared `.env` plus one per-chain file for each target network.

Shared `.env`:

```bash
PRIVATE_KEY=your_deployment_key
LOCALHOST_RPC_URL=http://127.0.0.1:8545
ETHERSCAN_API_KEY=your_explorer_key
VERIFY_MODE=auto
DEPLOY_CHAINS=base,sepolia,arbitrum
```

Per-chain `.env.<chain>` files, for example `.env.base`:

```bash
CHAIN_ID=8453
RPC_URL=https://your-base-rpc
ACTIVATE_BLOBS=False

# Optional per-chain overrides
ETHERSCAN_API_URL=
POSEIDON_T3_ADDRESS=
POSEIDON_T4_ADDRESS=
STATE_ROOT_LIB_ADDRESS=
PROCESS_ID_LIB_ADDRESS=
BLOBS_LIB_ADDRESS=
```

Then run:

```bash
./deploy_all_contracts_to_all_chains.sh
```

The wrapper:
1. Loads shared values from `.env`
2. Clears chain-scoped deployment variables
3. Loads `.env.<chain>` for each chain listed in `DEPLOY_CHAINS`
4. Falls back to `.env-<chain>` if the dotted filename does not exist
5. Calls `./deploy_all.sh` once per chain

This means old single-chain values left in `.env` will not bleed into multi-chain runs.

`deploy_all.sh` resolves libraries in this order:
1. `PoseidonT3`
2. `PoseidonT4`
3. `StateRootLib` (linked against Poseidon)
4. `ProcessIdLib`
5. `BlobsLib`

Then it deploys the main contracts with explicit linking for all of them.

Verification behavior is controlled by `VERIFY_MODE`:
- `auto`: disable verification on local chains (`31337`, `1337`), enable otherwise
- `true`: always attempt verification
- `false`: never verify

## 📚 Documentation

- [Whitepaper](https://whitepaper.vocdoni.io)
- [Introduction](docs/Intro.md)
- [ProcessRegistry Documentation](docs/ProcessRegistry.md)

## 🤝 Contributing

We welcome contributions!

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📜 License

This project is licensed under the GNU Affero General Public License v3.0 - see the [LICENSE](LICENSE.md) file for details.

## 🔗 Links

- [DAVINCI Website](https://davinci.vote)
- [Vocdoni Website](https://vocdoni.io)
- [Discord Community](https://chat.vocdoni.io)
- [Twitter](https://twitter.com/vocdoni)

## 🙏 Acknowledgments

Built with ❤️ by Vocdoni.
