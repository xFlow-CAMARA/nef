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

// NotificationType struct for NotificationType
type NotificationType struct {
	NotificationTypeAnyOf *NotificationTypeAnyOf
	string                *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NotificationType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NotificationTypeAnyOf
	err = json.Unmarshal(data, &dst.NotificationTypeAnyOf)
	if err == nil {
		jsonNotificationTypeAnyOf, _ := json.Marshal(dst.NotificationTypeAnyOf)
		if string(jsonNotificationTypeAnyOf) == "{}" { // empty struct
			dst.NotificationTypeAnyOf = nil
		} else {
			return nil // data stored in dst.NotificationTypeAnyOf, return on the first match
		}
	} else {
		dst.NotificationTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NotificationType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NotificationType) MarshalJSON() ([]byte, error) {
	if src.NotificationTypeAnyOf != nil {
		return json.Marshal(&src.NotificationTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNotificationType struct {
	value *NotificationType
	isSet bool
}

func (v NullableNotificationType) Get() *NotificationType {
	return v.value
}

func (v *NullableNotificationType) Set(val *NotificationType) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationType) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationType(val *NotificationType) *NullableNotificationType {
	return &NullableNotificationType{value: val, isSet: true}
}

func (v NullableNotificationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
