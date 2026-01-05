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
	"os"
)

type AppConfig struct {
	RedisSvc string
	Port     string
}

func AppConfigFromEnv() *AppConfig {
	return &AppConfig{
		RedisSvc: getEnvString("REDIS_SVC", "redis:6379"),
		Port:     getEnvString("HTTP_PORT", "8080"),
	}
}

func getEnvString(key string, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	} else {
		return defaultValue
	}
}
