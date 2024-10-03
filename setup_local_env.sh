#!/bin/bash

# REQUIREMENTS:
# 1. Foundry

set -e

ANVIL_LOG="anvil_output.log"

cleanup() {
  echo "Stopping anvil (PID: $ANVIL_PID)"
  kill $ANVIL_PID
  exit 0
}

trap cleanup SIGINT

if pgrep -x "anvil" > /dev/null; then
  echo "Anvil is already running."
else
  echo "Starting anvil in the background. Logging output to $ANVIL_LOG"
  anvil > "$ANVIL_LOG" 2>&1 &
  ANVIL_PID=$!
  echo "Anvil started with PID: $ANVIL_PID"

  sleep 2

  if ! ps -p $ANVIL_PID > /dev/null; then
    echo "Error: anvil failed to start"
    exit 1
  fi
fi

echo "Building contracts"
forge clean && forge build

echo "Deploying contracts"
forge script script/TEST_DeployAll.s.sol:TestDeployAllScript --fork-url http://localhost:8545 --broadcast

echo "Tailing anvil log. Press Ctrl+C to stop and shut down anvil."
tail -f "$ANVIL_LOG"
