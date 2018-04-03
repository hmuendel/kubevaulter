/*
 *    Copyright 2017 Hans Mündelein
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package kubevaulter

import (
	vault "github.com/hashicorp/vault/api"
)

type SecretData map[string]interface{}
type SecretDataMap map[string]Secret
type Secret struct {
	Name string
	Path string
	Data SecretData
}


//ApiWrapper is a wrapper around the official vault raw client as well as the more abstract logical api.
//It also holds a login forge for creating login requests to authenticate against a vault auth backend.
type ApiWrapper struct {
	loginForge LoginForge
	client     *vault.Client
	api        *vault.Logical
}

//NewApiWrapper creates an ApiWrapper with the specified LoginForge and vault server address.
func NewApiWrapper(loginForger LoginForge, addr string) (*ApiWrapper, error) {
	config := vault.DefaultConfig()
	config.Address = addr
	config.ConfigureTLS(&vault.TLSConfig{CACert:loginForger.CaCert()})
	vc, err := vault.NewClient(config)

	if err != nil {
		return nil,err
	}
	logical := vc.Logical()
	aw  := ApiWrapper{loginForge: loginForger,client:vc,api:logical}
	return &aw, nil
}

//KubeAuth performs
func (aw *ApiWrapper) KubeAuth() (*vault.Secret, error) {
	secret, err := aw.api.Write(aw.loginForge.GetPath(),aw.loginForge.ForgeRequest())
	if err != nil {
		return nil,err
	}
	aw.client.SetToken(secret.Auth.ClientToken)
	return secret, nil
}

func (aw *ApiWrapper) Read(path string) (*vault.Secret, error) {
	resp, err :=aw.api.Read(path)

	if err != nil {
		return nil,err
	}
	return  resp, nil
}


func (aw *ApiWrapper) Write(path string, data map[string]interface{}) (*vault.Secret, error) {
	resp, err :=aw.api.Write(path, data)

	if err != nil {
		return nil,err
	}
	return  resp, nil
}

func (aw *ApiWrapper) Populate()  {
	
}


