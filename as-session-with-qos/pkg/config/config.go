// Copyright 2025 EURECOM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Contributors:
//   Giulio CAROTA
//   Thomas DU
//   Adlen KSENTINI

package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	/* Basic configuration parameters */
	Sbi           SbiConfig `yaml:"sbi"`
	Nbi           NbiConfig `yaml:"nbi"`
	CapifSvc      string    `yaml:"capifSvc"`
	SupportedFeat string    `yaml:"supportedFeatures"`

	/* Custom configuration parameters */
	QosConf map[string]QosConfig `yaml:"qosConfig"`
}

type NbiConfig struct {
	HttpVersion uint16 `yaml:"httpVersion"`
	UseTLS      bool   `yaml:"useTLS"`
	Fqdn        string `yaml:"fqdn"`
	Port        uint16 `yaml:"port"`
}

type SbiConfig struct {
	NrfSvc      string `yaml:"nrfSvc"`
	UseNrf      bool   `yaml:"useNrf"`
	PcfSvc      string `yaml:"pcfSvc"`
	Httpversion int    `yaml:"httpVersion"`
}

type QosConfig struct {
	MarBwDl   string `yaml:"marBwDl"`
	MarBwUl   string `yaml:"marBwUl"`
	MediaType string `yaml:"mediaType"` /*allowed values CONTROL, AUDIO, VIDEO*/
}

func InitConfig(configPath string) *AppConfig {

	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("cannot read config file #%v ", err)
	}

	cfg := AppConfig{}
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return &cfg
}

func (cfg *AppConfig) Dumps() string {
	d, err := yaml.Marshal(&cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return string(d)

}
