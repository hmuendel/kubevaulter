# kubevaulter-executor
 
Executor executes a list of specified commands while providing

 
 ### Example Config
 ```yaml
executions:
  - command: "echo"
    failNonZero: true #default false
    args:
      - "-v"
      - "-Dpassword={{- .user.password }}"
      - "$NORMAL_ENV"
      - "$SECRET_ENV"
    env:
      SECRET_ENV: "{{ .env.secret }}"
      NORMAL_ENV: "Hello"
  - command: "cat"   
    args: 
      - "/var/templates/template1.tpl"
      

recursivePathList:
  - "/var/scripts"  
   

secretList:
  - name: env
    vaultPath: secret/env
  - name: user
    vaultPath: secret/user
```