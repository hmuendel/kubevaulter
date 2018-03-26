# kubevaulter-generator
 
 kubevaulter-recursive is meant to be used as a container inside 
 [Kubernetes](https://kubernetes.io/docs/concepts/workloads/pods/init-containers/)
 to create random string an store them as secrets in 
 [Vault](https://www.vaultproject.io/) 

 
 ### Example Config
 ```yaml
generatorList:
  - path: secret/foo/bar
    key: user_password
    override: true  #defaults to false
    length: 32 #defaults to 16
    allowedCharacters: "!@#$%^&*()-_=+,.?/:;{}[]~" #defaults to ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789
```
