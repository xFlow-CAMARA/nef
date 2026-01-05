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
	"fmt"
	"sync"

	"github.com/google/uuid"
	"gitlab.eurecom.fr/open-exposure/nef/as-session-with-qos/internal/models"
	"golang.org/x/exp/maps"
)

type AppFunctionCtx struct {
	afId    string
	subs    map[string]*AfSubscriptionCtx
	Mu      sync.RWMutex
	lastId  uint64
	TransId uint64
}

func NewAf(afId string) *AppFunctionCtx {
	return &AppFunctionCtx{
		afId:   afId,
		subs:   make(map[string]*AfSubscriptionCtx),
		lastId: 0,
		Mu:     sync.RWMutex{},
	}
}

func (afCtx *AppFunctionCtx) GetAfSubscriptions() []*AfSubscriptionCtx {
	return maps.Values(afCtx.subs)
}

func (afCtx *AppFunctionCtx) GetAfSubscription(subId string) *AfSubscriptionCtx {
	trInfSub, ok := afCtx.subs[subId]
	if ok {
		return trInfSub
	} else {
		return nil
	}
}

func (afCtx *AppFunctionCtx) NewAfSubscription(data *models.AsSessionWithQoSSubscription) (string, *AfSubscriptionCtx) {

	subId := afCtx.newSubscriptionId()
	loc := createAsSessionWithQosLocation(afCtx.afId, subId)
	sub := &AfSubscriptionCtx{
		subId: subId,
		Data:  data,
		loc:   loc,
	}
	data.Self = subId
	afCtx.subs[subId] = sub
	return loc, sub
}

func (afCtx *AppFunctionCtx) DeleteAfscription(subId string) error {
	_, ok := afCtx.subs[subId]
	if ok {
		delete(afCtx.subs, subId)
		return nil
	} else {
		return fmt.Errorf("could not find traffic influence subscription")
	}
}

func (afCtx *AppFunctionCtx) newSubscriptionId() string {
	return uuid.NewString()
}

func createAsSessionWithQosLocation(afId string, subId string) string {
	return "/3gpp-as-session-with-qos/v1/" + afId + "/subscriptions/" + subId
}
