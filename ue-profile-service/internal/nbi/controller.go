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

package nbi

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.eurecom.fr/open-exposure/nef/ue-profile-service/internal/cache"
)

type serverCtx interface {
	Cache() *cache.CacheManager
}

type UeProfileAPIController struct {
	serverCtx
}

func NewUeProfileAPIController(s serverCtx) *UeProfileAPIController {
	return &UeProfileAPIController{serverCtx: s}
}

func (ueApi *UeProfileAPIController) RetrieveIndividualUeProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imsi := vars["ue_id"]

	ueProfile, ok := ueApi.Cache().Get(imsi)
	if ok {
		bytes, err := json.Marshal(ueProfile)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			if _, err := w.Write(bytes); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")

			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func (ueApi *UeProfileAPIController) RetriveAllUeProfiles(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
