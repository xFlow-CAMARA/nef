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
	"encoding/json"
	"net/http"
	"strings"

	"gitlab.eurecom.fr/open-exposure/nef/monitoring-event/internal/models"

	"github.com/gorilla/mux"
)

// MonitoringEventSubscriptionsAPIController binds http requests to an api service and writes the service results to the http response
type MonitoringEventSubscriptionsAPIController struct {
	service      MonitoringEventSubscriptionsAPIServicer
	errorHandler models.ErrorHandler
}

// MonitoringEventSubscriptionsAPIOption for how the controller is set up.
type MonitoringEventSubscriptionsAPIOption func(*MonitoringEventSubscriptionsAPIController)

// WithMonitoringEventSubscriptionsAPIErrorHandler inject models.ErrorHandler into controller
func WithMonitoringEventSubscriptionsAPIErrorHandler(h models.ErrorHandler) MonitoringEventSubscriptionsAPIOption {
	return func(c *MonitoringEventSubscriptionsAPIController) {
		c.errorHandler = h
	}
}

// NewMonitoringEventSubscriptionsAPIController creates a default api controller
func NewMonitoringEventSubscriptionsAPIController(s MonitoringEventSubscriptionsAPIServicer, opts ...MonitoringEventSubscriptionsAPIOption) *MonitoringEventSubscriptionsAPIController {
	controller := &MonitoringEventSubscriptionsAPIController{
		service:      s,
		errorHandler: models.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// models.Routes returns all the api routes for the MonitoringEventSubscriptionsAPIController
func (c *MonitoringEventSubscriptionsAPIController) Routes() models.Routes {
	return models.Routes{
		"FetchAllMonitoringEventSubscriptions": models.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "/3gpp-monitoring-event/v1/{scsAsId}/subscriptions",
			HandlerFunc: c.FetchAllMonitoringEventSubscriptions,
		},
		"CreateMonitoringEventSubscription": models.Route{
			Method:      strings.ToUpper("Post"),
			Pattern:     "/3gpp-monitoring-event/v1/{scsAsId}/subscriptions",
			HandlerFunc: c.CreateMonitoringEventSubscription,
		},
	}
}

// FetchAllMonitoringEventSubscriptions - Read all or queried active subscriptions for the SCS/AS.
func (c *MonitoringEventSubscriptionsAPIController) FetchAllMonitoringEventSubscriptions(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.FetchAllMonitoringEventSubscriptions(r.Context(), scsAsIdParam, ipAddrsParam, ipDomainParam, macAddrsParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = models.EncodeJSONResponse(result.Body, &result.Code, w)
}

// CreateMonitoringEventSubscription - Creates a new subscription resource for monitoring event notification.
func (c *MonitoringEventSubscriptionsAPIController) CreateMonitoringEventSubscription(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	scsAsIdParam := params["scsAsId"]
	if scsAsIdParam == "" {
		c.errorHandler(w, r, &models.RequiredError{Field: "scsAsId"}, nil)
		return
	}
	monitoringEventSubscriptionParam := models.MonitoringEventSubscription{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&monitoringEventSubscriptionParam); err != nil {
		c.errorHandler(w, r, &models.ParsingError{Err: err}, nil)
		return
	}
	/*if err := models.AssertMonitoringEventSubscriptionRequired(monitoringEventSubscriptionParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertMonitoringEventSubscriptionConstraints(monitoringEventSubscriptionParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}*/
	result, err := c.service.CreateMonitoringEventSubscription(r.Context(), scsAsIdParam, &monitoringEventSubscriptionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	if len(result.Location) > 0 {
		w.Header().Set("Location", result.Location)
	}
	// If no error, encode the body and the result code
	_ = models.EncodeJSONResponse(result.Body, &result.Code, w)
}

// IndividualMonitoringEventSubscriptionAPIController binds http requests to an api service and writes the service results to the http response
type IndividualMonitoringEventSubscriptionAPIController struct {
	service      IndividualMonitoringEventSubscriptionAPIServicer
	errorHandler models.ErrorHandler
}

// IndividualMonitoringEventSubscriptionAPIOption for how the controller is set up.
type IndividualMonitoringEventSubscriptionAPIOption func(*IndividualMonitoringEventSubscriptionAPIController)

// WithIndividualMonitoringEventSubscriptionAPIErrorHandler inject models.ErrorHandler into controller
func WithIndividualMonitoringEventSubscriptionAPIErrorHandler(h models.ErrorHandler) IndividualMonitoringEventSubscriptionAPIOption {
	return func(c *IndividualMonitoringEventSubscriptionAPIController) {
		c.errorHandler = h
	}
}

// NewIndividualMonitoringEventSubscriptionAPIController creates a default api controller
func NewIndividualMonitoringEventSubscriptionAPIController(s IndividualMonitoringEventSubscriptionAPIServicer, opts ...IndividualMonitoringEventSubscriptionAPIOption) *IndividualMonitoringEventSubscriptionAPIController {
	controller := &IndividualMonitoringEventSubscriptionAPIController{
		service:      s,
		errorHandler: models.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// models.Routes returns all the api routes for the IndividualMonitoringEventSubscriptionAPIController
func (c *IndividualMonitoringEventSubscriptionAPIController) Routes() models.Routes {
	return models.Routes{
		"FetchIndMonitoringEventSubscription": models.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "/3gpp-monitoring-event/v1/{scsAsId}/subscriptions/{subscriptionId}",
			HandlerFunc: c.FetchIndMonitoringEventSubscription,
		},
		"UpdateIndMonitoringEventSubscription": models.Route{
			Method:      strings.ToUpper("Put"),
			Pattern:     "/3gpp-monitoring-event/v1/{scsAsId}/subscriptions/{subscriptionId}",
			HandlerFunc: c.UpdateIndMonitoringEventSubscription,
		},
		"DeleteIndMonitoringEventSubscription": models.Route{
			Method:      strings.ToUpper("Delete"),
			Pattern:     "/3gpp-monitoring-event/v1/{scsAsId}/subscriptions/{subscriptionId}",
			HandlerFunc: c.DeleteIndMonitoringEventSubscription,
		},
		"ModifyIndMonitoringEventSubscription": models.Route{
			Method:      strings.ToUpper("Patch"),
			Pattern:     "/3gpp-monitoring-event/v1/{scsAsId}/subscriptions/{subscriptionId}",
			HandlerFunc: c.ModifyIndMonitoringEventSubscription,
		},
	}
}

// FetchIndMonitoringEventSubscription - Read an active subscriptions for the SCS/AS and the subscription Id.
func (c *IndividualMonitoringEventSubscriptionAPIController) FetchIndMonitoringEventSubscription(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.FetchIndMonitoringEventSubscription(r.Context(), scsAsIdParam, subscriptionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = models.EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateIndMonitoringEventSubscription - Updates/replaces an existing subscription resource.
func (c *IndividualMonitoringEventSubscriptionAPIController) UpdateIndMonitoringEventSubscription(w http.ResponseWriter, r *http.Request) {
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
	monitoringEventSubscriptionParam := models.MonitoringEventSubscription{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&monitoringEventSubscriptionParam); err != nil {
		c.errorHandler(w, r, &models.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertMonitoringEventSubscriptionRequired(monitoringEventSubscriptionParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertMonitoringEventSubscriptionConstraints(monitoringEventSubscriptionParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateIndMonitoringEventSubscription(r.Context(), scsAsIdParam, subscriptionIdParam, &monitoringEventSubscriptionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = models.EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteIndMonitoringEventSubscription - Deletes an already existing monitoring event subscription.
func (c *IndividualMonitoringEventSubscriptionAPIController) DeleteIndMonitoringEventSubscription(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.DeleteIndMonitoringEventSubscription(r.Context(), scsAsIdParam, subscriptionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = models.EncodeJSONResponse(result.Body, &result.Code, w)
}

// ModifyIndMonitoringEventSubscription - Modifies an existing subscription of monitoring event.
func (c *IndividualMonitoringEventSubscriptionAPIController) ModifyIndMonitoringEventSubscription(w http.ResponseWriter, r *http.Request) {
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
	patchItemParam := []models.PatchItem{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&patchItemParam); err != nil {
		c.errorHandler(w, r, &models.ParsingError{Err: err}, nil)
		return
	}
	for _, el := range patchItemParam {
		if err := models.AssertPatchItemRequired(el); err != nil {
			c.errorHandler(w, r, err, nil)
			return
		}
	}
	result, err := c.service.ModifyIndMonitoringEventSubscription(r.Context(), scsAsIdParam, subscriptionIdParam, patchItemParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = models.EncodeJSONResponse(result.Body, &result.Code, w)
}
