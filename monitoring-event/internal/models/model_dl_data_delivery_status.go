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

// DlDataDeliveryStatus - Possible values are: - BUFFERED: The first downlink data is buffered with extended buffering matching the   source of the downlink traffic. - TRANSMITTED: The first downlink data matching the source of the downlink traffic is   transmitted after previous buffering or discarding of corresponding packet(s) because   the UE of the PDU Session becomes ACTIVE, and buffered data can be delivered to UE. - DISCARDED: The first downlink data matching the source of the downlink traffic is   discarded because the Extended Buffering time, as determined by the SMF, expires or   the amount of downlink data to be buffered is exceeded.
type DlDataDeliveryStatus string

const (
	// DL_DATA_DELIVERY_STATUS_BUFFERED - The first downlink data is buffered with extended
	// buffering matching the source of the downlink traffic.
	DlDataDeliveryStatus_BUFFERED DlDataDeliveryStatus = "BUFFERED"
	// DL_DATA_DELIVERY_STATUS_TRANSMITTED - The first downlink data matching the
	// source of the downlink traffic is transmitted after previous buffering or discarding
	// of corresponding packet(s) because the UE of the PDU Session becomes ACTIVE, and
	// buffered data can be delivered to UE.
	DlDataDeliveryStatus_TRANSMITTED DlDataDeliveryStatus = "TRANSMITTED"
	// DL_DATA_DELIVERY_STATUS_DISCARDED - The first downlink data matching the source
	// of the downlink traffic is discarded because the Extended Buffering time, as determined
	// by the SMF, expires or the amount of downlink data to be buffered is exceeded.
	DlDataDeliveryStatus_DISCARDED DlDataDeliveryStatus = "DISCARDED"
)

// AssertDlDataDeliveryStatusRequired checks if the required fields are not zero-ed
func AssertDlDataDeliveryStatusRequired(obj DlDataDeliveryStatus) error {
	return nil
}

// AssertDlDataDeliveryStatusConstraints checks if the values respects the defined constraints
func AssertDlDataDeliveryStatusConstraints(obj DlDataDeliveryStatus) error {
	return nil
}
