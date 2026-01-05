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

package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"gitlab.eurecom.fr/open-exposure/nef/ue-profile-service/internal/models"
	"gitlab.eurecom.fr/open-exposure/nef/ue-profile-service/internal/redis"
	"gitlab.eurecom.fr/open-exposure/nef/ue-profile-service/pkg/config"
)

type appCtx interface {
	Cfg() *config.AppConfig
	AppName() string
	Redis() *redis.RedisAdaptor
}

type CacheManager struct {
	appCtx
	mu      sync.RWMutex
	entries map[string]*models.UeProfile
	//expiresAt map[string]int64
	cancelFn map[string]context.CancelFunc
}

func NewCacheManager(app appCtx) *CacheManager {
	return &CacheManager{
		appCtx:   app,
		entries:  make(map[string]*models.UeProfile),
		cancelFn: make(map[string]context.CancelFunc),
	}
}

func (c *CacheManager) Get(imsi string) (*models.UeProfile, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	profile, ok := c.entries[imsi]
	if !ok {
		log.Printf("cache miss for %s", imsi)
		info, err := c.Redis().QueryUEInfo(imsi)
		if err != nil {
			return nil, false
		}
		profile = models.NewUeProfileFromUeInfo(info)
		if profile == nil {
			return nil, false
		}

		c.entries[imsi] = profile
		log.Printf("%s added to the cache TTL [+inf]", imsi)
		c.startListener(imsi)
	}

	return profile, true
}

func (c *CacheManager) startListener(imsi string) {
	// Cancel old one if exists
	if cancel, ok := c.cancelFn[imsi]; ok {
		cancel()
	}

	ctx, cancel := context.WithCancel(context.Background())
	c.cancelFn[imsi] = cancel

	go func() {
		sub := c.Redis().Instance.PSubscribe(ctx, fmt.Sprintf("user:%s:*", imsi))

		defer func() {
			_ = sub.Close()
		}()

		ch := sub.Channel()

		for {
			select {
			case msg := <-ch:
				patch := &models.UeInfoPatch{}
				if err := json.Unmarshal([]byte(msg.Payload), patch); err != nil {
					log.Printf("error in update: %s ", err.Error())

					continue
				}

				c.mu.RLock()
				profile, exists := c.entries[imsi]
				c.mu.RUnlock()

				if exists {
					profile.ApplyUpdate(patch)
				}
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (c *CacheManager) StopListener(imsi string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if cancel, ok := c.cancelFn[imsi]; ok {
		cancel()
		delete(c.cancelFn, imsi)
	}
}
