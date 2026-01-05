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

/*
 * 3gpp-traffic-influence
 *
 * API for AF traffic influence   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.3.0-alpha.4
 */

package nbi

import (
	"context"
	"errors"
	"net/http"

	models "gitlab.eurecom.fr/open-exposure/nef/traffic-influence/internal/models"
	"gitlab.eurecom.fr/open-exposure/nef/traffic-influence/internal/nbi/service"
)

type serverCtx interface {
	Service() *service.Service
}

// TrafficInfluenceSubscriptionAPIService is a service that implements the logic for the TrafficInfluenceSubscriptionAPIServicer
// This service should implement the business logic for every endpoint for the TrafficInfluenceSubscriptionAPI API.
// Include any external packages or services that will be required by this service.
type TrafficInfluenceSubscriptionAPIService struct {
	serverCtx
}

type IndividualTrafficInfluenceSubscriptionAPIService struct {
	serverCtx
}

// NewTrafficInfluenceSubscriptionAPIService creates a default api service
func NewTrafficInfluenceSubscriptionAPIService(srv serverCtx) *TrafficInfluenceSubscriptionAPIService {
	return &TrafficInfluenceSubscriptionAPIService{serverCtx: srv}
}

// NewIndividualTrafficInfluenceSubscriptionAPIService creates a default api service
func NewIndividualTrafficInfluenceSubscriptionAPIService(srv serverCtx) *IndividualTrafficInfluenceSubscriptionAPIService {
	return &IndividualTrafficInfluenceSubscriptionAPIService{serverCtx: srv}
}

// ReadAllSubscriptions - read all of the active subscriptions for the AF
func (s *TrafficInfluenceSubscriptionAPIService) ReadAllSubscriptions(ctx context.Context, afId string) (models.ImplResponse, error) {
	subs, status, err := s.Service().GetAllTrafficInfluenceSub(afId)
	return models.Response(status, subs, ""), err
}

// ReadAnSubscription - read an active subscriptions for the SCS/AS and the subscription Id
func (s *IndividualTrafficInfluenceSubscriptionAPIService) ReadAnSubscription(ctx context.Context, afId string, subscriptionId string) (models.ImplResponse, error) {
	sub, status, err := s.Service().GetIndividualTrafficInfluenceSub(afId, subscriptionId)
	return models.Response(status, sub, ""), err
}

// TODO CreateNewSubscription - Creates a new subscription resource
func (s *TrafficInfluenceSubscriptionAPIService) CreateNewSubscription(ctx context.Context, afId string, trafficInfluSub models.TrafficInfluSub) (models.ImplResponse, error) {
	loc, status, err := s.Service().CreateTrafficInfluenceSub(afId, &trafficInfluSub)
	if err != nil {
		return models.Response(status, nil, ""), errors.New("Error: " + err.Error())
	}
	return models.Response(status, trafficInfluSub, loc), nil
}

// DeleteAnSubscription - Deletes an already existing subscription
func (s *IndividualTrafficInfluenceSubscriptionAPIService) DeleteAnSubscription(ctx context.Context, afId string, subscriptionId string) (models.ImplResponse, error) {
	err := s.Service().DeleteTrafficInfluenceSub(afId, subscriptionId)
	if err == nil {
		return models.Response(http.StatusNoContent, nil, ""), nil
	}
	return models.Response(http.StatusNotFound, nil, ""), err
}

// TODO PartialUpdateAnSubscription - Partially updates/replaces an existing subscription resource
func (s *IndividualTrafficInfluenceSubscriptionAPIService) PartialUpdateAnSubscription(ctx context.Context, afId string, subscriptionId string, trafficInfluSubPatch models.TrafficInfluSubPatch) (models.ImplResponse, error) {
	return models.Response(http.StatusNotImplemented, nil, ""), errors.New("PartialUpdateAnSubscription method not implemented")
}

// FullyUpdateAnSubscription - Fully updates/replaces an existing subscription resource
func (s *IndividualTrafficInfluenceSubscriptionAPIService) FullyUpdateAnSubscription(ctx context.Context, afId string, subscriptionId string, trafficInfluSub models.TrafficInfluSub) (models.ImplResponse, error) {
	code, err := s.Service().UpdateTrafficInfluenceSub(afId, subscriptionId, &trafficInfluSub)
	if err == nil {
		return models.Response(code, nil, ""), nil
	}
	return models.Response(code, nil, ""), err
}
