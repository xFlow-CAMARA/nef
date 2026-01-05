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

/*
 * 3gpp-traffic-influence
 *
 * API for AF traffic influence   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.3.0-alpha.4
 */

package main

import (
	"log"
	"os"

	traffInflApp "gitlab.eurecom.fr/open-exposure/nef/traffic-influence/pkg/app"
)

var (
	gitBranch string
	gitCommit string
)

func main() {
	log.Default().Printf("================ traffic influence service for open-exposure ================")

	log.Printf("current git branch: %s", gitBranch)
	log.Printf("current git commit: %s", gitCommit)

	app, err := traffInflApp.InitApplication("traffic-influence", getConfigPath())
	if err != nil {
		log.Printf("cannot initialize application")
		return
	}

	app.Run()
}

func getConfigPath() string {

	if len(os.Args) > 2 && os.Args[1] == "-c" && len(os.Args[2]) > 0 {
		return os.Args[2]
	} else {
		return "./etc/config.yaml"
	}

}
