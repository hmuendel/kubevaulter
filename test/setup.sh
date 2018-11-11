#!/usr/bin/env bash
MY_PATH="`dirname \"$BASH_SOURCE\"`"              # relative
MY_PATH="`( cd \"$MY_PATH\" && pwd )`"  # absolutized and normalized

if minikube status | grep Running; then
    echo mininkube is running
else
    echo starting minikube
    minikube start --kubernetes-version v1.10.0
    sleep 40
    helm init
fi
kubectl get po --all-namespaces
echo configuring vault
source ${MY_PATH}/setup-vault.sh
export VAULT_ADDR=${VAULT_ADDR}
export VAULT_TOKEN=${VAULT_TOKEN}
kubectl get svc
echo Set env variable setup: ${VAULT_ADDR} ${VAULT_TOKEN}


