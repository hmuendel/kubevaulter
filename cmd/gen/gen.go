/*
 * Copyright 2018 Hans Mündelein
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hmuendel/kubevaulter"
	"github.com/hmuendel/kubevaulter/config"
	"github.com/hmuendel/kubevaulter/randstring"
	"github.com/hmuendel/kubevaulter/transformer"
)

var (
	VERSION string
	COMMIT string
)


func main() {
	defaults := make(map[string]interface{})
	defaults["configPath"] = "./config"
	defaults["configName"] = "config"

	defaults["logging.logLevel"] = "info"
	defaults["logging.logFormat"] = "text"

	defaults["vault.endpointUrl"] = "http://localhost:8200"
	defaults["vault.secretBackend"] = "secret"
	defaults["vault.role"] = "demo"
	defaults["vault.jwtPath"] = "/var/run/secrets/kubernetes.io/serviceaccount/token"
	defaults["vault.authPath"] = "auth/kubernetes/login"
	defaults["generator.override"] = false
	defaults["generator.length"] = 16
	defaults["generator.allowedCharacters"] = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	defaults["generator.transform"] = "NONE"


	config.Setup("kubevaulter gen", VERSION, COMMIT, "KV", defaults)
	loggingConfig := config.NewLogginConfig()
	err := loggingConfig.Init()
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("reading vault config")
	vaultConfig := config.NewVaultconfig()

	err = vaultConfig.Init()
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("reading target list config")
	targetList := config.NewTargetList()
	err = targetList.Init()
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("reading random strings config")
	randomStringsCfg := config.NewRandomStrings()
	err = randomStringsCfg.Init()
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("creating login forge with path: ", vaultConfig.JwtPath, ", Role: ", vaultConfig.Role,
		", caCert: ", vaultConfig.CaCert," and auth path ", vaultConfig.AuthPath )
	lf, err := kubevaulter.NewJwtLoginForge(vaultConfig.AuthPath, vaultConfig.JwtPath, vaultConfig.Role, vaultConfig.CaCert)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	log.Debug("creating api wrapper")
	vh, err := kubevaulter.NewApiWrapper(lf, vaultConfig.EndpointUrl)
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("authenticating against vault")
	_, err = vh.KubeAuth()
	if err != nil {
		log.Fatal(err)
	}

	randomStringMap := make(map[string]string)
	funcMap := transformer.DefaultFuncMap()
	for key, value := range randomStringsCfg {
		log.Debug("creating random string ", key)
		randomStringMap[key] = randstring.Create(value.Length,value.AllowedCharacters)
	}
	log.Debug("created random string map ")
	// iterating over all target entries in target list
	for _, target := range targetList  {
		log.Debug("handling data for target ", target.Path)
		payload := make(map[string]interface{})
		// reading target value from vault to check if it exists
		s,reterr := vh.Read(target.Path)
		// iterating over key value pairs in the desired secret
		for key, value := range target.Data {
			log.Debug("handling data ", key)
			//using the literal value if set
			if value.Lit != "" {
				log.Debug("using literal value ")
				if fn, ok := funcMap[value.Transform]; ok {
					payload[key] = fn(value.Lit)
				} else {
					log.Error("No function found with identifier: ", value.Transform)
				}
			} else if value.Ref != "" {
				log.Debug("using ref to random string ", value.Ref)
				var val interface{}
				ok := false
				//checking if the value at this key exists and can be retrieved
				if reterr == nil && s != nil && s.Data != nil {
					log.Debug("retrieving value from old secret")
					val, ok = s.Data[key]
				}
				// if value exists and override flag is not set prepare payload with old value

				if rString, exists := randomStringsCfg[value.Ref]; exists && ! rString.Override && ok {
					if ! exists {
						log.Warning("nothing found in random string cfg for ", value.Ref)
					}
					log.Debug("using old value")
					payload[key] = val
				} else {
					// otherwise prepare payload with new value
					log.Debug("writing new value")
					if  val, ok := randomStringMap[value.Ref]; ok {
						if fn, ok := funcMap[value.Transform]; ok {
							payload[key] = fn(val)
						} else {
							log.Error("No function found with identifier: ", value.Transform)
						}
					} else {
						log.Error("Referenced random string not find ",value.Ref )
					}
				}
			} else {
				log.Error("One of lit or ref must be set for ", key)
			}
		}
		log.Debug("writing data structure to vault")
		vh.Write(target.Path,payload)
	}
	log.Info("finished successfully")

}