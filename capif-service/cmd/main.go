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
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	capif "gitlab.eurecom.fr/open-exposure/nef/capif-service/internal/capif"
	handlers "gitlab.eurecom.fr/open-exposure/nef/capif-service/internal/handlers"
	"gitlab.eurecom.fr/open-exposure/nef/capif-service/internal/utils"
)

func main() {
	log.Default().Printf("================ capif provider for open-exposure ================")
	/* Initi Configuration */
	utils.InitConfig()

	/* Perform initialization routine */
	capif.InitializeDefaultProvider()
	if err := capif.RegisterDefaultProvider(); err != nil {
		log.Fatalf("initialization failed: %s", err.Error())
	}

	/* Define HTTP Router */
	r := mux.NewRouter()

	//r.HandleFunc("/provider", handlers.HandleMicroserviceRegistration)
	r.HandleFunc("/services", handlers.HandleServiceCreation)
	r.HandleFunc("/services/{serviceId}", handlers.HandleIndividualService)
	r.HandleFunc("/services/{serviceId}/logs", handlers.HandleLogEntry)
	r.HandleFunc("/services/{serviceId}/validateToken", handlers.HandleTokenValidation)
	r.HandleFunc("/version", handlers.HandleCapifServiceProfile)
	r.HandleFunc("/healthz", handlers.HandleHeartbeat)

	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Channel to listen for termination signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Goroutine to handle graceful shutdown
	go func() {
		<-signalChan
		log.Default().Printf("Termination signal received. Shutting down...")
		if err := capif.DestroyDefaultProvider(); err != nil {
			log.Printf("Error during provider destruction: %s", err.Error())
		}
		os.Exit(0)
	}()

	log.Default().Printf("Server Started")
	log.Fatal(srv.ListenAndServe())
}
