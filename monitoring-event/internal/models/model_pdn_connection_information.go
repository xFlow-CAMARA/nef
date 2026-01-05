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

package models

// PdnConnectionInformation - Represents the PDN connection information of the UE.
type PdnConnectionInformation struct {
	Status PdnConnectionStatus `json:"status"`

	// Identify the APN, it is depending on the SCEF local configuration whether or not this attribute is sent to the SCS/AS.
	Apn string `json:"apn,omitempty"`

	PdnType PdnType `json:"pdnType"`

	InterfaceInd *InterfaceIndication `json:"interfaceInd,omitempty"`

	// string identifying a Ipv4 address formatted in the \"dotted decimal\" notation as defined in IETF RFC 1166.
	Ipv4Addr string `json:"ipv4Addr,omitempty"`

	Ipv6Addrs []string `json:"ipv6Addrs,omitempty"`

	MacAddrs []string `json:"macAddrs,omitempty"`
}

// AssertPdnConnectionInformationRequired checks if the required fields are not zero-ed
func AssertPdnConnectionInformationRequired(obj PdnConnectionInformation) error {
	elements := map[string]interface{}{
		"status":  obj.Status,
		"pdnType": obj.PdnType,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertPdnConnectionStatusRequired(obj.Status); err != nil {
		return err
	}
	if err := AssertPdnTypeRequired(obj.PdnType); err != nil {
		return err
	}

	return nil
}

// AssertPdnConnectionInformationConstraints checks if the values respects the defined constraints
func AssertPdnConnectionInformationConstraints(obj PdnConnectionInformation) error {
	if err := AssertPdnConnectionStatusConstraints(obj.Status); err != nil {
		return err
	}
	if err := AssertPdnTypeConstraints(obj.PdnType); err != nil {
		return err
	}

	return nil
}
