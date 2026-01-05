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

package connector

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gitlab.eurecom.fr/open-exposure/nef/monitoring-event/pkg/config"
)

type app interface {
	Cfg() *config.AppConfig
}

type Connector struct {
	app
	redisClient *redis.Client
	ctx         context.Context
}

func NewConnector(app app) *Connector {

	svc := &Connector{
		app: app,
		ctx: context.Background(),
		redisClient: redis.NewClient(&redis.Options{
			Addr: app.Cfg().Sbi.RedisSvc,
		}),
	}
	return svc
}
