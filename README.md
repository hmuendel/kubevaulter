# kubevaulter

Kubevaulter are helper tools to handle secrets stored in vault
inside your kubernetes cluster. To authenticate against vault 
the kubernetes service account token is used, which is mounted 
to every pod by default.

## Available tools
 - [kubevaulter-init](/tree/master/cmd/init) an init container
 to write vault secret into templates in the pod filesystem

## Prerequisites

 To use these tools vault must be correctly configured and  
reachable from within the kubernetes cluster.
 Kubernetes must support 
 [RBAC](https://kubernetes.io/docs/admin/authorization/rbac/)
 

 
 ### Kubernetes with RBAC

 and the api-server must be started with the flags 
 `--authorization-mode=RBAC` and
 ` --service-account-lookup`. 

 Also the correct service accounts and clusterRoles and 
 RoleBindings must exist
 
 #### Kubernetes ClusterRole
 
 This ClusterRole should exist by default, if not it has to
 be created.
 
  ```yaml
  - apiVersion: rbac.authorization.k8s.io/v1
    kind: ClusterRole
    metadata:
      annotations:
        rbac.authorization.kubernetes.io/autoupdate: "true"
      creationTimestamp: null
      labels:
        kubernetes.io/bootstrapping: rbac-defaults
      name: system:auth-delegator
    rules:
    - apiGroups:
      - authentication.k8s.io
      resources:
      - tokenreviews
      verbs:
      - create
    - apiGroups:
      - authorization.k8s.io
      resources:
      - subjectaccessreviews
      verbs:
      - create
 ```
 
#### Create Service Accounts
  
One service account need to exist with which vault authenticates
against the kuberentes api. 

 ```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: vault-auth
```

For each pod that accesses secrets in vault, a service account
should exist to authenticate against vault. This can also be
the default namespace service account, depending on the needs.

In this example the



#### Create Role Binding 
 ```yaml
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: role-tokenreview-binding
  namespace: default
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:auth-delegator
subjects:
- kind: ServiceAccount
  name: vault-auth
  namespace: default
``` 


 ### Vault
 To use authentication via service account tokens, the 
  [Kubernetes Auth Backend](https://www.vaultproject.io/docs/auth/kubernetes.html)
  must be enabled and configured correctly. Also policies must
  exist to manage access of the role to vault secret paths
 
 #### Enable Kubernetes Auth in Vault
 ```bash
 $ vault auth-enable kubernetes
 Successfully enabled 'kubernetes' at 'kubernetes'!
 ```
 #### Create Role
 ```bash
 vault write auth/kubernetes/role/DEMO \
     bound_service_account_names=DEMO-SA\
     bound_service_account_namespaces=DEMO-NS \
     policies=DEMO-POLICY \
     ttl=1h
 ```
 #### Example Policy
 
 ```hcl
 path "secret/*" {
   capabilities = ["create"]
 }
 
 path "secret/foo" {
   capabilities = ["read"]
 }
 
 path "auth/token/lookup-self" {
   capabilities = ["read"]
 }
```


 
 

 

### General Kubevaulter Config

To configure Kubevaulter, a config file called config 
in yaml, toml or json format
must exist in . or ./config folder of the containers.
The location could be overwritten by specifying the 
environment variable KV_
the general config looks like this


```yaml
logging:
  logLevel: "debug" #defaults to "info"
  logFormat: "json" #default to "text"
  
vault:
  endpointUrl: "http://localhost:8200"
  secretBackend: "demo-secret" # defaults to "secret" 
  role: "DEMO"
  jwtPath:  "/run/secrets/namespace/token"  # defaults to "/run/secrets/namespace/token" 
```

