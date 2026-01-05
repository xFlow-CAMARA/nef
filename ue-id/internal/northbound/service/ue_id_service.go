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

package service

import (
	"fmt"
	"log"
	"net/http"

	"gitlab.eurecom.fr/open-exposure/nef/ue-id/internal/connector"
	"gitlab.eurecom.fr/open-exposure/nef/ue-id/internal/models"
	"gitlab.eurecom.fr/open-exposure/nef/ue-id/pkg/config"
)

type app interface {
	Cfg() *config.AppConfig
	Connector() *connector.Connector
}

type Service struct {
	app
}

func NewUeIdService(app app) *Service {
	svc := &Service{app: app}
	return svc
}

func (s *Service) RetrieveUeId(ueIdReq *models.UeIdReq) (int, *models.UeIdInfo, error) {

	afId := ueIdReq.AfId
	if ueIdReq.UeIpAddr == nil {
		return http.StatusBadRequest, nil, fmt.Errorf("UeIpv4 is a required field")
	}
	ueIpv4 := ueIdReq.UeIpAddr.Ipv4Addr

	/* call the ue-identity-service to generate an external id for this ue */
	extId, err := s.Connector().CreateExternalId(afId, ueIpv4)
	if err != nil {
		log.Printf("could not retrieve identity information")
		return http.StatusNotFound, nil, err
	}

	return http.StatusOK, &models.UeIdInfo{
		ExternalId: extId,
	}, nil

}
