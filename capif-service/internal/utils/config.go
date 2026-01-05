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

package utils

import (
	"fmt"
	"log"
	"os"
)

type Configuration struct {
	CapifFqdn       string
	CapifAuthPort   string
	CapifServiceUrl string
	CapifUsername   string
	CapifPassword   string
	IngressFqdn     string
	CertsPath       string
}

var Config *Configuration

func InitConfig() {

	Config = &Configuration{
		CapifFqdn:     os.Getenv("CAPIF_FQDN"),
		CapifAuthPort: os.Getenv("CAPIF_AUTH_PORT"),
		CapifUsername: os.Getenv("CAPIF_USERNAME"),
		CapifPassword: os.Getenv("CAPIF_PASSWORD"),
		IngressFqdn:   os.Getenv("INGRESS_FQDN"),
		CertsPath:     "./certs", //"/etc/capif-service/certs/",
	}

	if Config.CapifFqdn == "" || Config.CapifAuthPort == "" {
		log.Fatalf("Capif endpoints are not set")
		os.Exit(1)
	}

	Config.CapifServiceUrl = fmt.Sprintf("https://%s", Config.CapifFqdn)

}
