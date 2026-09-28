# syntax=docker/dockerfile:1

# Foundry toolchain with the contracts built in. docker-compose.yml runs the
# tests, a deployment and a local anvil chain from it.
#
#   docker build -t davinci-contracts .
#   docker run --rm davinci-contracts "forge test"
#
# The submodules under lib/ must be checked out
# (git submodule update --init --recursive).

ARG FOUNDRY_VERSION=v1.8.3
FROM ghcr.io/foundry-rs/foundry:${FOUNDRY_VERSION}

# jq for deploy_all.sh.
USER root
RUN apt-get update && \
    apt-get install --no-install-recommends -y jq && \
    rm -rf /var/lib/apt/lists/* && \
    install -d -o foundry -g foundry /app
USER foundry
WORKDIR /app

COPY --chown=foundry:foundry . .
# Fetches solc and compiles, so containers start with a warm cache.
RUN forge build

# The base image runs its arguments through /bin/sh -c.
CMD ["forge test"]
