# kubevaulter-init
 
 kubevaulter-init is meant to be used as an init container inside 
 [Kubernetes](https://kubernetes.io/docs/concepts/workloads/pods/init-containers/)
 to handle secrets from [Vault](https://www.vaultproject.io/) 
 before startup of the application.
 
 
 ### Example Config
 ```yaml
fileSecretList:
- templatePath: "/Users/hans/Documents/go/src/github.com/hmuendel/kubevaulter/test/tpl1"
  targetPath: "/Users/hans/Documents/go/src/github.com/hmuendel/kubevaulter/test/t1"
  secretPath: "demo/s1"
```
