#!/usr/bin/env bash
MY_PATH="`dirname \"$BASH_SOURCE\"`"              # relative
MY_PATH="`( cd \"$MY_PATH\" && pwd )`"  # absolutized and normalized

echo starting minikube
minikube start
echo configuring vault
${MY_PATH}/setup-vault.sh