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

package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/redis/go-redis/v9"
	"gitlab.eurecom.fr/open-exposure/nef/ue-profile-service/internal/models"
	"gitlab.eurecom.fr/open-exposure/nef/ue-profile-service/pkg/config"
)

type appCtx interface {
	Cfg() *config.AppConfig
	AppName() string
}

type RedisAdaptor struct {
	appCtx
	Instance *redis.Client
	ctx      context.Context
}

func NewRedisAdapter(ac appCtx) *RedisAdaptor {
	rds := redis.NewClient(&redis.Options{
		Addr: ac.Cfg().RedisSvc,
	})
	return &RedisAdaptor{
		appCtx:   ac,
		Instance: rds,
		ctx:      context.Background(),
	}
}

func (r *RedisAdaptor) QueryUEInfo(imsi string) (*models.UeInfo, error) {

	key := fmt.Sprintf("user:%s", imsi)
	result, err := r.Instance.Do(r.ctx, "JSON.GET", key).Text()
	if err != nil {
		return nil, fmt.Errorf("failed to get UE profile: %w", err)
	}
	// RedisJSON returns JSON array by default when using "$"
	ue := &models.UeInfo{}
	if err := json.Unmarshal([]byte(result), ue); err != nil {
		return nil, fmt.Errorf("failed to unmarshal UE profile: %w", err)
	}
	ue.Imsi = &imsi
	return ue, nil
}

func (r *RedisAdaptor) QueryUEsInfo() ([]*models.UeInfo, error) {

	keys, _, err := r.Instance.Scan(r.ctx, 0, "user:*", 0).Result()
	if err != nil {
		log.Fatalf("Failed to scan: %v", err)
	}

	var result []*models.UeInfo

	for _, key := range keys {
		key = strings.Split(key, ":")[1]
		ue, err := r.QueryUEInfo(key)
		if err == nil {
			result = append(result, ue)
		}
	}

	return result, nil
}
