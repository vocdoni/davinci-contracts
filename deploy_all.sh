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

forge script script/DeployAll.s.sol:DeployAllScript \
    --chain "$CHAIN_ID" \
    --rpc-url "$RPC_URL" \
    --broadcast \
    --slow \
    --optimize \
    --optimizer-runs 200 \
    --verify \
    -vvvv