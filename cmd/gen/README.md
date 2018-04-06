# kubevaulter-generator
 
 kubevaulter-generator is meant to be used as a container inside 
 [Kubernetes](https://kubernetes.io/docs/concepts/workloads/pods/init-containers/)
 to create random string an store them as secrets in 
 [Vault](https://www.vaultproject.io/) 

 
 ### Example Config
 ```yaml
targetList:
   - path: "secret/foo"
     data:
       keyName1: 
        ref: userPassword
        transform: # remove if cleartext wanted or one of
#       MD4,MD5,SHA1,SHA224,SHA256,SHA384,SHA512,
#       MD5SHA1,RIPEMD160,SHA3_224,SHA3_256,SHA3_384,
#       SHA3_512,SHA512_224,SHA512_256,BLAKE2s_256,
#       BLAKE2b_256,BLAKE2b_384,BLAKE2b_512
       keyName2: dbPassword
   - path: "secret/bar"
     data: 
       DB_PW: 
        ref: dbPassword
       DB_USER:
         lit: admin
       


randomStrings:
  userPassword: 
    override: true  #defaults to false, there might be the case, that
    allowedCharacters: "!@#$%^&*()-_=+,.?/:;{}[]~" #defaults to ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789
  dbPassword:
    length: 32 #defaults to 16
```
