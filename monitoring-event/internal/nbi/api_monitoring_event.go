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

	"gitlab.eurecom.fr/open-exposure/nef/monitoring-event/internal/models"
)

// IndividualMonitoringEventSubscriptionAPIRouter defines the required methods for binding the api requests to a responses for the IndividualMonitoringEventSubscriptionAPI
// The IndividualMonitoringEventSubscriptionAPIRouter implementation should parse necessary information from the http request,
// pass the data to a IndividualMonitoringEventSubscriptionAPIServicer to perform the required actions, then write the service results to the http response.
type IndividualMonitoringEventSubscriptionAPIRouter interface {
	FetchIndMonitoringEventSubscription(http.ResponseWriter, *http.Request)
	UpdateIndMonitoringEventSubscription(http.ResponseWriter, *http.Request)
	DeleteIndMonitoringEventSubscription(http.ResponseWriter, *http.Request)
	ModifyIndMonitoringEventSubscription(http.ResponseWriter, *http.Request)
}

// MonitoringEventSubscriptionsAPIRouter defines the required methods for binding the api requests to a responses for the MonitoringEventSubscriptionsAPI
// The MonitoringEventSubscriptionsAPIRouter implementation should parse necessary information from the http request,
// pass the data to a MonitoringEventSubscriptionsAPIServicer to perform the required actions, then write the service results to the http response.
type MonitoringEventSubscriptionsAPIRouter interface {
	FetchAllMonitoringEventSubscriptions(http.ResponseWriter, *http.Request)
	CreateMonitoringEventSubscription(http.ResponseWriter, *http.Request)
}

// IndividualMonitoringEventSubscriptionAPIServicer defines the api actions for the IndividualMonitoringEventSubscriptionAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type IndividualMonitoringEventSubscriptionAPIServicer interface {
	FetchIndMonitoringEventSubscription(context.Context, string, string) (models.ImplResponse, error)
	UpdateIndMonitoringEventSubscription(context.Context, string, string, *models.MonitoringEventSubscription) (models.ImplResponse, error)
	DeleteIndMonitoringEventSubscription(context.Context, string, string) (models.ImplResponse, error)
	ModifyIndMonitoringEventSubscription(context.Context, string, string, []models.PatchItem) (models.ImplResponse, error)
}

// MonitoringEventSubscriptionsAPIServicer defines the api actions for the MonitoringEventSubscriptionsAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type MonitoringEventSubscriptionsAPIServicer interface {
	FetchAllMonitoringEventSubscriptions(context.Context, string, []string, string, []string) (models.ImplResponse, error)
	CreateMonitoringEventSubscription(context.Context, string, *models.MonitoringEventSubscription) (models.ImplResponse, error)
}
