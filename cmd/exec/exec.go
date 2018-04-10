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

package exec

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hmuendel/kubevaulter"
	"github.com/hmuendel/kubevaulter/config"
	"text/template"
	"os"
	"path/filepath"
)



var (
	VERSION string
	COMMIT string
	secretMap map[string]secretData
	templatesPath = "/share"
)

type secretData map[string]interface{}

func main() {
	secretMap = make(map[string]secretData)
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

	log.Debug("reading secret config")
	secretList := config.NewSecretList()
	err = secretList.Init()
	if err != nil {
		log.Fatal(err)
	}
	log.Debug("creating login forge with path: ", vaultConfig.JwtPath, " and auth path ", vaultConfig.AuthPath )
	lf, err := kubevaulter.NewJwtLoginForge(vaultConfig.AuthPath, vaultConfig.JwtPath, vaultConfig.Role, vaultConfig.AuthPath)
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

	for _, secret := range secretList {
		log.Debug("getting secret from vault ", vaultConfig.SecretBackend+"/"+secret.VaultPath)
		s, err := vh.Read(vaultConfig.SecretBackend + "/" + secret.VaultPath)
		if err != nil {
			log.Error(err)
		}
		if s != nil && s.Data != nil {
			secretMap[secret.Name] = s.Data
		} else {
			if vaultConfig.FailOnEmptySecret {
				log.Fatal("Empty reply from secret ", vaultConfig.SecretBackend+"/"+secret.VaultPath)
			} else {
				log.Warning("Empty reply from secret ", vaultConfig.SecretBackend+"/"+secret.VaultPath)
			}
		}
	}
	err = filepath.Walk(templatesPath, walkFunc)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("finished successfully")
}


func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	log.Debug("reading template ", path)
	tpl, err := template.ParseFiles(path)
	if err != nil {
		return err
	}
	log.Debug("read template: ", tpl)
	log.Debug("openening for writing : ", path)
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		return err
	}
	log.Debug("rendering template: ", path)
	err = tpl.Execute(f,secretMap)
	if err != nil {
		return err
	}
	log.Debug("writing rendered template to file")
	f.Sync()
	return nil
}











