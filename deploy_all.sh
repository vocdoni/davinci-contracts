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
require_cmd cast
require_cmd jq
require_cmd python3

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

# PINS_FROM_REGISTRY inherits the verifier, pins and grace settings of a live registry;
# without it the pins are required here.
VERIFY_DEPLOYMENT_ARGS=()
if [ -n "${PINS_FROM_REGISTRY:-}" ]; then
    export PINS_FROM_REGISTRY
    VERIFY_DEPLOYMENT_ARGS+=(--pins-from "$PINS_FROM_REGISTRY")
else
    : "${BATCH_PROGRAM_VK:?BATCH_PROGRAM_VK is not set (env or .env)}"
    : "${RESULTS_PROGRAM_VK:?RESULTS_PROGRAM_VK is not set (env or .env)}"
    : "${ROOT_C_VADCOP_FINAL:?ROOT_C_VADCOP_FINAL is not set (env or .env)}"
    : "${BALLOT_VK_HASH:?BALLOT_VK_HASH is not set (env or .env)}"
    export BATCH_PROGRAM_VK RESULTS_PROGRAM_VK ROOT_C_VADCOP_FINAL BALLOT_VK_HASH
    VERIFY_DEPLOYMENT_ARGS+=(--batch-vk "$BATCH_PROGRAM_VK" --results-vk "$RESULTS_PROGRAM_VK"
        --root-c "$ROOT_C_VADCOP_FINAL" --ballot-vk-hash "$BALLOT_VK_HASH")
    if [ -n "${ZISK_VERIFIER:-}" ]; then
        VERIFY_DEPLOYMENT_ARGS+=(--verifier "$ZISK_VERIFIER")
    fi
fi
# The managers the adapters must point at; unset or empty expects the mode disabled.
VERIFY_DEPLOYMENT_ARGS+=(--dkg-manager "${DKG_MANAGER:-0x0}" --council-manager "${COUNCIL_MANAGER:-0x0}")

# Deploys ZiskVerifier, then ProcessRegistry (and the DKG adapter when DKG_MANAGER is set, the
# Council adapter when COUNCIL_MANAGER is set).
log_info "Deploying main contracts..."

forge script script/DeployAll.s.sol:DeployAllScript \
    --chain "$CHAIN_ID" \
    --rpc-url "$RPC_URL" \
    --broadcast \
    --slow \
    --optimize \
    --optimizer-runs 200 \
    "${VERIFY_ARGS[@]}" \
    -vvvv

# Checks the deployment against this build and the inputs, then writes its record.
RUN_FILE="broadcast/DeployAll.s.sol/$CHAIN_ID/run-latest.json"
REGISTRY="$(jq -r '[.transactions[] | select(.transactionType == "CREATE" and .contractName == "ProcessRegistry")][0].contractAddress' "$RUN_FILE")"
RECORD="deployments/$CHAIN_ID.json"
log_info "Verifying ProcessRegistry $REGISTRY and writing $RECORD..."
python3 script/verify_deployment.py --rpc "$RPC_URL" --chain-id "$CHAIN_ID" --registry "$REGISTRY" \
    "${VERIFY_DEPLOYMENT_ARGS[@]}" --cast "$(command -v cast)" --broadcast "$RUN_FILE" --record "$RECORD"

"$SCRIPT_DIR/helpers/write_contract_addresses.sh"
