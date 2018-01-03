/*
 * Copyright 2017 Hans Mündelein
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

package config

import (
	valid "github.com/asaskevich/govalidator"
	"errors"
	"github.com/spf13/viper"
)

type Vault struct {
	EndpointUrl string `valid:"url~Invalid EndpointUrl"`
	SecretBackend string `valid:"-"`
	Role string `valid:"-"`
	JwtPath string `valid:"matches(.+)~Invalid JwtPath"`
}

func NewVaultconfig() (*Vault) {
	vc := Vault{}
	return &vc
}

func (vc *Vault) Init() error {
	err := viper.UnmarshalKey("vault", &vc)
	if err != nil {
		return err
	}
	return nil
}

func (vc *Vault) Validate() error {
	ok, err := valid.ValidateStruct(vc)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("error validating vault config")
	}
	return nil
}

