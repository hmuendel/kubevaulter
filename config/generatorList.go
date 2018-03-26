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

package config

import (
"github.com/spf13/viper"
valid "github.com/asaskevich/govalidator"
"errors"
)

type Generator struct {
	Path              string
	Key               string
	Override          bool
	Length            int
	AllowedCharacters string
}

type GeneratorList []Generator

func NewGeneratorList() GeneratorList {
	gl := make(GeneratorList,0)
	return gl
}

func (gl *GeneratorList) Init() error {
	err := viper.UnmarshalKey("generatorList", &gl)
	if err != nil {
		return err
	}
	return nil
}

// Validate the values parsed from the config.
func (gl *GeneratorList) Validate() error {
	for g := range *gl {
		ok, err := valid.ValidateStruct(g)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("error validating generator list")
		}
	}
	return nil
}

