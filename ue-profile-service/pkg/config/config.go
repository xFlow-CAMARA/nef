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
	"os"
	"strconv"
)

type AppConfig struct {
	HttpVersion uint16
	UseTLS      bool
	NbiFqdn     string
	NbiPort     uint16
	CapifSvc    string
	NrfSvc      string
	RedisSvc    string
}

func InitConfig() *AppConfig {

	cfg := &AppConfig{
		HttpVersion: getenvUint16("HTTP_VER", 1),
		UseTLS:      getenvBool("HTTPS", false),
		NbiFqdn:     getenvStr("NBI_FQDN", "ue-profile"),
		NbiPort:     getenvUint16("NBI_PORT", 8080),
		RedisSvc:    getenvStr("REDIS_SVC", "redis:6379"),
	}

	return cfg
}

func getenvStr(key string, defaultValue string) string {
	val := os.Getenv(key)
	if len(val) > 0 {
		return val
	}
	return defaultValue
}

func getenvBool(key string, defaultValue bool) bool {
	val := os.Getenv(key)
	if len(val) > 0 {
		val, err := strconv.ParseBool(val)
		if err != nil {
			return val
		}
	}
	return defaultValue
}

func getenvUint16(key string, defaultValue uint16) uint16 {
	val := os.Getenv(key)
	if len(val) > 0 {
		val, err := strconv.ParseUint(val, 10, 16)
		if err != nil {
			return uint16(val)
		}
	}
	return defaultValue
}
