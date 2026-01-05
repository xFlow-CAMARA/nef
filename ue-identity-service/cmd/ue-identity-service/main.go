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
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"gitlab.eurecom.fr/open-exposure/nef/ue-identity-service/internal/redis"
	"gitlab.eurecom.fr/open-exposure/nef/ue-identity-service/internal/resolver"
	"gitlab.eurecom.fr/open-exposure/nef/ue-identity-service/internal/utils"
)

var (
	gitBranch string
	gitCommit string
)

func main() {
	log.Default().Printf("================ ue-identity service for open-exposure ================")

	log.Printf("current git branch: %s", gitBranch)
	log.Printf("current git commit: %s", gitCommit)

	// Init config
	config := utils.AppConfigFromEnv()

	// Connect to Redis
	rdb := redis.NewRedisAdapter(config.RedisSvc)
	key := []byte("qRUmErY0HLQ9UANsPr6FM6vMtwvxI2mF")

	resolver := resolver.NewResolver(rdb)

	// HTTP endpoint for testing
	http.HandleFunc("/lookup", func(w http.ResponseWriter, r *http.Request) {
		ip := r.URL.Query().Get("ip")
		afId := r.URL.Query().Get("afId")

		if ip == "" {
			http.Error(w, "Missing 'ip' query param", http.StatusBadRequest)
			return
		}
		supi, err := resolver.Lookup(ip)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		encSupi, err := utils.EncodeIMSIWithAfID(supi, afId, key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		err = json.NewEncoder(w).Encode(map[string]string{
			"Ip":         ip,
			"Supi":       supi,
			"ExternalId": fmt.Sprintf("%s:%s", afId, encSupi),
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// HTTP endpoint for testing
	http.HandleFunc("/resolve", func(w http.ResponseWriter, r *http.Request) {
		externalId := r.URL.Query().Get("externalId")

		if externalId == "" {
			http.Error(w, "missing 'externalId' query param", http.StatusBadRequest)
			return
		}

		var supi string
		var err error

		// Check if externalId is in format "afId:encodedIdentifier" or plain IMSI
		extId := strings.Split(externalId, ":")
		if len(extId) == 2 {
			// Format: "afId:encodedIdentifier" - decode it
			afId := extId[0]
			supiEncoded := extId[1]

			var afIdDecoded string
			supi, afIdDecoded, err = utils.DecodeIMSIWithAfID(supiEncoded, key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if afId != afIdDecoded {
				http.Error(w, "afId mismatch", http.StatusNotFound)
				return
			}
		} else {
			// Plain IMSI format - use directly as SUPI
			supi = externalId
		}

		err = json.NewEncoder(w).Encode(map[string]string{
			"Supi": supi,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// HTTP endpoint for MSISDN lookup by IP
	http.HandleFunc("/msisdn", func(w http.ResponseWriter, r *http.Request) {
		ip := r.URL.Query().Get("ip")

		if ip == "" {
			http.Error(w, "Missing 'ip' query param", http.StatusBadRequest)
			return
		}
		msisdn, err := resolver.LookupMsisdn(ip)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		err = json.NewEncoder(w).Encode(map[string]string{
			"Ip":     ip,
			"Msisdn": msisdn,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("Starting server on :", config.Port)
	if err := http.ListenAndServe(":"+config.Port, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
