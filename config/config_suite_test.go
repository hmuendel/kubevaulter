package config_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"path/filepath"
	"io/ioutil"
	"fmt"
	"os"
)


const (
	name = "setup-test"
	version = "0.1.0"
	commit = "00000000"
	realEnvPrefix = "KVTEST"
	dirPrefix = "kubevaulter-test-dir"
	realFilename = "test-config"
	realBrokenFilename = "realBrokenFileName"
	brokenFileContent ="	brokenContent"
)

var (
	realDirName string
	configFileContent string
)


var _ = BeforeSuite(func() {
	configFileContent =`
fileKey1: fileValue1
filekey2: fileValue2
keyMap:
  nestedKey: nestedValue
  nestedMap:
    nestednestedKey: nestednestedValue
`
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
	fb, err := os.Create(filepath.Join(realDirName,realBrokenFilename+".yaml"))
	if err != nil {
		Fail("Could not create temporary broken config file")
	}
	defer func() {
		err := fb.Close()
		if err != nil {
			fmt.Println("close:", err)
		}
	}()
	f.Write([]byte(configFileContent))
	fb.Write([]byte(brokenFileContent))

})

var _ = AfterSuite(func() {
	err := os.RemoveAll(realDirName)
	if err != nil {
		fmt.Println("remove:", err)
	}
})

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}
