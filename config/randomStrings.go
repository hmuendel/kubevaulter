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



type RandomString struct {
	Override          bool
	Length            int
	AllowedCharacters string
}

type RandomStrings map[string]*RandomString

func NewRandomStrings() RandomStrings {
	rs := make(RandomStrings)
	return rs
}

func (rs *RandomStrings) Init() error {
	err := viper.UnmarshalKey("randomStrings", &rs)
	if err != nil {
		return err
	}
	for _, value := range map[string]*RandomString(*rs) {
		if value.Length == 0 {
			value.Length = viper.GetInt("generator.length")
		}
		if value.AllowedCharacters == "" {
			value.AllowedCharacters = viper.GetString("generator.allowedCharacters")
		}
	}
	return nil
}

// Validate the values parsed from the config.
func (rs *RandomStrings) Validate() error {
	for g := range *rs {
		ok, err := valid.ValidateStruct(g)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("error validating random strings")
		}
	}
	return nil
}

