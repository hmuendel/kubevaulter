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
	"github.com/spf13/viper"
	valid "github.com/asaskevich/govalidator"
	"errors"
)

// Secret holds the configuration parsed from the config file regarding the file secret templating.
type Secret struct {
	VaultPath string
	Name string
}

// SecretList is a slice of Secret
type SecretList []Secret

// NewSecretList creates and returns a SecretList of length 0.
func NewSecretList() SecretList {
	fsl := make(SecretList,0)
	return fsl
}

// Init initializes the SecretList with values from the config.
func (fsl *SecretList) Init() error {
	err := viper.UnmarshalKey("secretList", &fsl)
	if err != nil {
		return err
	}
	return nil
}

// Validate the values parsed from the config.
func (fsl *SecretList) Validate() error {
	for fs := range *fsl {
		ok, err := valid.ValidateStruct(fs)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("error validating secretList config")
		}
	}
	return nil
}
