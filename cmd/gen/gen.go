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


	config.Setup("kubevaulter rec", VERSION, COMMIT, "KV", defaults)
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

	log.Debug("creating login forge with path: ", vaultConfig.JwtPath, " and auth path ", vaultConfig.AuthPath )
	lf, err := kubevaulter.NewJwtLoginForge(vaultConfig.AuthPath, vaultConfig.JwtPath, vaultConfig.Role)
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
	for key, value := range randomStringsCfg {
		randomStringMap[key] = randstring.Create(value.Length,value.AllowedCharacters)
	}
	log.Debug("created random string map ", randomStringMap)

	for _, target := range targetList  {
		payload := make(map[string]interface{})
		for key, value := range target.Data {
			payload[key] = randomStringMap[value.Ref]
		}
		vh.Write(target.Path,payload)
	}

}