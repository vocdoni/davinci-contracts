#!/usr/bin/env bash
set -euo pipefail

# Resolve script location and always run from repo root
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd "$SCRIPT_DIR"

# Load .env if available (explicitly exported environment variables still work)
if [ -f .env ]; then
    . .env
else
    echo "Warning: .env not found in $SCRIPT_DIR. Using current environment variables." >&2
fi

log_info() { echo "[deploy_all] $*"; }
log_warn() { echo "[deploy_all] Warning: $*" >&2; }
log_error() { echo "[deploy_all] Error: $*" >&2; }

require_cmd() {
    local cmd="$1"
    if ! command -v "$cmd" >/dev/null 2>&1; then
        log_error "$cmd is required but not installed."
        exit 1
    fi
}

require_cmd forge
require_cmd node

: "${CHAIN_ID:?CHAIN_ID is not set (env or .env)}"
: "${RPC_URL:?RPC_URL is not set (env or .env)}"
: "${PRIVATE_KEY:?PRIVATE_KEY is not set (env or .env)}"

ENABLE_VERIFY=true
VERIFY_MODE="${VERIFY_MODE:-auto}"
case "$VERIFY_MODE" in
    false|False|FALSE|0|no|NO)
        ENABLE_VERIFY=false
        ;;
    auto)
        if [ "$CHAIN_ID" = "31337" ] || [ "$CHAIN_ID" = "1337" ]; then
            ENABLE_VERIFY=false
        fi
        ;;
    true|True|TRUE|1|yes|YES)
        ENABLE_VERIFY=true
        ;;
    *)
        log_error "Invalid VERIFY_MODE='$VERIFY_MODE'. Use: auto|true|false."
        exit 1
        ;;
esac

VERIFY_ARGS=()
if [ "$ENABLE_VERIFY" = true ]; then
    VERIFY_ARGS+=(--verify)
    if [ -n "${ETHERSCAN_API_URL:-}" ]; then
        VERIFY_ARGS+=(--verifier-url "$ETHERSCAN_API_URL")
    fi
    if [ -n "${ETHERSCAN_API_KEY:-}" ]; then
        VERIFY_ARGS+=(--etherscan-api-key "$ETHERSCAN_API_KEY")
    fi
    if [ -z "${ETHERSCAN_API_KEY:-}" ]; then
        log_warn "Verification is enabled but ETHERSCAN_API_KEY is not set. Verification may fail on explorer-backed chains."
    fi
else
    log_info "Verification disabled (VERIFY_MODE=$VERIFY_MODE, CHAIN_ID=$CHAIN_ID)."
fi

if [ "${DEBUG_DEPLOY:-false}" = "true" ]; then
    set -x
fi

deploy_poseidon_library() {
    local deploy_module="$1"
    local target_address=""
    local output=""

    output=$(DEPLOY_MODULE="$deploy_module" RPC_URL="$RPC_URL" PRIVATE_KEY="$PRIVATE_KEY" node <<'NODE'
const { ethers } = require('ethers');

async function main() {
  const deployInfo = require(process.env.DEPLOY_MODULE);
  const provider = new ethers.providers.JsonRpcProvider(process.env.RPC_URL);
  const signer = new ethers.Wallet(process.env.PRIVATE_KEY, provider);

  if ((await provider.getCode(deployInfo.proxyAddress)) === '0x') {
    const fundTx = await signer.sendTransaction({
      to: deployInfo.from,
      value: ethers.BigNumber.from(deployInfo.gas),
    });
    await fundTx.wait();

    const proxyTx = await provider.sendTransaction(deployInfo.tx);
    await proxyTx.wait();
  }

  if ((await provider.getCode(deployInfo.address)) === '0x') {
    const deployTx = await signer.sendTransaction({
      to: deployInfo.proxyAddress,
      data: deployInfo.data,
    });
    await deployTx.wait();
  }

  process.stdout.write(`${deployInfo.address}\n`);
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
NODE
)

    target_address=$(echo "$output" | tail -n1 | tr -d '\r')
    if [ -z "$target_address" ]; then
        log_error "could not deploy $deploy_module"
        exit 1
    fi

    echo "$output" >&2
    echo "$target_address"
}

POSEIDON_T3_MODULE="./lib/poseidon-solidity/deploy/PoseidonT3.js"
POSEIDON_T4_MODULE="./lib/poseidon-solidity/deploy/PoseidonT4.js"
POSEIDON_T3_FQ="lib/poseidon-solidity/contracts/PoseidonT3.sol:PoseidonT3"
POSEIDON_T4_FQ="lib/poseidon-solidity/contracts/PoseidonT4.sol:PoseidonT4"

log_info "Resolving deployable library addresses..."

POSEIDON_T3_ADDRESS=$(deploy_poseidon_library "$POSEIDON_T3_MODULE")
export POSEIDON_T3_ADDRESS
log_info "PoseidonT3 deployed at: $POSEIDON_T3_ADDRESS"

POSEIDON_T4_ADDRESS=$(deploy_poseidon_library "$POSEIDON_T4_MODULE")
export POSEIDON_T4_ADDRESS
log_info "PoseidonT4 deployed at: $POSEIDON_T4_ADDRESS"

log_info "Deploying main contracts with linked libraries..."

forge script script/DeployAll.s.sol:DeployAllScript \
    --chain "$CHAIN_ID" \
    --rpc-url "$RPC_URL" \
    --broadcast \
    --slow \
    --libraries "$POSEIDON_T3_FQ:$POSEIDON_T3_ADDRESS" \
    --libraries "$POSEIDON_T4_FQ:$POSEIDON_T4_ADDRESS" \
    --optimize \
    --optimizer-runs 200 \
    "${VERIFY_ARGS[@]}" \
    -vvvv

"$SCRIPT_DIR/helpers/write_contract_addresses.sh"
