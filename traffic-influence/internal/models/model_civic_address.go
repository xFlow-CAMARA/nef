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

package models

// CivicAddress - Indicates a Civic address.
type CivicAddress struct {
	Country string `json:"country,omitempty"`

	A1 string `json:"A1,omitempty"`

	A2 string `json:"A2,omitempty"`

	A3 string `json:"A3,omitempty"`

	A4 string `json:"A4,omitempty"`

	A5 string `json:"A5,omitempty"`

	A6 string `json:"A6,omitempty"`

	PRD string `json:"PRD,omitempty"`

	POD string `json:"POD,omitempty"`

	STS string `json:"STS,omitempty"`

	HNO string `json:"HNO,omitempty"`

	HNS string `json:"HNS,omitempty"`

	LMK string `json:"LMK,omitempty"`

	LOC string `json:"LOC,omitempty"`

	NAM string `json:"NAM,omitempty"`

	PC string `json:"PC,omitempty"`

	BLD string `json:"BLD,omitempty"`

	UNIT string `json:"UNIT,omitempty"`

	FLR string `json:"FLR,omitempty"`

	ROOM string `json:"ROOM,omitempty"`

	PLC string `json:"PLC,omitempty"`

	PCN string `json:"PCN,omitempty"`

	POBOX string `json:"POBOX,omitempty"`

	ADDCODE string `json:"ADDCODE,omitempty"`

	SEAT string `json:"SEAT,omitempty"`

	RD string `json:"RD,omitempty"`

	RDSEC string `json:"RDSEC,omitempty"`

	RDBR string `json:"RDBR,omitempty"`

	RDSUBBR string `json:"RDSUBBR,omitempty"`

	PRM string `json:"PRM,omitempty"`

	POM string `json:"POM,omitempty"`

	UsageRules string `json:"usageRules,omitempty"`

	Method string `json:"method,omitempty"`

	ProvidedBy string `json:"providedBy,omitempty"`
}

// AssertCivicAddressRequired checks if the required fields are not zero-ed
func AssertCivicAddressRequired(obj CivicAddress) error {
	return nil
}

// AssertCivicAddressConstraints checks if the values respects the defined constraints
func AssertCivicAddressConstraints(obj CivicAddress) error {
	return nil
}
