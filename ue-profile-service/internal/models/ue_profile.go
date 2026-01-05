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

package models

import (
	"encoding/json"
	"log"
	"sync"
)

type pduSessionInfo struct {
	Id       int32
	Ipv4     string
	Snssai   Snssai
	DlStatus string
	Dnn      string
}

type UeProfile struct {
	mu   sync.RWMutex
	Imsi string
	Gpsi string // MSISDN in E.164 format (e.g., "+33612345678")
	//Ipv4               string
	RegistrationStatus string
	ConnectionStatus   string
	Plmn               PlmnId
	Tac                string
	NrCellId           string
	PduSessions        map[int32]*pduSessionInfo
}

func NewUeProfile() *UeProfile {
	ueProfile := &UeProfile{}
	ueProfile.mu = sync.RWMutex{}
	return ueProfile
}

func NewUeProfileFromUeInfo(ueInfo *UeInfo) *UeProfile {
	ueProfile := &UeProfile{}
	ueProfile.mu = sync.RWMutex{}
	ueProfile.PduSessions = make(map[int32]*pduSessionInfo)

	/* get identity*/
	if ueInfo.Imsi != nil {
		ueProfile.Imsi = *ueInfo.Imsi
	}

	/* get GPSI (MSISDN) */
	if ueInfo.Gpsi != nil {
		ueProfile.Gpsi = *ueInfo.Gpsi
	}

	/* get registration */
	if ueInfo.RegistrationInfo != nil {
		ueProfile.RegistrationStatus = string(ueInfo.RegistrationInfo.RmInfo.RmState)
	}

	/* get connectivity */
	if ueInfo.ConnectivityInfo != nil {
		ueProfile.ConnectionStatus = string(ueInfo.ConnectivityInfo.CmInfo.CmState)
	}

	/* get IP address */

	for pduId, pduSess := range ueInfo.PduSessEst {

		//check if address field is not nil
		if pduSess.AdIpv4Addr == nil {
			continue
		}
		if pduSess.PduSeId == nil {
			continue
		}

		if pduSess.Dnn == nil {
			continue
		}

		if pduSess.Snssai == nil {
			continue
		}

		// check that no corrispective release is present
		if ueInfo.PduSessRel[pduId] != nil {
			if ueInfo.PduSessRel[pduId].TimeStamp >= pduSess.TimeStamp {
				continue
			}
		}

		var dddStatus string

		if ueInfo.Ddds[pduId] != nil {
			dddStatus = string(*ueInfo.Ddds[pduId].DddStatus.DlDataDeliveryStatusAnyOf)
		} else {
			dddStatus = "UNKNOWN"
		}

		ueProfile.PduSessions[*pduSess.PduSeId] = &pduSessionInfo{
			Id:       *pduSess.PduSeId,
			Ipv4:     *pduSess.AdIpv4Addr,
			Dnn:      *pduSess.Dnn,
			Snssai:   *pduSess.Snssai,
			DlStatus: dddStatus,
		}
	}

	/* get Plmn */
	if ueInfo.Location != nil {
		nrLoc, ok := ueInfo.Location.UserLocation.GetNrLocationOk()
		if ok {
			ueProfile.Plmn = nrLoc.Tai.PlmnId
			ueProfile.Tac = nrLoc.Tai.Tac
			ueProfile.NrCellId = nrLoc.Ncgi.NrCellId
		}
	}

	return ueProfile
}

func (ue *UeProfile) ApplyUpdate(patch *UeInfoPatch) bool {

	if patch == nil {
		return false
	}
	ue.mu.Lock()
	defer ue.mu.Unlock()

	var jsonData []byte

	/* marshall interface into json */
	jsonData, err := json.Marshal(patch.Data)
	if err != nil {
		log.Printf("malformed patch data from redis")
	}

	switch patch.Type {
	case "REGISTRATION_STATE_REPORT":
		regInfo := rmInfo{}
		if err := json.Unmarshal(jsonData, &regInfo); err != nil {
			log.Printf("could not deserialize %s", patch.Type)
		}
		ue.RegistrationStatus = string(regInfo.RmInfo.RmState)

	case "CONNECTIVITY_STATE_REPORT":
		conInfo := cmInfo{}
		if err := json.Unmarshal(jsonData, &conInfo); err != nil {
			log.Printf("could not deserialize %s", patch.Type)
		}
		ue.ConnectionStatus = string(conInfo.CmInfo.CmState)

	case "LOCATION_REPORT":
		locInfo := location{}
		if err := json.Unmarshal(jsonData, &locInfo); err != nil {
			log.Printf("could not deserialize %s", patch.Type)
		}
		nrLoc, ok := locInfo.UserLocation.GetNrLocationOk()
		if ok {
			ue.Plmn = nrLoc.Tai.PlmnId
			ue.Tac = nrLoc.Tai.Tac
			ue.NrCellId = nrLoc.Ncgi.NrCellId
		}

	case "PDU_SES_EST":
		estInfo := pduSesEst{}
		if err := json.Unmarshal(jsonData, &estInfo); err != nil {
			log.Printf("could not deserialize %s", patch.Type)
		}

		// check if address field is not nil
		if estInfo.AdIpv4Addr == nil {
			return false
		}

		if estInfo.PduSeId == nil {
			return false
		}

		if estInfo.Dnn == nil {
			return false
		}

		if estInfo.Snssai == nil {
			return false
		}

		pduSess := ue.PduSessions[*estInfo.PduSeId]
		if pduSess != nil {
			// pdu session exists, than update fields

			pduSess.Id = *estInfo.PduSeId
			pduSess.Ipv4 = *estInfo.AdIpv4Addr
			pduSess.Dnn = *estInfo.Dnn
			pduSess.Snssai = *estInfo.Snssai
		} else {
			/*create new pdu session*/
			ue.PduSessions[*estInfo.PduSeId] = &pduSessionInfo{
				Id:       *estInfo.PduSeId,
				Ipv4:     *estInfo.AdIpv4Addr,
				Dnn:      *estInfo.Dnn,
				Snssai:   *estInfo.Snssai,
				DlStatus: "UNKNOWN",
			}
		}

	case "DDDS":
		dddsInfo := ddds{}
		if err := json.Unmarshal(jsonData, &dddsInfo); err != nil {
			log.Printf("could not deserialize %s", patch.Type)
		}
		// check if address field is not nil
		if dddsInfo.DddStatus == nil {
			return false
		}

		if dddsInfo.PduSeId == nil {
			return false
		}

		if ue.PduSessions[*dddsInfo.PduSeId] != nil {
			ue.PduSessions[*dddsInfo.PduSeId].DlStatus = string(*dddsInfo.DddStatus.DlDataDeliveryStatusAnyOf)
		}

	case "PDU_SES_REL":
		relInfo := pduSesRel{}
		if err := json.Unmarshal(jsonData, &relInfo); err != nil {
			log.Printf("could not deserialize %s", patch.Type)
		}

		if relInfo.PduSeId == nil {
			return false
		}
		if ue.PduSessions[*relInfo.PduSeId] != nil {
			delete(ue.PduSessions, *relInfo.PduSeId)
		}
	}
	return true
}
