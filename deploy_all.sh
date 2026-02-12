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

if ! command -v sed >/dev/null 2>&1; then
    log_error "sed is required but not installed."
    exit 1
fi

deploy_library() {
    local lib_path_name="$1"
    shift
    local extra_args=("$@")
    local output=""
    local address=""

    output=$(FOUNDRY_VIA_IR=false forge create "$lib_path_name" \
        --chain "$CHAIN_ID" \
        --rpc-url "$RPC_URL" \
        --private-key "$PRIVATE_KEY" \
        --broadcast \
        "${VERIFY_ARGS[@]}" \
        "${extra_args[@]}" \
        --optimize \
        --optimizer-runs 200 \
        -vvvv 2>&1)

    echo "$output" >&2

    address=$(echo "$output" | sed -n 's/.*Deployed to: \(0x[a-fA-F0-9]\{40\}\).*/\1/p' | tail -n1)
    if [ -z "$address" ]; then
        log_error "could not parse deployed address for $lib_path_name"
        exit 1
    fi

    echo "$address"
}

is_deployed_contract() {
    local address="${1:-}"
    local code=""

    if [[ ! "$address" =~ ^0x[a-fA-F0-9]{40}$ ]]; then
        return 1
    fi

    if ! code=$(cast code --rpc-url "$RPC_URL" "$address" 2>/dev/null); then
        return 1
    fi

    code="${code//[$'\t\r\n ']}"
    case "$code" in
        0x|0x0|0x00|"")
            return 1
            ;;
    esac

    return 0
}

resolve_or_deploy_library() {
    local var_name="$1"
    local library_name="$2"
    local lib_path_name="$3"
    shift 3
    local extra_args=("$@")
    local configured_address="${!var_name:-}"
    local resolved_address=""

    if [ -n "$configured_address" ]; then
        if is_deployed_contract "$configured_address"; then
            resolved_address="$configured_address"
            log_info "Reusing $library_name from $var_name: $resolved_address"
        else
            log_warn "$var_name is set to $configured_address but no bytecode was found. Deploying a new $library_name."
        fi
    else
        log_info "$var_name not set. Deploying $library_name."
    fi

    if [ -z "$resolved_address" ]; then
        resolved_address=$(deploy_library "$lib_path_name" "${extra_args[@]}")
        log_info "$library_name deployed at: $resolved_address"
    fi

    printf -v "$var_name" "%s" "$resolved_address"
    export "$var_name"

    log_info "Resolved $library_name address: $resolved_address"
    echo "export $var_name=$resolved_address"
}

POSEIDON_T3_FQ="lib/poseidon-solidity/contracts/PoseidonT3.sol:PoseidonT3"
POSEIDON_T4_FQ="lib/poseidon-solidity/contracts/PoseidonT4.sol:PoseidonT4"
STATE_ROOT_LIB_FQ="src/libraries/StateRootLib.sol:StateRootLib"
PROCESS_ID_LIB_FQ="src/libraries/ProcessIdLib.sol:ProcessIdLib"
BLOBS_LIB_FQ="src/libraries/BlobsLib.sol:BlobsLib"

log_info "Resolving deployable library addresses..."

resolve_or_deploy_library "POSEIDON_T3_ADDRESS" "PoseidonT3" "$POSEIDON_T3_FQ"
resolve_or_deploy_library "POSEIDON_T4_ADDRESS" "PoseidonT4" "$POSEIDON_T4_FQ"

state_root_create_args=(
    --libraries "$POSEIDON_T3_FQ:$POSEIDON_T3_ADDRESS"
    --libraries "$POSEIDON_T4_FQ:$POSEIDON_T4_ADDRESS"
)
resolve_or_deploy_library "STATE_ROOT_LIB_ADDRESS" "StateRootLib" "$STATE_ROOT_LIB_FQ" "${state_root_create_args[@]}"

resolve_or_deploy_library "PROCESS_ID_LIB_ADDRESS" "ProcessIdLib" "$PROCESS_ID_LIB_FQ"
resolve_or_deploy_library "BLOBS_LIB_ADDRESS" "BlobsLib" "$BLOBS_LIB_FQ"

log_info "Deploying main contracts with linked libraries..."

forge script script/DeployAll.s.sol:DeployAllScript \
    --chain "$CHAIN_ID" \
    --rpc-url "$RPC_URL" \
    --broadcast \
    --slow \
    --libraries "$POSEIDON_T3_FQ:$POSEIDON_T3_ADDRESS" \
    --libraries "$POSEIDON_T4_FQ:$POSEIDON_T4_ADDRESS" \
    --libraries "$STATE_ROOT_LIB_FQ:$STATE_ROOT_LIB_ADDRESS" \
    --libraries "$PROCESS_ID_LIB_FQ:$PROCESS_ID_LIB_ADDRESS" \
    --libraries "$BLOBS_LIB_FQ:$BLOBS_LIB_ADDRESS" \
    --optimize \
    --optimizer-runs 200 \
    "${VERIFY_ARGS[@]}" \
    -vvvv

"$SCRIPT_DIR/helpers/write_contract_addresses.sh"
