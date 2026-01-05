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
 * 3gpp-monitoring-event
 *
 * API for Monitoring Event.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.3.0-alpha.5
 */

package nbi

import (
	"context"
	"net/http"

	"log"

	"gitlab.eurecom.fr/open-exposure/nef/monitoring-event/internal/models"
	"gitlab.eurecom.fr/open-exposure/nef/monitoring-event/internal/nbi/service"

	"errors"
)

type serverCtx interface {
	Service() *service.Service
}

type MonitoringEventSubscriptionsAPIService struct {
	serverCtx
}

type IndividualMonitoringEventSubscriptionAPIService struct {
	serverCtx
}

func NewMonitoringEventSubscriptionsAPIService(srv serverCtx) *MonitoringEventSubscriptionsAPIService {
	return &MonitoringEventSubscriptionsAPIService{serverCtx: srv}
}

func NewIndividualMonitoringEventSubscriptionAPIService(srv serverCtx) *IndividualMonitoringEventSubscriptionAPIService {
	return &IndividualMonitoringEventSubscriptionAPIService{serverCtx: srv}
}

/* =========== collective endpoints ============ */

// FetchAllMonitoringEventSubscriptions - Read all or queried active subscriptions for the SCS/AS.
func (s *MonitoringEventSubscriptionsAPIService) FetchAllMonitoringEventSubscriptions(ctx context.Context, scsAsId string, ipAddrs []string, ipDomain string, macAddrs []string) (models.ImplResponse, error) {
	data, code, err := s.Service().GetAllMonitoringEventSubscriptions(scsAsId)
	if err == nil {
		log.Printf("FetchAllMonitoringEventSubscriptions: fetched %d subscriptions for %s", len(data), scsAsId)
		return models.Response(code, data), nil
	} else {
		log.Printf("FetchAllMonitoringEventSubscriptions: error fetching subscriptions for %s: %s", scsAsId, err.Error())
		return models.Response(code, models.ProblemDetails{
			Title:  "Subscription Fetch Error",
			Detail: err.Error(),
			Status: int32(code),
		}), nil
	}
}

// CreateMonitoringEventSubscription - Creates a new subscription resource for monitoring event notification.
func (s *MonitoringEventSubscriptionsAPIService) CreateMonitoringEventSubscription(ctx context.Context, scsAsId string, monitoringEventSubscription *models.MonitoringEventSubscription) (models.ImplResponse, error) {
	/* Define a report object to pass to PostMonitoringEventSub method, it will be filled with the current status of the UE*/
	immediateReport := &models.MonitoringEventReport{}

	loc, code, err := s.Service().PostMonitoringEventSubscription(scsAsId, monitoringEventSubscription, immediateReport)
	if err == nil {
		log.Printf("CreateMonitoringEventSubscription: created subscription for %s at %s", scsAsId, loc)
		return models.ResponseWithLocation(code, immediateReport, loc), nil
	} else {
		log.Printf("CreateMonitoringEventSubscription: error creating subscription for %s: %s", scsAsId, err.Error())
		return models.Response(code, models.ProblemDetails{
			Title:  "Subscription Creation Error",
			Detail: err.Error(),
			Status: int32(code),
		}), nil
	}
}

/* =========== individual endpoints ============ */

// FetchIndMonitoringEventSubscription - Read an active subscriptions for the SCS/AS and the subscription Id.
func (s *IndividualMonitoringEventSubscriptionAPIService) FetchIndMonitoringEventSubscription(ctx context.Context, scsAsId string, subscriptionId string) (models.ImplResponse, error) {
	data, code, err := s.Service().GetMonitoringEventSubscription(scsAsId, subscriptionId)
	if err == nil {
		log.Printf("FetchIndMonitoringEventSubscription: fetched subscription %s for %s", subscriptionId, scsAsId)
		return models.Response(code, data), nil
	} else {
		log.Printf("FetchIndMonitoringEventSubscription: error fetching subscription %s for %s: %s", subscriptionId, scsAsId, err.Error())
		return models.Response(code, models.ProblemDetails{
			Title:  "Subscription Fetch Error",
			Detail: err.Error(),
			Status: int32(code),
		}), nil
	}
}

// UpdateIndMonitoringEventSubscription - Updates/replaces an existing subscription resource.
func (s *IndividualMonitoringEventSubscriptionAPIService) UpdateIndMonitoringEventSubscription(ctx context.Context, scsAsId string, subscriptionId string, monitoringEventSubscription *models.MonitoringEventSubscription) (models.ImplResponse, error) {
	/* TODO: use models.ProblemDetails{} */
	return models.Response(http.StatusNotImplemented, nil), errors.New("UpdateIndMonitoringEventSubscription method not implemented")
}

// DeleteIndMonitoringEventSubscription - Deletes an already existing monitoring event subscription.
func (s *IndividualMonitoringEventSubscriptionAPIService) DeleteIndMonitoringEventSubscription(ctx context.Context, scsAsId string, subscriptionId string) (models.ImplResponse, error) {
	code, err := s.Service().DeleteMonitoringEventSubscription(scsAsId, subscriptionId)
	if err == nil {
		log.Printf("DeleteIndMonitoringEventSubscription: deleted subscription %s for %s", subscriptionId, scsAsId)
		return models.Response(code, nil), nil
	} else {
		log.Printf("DeleteIndMonitoringEventSubscription: error deleting subscription %s for %s: %s", subscriptionId, scsAsId, err.Error())
		return models.Response(code, models.ProblemDetails{
			Title:  "Subscription Deletion Error",
			Detail: err.Error(),
			Status: int32(code),
		}), nil
	}
}

// ModifyIndMonitoringEventSubscription - Modifies an existing subscription of monitoring event.
func (s *IndividualMonitoringEventSubscriptionAPIService) ModifyIndMonitoringEventSubscription(ctx context.Context, scsAsId string, subscriptionId string, patchItem []models.PatchItem) (models.ImplResponse, error) {
	/* TODO: use models.ProblemDetails{} */
	return models.Response(http.StatusNotImplemented, nil), errors.New("ModifyIndMonitoringEventSubscription method not implemented")
}
