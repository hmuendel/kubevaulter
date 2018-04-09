# kubevaulter-executor
 
Executor executes a list of specified commands while providing

 
 ### Example Config
 ```yaml
executions:
  - command: "echo"
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
      
templateList:
  - "/var/templates/template1.tpl" 
  
recursivePathList:
  - "/var/scripts"  
   
secretList:
  - name: demo
    vaultPath: demo/secret
```
