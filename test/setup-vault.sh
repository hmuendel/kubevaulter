#!/usr/bin/env bash -x
MY_PATH="`dirname \"$BASH_SOURCE\"`"              # relative
MY_PATH="`( cd \"$MY_PATH\" && pwd )`"  # absolutized and normalized

name=kubevaulter-test
echo installing vault chart
helm install  --wait --name ${name} ${MY_PATH}/vault

echo Creating role, rolebinding and service account in k8s
kubectl apply -f ${MY_PATH}/vault-auth.yaml
kubectl create configmap ${name} --from-file=${MY_PATH}/certs
echo waiting for pod to become ready
sleep 5
echo continueing..
echo getting entities from kubernetes
pod=$(kubectl get po -l release=${name} -l app=vault -o go-template="{{ (index .items 0 ).metadata.name }}")
port=$(kubectl get svc  -l release=${name} -l app=vault -o go-template='{{ index (index (index .items 0 ).spec.ports 0) "nodePort" }}')
ip=$(minikube ip)
secret=$(kubectl get secrets -o go-template --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}' | grep vault-auth)
ca=$(kubectl get secret ${secret} -o go-template --template '{{ index .data "ca.crt" }}' | base64 -D)
jwt=$(kubectl get secret ${secret} -o go-template --template '{{ .data.token }}' | base64 -D)

#echo vault-root-token: ${token}
echo url: ${ip}:${port} pod: ${pod} secret: ${secret}
echo CA-Cert: ${ca}
echo kube-auth-service-account-token: ${jwt}

echo Writing CA cert to disk
#echo ${ca} > ca.crt

echo setting vault env variables
export VAULT_ADDR=https://${ip}:${port}

echo initializing vault
for i in {1..100}; do
   vault operator init -key-shares=1 -key-threshold=1 -ca-cert=${MY_PATH}/certs/root.pem | tee init.log
   token=$(cat init.log | ack -o '(?<=Initial Root Token:\s)[a-zA-Z0-9-]+')
   unseal=$(cat init.log | ack -o '(?<=Unseal Key 1:\s)[a-zA-Z0-9-\=\+\/]+')
   ret=$?
   if [ $ret -eq 0 ]; then
    echo return code was $ret, initialization sucessful
    break
   fi
   echo return code was $ret, retrying...
   sleep 1
done





export VAULT_TOKEN=${token}
echo unseal key: $unseal token: $token

echo unsealing vault
vault operator unseal -ca-cert=${MY_PATH}/certs/root.pem $unseal
vault login -ca-cert=${MY_PATH}/certs/root.pem  $token
echo configuring vault for k8s authentication
vault auth enable -ca-cert=${MY_PATH}/certs/root.pem kubernetes
vault write -ca-cert=${MY_PATH}/certs/root.pem auth/kubernetes/config \
    token_reviewer_jwt="${jwt}" \
    kubernetes_host=https://kubernetes:443 \
    kubernetes_ca_cert=@${MY_PATH}/ca.crt

echo writing vault policy
vault policy write -ca-cert=${MY_PATH}/certs/root.pem default ${MY_PATH}/default-policy.hcl

echo writing kubernetes auth role
vault write -ca-cert=${MY_PATH}/certs/root.pem auth/kubernetes/role/default @${MY_PATH}/default-role.json


echo cleaning up ca.crt
#rm ca.crt
