# Vocdoni DAVINCI Contracts

**DISCLAIMER**: The **code** in this repository is a **work-in-progress** and it is not meant to be used in production environments.

[![License: AGPL v3](https://img.shields.io/badge/License-AGPL%20v3-blue.svg)](https://www.gnu.org/licenses/agpl-3.0)
[![Version](https://img.shields.io/badge/version-0.0.7-brightgreen.svg)](https://github.com/vocdoni/contracts-z/releases)

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

1. **OrganizationRegistry**: Manages creation and administration of voting organizations
2. **ProcessRegistry**: Handles voting process lifecycle, state transitions, and results
3. **ZK Verifiers**: On-chain verification of zkSNARK proofs for state transitions and results
4. **Process ID Library**: Utilities for generating unique process identifiers

### Deployed libraries

- **Sepolia**
    - PoseidonT3: `0x1464bD48D1635E9B9F65cFd629d8E9f507A952dD`
    - PoseidonT4: `0xd747896B912C1585b04007c103D10A04e71bfb25`

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

1. Configure network in `.env`:

```bash
PRIVATE_KEY=your_deployment_key
RPC_URL=your_rpc_endpoint
```

2. Deploy:

```bash
forge script script/DeployAll.s.sol --rpc-url $RPC_URL --broadcast --verify
```

## 📚 Documentation

- [Whitepaper](https://whitepaper.vocdoni.io)
- [Introduction](docs/Intro.md)
- [OrganizationRegistry Documentation](docs/OrganizationRegistry.md)
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
