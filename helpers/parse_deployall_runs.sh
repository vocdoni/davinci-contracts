#!/usr/bin/env bash
set -euo pipefail

BROADCAST_ROOT="${1:-broadcast/DeployAll.s.sol}"

TMP_FILE="$(mktemp)"

log_stderr(){ echo $@ >&2 ;}

if ! command -v jq >/dev/null 2>&1; then
  log_stderr "Error: jq is required. Install it (apt install jq / brew install jq)."
  exit 1
fi

echo "{}" > "$TMP_FILE"

mapfile -t RUN_FILES < <(find "$BROADCAST_ROOT" -name run-latest.json | sort)

if [[ ${#RUN_FILES[@]} -eq 0 ]]; then
  log_stderr "Error: no run-latest.json found under $BROADCAST_ROOT"
  exit 1
fi

chain_to_network() {
  local chain="$1"
  case "$chain" in
    11155111) echo "sepolia" ;;
    1)        echo "mainnet" ;;
    8453)     echo "base" ;;
    42220)    echo "celo" ;;
    710)      echo "uzh" ;;
    *)        echo "$chain" ;;
  esac
}

for run in "${RUN_FILES[@]}"; do
  CHAIN_ID="$(jq -r '.chain // empty' "$run")"
  if [[ -z "$CHAIN_ID" || "$CHAIN_ID" == "null" ]]; then
    CHAIN_ID="$(basename "$(dirname "$run")")"
  fi

  NETWORK="$(chain_to_network "$CHAIN_ID")"
  log_stderr "Processing $run  (chain=$CHAIN_ID -> net=$NETWORK)"

  jq \
    --arg net "$NETWORK" \
    --slurpfile runfile "$run" \
    '
    # safe lowerCamel: just lower first char, no explode-y regex
    def lower_first:
      tostring
      | if length == 0 then .
        else (.[0:1] | ascii_downcase) + .[1:]
        end;

    def created_map:
      ($runfile[0].transactions // [])
      | (if type=="array" then . else [] end)
      | map(select(.transactionType=="CREATE"))
      | map(
          try {
            key: (.contractName | lower_first),
            addr: (.contractAddress | tostring | ascii_downcase)
          } catch empty
        )
      | reduce .[] as $c ({}; .[$c.key] = $c.addr);

    . as $existing
    | created_map as $new
    | reduce ($new | to_entries[]) as $e ($existing;
        .[$e.key] = (
          (.[$e.key] | if type=="object" then . else {} end)
          | .[$net] = ($e.value | gsub("\\s+";""))
        )
      )
    ' \
    "$TMP_FILE" > "$TMP_FILE.next"

  mv "$TMP_FILE.next" "$TMP_FILE"
done

log_stderr "✅ Wrote $TMP_FILE from ${#RUN_FILES[@]} run-latest.json files."

cat $TMP_FILE
