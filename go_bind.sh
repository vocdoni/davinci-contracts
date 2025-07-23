#!/bin/bash

set -euo pipefail

rm -rf ./golang-types
mkdir -p ./golang-types


sed_in_place() {
    local pattern="$1"
    local file="$2"
    if [[ "$(uname)" == "Darwin" ]]; then
        # macOS
        sed -i '' "$pattern" "$file"
    else
        # Linux
        sed -i "$pattern" "$file"
    fi
}

abi() {
    local json_file="$1"
    local pkg_name="$2"
    local output_file="$3"
    local bin_file="./golang-types/${pkg_name}.bin"

    if [[ ! -f "$json_file" ]]; then
        echo "Error: File '$json_file' does not exist." >&2
        exit 1
    fi

    mkdir -p "$(dirname "$output_file")"

    jq -r '.bytecode' "$json_file" > "$bin_file"

    # Generate the Go bindings directly from the JSON ABI
    jq -r '.abi' "$json_file" | abigen --abi=/dev/stdin --pkg="$pkg_name" --out="$output_file" --bin="$bin_file"
    echo "Successfully generated Go bindings for '$pkg_name' contract."

    rm -f "$bin_file"

    # Replace the package name in the generated file with "contracts"
    sed_in_place "s/^package $pkg_name/package contracts/" "$output_file"
}

abi "./artifacts/src/verifiers/StateTransitionVerifierGroth16.sol/StateTransitionVerifierGroth16.json" \
    "StateTransitionVerifierGroth16" \
    "./golang-types/verifiers/StateTransitionVerifierGroth16.go"

abi "./artifacts/src/verifiers/ResultsVerifierGroth16.sol/ResultsVerifierGroth16.json" \
    "ResultsVerifierGroth16" \
    "./golang-types/verifiers/ResultsVerifierGroth16.go"

abi "./artifacts/src/OrganizationRegistry.sol/OrganizationRegistry.json" \
    "OrganizationRegistry" \
    "./golang-types/OrganizationRegistry.go"

abi "./artifacts/src/ProcessRegistry.sol/ProcessRegistry.json" \
    "ProcessRegistry" \
    "./golang-types/ProcessRegistry.go"
