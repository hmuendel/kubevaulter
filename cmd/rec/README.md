# kubevaulter-recursive
 
 kubevaulter-recursive is meant to be used as an init container inside 
 [Kubernetes](https://kubernetes.io/docs/concepts/workloads/pods/init-containers/)
 to handle secrets from [Vault](https://www.vaultproject.io/) 
 before startup of the application. It steps recursively through a directory
 replacing rendering all templates in place 

 
 ### Example Config
 ```yaml
templateMountPath : /share
secretList:
- name: demo
  vaultPath: demo/secret
```
