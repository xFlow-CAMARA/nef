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

// PdnType - Represents the PDN connection type.   Possible values are - IPV4: PDN connection of IPv4 type. - IPV6: PDN connection of IPv6 type. - IPV4V6: PDN connection of IPv4v6 type. - NON_IP: PDN connection of non-IP type. - ETHERNET: PDN connection of Ethernet type.
type PdnType string

const (
	// PDN_TYPE_IPV4 - PDN connection of IPv4 type.
	PdnType_IPV4 PdnType = "IPV4"
	// PDN_TYPE_IPV6 - PDN connection of IPv6 type.
	PdnType_IPV6 PdnType = "IPV6"
	// PDN_TYPE_IPV4V6 - PDN connection of IPv4v6 type
	PdnType_IPV4V6 PdnType = "IPV4V6"
	// PDN_TYPE_NON_IP - PDN connection of non-IP type.
	PdnType_NON_IP PdnType = "NON_IP"
	// PDN_TYPE_ETHERNET - PDN connection of Ethernet type.
	PdnType_ETHERNET PdnType = "ETHERNET"
)

// AssertPdnTypeRequired checks if the required fields are not zero-ed
func AssertPdnTypeRequired(obj PdnType) error {
	return nil
}

// AssertPdnTypeConstraints checks if the values respects the defined constraints
func AssertPdnTypeConstraints(obj PdnType) error {
	return nil
}
