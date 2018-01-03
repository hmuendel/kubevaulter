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
	"io/ioutil"
)

const k8sAuthPath = "auth/kubernetes/login"


type JwtLoginForge struct {
	Payload map[string]interface{}
}

func NewJwtLoginForge(path, role string)  (*JwtLoginForge, error) {
	ka := JwtLoginForge{make(map[string]interface{})}
	err := ka.ReadToken(path)
	if err != nil {
		return nil, err
	}
	ka.SetRole(role)
	return &ka,nil
}

func (ka *JwtLoginForge) ReadToken(path string) error {
	token, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	ka.Payload["jwt"] = string(token)
	return  nil
}

func (ka *JwtLoginForge) SetRole(role string) {
	ka.Payload["role"] = role
}

func (ka *JwtLoginForge) GetPath() string {
	return k8sAuthPath
}

func (ka *JwtLoginForge) ForgeRequest() (map[string]interface{}) {
	return ka.Payload
}