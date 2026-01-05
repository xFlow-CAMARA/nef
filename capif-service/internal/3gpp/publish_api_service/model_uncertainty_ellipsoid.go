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
CAPIF_Publish_Service_API

API for publishing service APIs.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.5
*/

package publish_api_service

import (
	"bytes"
	"encoding/json"
)

// checks if the UncertaintyEllipsoid type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UncertaintyEllipsoid{}

// UncertaintyEllipsoid Ellipsoid with uncertainty
type UncertaintyEllipsoid struct {
	// Indicates value of uncertainty.
	SemiMajor float32 `json:"semiMajor"`
	// Indicates value of uncertainty.
	SemiMinor float32 `json:"semiMinor"`
	// Indicates value of uncertainty.
	Vertical float32 `json:"vertical"`
	// Indicates value of orientation angle.
	OrientationMajor int32 `json:"orientationMajor"`
}

type _UncertaintyEllipsoid UncertaintyEllipsoid

// NewUncertaintyEllipsoid instantiates a new UncertaintyEllipsoid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUncertaintyEllipsoid(semiMajor float32, semiMinor float32, vertical float32, orientationMajor int32) *UncertaintyEllipsoid {
	this := UncertaintyEllipsoid{}
	this.SemiMajor = semiMajor
	this.SemiMinor = semiMinor
	this.Vertical = vertical
	this.OrientationMajor = orientationMajor
	return &this
}

// NewUncertaintyEllipsoidWithDefaults instantiates a new UncertaintyEllipsoid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUncertaintyEllipsoidWithDefaults() *UncertaintyEllipsoid {
	this := UncertaintyEllipsoid{}
	return &this
}

// GetSemiMajor returns the SemiMajor field value
func (o *UncertaintyEllipsoid) GetSemiMajor() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SemiMajor
}

// GetSemiMajorOk returns a tuple with the SemiMajor field value
// and a boolean to check if the value has been set.
func (o *UncertaintyEllipsoid) GetSemiMajorOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SemiMajor, true
}

// SetSemiMajor sets field value
func (o *UncertaintyEllipsoid) SetSemiMajor(v float32) {
	o.SemiMajor = v
}

// GetSemiMinor returns the SemiMinor field value
func (o *UncertaintyEllipsoid) GetSemiMinor() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SemiMinor
}

// GetSemiMinorOk returns a tuple with the SemiMinor field value
// and a boolean to check if the value has been set.
func (o *UncertaintyEllipsoid) GetSemiMinorOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SemiMinor, true
}

// SetSemiMinor sets field value
func (o *UncertaintyEllipsoid) SetSemiMinor(v float32) {
	o.SemiMinor = v
}

// GetVertical returns the Vertical field value
func (o *UncertaintyEllipsoid) GetVertical() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Vertical
}

// GetVerticalOk returns a tuple with the Vertical field value
// and a boolean to check if the value has been set.
func (o *UncertaintyEllipsoid) GetVerticalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vertical, true
}

// SetVertical sets field value
func (o *UncertaintyEllipsoid) SetVertical(v float32) {
	o.Vertical = v
}

// GetOrientationMajor returns the OrientationMajor field value
func (o *UncertaintyEllipsoid) GetOrientationMajor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrientationMajor
}

// GetOrientationMajorOk returns a tuple with the OrientationMajor field value
// and a boolean to check if the value has been set.
func (o *UncertaintyEllipsoid) GetOrientationMajorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrientationMajor, true
}

// SetOrientationMajor sets field value
func (o *UncertaintyEllipsoid) SetOrientationMajor(v int32) {
	o.OrientationMajor = v
}

func (o UncertaintyEllipsoid) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UncertaintyEllipsoid) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["semiMajor"] = o.SemiMajor
	toSerialize["semiMinor"] = o.SemiMinor
	toSerialize["vertical"] = o.Vertical
	toSerialize["orientationMajor"] = o.OrientationMajor
	return toSerialize, nil
}

func (o *UncertaintyEllipsoid) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"semiMajor",
		"semiMinor",
		"vertical",
		"orientationMajor",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return nil
		}
	}

	varUncertaintyEllipsoid := _UncertaintyEllipsoid{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUncertaintyEllipsoid)

	if err != nil {
		return err
	}

	*o = UncertaintyEllipsoid(varUncertaintyEllipsoid)

	return err
}

type NullableUncertaintyEllipsoid struct {
	value *UncertaintyEllipsoid
	isSet bool
}

func (v NullableUncertaintyEllipsoid) Get() *UncertaintyEllipsoid {
	return v.value
}

func (v *NullableUncertaintyEllipsoid) Set(val *UncertaintyEllipsoid) {
	v.value = val
	v.isSet = true
}

func (v NullableUncertaintyEllipsoid) IsSet() bool {
	return v.isSet
}

func (v *NullableUncertaintyEllipsoid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUncertaintyEllipsoid(val *UncertaintyEllipsoid) *NullableUncertaintyEllipsoid {
	return &NullableUncertaintyEllipsoid{value: val, isSet: true}
}

func (v NullableUncertaintyEllipsoid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUncertaintyEllipsoid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
