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
	"encoding/json"
	"log"
	"net/http"
	"strings"

	models "gitlab.eurecom.fr/open-exposure/nef/as-session-with-qos/internal/models"

	"github.com/gorilla/mux"
)

// ASSessionWithRequiredQoSSubscriptionsAPIController binds http requests to an api service and writes the service results to the http response
type ASSessionWithRequiredQoSSubscriptionsAPIController struct {
	service      ASSessionWithRequiredQoSSubscriptionsAPIServicer
	errorHandler models.ErrorHandler
}

// ASSessionWithRequiredQoSSubscriptionsAPIOption for how the controller is set up.
type ASSessionWithRequiredQoSSubscriptionsAPIOption func(*ASSessionWithRequiredQoSSubscriptionsAPIController)

// WithASSessionWithRequiredQoSSubscriptionsAPIErrorHandler inject ErrorHandler into controller
func WithASSessionWithRequiredQoSSubscriptionsAPIErrorHandler(h models.ErrorHandler) ASSessionWithRequiredQoSSubscriptionsAPIOption {
	return func(c *ASSessionWithRequiredQoSSubscriptionsAPIController) {
		c.errorHandler = h
	}
}

// NewASSessionWithRequiredQoSSubscriptionsAPIController creates a default api controller
func NewASSessionWithRequiredQoSSubscriptionsAPIController(s ASSessionWithRequiredQoSSubscriptionsAPIServicer, opts ...ASSessionWithRequiredQoSSubscriptionsAPIOption) *ASSessionWithRequiredQoSSubscriptionsAPIController {
	controller := &ASSessionWithRequiredQoSSubscriptionsAPIController{
		service:      s,
		errorHandler: models.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ASSessionWithRequiredQoSSubscriptionsAPIController
func (c *ASSessionWithRequiredQoSSubscriptionsAPIController) Routes() models.Routes {
	return models.Routes{
		"FetchAllASSessionWithQoSSubscriptions": models.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "/3gpp-as-session-with-qos/v1/{scsAsId}/subscriptions",
			HandlerFunc: c.FetchAllASSessionWithQoSSubscriptions,
		},
		"CreateASSessionWithQoSSubscription": models.Route{
			Method:      strings.ToUpper("Post"),
			Pattern:     "/3gpp-as-session-with-qos/v1/{scsAsId}/subscriptions",
			HandlerFunc: c.CreateASSessionWithQoSSubscription,
		},
	}
}

// FetchAllASSessionWithQoSSubscriptions - Read all or queried active subscriptions for the SCS/AS.
func (c *ASSessionWithRequiredQoSSubscriptionsAPIController) FetchAllASSessionWithQoSSubscriptions(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query, err := models.ParseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &models.ParsingError{Err: err}, nil)
		return
	}
	scsAsIdParam := params["scsAsId"]
	if scsAsIdParam == "" {
		c.errorHandler(w, r, &models.RequiredError{Field: "scsAsId"}, nil)
		return
	}
	var ipAddrsParam []string
	if query.Has("ip-addrs") {
		ipAddrsParam = strings.Split(query.Get("ip-addrs"), ",")
	}
	var ipDomainParam string
	if query.Has("ip-domain") {
		param := query.Get("ip-domain")

		ipDomainParam = param
	}
	var macAddrsParam []string
	if query.Has("mac-addrs") {
		macAddrsParam = strings.Split(query.Get("mac-addrs"), ",")
	}
	result, err := c.service.FetchAllASSessionWithQoSSubscriptions(r.Context(), scsAsIdParam, ipAddrsParam, ipDomainParam, macAddrsParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = models.EncodeJSONResponse(result.Body, &result.Code, w)
}

