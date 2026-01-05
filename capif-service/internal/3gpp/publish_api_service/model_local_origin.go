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
	"encoding/json"
)

// checks if the LocalOrigin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocalOrigin{}

// LocalOrigin Indicates a Local origin in a reference system
type LocalOrigin struct {
	CoordinateId *string                  `json:"coordinateId,omitempty"`
	Point        *GeographicalCoordinates `json:"point,omitempty"`
}

// NewLocalOrigin instantiates a new LocalOrigin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalOrigin() *LocalOrigin {
	this := LocalOrigin{}
	return &this
}

// NewLocalOriginWithDefaults instantiates a new LocalOrigin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalOriginWithDefaults() *LocalOrigin {
	this := LocalOrigin{}
	return &this
}

// GetCoordinateId returns the CoordinateId field value if set, zero value otherwise.
func (o *LocalOrigin) GetCoordinateId() string {
	if o == nil || IsNil(o.CoordinateId) {
		var ret string
		return ret
	}
	return *o.CoordinateId
}

// GetCoordinateIdOk returns a tuple with the CoordinateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalOrigin) GetCoordinateIdOk() (*string, bool) {
	if o == nil || IsNil(o.CoordinateId) {
		return nil, false
	}
	return o.CoordinateId, true
}

// HasCoordinateId returns a boolean if a field has been set.
func (o *LocalOrigin) HasCoordinateId() bool {
	if o != nil && !IsNil(o.CoordinateId) {
		return true
	}

	return false
}

// SetCoordinateId gets a reference to the given string and assigns it to the CoordinateId field.
func (o *LocalOrigin) SetCoordinateId(v string) {
	o.CoordinateId = &v
}

// GetPoint returns the Point field value if set, zero value otherwise.
func (o *LocalOrigin) GetPoint() GeographicalCoordinates {
	if o == nil || IsNil(o.Point) {
		var ret GeographicalCoordinates
		return ret
	}
	return *o.Point
}

// GetPointOk returns a tuple with the Point field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalOrigin) GetPointOk() (*GeographicalCoordinates, bool) {
	if o == nil || IsNil(o.Point) {
		return nil, false
	}
	return o.Point, true
}

// HasPoint returns a boolean if a field has been set.
func (o *LocalOrigin) HasPoint() bool {
	if o != nil && !IsNil(o.Point) {
		return true
	}

	return false
}

// SetPoint gets a reference to the given GeographicalCoordinates and assigns it to the Point field.
func (o *LocalOrigin) SetPoint(v GeographicalCoordinates) {
	o.Point = &v
}

func (o LocalOrigin) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocalOrigin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CoordinateId) {
		toSerialize["coordinateId"] = o.CoordinateId
	}
	if !IsNil(o.Point) {
		toSerialize["point"] = o.Point
	}
	return toSerialize, nil
}

type NullableLocalOrigin struct {
	value *LocalOrigin
	isSet bool
}

func (v NullableLocalOrigin) Get() *LocalOrigin {
	return v.value
}

func (v *NullableLocalOrigin) Set(val *LocalOrigin) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalOrigin) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalOrigin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalOrigin(val *LocalOrigin) *NullableLocalOrigin {
	return &NullableLocalOrigin{value: val, isSet: true}
}

func (v NullableLocalOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalOrigin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
