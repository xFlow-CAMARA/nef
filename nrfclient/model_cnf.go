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

// Cnf struct for Cnf
type Cnf struct {
	CnfUnits []CnfUnit `json:"cnfUnits"`
}

// NewCnf instantiates a new Cnf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCnf(cnfUnits []CnfUnit) *Cnf {
	this := Cnf{}
	this.CnfUnits = cnfUnits
	return &this
}

// NewCnfWithDefaults instantiates a new Cnf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCnfWithDefaults() *Cnf {
	this := Cnf{}
	return &this
}

// GetCnfUnits returns the CnfUnits field value
func (o *Cnf) GetCnfUnits() []CnfUnit {
	if o == nil {
		var ret []CnfUnit
		return ret
	}

	return o.CnfUnits
}

// GetCnfUnitsOk returns a tuple with the CnfUnits field value
// and a boolean to check if the value has been set.
func (o *Cnf) GetCnfUnitsOk() ([]CnfUnit, bool) {
	if o == nil {
		return nil, false
	}
	return o.CnfUnits, true
}

// SetCnfUnits sets field value
func (o *Cnf) SetCnfUnits(v []CnfUnit) {
	o.CnfUnits = v
}

func (o Cnf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cnfUnits"] = o.CnfUnits
	}
	return json.Marshal(toSerialize)
}

type NullableCnf struct {
	value *Cnf
	isSet bool
}

func (v NullableCnf) Get() *Cnf {
	return v.value
}

func (v *NullableCnf) Set(val *Cnf) {
	v.value = val
	v.isSet = true
}

func (v NullableCnf) IsSet() bool {
	return v.isSet
}

func (v *NullableCnf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCnf(val *Cnf) *NullableCnf {
	return &NullableCnf{value: val, isSet: true}
}

func (v NullableCnf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCnf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
