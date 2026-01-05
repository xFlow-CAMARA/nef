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

type pduSessionInfo struct {
	Id       int32
	Ipv4     string
	Snssai   map[string]interface{} // not necessary
	DlStatus string
	Dnn      string
}

type UeProfile struct {
	Imsi               string
	RegistrationStatus string
	ConnectionStatus   string
	Plmn               map[string]interface{} // not necessary
	Tac                string
	NrCellId           string
	PduSessions        map[int32]*pduSessionInfo
}

func NewUeProfile() *UeProfile {
	ueProfile := &UeProfile{}
	return ueProfile
}
