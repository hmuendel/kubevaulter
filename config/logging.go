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
	log "github.com/Sirupsen/logrus"
	valid "github.com/asaskevich/govalidator"
	"github.com/fsnotify/fsnotify"
	"errors"
)

type Logging struct {
	LogLevel string `valid:"in(debug|info|warning|error)~Invalid logLevel"`
	LogFormat string `valid:"in(text|json)~Invalid logFormat"`
}

func NewLogginConfig() (*Logging) {
	lc := Logging{}
	return &lc
}

func (lc *Logging) Init() error {
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Debug("LoggingConfig file changed:", e.Name)
		err := lc.configure()
		if err != nil {
			log.Warning("changed config produce err ", err, " falling back to previous")
		}
	})
	err := lc.configure()
	if err != nil {
		return err
	}
	return nil
}

func (lc *Logging) configure() error {
	err := viper.UnmarshalKey("logging", &lc)
	if err != nil {
		return err
	}
	err = lc.Validate()
	if err != nil {
		return err
	}
	err = lc.configureLogrus()
	if err != nil {
		return err
	}
	return nil
}

func (lc *Logging) Validate() error {
	ok, err := valid.ValidateStruct(lc)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("error validating logging config")
	}
	return nil
}


func (lc *Logging) configureLogrus() error {

	if logrusLevel, err := log.ParseLevel(lc.LogLevel); err == nil {
		log.SetLevel(logrusLevel)
		log.Debug("log level set to ", logrusLevel)
	} else {
		return errors.New("error setting log level")
	}
	switch lc.LogFormat {
	case "text":
		log.SetFormatter(&log.TextFormatter{})
		log.Debug("logging format set to text")
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
		log.Debug("logging format set to json")
	}
	return nil
}