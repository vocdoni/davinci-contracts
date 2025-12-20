# Makefile for Davinci Contracts

# ==============================================================================
# Environment Setup
# ==============================================================================

# Load .env file if it exists
ifneq (,$(wildcard .env))
    include .env
    export
endif

# Defaults
ACTIVATE_BLOBS ?= false
ETHERSCAN_API_KEY ?= ""

# Check required commands
REQUIRED_BINS := forge npm git jq
$(foreach bin,$(REQUIRED_BINS),	$(if $(shell command -v $(bin) 2> /dev/null),,$(error Please install `$(bin)`)))

# Configure Verification Flag
VERIFY_FLAG :=
ifneq ($(ETHERSCAN_API_KEY),)
    VERIFY_FLAG := --verify --etherscan-api-key $(ETHERSCAN_API_KEY)
endif

# ==============================================================================
# Development
# ==============================================================================

.PHONY: install build clean test test-v lint go-bind

install:
	nnpm install
	forge install

build:
	forge build

clean:
	forge clean
	rm -rf cache out broadcast artifacts typechain-types golang-types/addresses.go

test:
	forge test

test-v:
	forge test -vvv

gas-report:
	forge test --gas-report

lint:
	npm run lint:sol
	npm run prettier

go-bind:
	./go_bind.sh

# ==============================================================================
# Deployment
# ==============================================================================

.PHONY: deploy deploy-local update-addresses

# Generic deploy target using environment variables
deploy:
	@echo "----------------------------------------------------------------"
	@echo "Deploying Contracts..."
	@if [ -z "$(RPC_URL)" ]; then echo "Error: RPC_URL is not set"; exit 1; fi
	@if [ -z "$(PRIVATE_KEY)" ]; then echo "Error: PRIVATE_KEY is not set"; exit 1; fi
	@if [ -z "$(CHAIN_ID)" ]; then echo "Error: CHAIN_ID is not set"; exit 1; fi
	@echo "Network RPC: $(RPC_URL)"
	@echo "Chain ID:    $(CHAIN_ID)"
	@echo "Blobs DA:    $(ACTIVATE_BLOBS)"
	@echo "Verifying:   $(if $(VERIFY_FLAG),Yes,No)"
	@echo "----------------------------------------------------------------"
	@CHAIN_ID=$(CHAIN_ID) ACTIVATE_BLOBS=$(ACTIVATE_BLOBS) forge script script/DeployAll.s.sol \
		--rpc-url $(RPC_URL) \
		--broadcast \
		--private-key $(PRIVATE_KEY) \
		$(VERIFY_FLAG)

# Short-cut for local deployment (overrides environment variables)
deploy-local:
	@$(MAKE) deploy \
		RPC_URL=http://localhost:8545 \
		CHAIN_ID=1337 \
		ACTIVATE_BLOBS=false \
		PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
		ETHERSCAN_API_KEY=""

# Updates Go address constants from broadcast artifacts
update-addresses:
	./helpers/write_contract_addresses.sh

# ==============================================================================
# Workflows
# ==============================================================================

.PHONY: all all-local

# Generic full workflow: Build -> Deploy (using .env) -> Update Go Addresses
all: build deploy update-addresses
	@echo "Deployment workflow complete!"

# Local full workflow
all-local: build deploy-local update-addresses
	@echo "Local deployment workflow complete!"