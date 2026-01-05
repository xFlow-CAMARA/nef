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

// checks if the GADShape type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GADShape{}

// GADShape Common base type for GAD shapes.
type GADShape struct {
	Shape SupportedGADShapes `json:"shape"`
}

type _GADShape GADShape

// NewGADShape instantiates a new GADShape object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGADShape(shape SupportedGADShapes) *GADShape {
	this := GADShape{}
	this.Shape = shape
	return &this
}

// NewGADShapeWithDefaults instantiates a new GADShape object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGADShapeWithDefaults() *GADShape {
	this := GADShape{}
	return &this
}

// GetShape returns the Shape field value
func (o *GADShape) GetShape() SupportedGADShapes {
	if o == nil {
		var ret SupportedGADShapes
		return ret
	}

	return o.Shape
}

// GetShapeOk returns a tuple with the Shape field value
// and a boolean to check if the value has been set.
func (o *GADShape) GetShapeOk() (*SupportedGADShapes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Shape, true
}

// SetShape sets field value
func (o *GADShape) SetShape(v SupportedGADShapes) {
	o.Shape = v
}

func (o GADShape) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GADShape) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shape"] = o.Shape
	return toSerialize, nil
}

func (o *GADShape) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"shape",
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

	varGADShape := _GADShape{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGADShape)

	if err != nil {
		return err
	}

	*o = GADShape(varGADShape)

	return err
}

type NullableGADShape struct {
	value *GADShape
	isSet bool
}

func (v NullableGADShape) Get() *GADShape {
	return v.value
}

func (v *NullableGADShape) Set(val *GADShape) {
	v.value = val
	v.isSet = true
}

func (v NullableGADShape) IsSet() bool {
	return v.isSet
}

func (v *NullableGADShape) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGADShape(val *GADShape) *NullableGADShape {
	return &NullableGADShape{value: val, isSet: true}
}

func (v NullableGADShape) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGADShape) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
