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
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	valid "github.com/asaskevich/govalidator"
	"os"
)


// Setup is initializing the basic configuration like config path and name and the validation framework
// as well as setting the passed defaults
func Setup(name, version, commit ,envPrefix string, defaults map[string]interface{})  {
	log.Printf("Starting %s in version: %s commit: %s", name, version, commit )
	setDefaults(defaults)
	viper.SetEnvPrefix(envPrefix)
	viper.BindEnv("configPath", envPrefix + "_CONFIG_PATH")
	viper.BindEnv("configName", envPrefix + "_CONFIG_NAME")
	log.Printf("Config path env variable ",envPrefix,"_CONFIG_PATH:", os.Getenv(envPrefix + "_CONFIG_PATH"))
	log.Printf("Config name env variable ",envPrefix,"_CONFIG_NAME:", os.Getenv(envPrefix + "_CONFIG_NAME"))
	log.Printf("Reading config from ", viper.GetString("configPath"),"/",viper.GetString("configName"))
	viper.SetConfigName(viper.GetString("configName"))
	viper.AddConfigPath(viper.GetString("configPath"))

	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		log.Panic("Fatal error config file: %s \n", err)
	}
	viper.AutomaticEnv()
	viper.WatchConfig()

	//Initializing the validation framework
	valid.SetFieldsRequiredByDefault(true)
}

func setDefaults(defaults map[string]interface{})  {
	log.Debug("Setting defaults ", defaults)
	for k, v := range defaults {
		viper.SetDefault(k,v)
	}
}


