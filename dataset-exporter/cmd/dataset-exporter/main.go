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

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"gitlab.eurecom.fr/open-exposure/nef/dataset-exporter/internal/controller"
	"gitlab.eurecom.fr/open-exposure/nef/dataset-exporter/internal/service"
	"gitlab.eurecom.fr/open-exposure/nef/dataset-exporter/internal/utils"
)

var (
	gitBranch string
	gitCommit string
)

func main() {
	log.Default().Printf("================ dataset-exporter service for open-exposure ================")

	log.Printf("current git branch: %s", gitBranch)
	log.Printf("current git commit: %s", gitCommit)

	config := utils.AppConfigFromEnv()

	redisClient := redis.NewClient(&redis.Options{
		Addr: config.RedisSvc,
	},
	)

	captureService := service.NewCaptureService(redisClient)

	router := mux.NewRouter()
	router.Handle("/dataset-exporter/v1/captures", controller.StartCaptureHandler(captureService))
	router.Handle("/dataset-exporter/v1/captures/status", controller.GetCaptureStatusHandler(captureService))
	router.Handle("/dataset-exporter/v1/captures/download", controller.DownloadCaptureHandler(captureService))

	log.Printf("northbound server started")
	addr := fmt.Sprintf("0.0.0.0:%s", config.Port)

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("failed to start northbound server: %s", err.Error())
	}
}
