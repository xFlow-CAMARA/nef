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

package northbound

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"gitlab.eurecom.fr/open-exposure/nef/libcapif"
	"gitlab.eurecom.fr/open-exposure/nef/ue-id/internal/models"
	"gitlab.eurecom.fr/open-exposure/nef/ue-id/internal/northbound/service"
	"gitlab.eurecom.fr/open-exposure/nef/ue-id/pkg/config"
)

type appCtx interface {
	Cfg() *config.AppConfig
	Service() *service.Service
	AppName() string
}

type NbiServer struct {
	appCtx
	router   *mux.Router
	capifCtx *libcapif.CapifConnector
	server   *http.Server
}

func NewNorthbound(app appCtx) (*NbiServer, error) {
	nbi := &NbiServer{appCtx: app}

	UeIdAPIService := NewDefaultAPIService(nbi)
	UeIdAPIController := NewDefaultAPIController(UeIdAPIService)

	nbi.router = models.NewRouter(UeIdAPIController)

	if len(nbi.Cfg().CapifSvc) > 0 {
		nbi.capifCtx = libcapif.NewConnector(nbi.Cfg().CapifSvc)
		nbi.capifCtx.InstantiateConnector(nbi.AppName(), "v1", "HTTP_1_1")
		nbi.router.Use(nbi.capifCtx.CapifMiddleware)

		/*CAPIF Related code */
		/* Register routes to CAPIF */
		for rName, route := range UeIdAPIController.Routes() {
			nbi.capifCtx.AddEndpoint(route.Pattern, "REQUEST_RESPONSE", []string{route.Method}, rName)
		}

		nbi.capifCtx.AddInterface(nbi.Cfg().Nbi.Fqdn, int32(nbi.Cfg().Nbi.Port))
	}

	nbi.server = &http.Server{Addr: ":" + strconv.FormatUint(uint64(nbi.Cfg().Nbi.Port), 10), Handler: nbi.router}
	return nbi, nil
}

func (n *NbiServer) startListening(wg *sync.WaitGroup) {
	defer func() {
		_ = recover()
		wg.Done()
	}()

	if n.capifCtx != nil {
		n.capifCtx.RegisterService()
	} else {
		log.Default().Printf("capif svc not defined, disabling feature")
	}

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

	if n.capifCtx != nil {
		_ = n.capifCtx.DeleteService()
	}

}

func (n *NbiServer) Run(wg *sync.WaitGroup) {
	wg.Add(1)
	go n.startListening(wg)

}
