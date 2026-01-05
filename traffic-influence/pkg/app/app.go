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

package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/google/uuid"
	"gitlab.eurecom.fr/open-exposure/nef/traffic-influence/internal/connector"
	contexts "gitlab.eurecom.fr/open-exposure/nef/traffic-influence/internal/context"
	"gitlab.eurecom.fr/open-exposure/nef/traffic-influence/internal/nbi"
	"gitlab.eurecom.fr/open-exposure/nef/traffic-influence/internal/nbi/service"
	"gitlab.eurecom.fr/open-exposure/nef/traffic-influence/pkg/config"
)

type AppCtx struct {
	ctx context.Context
	wg  sync.WaitGroup

	config  *config.AppConfig
	appId   string
	appName string
	//components
	server    *nbi.NbiServer
	connector *connector.Connector
	service   *service.Service

	//contexts
	traffInflCtx *contexts.TraffInflAppCtx
}

func InitApplication(name string, configPath string) (*AppCtx, error) {

	appInstance := &AppCtx{
		appName: name,
		config:  config.InitConfig(configPath),
		appId:   uuid.New().String(),
	}

	appInstance.server, _ = nbi.NewNorthbound(appInstance)
	appInstance.service = service.NewTraffInflService(appInstance)
	appInstance.connector = connector.NewConnector(appInstance)

	appInstance.traffInflCtx = contexts.NewTraffInflCtx(appInstance)

	return appInstance, nil
}

func (app *AppCtx) Cfg() *config.AppConfig {
	return app.config
}

func (app *AppCtx) Service() *service.Service {
	return app.service
}

func (app *AppCtx) Connector() *connector.Connector {
	return app.connector
}

func (app *AppCtx) Ctx() *contexts.TraffInflAppCtx {
	return app.traffInflCtx
}

func (app *AppCtx) Nbi() *nbi.NbiServer {
	return app.server
}

func (app *AppCtx) Run() {
	var cancel context.CancelFunc
	app.ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	app.wg.Add(1)
	go app.listenShutdownEvent()

	log.Printf("starting: %s (%s)", app.appName, app.appId)
	log.Printf("running config: \n%s", app.config.Dumps())

	if app.server != nil {
		app.server.Run(&app.wg)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	log.Printf("terminating...")

	cancel()
	app.wg.Wait()

}

func (app *AppCtx) listenShutdownEvent() {
	defer func() {
		_ = recover()
		app.wg.Done()
	}()

	<-app.ctx.Done()
	app.server.Stop()
}
