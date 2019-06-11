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
	"github.com/hmuendel/kubevaulter/templater"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

var (
	VERSION       string
	COMMIT        string
	secretMap     map[string]secretData
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

	config.Setup("kubevaulter exec", VERSION, COMMIT, "KV", defaults)
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

	log.Debug("reading execution config")
	executions := config.NewExecutions()
	err = executions.Init()
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("reading recursive pathes config")
	recPathes := config.NewRecursivePathList()
	err = recPathes.Init()
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("creating login forge with path: ", vaultConfig.JwtPath, " and auth path ", vaultConfig.AuthPath)
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
	log.Debug("populating secret map")
	secretMap, err := vh.Populate(vaultConfig.SecretBackend, secretList)
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("rendering pathes: ", recPathes)
	for _, recPath := range []string(recPathes) {
		log.Debug("path: ", recPath)
		templater.Walk(recPath, secretMap)
	}

	for _, execution := range executions {
		renderedArgs := make([]string, len(execution.Args))
		for idx, arg := range execution.Args {
			tpl := template.New("arg")
			_, err := tpl.Parse(os.ExpandEnv(arg))
			if err != nil {
				log.Error(err)
			}
			buf := strings.Builder{}
			err = tpl.Execute(&buf, secretMap)
			if err != nil {
				log.Error(err)
			}
			renderedArgs[idx] = buf.String()
		}

		cmd := exec.Command(execution.Command, renderedArgs...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		//cmd.Args = renderedArgs

		renderedEnv := make([]string, len(execution.Env))
		idx := 0
		for name, value := range execution.Env {
			tpl := template.New("env")
			_, err := tpl.Parse(value)
			if err != nil {
				log.Error(err)
			}
			buf := strings.Builder{}
			buf.WriteString(name)
			buf.WriteString("=")
			err = tpl.Execute(&buf, secretMap)
			if err != nil {
				log.Error(err)
			}
			renderedEnv[idx] = buf.String()
			idx += 1
		}
		cmd.Env = append(os.Environ(), renderedEnv...)
		log.Debug("runnning command: ", cmd.Path)
		log.Debug("args: ", execution.Args)
		log.Debug("env: ", execution.Env)

		err := cmd.Run()
		//out,err:= cmd.Output()
		//log.Info(string(out))
		if err != nil {
			log.Error(err)
			if execution.FailNonZero {
				log.Error("fail flag set, exiting with non zero")
				os.Exit(1)
			}
			log.Error("fail flag unset, continuing")

		}
	}
	log.Info("finished successfully")
}
