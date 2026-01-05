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

// checks if the RelativeCartesianLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelativeCartesianLocation{}

// RelativeCartesianLocation Relative Cartesian Location
type RelativeCartesianLocation struct {
	// string with format 'float' as defined in OpenAPI.
	X float32 `json:"x"`
	// string with format 'float' as defined in OpenAPI.
	Y float32 `json:"y"`
	// string with format 'float' as defined in OpenAPI.
	Z *float32 `json:"z,omitempty"`
}

type _RelativeCartesianLocation RelativeCartesianLocation

// NewRelativeCartesianLocation instantiates a new RelativeCartesianLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelativeCartesianLocation(x float32, y float32) *RelativeCartesianLocation {
	this := RelativeCartesianLocation{}
	this.X = x
	this.Y = y
	return &this
}

// NewRelativeCartesianLocationWithDefaults instantiates a new RelativeCartesianLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelativeCartesianLocationWithDefaults() *RelativeCartesianLocation {
	this := RelativeCartesianLocation{}
	return &this
}

// GetX returns the X field value
func (o *RelativeCartesianLocation) GetX() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *RelativeCartesianLocation) GetXOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *RelativeCartesianLocation) SetX(v float32) {
	o.X = v
}

// GetY returns the Y field value
func (o *RelativeCartesianLocation) GetY() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *RelativeCartesianLocation) GetYOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *RelativeCartesianLocation) SetY(v float32) {
	o.Y = v
}

// GetZ returns the Z field value if set, zero value otherwise.
func (o *RelativeCartesianLocation) GetZ() float32 {
	if o == nil || IsNil(o.Z) {
		var ret float32
		return ret
	}
	return *o.Z
}

// GetZOk returns a tuple with the Z field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelativeCartesianLocation) GetZOk() (*float32, bool) {
	if o == nil || IsNil(o.Z) {
		return nil, false
	}
	return o.Z, true
}

// HasZ returns a boolean if a field has been set.
func (o *RelativeCartesianLocation) HasZ() bool {
	if o != nil && !IsNil(o.Z) {
		return true
	}

	return false
}

// SetZ gets a reference to the given float32 and assigns it to the Z field.
func (o *RelativeCartesianLocation) SetZ(v float32) {
	o.Z = &v
}

func (o RelativeCartesianLocation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelativeCartesianLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y
	if !IsNil(o.Z) {
		toSerialize["z"] = o.Z
	}
	return toSerialize, nil
}

func (o *RelativeCartesianLocation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"x",
		"y",
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

	varRelativeCartesianLocation := _RelativeCartesianLocation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRelativeCartesianLocation)

	if err != nil {
		return err
	}

	*o = RelativeCartesianLocation(varRelativeCartesianLocation)

	return err
}

type NullableRelativeCartesianLocation struct {
	value *RelativeCartesianLocation
	isSet bool
}

func (v NullableRelativeCartesianLocation) Get() *RelativeCartesianLocation {
	return v.value
}

func (v *NullableRelativeCartesianLocation) Set(val *RelativeCartesianLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableRelativeCartesianLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableRelativeCartesianLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelativeCartesianLocation(val *RelativeCartesianLocation) *NullableRelativeCartesianLocation {
	return &NullableRelativeCartesianLocation{value: val, isSet: true}
}

func (v NullableRelativeCartesianLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelativeCartesianLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
