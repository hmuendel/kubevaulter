#!/usr/bin/env bash
MY_PATH="`dirname \"$BASH_SOURCE\"`"              # relative
MY_PATH="`( cd \"$MY_PATH\" && pwd )`"  # absolutized and normalized

source ${MY_PATH}/setup.sh
echo Set env variable global: ${VAULT_ADDR} ${VAULT_TOKEN}
export VAULT_ADDR=${VAULT_ADDR}
export VAULT_TOKEN=${VAULT_TOKEN}
${MY_PATH}/gen/run-tests.sh
${MY_PATH}/exec/run-tests.sh

${MY_PATH}/cleanup.sh