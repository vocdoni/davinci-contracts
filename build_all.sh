#!/bin/bash
forge clean
npx hardhat clean

forge build
npx hardhat typechain

./go_bind.sh
