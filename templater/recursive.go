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

package templater

import (
	"github.com/hmuendel/kubevaulter"
	"os"
	"path/filepath"
	"text/template"
)


//Walk recursively walks through all files under the path directory and rendering
// the templates in place with values from the secretMap
func Walk(path string, secretMap map[string]kubevaulter.SecretData) error {

	err := filepath.Walk(path,walkClosure(&secretMap))
	if err != nil {
		return err
	}
	return nil
}

func walkClosure(secretMap *map[string]kubevaulter.SecretData) func(string, os.FileInfo, error) error {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		tpl, err := template.ParseFiles(path)
		if err != nil {
			return err
		}
		f, err := os.Create(path)
		defer f.Close()
		if err != nil {
			return err
		}
		err = tpl.Execute(f,*secretMap)
		if err != nil {
			return err
		}
		f.Sync()
		return nil
	}
}

