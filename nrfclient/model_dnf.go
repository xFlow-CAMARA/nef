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

// Dnf struct for Dnf
type Dnf struct {
	DnfUnits []DnfUnit `json:"dnfUnits"`
}

// NewDnf instantiates a new Dnf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnf(dnfUnits []DnfUnit) *Dnf {
	this := Dnf{}
	this.DnfUnits = dnfUnits
	return &this
}

// NewDnfWithDefaults instantiates a new Dnf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnfWithDefaults() *Dnf {
	this := Dnf{}
	return &this
}

// GetDnfUnits returns the DnfUnits field value
func (o *Dnf) GetDnfUnits() []DnfUnit {
	if o == nil {
		var ret []DnfUnit
		return ret
	}

	return o.DnfUnits
}

// GetDnfUnitsOk returns a tuple with the DnfUnits field value
// and a boolean to check if the value has been set.
func (o *Dnf) GetDnfUnitsOk() ([]DnfUnit, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnfUnits, true
}

// SetDnfUnits sets field value
func (o *Dnf) SetDnfUnits(v []DnfUnit) {
	o.DnfUnits = v
}

func (o Dnf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dnfUnits"] = o.DnfUnits
	}
	return json.Marshal(toSerialize)
}

type NullableDnf struct {
	value *Dnf
	isSet bool
}

func (v NullableDnf) Get() *Dnf {
	return v.value
}

func (v *NullableDnf) Set(val *Dnf) {
	v.value = val
	v.isSet = true
}

func (v NullableDnf) IsSet() bool {
	return v.isSet
}

func (v *NullableDnf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnf(val *Dnf) *NullableDnf {
	return &NullableDnf{value: val, isSet: true}
}

func (v NullableDnf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
