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
 * 3gpp-ueid
 *
 * API for UE ID service. © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.1.0-alpha.2
 */

package northbound

import (
	"encoding/json"
	"net/http"
	"strings"

	"gitlab.eurecom.fr/open-exposure/nef/ue-id/internal/models"
)

// DefaultAPIController binds http requests to an api service and writes the service results to the http response
type DefaultAPIController struct {
	service      DefaultAPIServicer
	ErrorHandler models.ErrorHandler
}

// DefaultAPIOption for how the controller is set up.
type DefaultAPIOption func(*DefaultAPIController)

// WithDefaultAPIErrorHandler inject models.ErrorHandler into controller
func WithDefaultAPIErrorHandler(h models.ErrorHandler) DefaultAPIOption {
	return func(c *DefaultAPIController) {
		c.ErrorHandler = h
	}
}

// NewDefaultAPIController creates a default api controller
func NewDefaultAPIController(s DefaultAPIServicer, opts ...DefaultAPIOption) *DefaultAPIController {
	controller := &DefaultAPIController{
		service:      s,
		ErrorHandler: models.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the DefaultAPIController
func (c *DefaultAPIController) Routes() models.Routes {
	return models.Routes{
		"RetrieveUEId": models.Route{
			Method:      strings.ToUpper("Post"),
			Pattern:     "/3gpp-ueid/v1/retrieve",
			HandlerFunc: c.RetrieveUEId,
		},
	}
}

// RetrieveUEId - Retrieve AF specific UE ID.
func (c *DefaultAPIController) RetrieveUEId(w http.ResponseWriter, r *http.Request) {
	ueIdReqParam := models.UeIdReq{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&ueIdReqParam); err != nil {
		c.ErrorHandler(w, r, &models.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertUeIdReqRequired(ueIdReqParam); err != nil {
		c.ErrorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertUeIdReqConstraints(ueIdReqParam); err != nil {
		c.ErrorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.RetrieveUEId(r.Context(), &ueIdReqParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.ErrorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = models.EncodeJSONResponse(result.Body, &result.Code, w)
}
