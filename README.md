# kubevaulter

Kubevaulter are helper tools to handle secrets stored in vault inside 
inside your kubernetes cluster

## Prerequisites

### General Config
```yaml
logging:
  logLevel: "debug" #defaults to
  logFormat: "text"

vault:
  endpointUrl: "http://localhost:8200"
  secretBackend: "secret"
  role: "demo"
  jwtPath: "/Users/hans/Documents/go/src/github.com/hmuendel/kubevaulter/local/default_token"
```

## Available tools
 - [kubevaulter-init](#/cmd/init)