// CreateASSessionWithQoSSubscription - Creates a new subscription resource.
func (c *ASSessionWithRequiredQoSSubscriptionsAPIController) CreateASSessionWithQoSSubscription(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	scsAsIdParam := params["scsAsId"]
	if scsAsIdParam == "" {
		c.errorHandler(w, r, &models.RequiredError{Field: "scsAsId"}, nil)
		return
	}
	asSessionWithQoSSubscriptionParam := models.AsSessionWithQoSSubscription{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&asSessionWithQoSSubscriptionParam); err != nil {
		log.Print("here")
		c.errorHandler(w, r, &models.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertAsSessionWithQoSSubscriptionRequired(asSessionWithQoSSubscriptionParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertAsSessionWithQoSSubscriptionConstraints(asSessionWithQoSSubscriptionParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateASSessionWithQoSSubscription(r.Context(), scsAsIdParam, asSessionWithQoSSubscriptionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	w.Header().Add("Location", result.Location)
	// If no error, encode the body and the result code
	_ = models.EncodeJSONResponse(result.Body, &result.Code, w)
}

// IndividualASSessionWithRequiredQoSSubscriptionAPIController binds http requests to an api service and writes the service results to the http response
type IndividualASSessionWithRequiredQoSSubscriptionAPIController struct {
	service      IndividualASSessionWithRequiredQoSSubscriptionAPIServicer
	errorHandler models.ErrorHandler
}

// IndividualASSessionWithRequiredQoSSubscriptionAPIOption for how the controller is set up.
type IndividualASSessionWithRequiredQoSSubscriptionAPIOption func(*IndividualASSessionWithRequiredQoSSubscriptionAPIController)

// WithIndividualASSessionWithRequiredQoSSubscriptionAPIErrorHandler injectmodels.ErrorHandler into controller
func WithIndividualASSessionWithRequiredQoSSubscriptionAPIErrorHandler(h models.ErrorHandler) IndividualASSessionWithRequiredQoSSubscriptionAPIOption {
	return func(c *IndividualASSessionWithRequiredQoSSubscriptionAPIController) {
		c.errorHandler = h
	}
}

// NewIndividualASSessionWithRequiredQoSSubscriptionAPIController creates a default api controller
func NewIndividualASSessionWithRequiredQoSSubscriptionAPIController(s IndividualASSessionWithRequiredQoSSubscriptionAPIServicer, opts ...IndividualASSessionWithRequiredQoSSubscriptionAPIOption) *IndividualASSessionWithRequiredQoSSubscriptionAPIController {
	controller := &IndividualASSessionWithRequiredQoSSubscriptionAPIController{
		service:      s,
		errorHandler: models.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the IndividualASSessionWithRequiredQoSSubscriptionAPIController
func (c *IndividualASSessionWithRequiredQoSSubscriptionAPIController) Routes() models.Routes {
	return models.Routes{
		"FetchIndASSessionWithQoSSubscription": models.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "/3gpp-as-session-with-qos/v1/{scsAsId}/subscriptions/{subscriptionId}",
			HandlerFunc: c.FetchIndASSessionWithQoSSubscription,
		},
		"UpdateIndASSessionWithQoSSubscription": models.Route{
			Method:      strings.ToUpper("Put"),
			Pattern:     "/3gpp-as-session-with-qos/v1/{scsAsId}/subscriptions/{subscriptionId}",
			HandlerFunc: c.UpdateIndASSessionWithQoSSubscription,
		},
		"DeleteIndASSessionWithQoSSubscription": models.Route{
			Method:      strings.ToUpper("Delete"),
			Pattern:     "/3gpp-as-session-with-qos/v1/{scsAsId}/subscriptions/{subscriptionId}",
			HandlerFunc: c.DeleteIndASSessionWithQoSSubscription,
		},
		"ModifyIndASSessionWithQoSSubscription": models.Route{
			Method:      strings.ToUpper("Patch"),
			Pattern:     "/3gpp-as-session-with-qos/v1/{scsAsId}/subscriptions/{subscriptionId}",
			HandlerFunc: c.ModifyIndASSessionWithQoSSubscription,
		},
	}
}

// FetchIndASSessionWithQoSSubscription - Read an active subscriptions for the SCS/AS and the subscription Id.
func (c *IndividualASSessionWithRequiredQoSSubscriptionAPIController) FetchIndASSessionWithQoSSubscription(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	scsAsIdParam := params["scsAsId"]
	if scsAsIdParam == "" {
		c.errorHandler(w, r, &models.RequiredError{Field: "scsAsId"}, nil)
		return
	}
	subscriptionIdParam := params["subscriptionId"]
	if subscriptionIdParam == "" {
		c.errorHandler(w, r, &models.RequiredError{Field: "subscriptionId"}, nil)
		return
	}
	result, err := c.service.FetchIndASSessionWithQoSSubscription(r.Context(), scsAsIdParam, subscriptionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = models.EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateIndASSessionWithQoSSubscription - Updates/replaces an existing subscription resource.
func (c *IndividualASSessionWithRequiredQoSSubscriptionAPIController) UpdateIndASSessionWithQoSSubscription(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	scsAsIdParam := params["scsAsId"]
	if scsAsIdParam == "" {
		c.errorHandler(w, r, &models.RequiredError{Field: "scsAsId"}, nil)
		return
	}
	subscriptionIdParam := params["subscriptionId"]
	if subscriptionIdParam == "" {
		c.errorHandler(w, r, &models.RequiredError{Field: "subscriptionId"}, nil)
		return
	}
	asSessionWithQoSSubscriptionParam := models.AsSessionWithQoSSubscription{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&asSessionWithQoSSubscriptionParam); err != nil {
		c.errorHandler(w, r, &models.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertAsSessionWithQoSSubscriptionRequired(asSessionWithQoSSubscriptionParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertAsSessionWithQoSSubscriptionConstraints(asSessionWithQoSSubscriptionParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateIndASSessionWithQoSSubscription(r.Context(), scsAsIdParam, subscriptionIdParam, asSessionWithQoSSubscriptionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = models.EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteIndASSessionWithQoSSubscription - Deletes an already existing subscription.
func (c *IndividualASSessionWithRequiredQoSSubscriptionAPIController) DeleteIndASSessionWithQoSSubscription(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	scsAsIdParam := params["scsAsId"]
	if scsAsIdParam == "" {
		c.errorHandler(w, r, &models.RequiredError{Field: "scsAsId"}, nil)
		return
	}
	subscriptionIdParam := params["subscriptionId"]
	if subscriptionIdParam == "" {
		c.errorHandler(w, r, &models.RequiredError{Field: "subscriptionId"}, nil)
		return
	}
	result, err := c.service.DeleteIndASSessionWithQoSSubscription(r.Context(), scsAsIdParam, subscriptionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = models.EncodeJSONResponse(result.Body, &result.Code, w)
}

// ModifyIndASSessionWithQoSSubscription - Updates/replaces an existing subscription resource.
func (c *IndividualASSessionWithRequiredQoSSubscriptionAPIController) ModifyIndASSessionWithQoSSubscription(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	scsAsIdParam := params["scsAsId"]
	if scsAsIdParam == "" {
		c.errorHandler(w, r, &models.RequiredError{Field: "scsAsId"}, nil)
		return
	}
	subscriptionIdParam := params["subscriptionId"]
	if subscriptionIdParam == "" {
		c.errorHandler(w, r, &models.RequiredError{Field: "subscriptionId"}, nil)
		return
	}
	asSessionWithQoSSubscriptionPatchParam := models.AsSessionWithQoSSubscriptionPatch{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&asSessionWithQoSSubscriptionPatchParam); err != nil {
		c.errorHandler(w, r, &models.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertAsSessionWithQoSSubscriptionPatchRequired(asSessionWithQoSSubscriptionPatchParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertAsSessionWithQoSSubscriptionPatchConstraints(asSessionWithQoSSubscriptionPatchParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.ModifyIndASSessionWithQoSSubscription(r.Context(), scsAsIdParam, subscriptionIdParam, asSessionWithQoSSubscriptionPatchParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = models.EncodeJSONResponse(result.Body, &result.Code, w)
}
