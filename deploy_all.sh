#!/usr/bin/env bash
set -euo pipefail

# Load .env
if [ ! -f .env ]; then
    echo "Error: .env not found in $(pwd)" >&2
    exit 1
fi

. .env

: "${CHAIN_ID:?CHAIN_ID is not set in .env}"
: "${RPC_URL:?RPC_URL is not set in .env}"
: "${PRIVATE_KEY:?PRIVATE_KEY is not set in .env}"

CREATE_VERIFY_ARGS=(--verify)
if [ -n "${ETHERSCAN_API_URL:-}" ]; then
    CREATE_VERIFY_ARGS+=(--verifier-url "$ETHERSCAN_API_URL")
fi
if [ -n "${ETHERSCAN_API_KEY:-}" ]; then
    CREATE_VERIFY_ARGS+=(--etherscan-api-key "$ETHERSCAN_API_KEY")
fi

deploy_library() {
    local lib_path_name="$1"
    local output=""
    local address=""

    output=$(FOUNDRY_VIA_IR=false forge create "$lib_path_name" \
        --chain "$CHAIN_ID" \
        --rpc-url "$RPC_URL" \
        --private-key "$PRIVATE_KEY" \
        --broadcast \
        "${CREATE_VERIFY_ARGS[@]}" \
        --optimize \
        --optimizer-runs 200 \
        -vvvv 2>&1)

    echo "$output" >&2

    address=$(echo "$output" | sed -n 's/.*Deployed to: \(0x[a-fA-F0-9]\{40\}\).*/\1/p' | tail -n1)
    if [ -z "$address" ]; then
        echo "Error: could not parse deployed address for $lib_path_name" >&2
        exit 1
    fi

    echo "$address"
}

echo "Deploying Poseidon libraries..."
POSEIDON_T3_ADDRESS=$(deploy_library "lib/poseidon-solidity/contracts/PoseidonT3.sol:PoseidonT3")
echo "PoseidonT3 deployed at: $POSEIDON_T3_ADDRESS"

POSEIDON_T4_ADDRESS=$(deploy_library "lib/poseidon-solidity/contracts/PoseidonT4.sol:PoseidonT4")
echo "PoseidonT4 deployed at: $POSEIDON_T4_ADDRESS"

echo "Deploying main contracts with linked Poseidon libraries..."

forge script script/DeployAll.s.sol:DeployAllScript \
    --chain "$CHAIN_ID" \
    --rpc-url "$RPC_URL" \
    --broadcast \
    --slow \
    --libraries "lib/poseidon-solidity/contracts/PoseidonT3.sol:PoseidonT3:$POSEIDON_T3_ADDRESS" \
    --libraries "lib/poseidon-solidity/contracts/PoseidonT4.sol:PoseidonT4:$POSEIDON_T4_ADDRESS" \
    --optimize \
    --optimizer-runs 200 \
    --verify \
    -vvvv

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
$SCRIPT_DIR/helpers/write_contract_addresses.sh
