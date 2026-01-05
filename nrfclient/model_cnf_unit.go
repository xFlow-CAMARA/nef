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
)

// CnfUnit struct for CnfUnit
type CnfUnit struct {
	CnfUnit []Atom `json:"cnfUnit"`
}

// NewCnfUnit instantiates a new CnfUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCnfUnit(cnfUnit []Atom) *CnfUnit {
	this := CnfUnit{}
	this.CnfUnit = cnfUnit
	return &this
}

// NewCnfUnitWithDefaults instantiates a new CnfUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCnfUnitWithDefaults() *CnfUnit {
	this := CnfUnit{}
	return &this
}

// GetCnfUnit returns the CnfUnit field value
func (o *CnfUnit) GetCnfUnit() []Atom {
	if o == nil {
		var ret []Atom
		return ret
	}

	return o.CnfUnit
}

// GetCnfUnitOk returns a tuple with the CnfUnit field value
// and a boolean to check if the value has been set.
func (o *CnfUnit) GetCnfUnitOk() ([]Atom, bool) {
	if o == nil {
		return nil, false
	}
	return o.CnfUnit, true
}

// SetCnfUnit sets field value
func (o *CnfUnit) SetCnfUnit(v []Atom) {
	o.CnfUnit = v
}

func (o CnfUnit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cnfUnit"] = o.CnfUnit
	}
	return json.Marshal(toSerialize)
}

type NullableCnfUnit struct {
	value *CnfUnit
	isSet bool
}

func (v NullableCnfUnit) Get() *CnfUnit {
	return v.value
}

func (v *NullableCnfUnit) Set(val *CnfUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableCnfUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableCnfUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCnfUnit(val *CnfUnit) *NullableCnfUnit {
	return &NullableCnfUnit{value: val, isSet: true}
}

func (v NullableCnfUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCnfUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
