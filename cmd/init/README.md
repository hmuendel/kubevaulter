# kubevaulter-init
 
 kubevaulter-init is meant to be used as an init container inside 
 [Kubernetes](https://kubernetes.io/docs/concepts/workloads/pods/init-containers/)
 to handle secrets from [Vault](https://www.vaultproject.io/) 
 before startup of the application.
 
 ## Authentication
 To authenticate against vault, kubevaulter-init uses the Kubernetes 
 [service account](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/)
 token mounted by default into each pod by Kubernetes 
 automatically. The token should reside under
 _/var/run/secrets/kubernetes.io/serviceaccount/token_ and is 
 signed by the Kubernetes signing CA
 
 ## Setup
 To use this init container vault must be configured
 with the 
 [Kubernetes Auth Backend](https://www.vaultproject.io/docs/auth/kubernetes.html)
 and Kubernetes must support 
 [RBAC]()
 and has a service account for vault to allow the token lookup
 
 
 ### Vault
 In vault the 
 #### Enable Kubernetes Auth in Vault
 ```bash
 $ vault auth-enable kubernetes
 Successfully enabled 'kubernetes' at 'kubernetes'!
 ```
 #### Create Role
 ```bash
 vault write auth/kubernetes/role/demo \
     bound_service_account_names=vault-auth \
     bound_service_account_namespaces=default \
     policies=default \
     ttl=1h
 ```
 ### Kubernetes
 
 --service-account-lookup
 #### Kubernetes RBAC
 
 #### Create Service Account
 
 #### Create Role Binding 

 ## Templating
 
 ### Template Path
 
 ### Secret Name
 
 ### Place Holder
 
 ### Example Config
 ```yaml
fileSecretList:
- templatePath: "/Users/hans/Documents/go/src/github.com/hmuendel/kubevaulter/test/tpl1"
  targetPath: "/Users/hans/Documents/go/src/github.com/hmuendel/kubevaulter/test/t1"
  secretPath: "demo/s1"
```
