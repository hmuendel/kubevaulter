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

package templater

import (
	"errors"
	log "github.com/Sirupsen/logrus"
)

//FileTemplate contains the description
type FileTemplate struct {
	TemplatePath string
	TargetPath string
	Template map[interface{}]interface{}
}

func (ft *FileTemplate) CastFromMap(inputMap map[interface{}]interface{}) error {
	path, ok := inputMap["templatePath"].(string)
	if ok {
		ft.TemplatePath = path
	} else {
		return errors.New("error casting templatePath to string")
	}
	target, ok := inputMap["targetPath"].(string)
	if ok {
		ft.TargetPath = target
	} else {
		return errors.New("error casting templatePath to string")
	}
	template, ok := inputMap["template"].(map[interface{}]interface{})
	if ok {
		ft.Template = template
	} else {
		return errors.New("error casting template to map")
	}
	return nil
}

func NewFileTemplate(input interface{}) (*FileTemplate, error) {

	inputMap, ok := input.(map[interface{}]interface{})
	if !ok {
		return nil, errors.New("error casting input to file map")
	}
	ft := FileTemplate{}
	err := ft.CastFromMap(inputMap)
	if err != nil {
		return nil, err
	} else {
		return &ft, nil
	}
}

type FileTemplates []*FileTemplate

func CastToFileTemplates(input interface{}) (FileTemplates,error) {
	inputSlice, ok := input.([]interface{})
	log.Println(inputSlice)
	fts := make(FileTemplates, len(inputSlice))
	if ok {
		for idx, elem := range inputSlice {
			log.Println("handling", elem)
			ft, err := NewFileTemplate(elem)
			if err != nil {
				return nil, err
			}
			fts[idx] = ft
		}
	return fts, nil
	} else {
		return nil, errors.New("error casting to array")
	}
}


