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
	"strings"

	"gitlab.eurecom.fr/open-exposure/nef/ue-address/internal/connector"
	"gitlab.eurecom.fr/open-exposure/nef/ue-address/internal/models"
	"gitlab.eurecom.fr/open-exposure/nef/ue-address/pkg/config"
)

type app interface {
	Cfg() *config.AppConfig
	Connector() *connector.Connector
}

type Service struct {
	app
}

func NewUeAddressService(app app) *Service {
	svc := &Service{app: app}
	return svc
}

func (s *Service) RetrieveUeAddress(ueIpReq *models.UeAddressReq) (int, *models.UeAddressInfo, error) {
	/* convert extid/gpsi to supi */
	gpsi := ueIpReq.Gpsi
	isExtId, output, err := parseGpsi(gpsi)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var supi string

	if isExtId {
		log.Printf("%s", output)
		supi, err = s.Connector().LookupExternalId(ueIpReq.AfId, output)
		if err != nil {
			return http.StatusNotFound, nil, err
		}
	} else {
		return http.StatusNotImplemented, nil, fmt.Errorf("msisdn is not supported yet")
	}

	ueIp, err := s.Connector().GetCurrentIps(supi)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, &models.UeAddressInfo{UeIpAddrs: ueIp}, nil
}
func parseGpsi(gpsi string) (bool, string, error) {
	if strings.HasPrefix(gpsi, "extid-") {
		return true, strings.TrimPrefix(gpsi, "extid-"), nil
	}
	if strings.HasPrefix(gpsi, "msisdn-") {
		return false, strings.TrimPrefix(gpsi, "msisdn-"), nil
	}
	return false, "", fmt.Errorf("invalid gpsi type, must be msisdn or extid")
}
