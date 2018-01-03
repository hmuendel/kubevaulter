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
	"text/template"
	"os"
)

type Transformer struct {
	Fts FileTemplates
}

func (t *Transformer) Apply(idx int) error {
	temp, err := template.ParseFiles(t.Fts[idx].TemplatePath)
	if err != nil {
		return err
	}
	os.Remove(t.Fts[idx].TargetPath)
	f,err:= os.Create(t.Fts[idx].TargetPath)
	defer f.Close()
	if err != nil {
		return err
	}
	temp.Execute(f,t.Fts[idx].Template)
	f.Sync()
	return nil
}
