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
