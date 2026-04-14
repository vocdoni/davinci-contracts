#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd "$SCRIPT_DIR"

log_info() { echo "[deploy_all_all_chains] $*"; }
log_warn() { echo "[deploy_all_all_chains] Warning: $*" >&2; }
log_error() { echo "[deploy_all_all_chains] Error: $*" >&2; }

require_cmd() {
    local cmd="$1"
    if ! command -v "$cmd" >/dev/null 2>&1; then
        log_error "$cmd is required but not installed."
        exit 1
    fi
}

require_cmd bash

SHARED_ENV_FILE="${SHARED_ENV_FILE:-$SCRIPT_DIR/.env}"
CHAIN_ENV_DIR="${CHAIN_ENV_DIR:-$SCRIPT_DIR}"
DEPLOY_SINGLE_CHAIN_CMD="${DEPLOY_SINGLE_CHAIN_CMD:-$SCRIPT_DIR/deploy_all.sh}"

if [ ! -f "$SHARED_ENV_FILE" ]; then
    log_error "Shared env file not found: $SHARED_ENV_FILE"
    exit 1
fi

if [ ! -x "$DEPLOY_SINGLE_CHAIN_CMD" ]; then
    log_error "Single-chain deploy command is not executable: $DEPLOY_SINGLE_CHAIN_CMD"
    exit 1
fi

chain_env_file_for() {
    local chain="$1"
    local dot_file="$CHAIN_ENV_DIR/.env.$chain"
    local dash_file="$CHAIN_ENV_DIR/.env-$chain"

    if [ -f "$dot_file" ]; then
        echo "$dot_file"
        return 0
    fi

    if [ -f "$dash_file" ]; then
        echo "$dash_file"
        return 0
    fi

    return 1
}

unset_chain_specific_vars() {
    unset CHAIN_ID || true
    unset RPC_URL || true
    unset ACTIVATE_BLOBS || true
    unset ETHERSCAN_API_URL || true
    unset POSEIDON_T3_ADDRESS || true
    unset POSEIDON_T4_ADDRESS || true
    unset STATE_ROOT_LIB_ADDRESS || true
    unset PROCESS_ID_LIB_ADDRESS || true
    unset BLOBS_LIB_ADDRESS || true
}

source_env_file() {
    local env_file="$1"
    set -a
    # shellcheck disable=SC1090
    . "$env_file"
    set +a
}

load_chain_env() {
    local chain="$1"
    local chain_env_file=""

    source_env_file "$SHARED_ENV_FILE"
    unset_chain_specific_vars

    if ! chain_env_file=$(chain_env_file_for "$chain"); then
        log_error "No env file found for chain '$chain'. Expected $CHAIN_ENV_DIR/.env.$chain or $CHAIN_ENV_DIR/.env-$chain"
        exit 1
    fi

    source_env_file "$chain_env_file"
}

validate_loaded_env() {
    : "${PRIVATE_KEY:?PRIVATE_KEY is not set after loading env files}"
    : "${CHAIN_ID:?CHAIN_ID is not set after loading env files}"
    : "${RPC_URL:?RPC_URL is not set after loading env files}"
}

parse_deploy_chains() {
    local raw="${DEPLOY_CHAINS:-}"
    local normalized=""

    if [ -z "$raw" ]; then
        log_error "DEPLOY_CHAINS is not set in $SHARED_ENV_FILE"
        exit 1
    fi

    normalized="${raw//,/ }"
    read -r -a DEPLOY_CHAIN_LIST <<< "$normalized"

    if [ "${#DEPLOY_CHAIN_LIST[@]}" -eq 0 ]; then
        log_error "DEPLOY_CHAINS is empty in $SHARED_ENV_FILE"
        exit 1
    fi
}

source_env_file "$SHARED_ENV_FILE"
parse_deploy_chains

for chain in "${DEPLOY_CHAIN_LIST[@]}"; do
    log_info "Preparing deployment for chain '$chain'"
    load_chain_env "$chain"
    validate_loaded_env
    log_info "Deploying to chain '$chain' (CHAIN_ID=$CHAIN_ID, RPC_URL=$RPC_URL)"
    "$DEPLOY_SINGLE_CHAIN_CMD"
done

log_info "Finished deployments for: ${DEPLOY_CHAIN_LIST[*]}"
