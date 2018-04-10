package config_test

import (
	"fmt"
	"github.com/hmuendel/kubevaulter/config"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
	"strings"

	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
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
	configFileContent =`fileKey1: fileValue1
filekey2: fileValue2
keyMap:
  nestedKey: nestedValue
  nestedMap:
    nestednestedKey: nestednestedValue`
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

var _ = Describe("Setup", func() {
	var (
		setEnvPrefix string
		defaults          map[string]interface{}
		setDirName string
		setFileName string
		testStringKey = "string"
		testStringValue = "test-string"
		testIntKey = "int"
		testIntValue = 42
		testBoolKey = "bool"
		testBoolValue = "true"
		envParamName = "envparam"
	)

	BeforeEach(func() {
		setEnvPrefix = realEnvPrefix
		setDirName = realDirName
		setFileName = realFilename
		defaults = make(map[string]interface{})
		defaults[testStringKey] = testStringValue
		defaults[testIntKey] = testIntValue
		defaults[testBoolKey] = testBoolValue
	})

	JustBeforeEach(func() {
		viper.Reset()
		os.Setenv(setEnvPrefix + "_CONFIG_PATH", setDirName)
		os.Setenv(setEnvPrefix + "_CONFIG_NAME", setFileName)
		os.Setenv(setEnvPrefix + "_" + strings.ToUpper(envParamName), testStringValue)
	})



	Describe("Getting config from the environment", func() {
		Context("with prefix set correctly", func() {
			It("should read the config location correctly from env", func() {
				config.Setup(name, version, commit, setEnvPrefix,defaults)
				Expect(viper.GetString("configName")).To(Equal(setFileName))
				Expect(viper.GetString("configPath")).To(Equal(setDirName))
			})
			It("should read a value correctly from env", func() {
				config.Setup(name, version, commit, setEnvPrefix,defaults)
				Expect(viper.GetString(envParamName)).To(Equal(testStringValue))
			})
		})
		Context("with prefix set incorrectly ", func() {
			It("the missing config location should cause a panic", func() {
				setEnvPrefix = "WrongPrefix"
				fn := func() {
					config.Setup(name, version, commit, setEnvPrefix,defaults)
				}
				Expect(fn).To(Panic())
			})
		})


	Describe("Reading values from config file", func() {
		Context("with parsable config", func() {
			It("should result in the correct value", func() {
				config.Setup(name, version, commit, setEnvPrefix, defaults)
				Expect(viper.GetString("fileKey1")).To(Equal("fileValue1"))
				Expect(viper.GetString("fileKey2")).To(Equal("fileValue2"))
				Expect(viper.GetString("keyMap.nestedKey")).To(Equal("nestedValue"))
				Expect(viper.GetString("keymap.nestedMap.nestednestedKey")).To(Equal("nestednestedValue"))
			})
		})
	})
	Describe("Checking default settings", func() {
		It("should use defaults, if no value is st", func() {
			config.Setup(name, version, commit, setEnvPrefix, defaults)
			Expect(viper.GetString(testBoolKey)).To(Equal(testStringValue))
		})
	})

	})
})
