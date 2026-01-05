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

// DataSetId struct for DataSetId
type DataSetId struct {
	DataSetIdAnyOf *DataSetIdAnyOf
	string         *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DataSetId) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into DataSetIdAnyOf
	err = json.Unmarshal(data, &dst.DataSetIdAnyOf)
	if err == nil {
		jsonDataSetIdAnyOf, _ := json.Marshal(dst.DataSetIdAnyOf)
		if string(jsonDataSetIdAnyOf) == "{}" { // empty struct
			dst.DataSetIdAnyOf = nil
		} else {
			return nil // data stored in dst.DataSetIdAnyOf, return on the first match
		}
	} else {
		dst.DataSetIdAnyOf = nil
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

	return fmt.Errorf("Data failed to match schemas in anyOf(DataSetId)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DataSetId) MarshalJSON() ([]byte, error) {
	if src.DataSetIdAnyOf != nil {
		return json.Marshal(&src.DataSetIdAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDataSetId struct {
	value *DataSetId
	isSet bool
}

func (v NullableDataSetId) Get() *DataSetId {
	return v.value
}

func (v *NullableDataSetId) Set(val *DataSetId) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSetId) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSetId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSetId(val *DataSetId) *NullableDataSetId {
	return &NullableDataSetId{value: val, isSet: true}
}

func (v NullableDataSetId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSetId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
