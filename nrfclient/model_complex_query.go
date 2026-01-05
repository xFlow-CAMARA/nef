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

// ComplexQuery - struct for ComplexQuery
type ComplexQuery struct {
	Cnf *Cnf
	Dnf *Dnf
}

// CnfAsComplexQuery is a convenience function that returns Cnf wrapped in ComplexQuery
func CnfAsComplexQuery(v *Cnf) ComplexQuery {
	return ComplexQuery{
		Cnf: v,
	}
}

// DnfAsComplexQuery is a convenience function that returns Dnf wrapped in ComplexQuery
func DnfAsComplexQuery(v *Dnf) ComplexQuery {
	return ComplexQuery{
		Dnf: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ComplexQuery) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Cnf
	err = newStrictDecoder(data).Decode(&dst.Cnf)
	if err == nil {
		jsonCnf, _ := json.Marshal(dst.Cnf)
		if string(jsonCnf) == "{}" { // empty struct
			dst.Cnf = nil
		} else {
			match++
		}
	} else {
		dst.Cnf = nil
	}

	// try to unmarshal data into Dnf
	err = newStrictDecoder(data).Decode(&dst.Dnf)
	if err == nil {
		jsonDnf, _ := json.Marshal(dst.Dnf)
		if string(jsonDnf) == "{}" { // empty struct
			dst.Dnf = nil
		} else {
			match++
		}
	} else {
		dst.Dnf = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Cnf = nil
		dst.Dnf = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(ComplexQuery)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ComplexQuery)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ComplexQuery) MarshalJSON() ([]byte, error) {
	if src.Cnf != nil {
		return json.Marshal(&src.Cnf)
	}

	if src.Dnf != nil {
		return json.Marshal(&src.Dnf)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ComplexQuery) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Cnf != nil {
		return obj.Cnf
	}

	if obj.Dnf != nil {
		return obj.Dnf
	}

	// all schemas are nil
	return nil
}

type NullableComplexQuery struct {
	value *ComplexQuery
	isSet bool
}

func (v NullableComplexQuery) Get() *ComplexQuery {
	return v.value
}

func (v *NullableComplexQuery) Set(val *ComplexQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableComplexQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableComplexQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComplexQuery(val *ComplexQuery) *NullableComplexQuery {
	return &NullableComplexQuery{value: val, isSet: true}
}

func (v NullableComplexQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComplexQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
