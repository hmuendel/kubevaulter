#!/usr/bin/env bash
MY_PATH="`dirname \"$BASH_SOURCE\"`"              # relative
MY_PATH="`( cd \"$MY_PATH\" && pwd )`"  # absolutized and normalized


echo Set env variable gen tests: ${VAULT_ADDR} ${VAULT_TOKEN}
export VAULT_ADDR=${VAULT_ADDR}
export VAULT_TOKEN=${VAULT_TOKEN}

echo creating configmap for gen tests
kubectl apply -f  ${MY_PATH}/config.yaml
echo creating gen test pod
kubectl apply -f ${MY_PATH}/pod.yaml
kubectl get po
sleep 15
kubectl get cm kubevaulter-gen -o yaml
kubectl logs -f kubevaulter-gen

vault list -ca-cert=$(dirname ${MY_PATH})/certs/root.pem secret
vault kv list -ca-cert=$(dirname ${MY_PATH})/certs/root.pem secret
vault kv get -ca-cert=$(dirname ${MY_PATH})/certs/root.pem secret/write/foo
vault kv get -ca-cert=$(dirname ${MY_PATH})/certs/root.pem secret/write/bar

echo deleting gen test pod
kubectl delete -f ${MY_PATH}/config.yaml
echo deleting configmap for gen tests
kubectl delete -f ${MY_PATH}/pod.yaml --now

