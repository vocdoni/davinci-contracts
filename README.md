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

2. Install dependencies using the Makefile:

```bash
make install
```

3. Set up environment variables:

```bash
cp .env.example .env
# Edit .env with your configuration
```

## 🛠️ Development

This project uses a `Makefile` to streamline common development tasks.

### Building

Clean artifacts and build the project:

```bash
make build
```

To clean the project (remove artifacts, cache, and generated bindings):

```bash
make clean
```

### Code Quality

Run linters (Solhint and Prettier):

```bash
make lint
```

### Go Bindings

Generate Go bindings for contract integration (requires `abigen`):

```bash
make go-bind
```

## 🧪 Testing

Run the comprehensive test suite:

```bash
# Run all tests
make test

# Run with verbosity (equivalent to -vvv)
make test-v

# Generate gas report
make gas-report
```

## 🚢 Deployment

The deployment process is managed via the `Makefile` and uses configuration from your `.env` file.

### 1. Configuration

Ensure your `.env` file is configured for your target network.

**For Sepolia/Testnets/Mainnet:**
```dotenv
RPC_URL=https://rpc.sepolia.org
PRIVATE_KEY=0xYourPrivateKey...
CHAIN_ID=11155111
ACTIVATE_BLOBS=false
ETHERSCAN_API_KEY=YourEtherscanKey... # Optional, for verification
```

### 2. Deploy

**Deploy to the configured network:**

```bash
make deploy
```
This command will:
1. Load variables from `.env`.
2. Deploy contracts using `forge script`.
3. Verify contracts on Etherscan (if `ETHERSCAN_API_KEY` is provided).

**Deploy to Local Node (Anvil):**

Start a local node in a separate terminal:
```bash
anvil
```

Then run:
```bash
make deploy-local
```
This uses default keys and endpoints for a standard Anvil instance.

### 3. Update Go Bindings

After deployment, update the hardcoded address constants in the Go package:

```bash
make update-addresses
```

### Full Workflow

To build, deploy, and update addresses in one go:

```bash
# For your configured network (.env)
make all

# For local development
make all-local
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