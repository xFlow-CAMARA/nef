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

package contexts

import (
	"gitlab.eurecom.fr/open-exposure/nef/monitoring-event/internal/handlers"
	"gitlab.eurecom.fr/open-exposure/nef/monitoring-event/internal/models"
)

type AfSubscriptionCtx struct {
	subId string
	data  *models.MonitoringEventSubscription //replace with subscription model
	loc   string

	/*add custom data here*/
	notifUri     string
	eventType    models.MonitoringType
	notifHandler *handlers.NotificationHandler
	supi         string
}

func NewAfSubscriptionCtx(subId string, loc string, supi string, data *models.MonitoringEventSubscription) *AfSubscriptionCtx {
	/* derive other parameters from the subscription data */

	if len(supi) == 0 || len(loc) == 0 || data == nil || len(data.NotificationDestination) == 0 || len(data.MonitoringType) == 0 {
		return nil
	}

	data.Self = subId

	return &AfSubscriptionCtx{
		subId:        subId,
		data:         data,
		supi:         supi,
		loc:          loc,
		notifUri:     data.NotificationDestination,
		eventType:    data.MonitoringType,
		notifHandler: nil,
	}
}

func (subCtx *AfSubscriptionCtx) SetNotificationHandler(handler *handlers.NotificationHandler) {
	subCtx.notifHandler = handler
}

func (subCtx *AfSubscriptionCtx) GetNotificationHandler() *handlers.NotificationHandler {
	return subCtx.notifHandler
}

func (subCtx *AfSubscriptionCtx) GetEventType() models.MonitoringType {
	return subCtx.eventType
}

func (subCtx *AfSubscriptionCtx) GetNotificationUri() string {
	return subCtx.notifUri
}

func (subCtx *AfSubscriptionCtx) GetLocation() string {
	return subCtx.loc
}

func (subCtx *AfSubscriptionCtx) GetSubscriptionData() models.MonitoringEventSubscription {
	/*return a copy of the subscription data*/
	if subCtx.data == nil {
		return models.MonitoringEventSubscription{}
	} else {
		return *subCtx.data
	}
}
