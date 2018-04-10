#!/usr/bin/env bash
MY_PATH="`dirname \"$BASH_SOURCE\"`"              # relative
MY_PATH="`( cd \"$MY_PATH\" && pwd )`"  # absolutized and normalized

name=kubevaulter-test

echo Adding helm repo for vault chart
helm repo add incubator https://storage.googleapis.com/kubernetes-charts-incubator
helm install --name ${name} -f vault-values.yaml incubator/vault

echo Creating role, rolebinding and service account in k8s
kubectl apply -f vault-auth.yaml

echo getting entities from kubernetes
pod=$(kubectl get po -l release=${name} -o go-template="{{ (index .items 0 ).metadata.name }}")
port=$(kubectl get svc  -l release=${name} -o go-template='{{ index (index (index .items 0 ).spec.ports 0) "nodePort" }}')
ip=$(minikube ip)
secret=$(kubectl get secrets -o go-template --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}' | grep vault-auth)
ca=$(kubectl get secret ${secret} -o go-template --template '{{ index .data "ca.crt" }}' | base64 -D)
jwt=$(kubectl get secret ${secret} -o go-template --template '{{ .data.token }}' | base64 -D)
token=$(kubectl logs ${pod} | ack -o '(?<=Root Token:\s)[a-zA-Z0-9-]+')
echo vault-root-token: ${token}
echo url: ${ip}:${port} pod: ${pod} secret: ${secret}
echo CA-Cert: ${ca}
echo kube-auth-service-account-token: ${jwt}

echo Writing CA cert to disk
echo ${ca} > ca.crt

echo setting vault env variables

export VAULT_TOKEN=${token}
export VAULT_ADDR=http://${ip}:${port}

echo configuring vault for k8s authentication
vault auth enable kubernetes || true
vault write auth/kubernetes/config \
    token_reviewer_jwt="${jwt}" \
    kubernetes_host=https://${ip}:8443 \
    kubernetes_ca_cert=@ca.crt

echo writing vault policy
vault policy write default ${MY_PATH}/default-policy.hcl

echo writing kubernetes auth role
vault write auth/kubernetes/role/default @${MY_PATH}/default-role.json


echo cleaning up ca.crt
rm ca.crt
