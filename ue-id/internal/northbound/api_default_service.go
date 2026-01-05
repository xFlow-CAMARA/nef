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
	"context"

	"gitlab.eurecom.fr/open-exposure/nef/ue-id/internal/models"
	"gitlab.eurecom.fr/open-exposure/nef/ue-id/internal/northbound/service"
	//"net/http"
	//"errors"
)

type serverCtx interface {
	Service() *service.Service
}

// DefaultAPIService is a service that implements the logic for the DefaultAPIServicer
// This service should implement the business logic for every endpoint for the DefaultAPI API.
// Include any external packages or services that will be required by this service.
type DefaultAPIService struct {
	serverCtx
}

// NewDefaultAPIService creates a default api service
func NewDefaultAPIService(serverCtx serverCtx) *DefaultAPIService {
	return &DefaultAPIService{serverCtx: serverCtx}
}

// RetrieveUEId - Retrieve AF specific UE ID.
func (s *DefaultAPIService) RetrieveUEId(ctx context.Context, ueIdReq *models.UeIdReq) (models.ImplResponse, error) {
	code, ueInfo, err := s.Service().RetrieveUeId(ueIdReq)
	if err != nil {
		return models.Response(code, models.ProblemDetails{Cause: err.Error()}), nil

	}
	// TODO: Uncomment the next line to return response Response(404, ProblemDetails{}) or use other options such as http.Ok ...
	return models.Response(code, ueInfo), nil
}
