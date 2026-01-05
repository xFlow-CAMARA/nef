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

package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.eurecom.fr/open-exposure/nef/capif-service/internal/3gpp/publish_api_service"
	"gitlab.eurecom.fr/open-exposure/nef/capif-service/internal/capif"
)

func HandleServiceCreation(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body := &publish_api_service.ServiceAPIDescription{}
		err := json.NewDecoder(r.Body).Decode(body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		sid, err := capif.PublishNewApi(body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// todo: add go routine to monitor the service status

		w.WriteHeader(http.StatusCreated)
		if _, err := w.Write([]byte(sid)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleIndividualService(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusNotImplemented)
		return
	case http.MethodDelete:
		serviceId := mux.Vars(r)["serviceId"]
		err := capif.DeletePublishedApi(serviceId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func HandleLogEntry(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body := &capif.LogInfo{}
		err := json.NewDecoder(r.Body).Decode(body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		serviceId := mux.Vars(r)["serviceId"]
		err = capif.NewLogEntry(serviceId, body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleTokenValidation(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body := &capif.OauthInfo{}
		err := json.NewDecoder(r.Body).Decode(body)
		if err != nil {
			log.Default().Printf("cannot validate request body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		serviceId := mux.Vars(r)["serviceId"]
		err = capif.ValidateToken(serviceId, body)
		if err != nil {
			log.Printf("could not validate token due to: %s", err.Error())
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusOK)
		}

		out := &bytes.Buffer{}
		err = json.NewEncoder(out).Encode(body)
		if err != nil {
			http.Error(w, "could not encode response body", http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(out.Bytes()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleCallbacks(w http.ResponseWriter, r *http.Request)           {}
func HandleCapifServiceProfile(w http.ResponseWriter, r *http.Request) {}

func HandleHeartbeat(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
