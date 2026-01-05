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

package nrfclient

import (
	"encoding/json"
	"fmt"
)

// NwdafEvent Possible values are - SLICE_LOAD_LEVEL: Indicates that the event subscribed is load level information of Network Slice - NETWORK_PERFORMANCE: Indicates that the event subscribed is network performance information. - NF_LOAD: Indicates that the event subscribed is load level and status of one or several Network Functions. - SERVICE_EXPERIENCE: Indicates that the event subscribed is service experience. - UE_MOBILITY: Indicates that the event subscribed is UE mobility information. - UE_COMMUNICATION: Indicates that the event subscribed is UE communication information. - QOS_SUSTAINABILITY: Indicates that the event subscribed is QoS sustainability. - ABNORMAL_BEHAVIOUR: Indicates that the event subscribed is abnormal behaviour. - USER_DATA_CONGESTION: Indicates that the event subscribed is user data congestion information. - NSI_LOAD_LEVEL: Indicates that the event subscribed is load level information of Network Slice and the optionally associated Network Slice Instance
type NwdafEvent struct {
	NwdafEventAnyOf *NwdafEventAnyOf
	string          *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NwdafEvent) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NwdafEventAnyOf
	err = json.Unmarshal(data, &dst.NwdafEventAnyOf)
	if err == nil {
		jsonNwdafEventAnyOf, _ := json.Marshal(dst.NwdafEventAnyOf)
		if string(jsonNwdafEventAnyOf) == "{}" { // empty struct
			dst.NwdafEventAnyOf = nil
		} else {
			return nil // data stored in dst.NwdafEventAnyOf, return on the first match
		}
	} else {
		dst.NwdafEventAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NwdafEvent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NwdafEvent) MarshalJSON() ([]byte, error) {
	if src.NwdafEventAnyOf != nil {
		return json.Marshal(&src.NwdafEventAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNwdafEvent struct {
	value *NwdafEvent
	isSet bool
}

func (v NullableNwdafEvent) Get() *NwdafEvent {
	return v.value
}

func (v *NullableNwdafEvent) Set(val *NwdafEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableNwdafEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableNwdafEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNwdafEvent(val *NwdafEvent) *NullableNwdafEvent {
	return &NullableNwdafEvent{value: val, isSet: true}
}

func (v NullableNwdafEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNwdafEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
