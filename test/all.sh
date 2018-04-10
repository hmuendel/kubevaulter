#!/usr/bin/env bash
MY_PATH="`dirname \"$BASH_SOURCE\"`"              # relative
MY_PATH="`( cd \"$MY_PATH\" && pwd )`"  # absolutized and normalized

${MY_PATH}/setup.sh
${MY_PATH}/gen/run-tests.sh

${MY_PATH}/cleanup.sh