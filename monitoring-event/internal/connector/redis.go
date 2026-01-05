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
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/redis/go-redis/v9"
	"gitlab.eurecom.fr/open-exposure/nef/monitoring-event/internal/models"
)

func (r *Connector) QueryUEInfo(imsi string) (*models.UeInfo, error) {

	key := fmt.Sprintf("user:%s", imsi)
	log.Printf("QueryUEInfo: Querying Redis for key=%s", key)
	result, err := r.redisClient.Do(r.ctx, "JSON.GET", key).Text()
	if err != nil {
		log.Printf("QueryUEInfo: Redis query failed for key=%s, error=%v", key, err)
		return nil, fmt.Errorf("failed to get UE profile: %w", err)
	}
	log.Printf("QueryUEInfo: Got Redis result for key=%s, length=%d bytes", key, len(result))
	// RedisJSON returns JSON array by default when using "$"
	ue := &models.UeInfo{}
	if err := json.Unmarshal([]byte(result), ue); err != nil {
		return nil, fmt.Errorf("failed to unmarshal UE profile: %w", err)
	}
	ue.Imsi = &imsi
	return ue, nil
}

func (r *Connector) QueryUEsInfo() ([]*models.UeInfo, error) {

	keys, _, err := r.redisClient.Scan(r.ctx, 0, "user:*", 0).Result()
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

func (r *Connector) SubscribeUserInfo(imsi string) (*redis.PubSub, error) {
	sub := r.redisClient.Subscribe(r.ctx, fmt.Sprintf("user:%s", imsi))
	if sub == nil {
		return nil, fmt.Errorf("failed to subscribe to user info for %s", imsi)
	}

	return sub, nil
}

func (r *Connector) SubscribeUserEvent(imsi string, eventType string) (*redis.PubSub, error) {
	sub := r.redisClient.Subscribe(r.ctx, fmt.Sprintf("user:%s:%s", imsi, eventType))
	if sub == nil {
		return nil, fmt.Errorf("failed to subscribe to user info for %s", imsi)
	}

	return sub, nil
}

func (r *Connector) SubscribeEventType(eventType string) (*redis.PubSub, error) {
	sub := r.redisClient.Subscribe(r.ctx, fmt.Sprintf("BROADCAST:%s", eventType))
	if sub == nil {
		return nil, fmt.Errorf("failed to subscribe to %s", eventType)
	}

	return sub, nil
}
