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
	"net/http"

	models "gitlab.eurecom.fr/open-exposure/nef/traffic-influence/internal/models"
)

// IndividualTrafficInfluenceSubscriptionAPIRouter defines the required methods for binding the api requests to a responses for the IndividualTrafficInfluenceSubscriptionAPI
// The IndividualTrafficInfluenceSubscriptionAPIRouter implementation should parse necessary information from the http request,
// pass the data to a IndividualTrafficInfluenceSubscriptionAPIServicer to perform the required actions, then write the service results to the http response.
type IndividualTrafficInfluenceSubscriptionAPIRouter interface {
	ReadAnSubscription(http.ResponseWriter, *http.Request)
	FullyUpdateAnSubscription(http.ResponseWriter, *http.Request)
	DeleteAnSubscription(http.ResponseWriter, *http.Request)
	PartialUpdateAnSubscription(http.ResponseWriter, *http.Request)
}

// TrafficInfluenceSubscriptionAPIRouter defines the required methods for binding the api requests to a responses for the TrafficInfluenceSubscriptionAPI
// The TrafficInfluenceSubscriptionAPIRouter implementation should parse necessary information from the http request,
// pass the data to a TrafficInfluenceSubscriptionAPIServicer to perform the required actions, then write the service results to the http response.
type TrafficInfluenceSubscriptionAPIRouter interface {
	ReadAllSubscriptions(http.ResponseWriter, *http.Request)
	CreateNewSubscription(http.ResponseWriter, *http.Request)
}

// IndividualTrafficInfluenceSubscriptionAPIServicer defines the api actions for the IndividualTrafficInfluenceSubscriptionAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type IndividualTrafficInfluenceSubscriptionAPIServicer interface {
	ReadAnSubscription(context.Context, string, string) (models.ImplResponse, error)
	FullyUpdateAnSubscription(context.Context, string, string, models.TrafficInfluSub) (models.ImplResponse, error)
	DeleteAnSubscription(context.Context, string, string) (models.ImplResponse, error)
	PartialUpdateAnSubscription(context.Context, string, string, models.TrafficInfluSubPatch) (models.ImplResponse, error)
}

// TrafficInfluenceSubscriptionAPIServicer defines the api actions for the TrafficInfluenceSubscriptionAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type TrafficInfluenceSubscriptionAPIServicer interface {
	ReadAllSubscriptions(context.Context, string) (models.ImplResponse, error)
	CreateNewSubscription(context.Context, string, models.TrafficInfluSub) (models.ImplResponse, error)
}
