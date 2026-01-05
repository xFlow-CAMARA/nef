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
	"encoding/json"
	"net/http"
	"strings"

	models "gitlab.eurecom.fr/open-exposure/nef/traffic-influence/internal/models"

	"github.com/gorilla/mux"
)

// TrafficInfluenceSubscriptionAPIController binds http requests to an api service and writes the service results to the http response
type TrafficInfluenceSubscriptionAPIController struct {
	service      TrafficInfluenceSubscriptionAPIServicer
	errorHandler models.ErrorHandler
}

// TrafficInfluenceSubscriptionAPIOption for how the controller is set up.
type TrafficInfluenceSubscriptionAPIOption func(*TrafficInfluenceSubscriptionAPIController)

// WithTrafficInfluenceSubscriptionAPIErrorHandler inject models.ErrorHandler into controller
func WithTrafficInfluenceSubscriptionAPIErrorHandler(h models.ErrorHandler) TrafficInfluenceSubscriptionAPIOption {
	return func(c *TrafficInfluenceSubscriptionAPIController) {
		c.errorHandler = h
	}
}

// NewTrafficInfluenceSubscriptionAPIController creates a default api controller
func NewTrafficInfluenceSubscriptionAPIController(s TrafficInfluenceSubscriptionAPIServicer, opts ...TrafficInfluenceSubscriptionAPIOption) *TrafficInfluenceSubscriptionAPIController {
	controller := &TrafficInfluenceSubscriptionAPIController{
		service:      s,
		errorHandler: models.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// models.Route returns all the api routes for the TrafficInfluenceSubscriptionAPIController
func (c *TrafficInfluenceSubscriptionAPIController) Routes() models.Routes {
	return models.Routes{
		"ReadAllSubscriptions": models.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "/3gpp-traffic-influence/v1/{afId}/subscriptions",
			HandlerFunc: c.ReadAllSubscriptions,
		},
		"CreateNewSubscription": models.Route{
			Method:      strings.ToUpper("Post"),
			Pattern:     "/3gpp-traffic-influence/v1/{afId}/subscriptions",
			HandlerFunc: c.CreateNewSubscription,
		},
	}
}

// ReadAllSubscriptions - read all of the active subscriptions for the AF
func (c *TrafficInfluenceSubscriptionAPIController) ReadAllSubscriptions(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	afIdParam := params["afId"]
	if afIdParam == "" {
		c.errorHandler(w, r, &models.RequiredError{Field: "afId"}, nil)
		return
	}
	result, err := c.service.ReadAllSubscriptions(r.Context(), afIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = models.EncodeJSONResponse(result.Body, &result.Code, w)
}

// CreateNewSubscription - Creates a new subscription resource
func (c *TrafficInfluenceSubscriptionAPIController) CreateNewSubscription(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	afIdParam := params["afId"]
	if afIdParam == "" {
		c.errorHandler(w, r, &models.RequiredError{Field: "afId"}, nil)
		return
	}
	trafficInfluSubParam := models.TrafficInfluSub{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&trafficInfluSubParam); err != nil {
		c.errorHandler(w, r, &models.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertTrafficInfluSubRequired(trafficInfluSubParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertTrafficInfluSubConstraints(trafficInfluSubParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateNewSubscription(r.Context(), afIdParam, trafficInfluSubParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	w.Header().Add("Location", result.Location)
	// If no error, encode the body and the result code
	_ = models.EncodeJSONResponse(result.Body, &result.Code, w)
}

// IndividualTrafficInfluenceSubscriptionAPIController binds http requests to an api service and writes the service results to the http response
type IndividualTrafficInfluenceSubscriptionAPIController struct {
	service      IndividualTrafficInfluenceSubscriptionAPIServicer
	errorHandler models.ErrorHandler
}

// IndividualTrafficInfluenceSubscriptionAPIOption for how the controller is set up.
type IndividualTrafficInfluenceSubscriptionAPIOption func(*IndividualTrafficInfluenceSubscriptionAPIController)

// WithIndividualTrafficInfluenceSubscriptionAPIErrorHandler inject models.ErrorHandler into controller
func WithIndividualTrafficInfluenceSubscriptionAPIErrorHandler(h models.ErrorHandler) IndividualTrafficInfluenceSubscriptionAPIOption {
	return func(c *IndividualTrafficInfluenceSubscriptionAPIController) {
		c.errorHandler = h
	}
}

// NewIndividualTrafficInfluenceSubscriptionAPIController creates a default api controller
func NewIndividualTrafficInfluenceSubscriptionAPIController(s IndividualTrafficInfluenceSubscriptionAPIServicer, opts ...IndividualTrafficInfluenceSubscriptionAPIOption) *IndividualTrafficInfluenceSubscriptionAPIController {
	controller := &IndividualTrafficInfluenceSubscriptionAPIController{
		service:      s,
		errorHandler: models.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the IndividualTrafficInfluenceSubscriptionAPIController
func (c *IndividualTrafficInfluenceSubscriptionAPIController) Routes() models.Routes {
	return models.Routes{
		"ReadAnSubscription": models.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "/3gpp-traffic-influence/v1/{afId}/subscriptions/{subscriptionId}",
			HandlerFunc: c.ReadAnSubscription,
		},
		"FullyUpdateAnSubscription": models.Route{
			Method:      strings.ToUpper("Put"),
			Pattern:     "/3gpp-traffic-influence/v1/{afId}/subscriptions/{subscriptionId}",
			HandlerFunc: c.FullyUpdateAnSubscription,
		},
		"DeleteAnSubscription": models.Route{
			Method:      strings.ToUpper("Delete"),
			Pattern:     "/3gpp-traffic-influence/v1/{afId}/subscriptions/{subscriptionId}",
			HandlerFunc: c.DeleteAnSubscription,
		},
		"PartialUpdateAnSubscription": models.Route{
			Method:      strings.ToUpper("Patch"),
			Pattern:     "/3gpp-traffic-influence/v1/{afId}/subscriptions/{subscriptionId}",
			HandlerFunc: c.PartialUpdateAnSubscription,
		},
	}
}

// ReadAnSubscription - read an active subscriptions for the SCS/AS and the subscription Id
func (c *IndividualTrafficInfluenceSubscriptionAPIController) ReadAnSubscription(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	afIdParam := params["afId"]
	if afIdParam == "" {
		c.errorHandler(w, r, &models.RequiredError{Field: "afId"}, nil)
		return
	}
	subscriptionIdParam := params["subscriptionId"]
	if subscriptionIdParam == "" {
		c.errorHandler(w, r, &models.RequiredError{Field: "subscriptionId"}, nil)
		return
	}
	result, err := c.service.ReadAnSubscription(r.Context(), afIdParam, subscriptionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = models.EncodeJSONResponse(result.Body, &result.Code, w)
}

// FullyUpdateAnSubscription - Fully updates/replaces an existing subscription resource
func (c *IndividualTrafficInfluenceSubscriptionAPIController) FullyUpdateAnSubscription(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	afIdParam := params["afId"]
	if afIdParam == "" {
		c.errorHandler(w, r, &models.RequiredError{Field: "afId"}, nil)
		return
	}
	subscriptionIdParam := params["subscriptionId"]
	if subscriptionIdParam == "" {
		c.errorHandler(w, r, &models.RequiredError{Field: "subscriptionId"}, nil)
		return
	}
	trafficInfluSubParam := models.TrafficInfluSub{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&trafficInfluSubParam); err != nil {
		c.errorHandler(w, r, &models.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertTrafficInfluSubRequired(trafficInfluSubParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertTrafficInfluSubConstraints(trafficInfluSubParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.FullyUpdateAnSubscription(r.Context(), afIdParam, subscriptionIdParam, trafficInfluSubParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = models.EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteAnSubscription - Deletes an already existing subscription
func (c *IndividualTrafficInfluenceSubscriptionAPIController) DeleteAnSubscription(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	afIdParam := params["afId"]
	if afIdParam == "" {
		c.errorHandler(w, r, &models.RequiredError{Field: "afId"}, nil)
		return
	}
	subscriptionIdParam := params["subscriptionId"]
	if subscriptionIdParam == "" {
		c.errorHandler(w, r, &models.RequiredError{Field: "subscriptionId"}, nil)
		return
	}
	result, err := c.service.DeleteAnSubscription(r.Context(), afIdParam, subscriptionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = models.EncodeJSONResponse(result.Body, &result.Code, w)
}

// PartialUpdateAnSubscription - Partially updates/replaces an existing subscription resource
func (c *IndividualTrafficInfluenceSubscriptionAPIController) PartialUpdateAnSubscription(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	afIdParam := params["afId"]
	if afIdParam == "" {
		c.errorHandler(w, r, &models.RequiredError{Field: "afId"}, nil)
		return
	}
	subscriptionIdParam := params["subscriptionId"]
	if subscriptionIdParam == "" {
		c.errorHandler(w, r, &models.RequiredError{Field: "subscriptionId"}, nil)
		return
	}
	trafficInfluSubPatchParam := models.TrafficInfluSubPatch{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&trafficInfluSubPatchParam); err != nil {
		c.errorHandler(w, r, &models.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertTrafficInfluSubPatchRequired(trafficInfluSubPatchParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertTrafficInfluSubPatchConstraints(trafficInfluSubPatchParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PartialUpdateAnSubscription(r.Context(), afIdParam, subscriptionIdParam, trafficInfluSubPatchParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = models.EncodeJSONResponse(result.Body, &result.Code, w)
}
