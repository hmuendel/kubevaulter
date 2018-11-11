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

type Execution struct {
	Command string
	Args []string
	Env map[string]string
}

type Executions []Execution

func NewExecutions() Executions {
	ex := make(Executions,0)
	return ex
}

func (ex *Executions) Init() error {
	err := viper.UnmarshalKey("executions", &ex)
	if err != nil {
		return err
	}
	return nil
}

func (ex *Executions) Validate() error {
	for e := range *ex {
		ok, err := valid.ValidateStruct(e)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("error validating executions config")
		}
	}
	return nil
}
