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
 * 3gpp-as-session-with-qos
 *
 * API for setting us an AS session with required QoS.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.3.0-alpha.5
 */

package northbound

import (
	"context"
	"errors"
	"net/http"

	models "gitlab.eurecom.fr/open-exposure/nef/as-session-with-qos/internal/models"
	"gitlab.eurecom.fr/open-exposure/nef/as-session-with-qos/internal/northbound/service"
)

type serverCtx interface {
	Service() *service.Service
}

type ASSessionWithRequiredQoSSubscriptionsAPIService struct {
	serverCtx
}
type IndividualASSessionWithRequiredQoSSubscriptionAPIService struct {
	serverCtx
}

// NewASSessionWithRequiredQoSSubscriptionsAPIService creates a default api service
func NewASSessionWithRequiredQoSSubscriptionsAPIService(srv serverCtx) *ASSessionWithRequiredQoSSubscriptionsAPIService {
	return &ASSessionWithRequiredQoSSubscriptionsAPIService{serverCtx: srv}
}

// NewIndividualASSessionWithRequiredQoSSubscriptionAPIService creates a default api service
func NewIndividualASSessionWithRequiredQoSSubscriptionAPIService(srv serverCtx) *IndividualASSessionWithRequiredQoSSubscriptionAPIService {
	return &IndividualASSessionWithRequiredQoSSubscriptionAPIService{serverCtx: srv}
}

// FetchAllASSessionWithQoSSubscriptions - Read all or queried active subscriptions for the SCS/AS.
func (s *ASSessionWithRequiredQoSSubscriptionsAPIService) FetchAllASSessionWithQoSSubscriptions(ctx context.Context, scsAsId string, ipAddrs []string, ipDomain string, macAddrs []string) (models.ImplResponse, error) {
	return models.Response(http.StatusNotImplemented, nil, ""), errors.New("FetchAllASSessionWithQoSSubscriptions method not implemented")
}

// CreateASSessionWithQoSSubscription - Creates a new subscription resource.
func (s *ASSessionWithRequiredQoSSubscriptionsAPIService) CreateASSessionWithQoSSubscription(ctx context.Context, scsAsId string, asSessionWithQoSSubscription models.AsSessionWithQoSSubscription) (models.ImplResponse, error) {
	loc, status, err := s.Service().PostSessionWithQoSSubscription(scsAsId, &asSessionWithQoSSubscription)

	if err != nil {
		return models.Response(status, nil, ""), errors.New("Error: " + err.Error())
	}

	return models.Response(status, asSessionWithQoSSubscription, loc), nil
}

// FetchIndASSessionWithQoSSubscription - Read an active subscriptions for the SCS/AS and the subscription Id.
func (s *IndividualASSessionWithRequiredQoSSubscriptionAPIService) FetchIndASSessionWithQoSSubscription(ctx context.Context, scsAsId string, subscriptionId string) (models.ImplResponse, error) {
	sub, status, err := s.Service().FetchSessionWithQoSSubscription(scsAsId, subscriptionId)
	if err != nil {
		return models.Response(status, nil, ""), errors.New("Error: " + err.Error())
	}
	return models.Response(status, sub, ""), nil
}

// UpdateIndASSessionWithQoSSubscription - Updates/replaces an existing subscription resource.
func (s *IndividualASSessionWithRequiredQoSSubscriptionAPIService) UpdateIndASSessionWithQoSSubscription(ctx context.Context, scsAsId string, subscriptionId string, asSessionWithQoSSubscription models.AsSessionWithQoSSubscription) (models.ImplResponse, error) {

	return models.Response(http.StatusNotImplemented, nil, ""), errors.New("UpdateIndASSessionWithQoSSubscription method not implemented")
}

// DeleteIndASSessionWithQoSSubscription - Deletes an already existing subscription.
func (s *IndividualASSessionWithRequiredQoSSubscriptionAPIService) DeleteIndASSessionWithQoSSubscription(ctx context.Context, scsAsId string, subscriptionId string) (models.ImplResponse, error) {
	status, err := s.Service().DeleteSessionWithQoSSubscription(scsAsId, subscriptionId)
	return models.Response(status, nil, ""), err
}

// ModifyIndASSessionWithQoSSubscription - Updates/replaces an existing subscription resource.
func (s *IndividualASSessionWithRequiredQoSSubscriptionAPIService) ModifyIndASSessionWithQoSSubscription(ctx context.Context, scsAsId string, subscriptionId string, asSessionWithQoSSubscriptionPatch models.AsSessionWithQoSSubscriptionPatch) (models.ImplResponse, error) {

	return models.Response(http.StatusNotImplemented, nil, ""), errors.New("ModifyIndASSessionWithQoSSubscription method not implemented")
}
