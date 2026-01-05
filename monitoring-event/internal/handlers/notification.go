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
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
	"gitlab.eurecom.fr/open-exposure/nef/monitoring-event/internal/models"
)

type callbackFun func(subscriptionLocation string, userId string, notificationUri string, patch *models.UeInfoPatch) error

type NotificationHandler struct {
	active               bool
	userId               string
	ctx                  context.Context
	notificationUri      string
	redisSubscription    *redis.PubSub
	callback             callbackFun
	cancelFunc           context.CancelFunc
	subscriptionLocation string
}

func NewNotificationHandler(subscriptionLocation string, identiy string, notificationUri string, sub *redis.PubSub) *NotificationHandler {

	return &NotificationHandler{
		active:               false,
		ctx:                  nil,
		userId:               identiy,
		notificationUri:      notificationUri,
		redisSubscription:    sub,
		callback:             nil,
		subscriptionLocation: subscriptionLocation,
	}
}

func (notifHandler *NotificationHandler) Start() bool {
	if notifHandler.callback == nil {
		log.Printf("NotificationHandler: callback is not set")
		return false
	}

	notifHandler.ctx, notifHandler.cancelFunc = context.WithCancel(context.Background())

	go func() {
		if notifHandler.redisSubscription == nil {
			log.Printf("NotificationHandler: redis subscription is nil")
			return
		}

		defer func(sub *redis.PubSub) {
			if err := sub.Close(); err != nil {
				log.Printf("could not close redis channel correctly")
			}
		}(notifHandler.redisSubscription)

		ch := notifHandler.redisSubscription.Channel()

		for {
			select {
			case msg := <-ch:
				patch := &models.UeInfoPatch{}
				if err := json.Unmarshal([]byte(msg.Payload), patch); err != nil {
					log.Printf("error in update: %s ", err.Error())
					continue
				}
				err := notifHandler.callback(notifHandler.subscriptionLocation, notifHandler.userId, notifHandler.notificationUri, patch)
				if err != nil {
					log.Printf("callback returned: %s", err.Error())
					continue
				}

			case <-notifHandler.ctx.Done():
				return
			}
		}
	}()

	return true
}

func (notifHandler *NotificationHandler) Stop() bool {
	if notifHandler.cancelFunc != nil {
		notifHandler.cancelFunc()
		return true
	}
	return false
}

func (notifHandler *NotificationHandler) SetEventCallback(callback callbackFun) {
	if callback == nil {
		log.Printf("NotificationHandler: callback is nil")
		return
	}

	notifHandler.callback = callback
	log.Printf("NotificationHandler: callback set for %s", notifHandler.userId)
}
