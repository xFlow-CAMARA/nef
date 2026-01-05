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
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"gitlab.eurecom.fr/open-exposure/nef/ue-profile-service/internal/cache"
	"gitlab.eurecom.fr/open-exposure/nef/ue-profile-service/pkg/config"
)

type appCtx interface {
	Cfg() *config.AppConfig
	Cache() *cache.CacheManager
	AppName() string
}

type NbiServer struct {
	appCtx
	router *mux.Router
	server *http.Server
}

func NewServerInstance(app appCtx) (*NbiServer, error) {
	nbi := &NbiServer{appCtx: app}

	controller := NewUeProfileAPIController(nbi)

	nbi.router = mux.NewRouter()
	nbi.router.HandleFunc("/ue-profile/v1/profiles/{ue_id}", controller.RetrieveIndividualUeProfile)
	nbi.router.HandleFunc("/ue-profile/v1/profiles", controller.RetriveAllUeProfiles)

	nbi.router.Use(LoggerMiddleware)

	nbi.server = &http.Server{Addr: ":" + strconv.FormatUint(uint64(nbi.Cfg().NbiPort), 10), Handler: nbi.router}
	return nbi, nil
}

func (n *NbiServer) startListening(wg *sync.WaitGroup) {
	defer func() {
		_ = recover()
		wg.Done()
	}()

	log.Printf("nbi http(s) server started")
	err := n.server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		log.Default().Printf("could not start nbi http(s) server")
	}

}

func (n *NbiServer) Stop() {
	if n.server != nil {
		err := n.server.Close()
		if err != nil {
			log.Default().Printf("could not stop nbi server")
		}
	}

}

func (n *NbiServer) Run(wg *sync.WaitGroup) {
	wg.Add(1)
	go n.startListening(wg)

}
