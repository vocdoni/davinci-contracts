#!/bin/bash

set -euo pipefail

abi() {
    local json_file="$1"
    local pkg_name="$2"
    local output_file="$3"

    if [[ ! -f "$json_file" ]]; then
        echo "Error: File '$json_file' does not exist." >&2
        exit 1
    fi

    local temp_abi_file
    temp_abi_file=$(mktemp)

    jq -r '.abi' "$json_file" > "$temp_abi_file"
    if [[ ! -s "$temp_abi_file" ]]; then
        echo "Error: ABI extraction failed." >&2
        rm -f "$temp_abi_file"
        exit 1
    fi

    abigen --abi="$temp_abi_file" --pkg="$pkg_name" --out="$output_file"

    rm -f "$temp_abi_file"
}

abi "./artifacts/src/OrganizationRegistry.sol/OrganizationRegistry.json" \
    "OrganizationRegistry" \
    "./golang-types/OrganizationRegistry.go"

abi "./artifacts/src/ProposalRegistry.sol/ProposalRegistry.json" \
    "ProposalRegistry" \
    "./golang-types/ProposalRegistry.go"

# abi "./artifacts/src/SequencerRegistry.sol/SequencerRegistry.json" \
#     "SequencerRegistry" \
#     "./golang-types/SequencerRegistry.go"

abi "./artifacts/src/ZKRegistry.sol/ZKRegistry.json" \
    "ZKRegistry" \
    "./golang-types/ZKRegistry.go"
