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


type FileSecret struct {
	TemplatePath string
	TargetPath string
	SecretPath string
}

type FileSecretList []FileSecret

func NewFileSecretList() FileSecretList {
	fsl := make(FileSecretList,0)
	return fsl
}

func (fsl *FileSecretList) Init() error {
	err := viper.UnmarshalKey("fileSecretList", &fsl)
	if err != nil {
		return err
	}
	return nil
}

func (fsl *FileSecretList) Validate() error {
	for fs := range *fsl {
		ok, err := valid.ValidateStruct(fs)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("error validating vault config")
		}
	}
	return nil
}
