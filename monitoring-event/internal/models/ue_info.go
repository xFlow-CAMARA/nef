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

type UeInfoPatch struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
	Imsi string                 `json:"imsi"`
}

type UeInfo struct {
	Imsi               *string
	PduSessEst         map[string]*PduSesEst `json:"PDU_SES_EST"`
	PduSessRel         map[string]*PduSesRel `json:"PDU_SES_REL"`
	Ddds               map[string]*Ddds      `json:"DDDS"`
	RegistrationInfo   *RegInfo              `json:"REGISTRATION_STATE_REPORT"`
	ConnectivityInfo   *ConnInfo             `json:"CONNECTIVITY_STATE_REPORT"`
	LossOfConnectivity *LossOfConnectReason  `json:"LOSS_OF_CONNECTIVITY"`
	Location           *Location             `json:"LOCATION_REPORT"`
}

type PduSesEst struct {
	AdIpv4Addr  *string         `json:"AdIpv4Addr"`
	Dnn         *string         `json:"Dnn"`
	PduSeId     *int32          `json:"PduSeId"`
	PduSessType *PduSessionType `json:"PduSessType"`
	Snssai      *Snssai         `json:"Snssai"`
	TimeStamp   int64           `json:"TimeStamp"`
}
type PduSesRel struct {
	Ipv4Addr    *string         `json:"Ipv4Addr"`
	Dnn         *string         `json:"Dnn"`
	PduSeId     *int32          `json:"PduSeId"`
	PduSessType *PduSessionType `json:"PduSessType"`
	Snssai      *Snssai         `json:"Snssai"`
	TimeStamp   int64           `json:"TimeStamp"`
}

type UeIpCh struct {
	AdIpv4Addr *string `json:"AdIpv4Addr"`
	PduSeId    *int32  `json:"PduSeId"`
	TimeStamp  int64   `json:"TimeStamp"`
}

type Ddds struct {
	DddStatus  *DlDataDeliveryStatus `json:"DddStatus"`
	PduSeId    *int32                `json:"PduSeId"`
	Dnn        *string               `json:"Dnn"`
	Snssai     *Snssai               `json:"Snssai"`
	UeIpv4Addr *string               `json:"UeIpv4Addr"`
	TimeStamp  int64                 `json:"TimeStamp"`
}

type RegInfo struct {
	RmInfo    RmInfo `json:"RmInfo"`
	TimeStamp int64  `json:"TimeStamp"`
}

type ConnInfo struct {
	CmInfo    CmInfo `json:"CmInfo"`
	TimeStamp int64  `json:"TimeStamp"`
}

type Location struct {
	UserLocation   UserLocation    `json:"UserLocation"`
	GeographicArea *GeographicArea `json:"GeographicArea,omitempty"`
	TimeStamp      int64           `json:"TimeStamp"`
}

type LossOfConnectReason struct {
	LossOfConnectReason LossOfConnectivityReasonAnyOf `json:"LossOfConnectReason"`
	TimeStamp           int64                         `json:"TimeStamp"`
}
