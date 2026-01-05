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
NRF NFDiscovery Service

NRF NFDiscovery Service. © 2019, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.0.alpha-1
*/

package nrfclient

import (
	"encoding/json"
	"fmt"
)

// EventId Possible values are - LOAD_LEVEL_INFORMATION: Represent the analytics of load level information of corresponding network slice. - NETWORK_PERFORMANCE: Represent the analytics of network performance information. - NF_LOAD: Indicates that the event subscribed is NF Load. - SERVICE_EXPERIENCE: Represent the analytics of service experience information of the specific applications. - UE_MOBILITY: Represent the analytics of UE mobility. - UE_COMMUNICATION: Represent the analytics of UE communication. - QOS_SUSTAINABILITY: Represent the analytics of QoS sustainability information in the certain area.  - ABNORMAL_BEHAVIOUR: Indicates that the event subscribed is abnormal behaviour information. - USER_DATA_CONGESTION: Represent the analytics of the user data congestion in the certain area. - NSI_LOAD_LEVEL: Represent the analytics of Network Slice and the optionally associated Network Slice Instance.
type EventId struct {
	EventIdAnyOf *EventIdAnyOf
	string       *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EventId) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into EventIdAnyOf
	err = json.Unmarshal(data, &dst.EventIdAnyOf)
	if err == nil {
		jsonEventIdAnyOf, _ := json.Marshal(dst.EventIdAnyOf)
		if string(jsonEventIdAnyOf) == "{}" { // empty struct
			dst.EventIdAnyOf = nil
		} else {
			return nil // data stored in dst.EventIdAnyOf, return on the first match
		}
	} else {
		dst.EventIdAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string)
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(EventId)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EventId) MarshalJSON() ([]byte, error) {
	if src.EventIdAnyOf != nil {
		return json.Marshal(&src.EventIdAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEventId struct {
	value *EventId
	isSet bool
}

func (v NullableEventId) Get() *EventId {
	return v.value
}

func (v *NullableEventId) Set(val *EventId) {
	v.value = val
	v.isSet = true
}

func (v NullableEventId) IsSet() bool {
	return v.isSet
}

func (v *NullableEventId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventId(val *EventId) *NullableEventId {
	return &NullableEventId{value: val, isSet: true}
}

func (v NullableEventId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
