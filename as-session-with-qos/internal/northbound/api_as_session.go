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
	"net/http"

	models "gitlab.eurecom.fr/open-exposure/nef/as-session-with-qos/internal/models"
)

// ASSessionWithRequiredQoSSubscriptionsAPIRouter defines the required methods for binding the api requests to a responses for the ASSessionWithRequiredQoSSubscriptionsAPI
// The ASSessionWithRequiredQoSSubscriptionsAPIRouter implementation should parse necessary information from the http request,
// pass the data to a ASSessionWithRequiredQoSSubscriptionsAPIServicer to perform the required actions, then write the service results to the http response.
type ASSessionWithRequiredQoSSubscriptionsAPIRouter interface {
	FetchAllASSessionWithQoSSubscriptions(http.ResponseWriter, *http.Request)
	CreateASSessionWithQoSSubscription(http.ResponseWriter, *http.Request)
}

// IndividualASSessionWithRequiredQoSSubscriptionAPIRouter defines the required methods for binding the api requests to a responses for the IndividualASSessionWithRequiredQoSSubscriptionAPI
// The IndividualASSessionWithRequiredQoSSubscriptionAPIRouter implementation should parse necessary information from the http request,
// pass the data to a IndividualASSessionWithRequiredQoSSubscriptionAPIServicer to perform the required actions, then write the service results to the http response.
type IndividualASSessionWithRequiredQoSSubscriptionAPIRouter interface {
	FetchIndASSessionWithQoSSubscription(http.ResponseWriter, *http.Request)
	UpdateIndASSessionWithQoSSubscription(http.ResponseWriter, *http.Request)
	DeleteIndASSessionWithQoSSubscription(http.ResponseWriter, *http.Request)
	ModifyIndASSessionWithQoSSubscription(http.ResponseWriter, *http.Request)
}

// ASSessionWithRequiredQoSSubscriptionsAPIServicer defines the api actions for the ASSessionWithRequiredQoSSubscriptionsAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ASSessionWithRequiredQoSSubscriptionsAPIServicer interface {
	FetchAllASSessionWithQoSSubscriptions(context.Context, string, []string, string, []string) (models.ImplResponse, error)
	CreateASSessionWithQoSSubscription(context.Context, string, models.AsSessionWithQoSSubscription) (models.ImplResponse, error)
}

// IndividualASSessionWithRequiredQoSSubscriptionAPIServicer defines the api actions for the IndividualASSessionWithRequiredQoSSubscriptionAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type IndividualASSessionWithRequiredQoSSubscriptionAPIServicer interface {
	FetchIndASSessionWithQoSSubscription(context.Context, string, string) (models.ImplResponse, error)
	UpdateIndASSessionWithQoSSubscription(context.Context, string, string, models.AsSessionWithQoSSubscription) (models.ImplResponse, error)
	DeleteIndASSessionWithQoSSubscription(context.Context, string, string) (models.ImplResponse, error)
	ModifyIndASSessionWithQoSSubscription(context.Context, string, string, models.AsSessionWithQoSSubscriptionPatch) (models.ImplResponse, error)
}
