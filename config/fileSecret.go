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

// FileSecret holds the configuration parsed from the config file regarding the file secret templating.
type FileSecret struct {
	TemplatePath string
	TargetPath string
	SecretPath string
}

// FileSecretList is a slice of FileSecret
type FileSecretList []FileSecret

// NewFileSecretList creates and returns a FileSecretList of length 0.
func NewFileSecretList() FileSecretList {
	fsl := make(FileSecretList,0)
	return fsl
}

// Init initializes the FileSecretList with values from the config.
func (gl *FileSecretList) Init() error {
	err := viper.UnmarshalKey("fileSecretList", &gl)
	if err != nil {
		return err
	}
	return nil
}

// Validate the values parsed from the config.
func (gl *FileSecretList) Validate() error {
	for fs := range *gl {
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
