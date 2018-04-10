package config_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	"io/ioutil"
	"os"
	"path/filepath"
)


var _ = BeforeSuite(func() {
	configFileContent =`executions:
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
      

recursivePathList:
  - "/var/scripts"  
   
secretList:
  - name: env
    vaultPath: secret/env
  - name: user
    vaultPath: secret/user`

	var err error
	realDirName, err = ioutil.TempDir("", dirPrefix)
	if err != nil {
		Fail("Could not create temp dir")
	}
	f, err := os.Create(filepath.Join(realDirName,realFilename+".yaml"))
	if err != nil {
		Fail("Could not create temporary config file")
	}
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Println("close:", err)
		}
	} ()
	f.Write([]byte(configFileContent))

})

var _ = AfterSuite(func() {
	err := os.RemoveAll(realDirName)
	if err != nil {
		fmt.Println("remove:", err)
	}
})

var _ = Describe("Executions", func() {

})
