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
 * 3gpp-as-session-with-qos
 *
 * API for setting us an AS session with required QoS.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.3.0-alpha.5
 */

package models

// ProtocolDescription - ProtocolDescription contains information to derive PDU set information.
type ProtocolDescription struct {
	TransportProto MediaTransportProto `json:"transportProto,omitempty"`

	RtpHeaderExtInfo RtpHeaderExtInfo `json:"rtpHeaderExtInfo,omitempty"`

	RtpPayloadInfoList []RtpPayloadInfo `json:"rtpPayloadInfoList,omitempty"`
}

// AssertProtocolDescriptionRequired checks if the required fields are not zero-ed
func AssertProtocolDescriptionRequired(obj ProtocolDescription) error {
	if err := AssertMediaTransportProtoRequired(obj.TransportProto); err != nil {
		return err
	}
	if err := AssertRtpHeaderExtInfoRequired(obj.RtpHeaderExtInfo); err != nil {
		return err
	}
	for _, el := range obj.RtpPayloadInfoList {
		if err := AssertRtpPayloadInfoRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertProtocolDescriptionConstraints checks if the values respects the defined constraints
func AssertProtocolDescriptionConstraints(obj ProtocolDescription) error {
	if err := AssertMediaTransportProtoConstraints(obj.TransportProto); err != nil {
		return err
	}
	if err := AssertRtpHeaderExtInfoConstraints(obj.RtpHeaderExtInfo); err != nil {
		return err
	}
	for _, el := range obj.RtpPayloadInfoList {
		if err := AssertRtpPayloadInfoConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
