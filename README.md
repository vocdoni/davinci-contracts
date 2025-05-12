# Set up local env

1. Installed foundry -> `curl -L https://foundry.paradigm.xyz | bash`
2. `npm i`
3. `cp example.env .env`
4. `source .env`
5. `./setup_local_env.sh`

## Foundry

### Install

```bash
curl -L https://foundry.paradigm.xyz | bash
```

### Setup

`npm i`

**Foundry is a blazing fast, portable and modular toolkit for Ethereum application development written in Rust.**

Foundry consists of:

-   **Forge**: Ethereum testing framework (like Truffle, Hardhat and DappTools).
-   **Cast**: Swiss army knife for interacting with EVM smart contracts, sending transactions and getting chain data.
-   **Anvil**: Local Ethereum node, akin to Ganache, Hardhat Network.
-   **Chisel**: Fast, utilitarian, and verbose solidity REPL.

## Documentation

https://book.getfoundry.sh/

## Usage

### Build

```shell
$ forge build
```

### Test

```shell
$ forge test
```

### Format

```shell
$ forge fmt
```

### Gas Snapshots

```shell
$ forge snapshot
```

### Anvil

```shell
$ anvil
```

### Deploy

```shell
$ forge script script/Counter.s.sol:CounterScript --rpc-url <your_rpc_url> --private-key <your_private_key>
```

### Cast

```shell
$ cast <subcommand>
```

### Help

```shell
$ forge --help
$ anvil --help
$ cast --help
```

### Typechain

```shell
$ npx hardhat compile
```

### Golang bindings

```shell
$ ./go_bind.sh
```

### Deployments roadmap

- [x] Sepolia
- [ ] Mainnet
- [ ] [Celo](https://celo.org) L2
