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

// checks if the PointUncertaintyCircle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PointUncertaintyCircle{}

// PointUncertaintyCircle Ellipsoid point with uncertainty circle.
type PointUncertaintyCircle struct {
	GADShape
	Point GeographicalCoordinates `json:"point"`
	// Indicates value of uncertainty.
	Uncertainty float32 `json:"uncertainty"`
}

type _PointUncertaintyCircle PointUncertaintyCircle

// NewPointUncertaintyCircle instantiates a new PointUncertaintyCircle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPointUncertaintyCircle(point GeographicalCoordinates, uncertainty float32, shape SupportedGADShapes) *PointUncertaintyCircle {
	this := PointUncertaintyCircle{}
	this.Shape = shape
	this.Point = point
	this.Uncertainty = uncertainty
	return &this
}

// NewPointUncertaintyCircleWithDefaults instantiates a new PointUncertaintyCircle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPointUncertaintyCircleWithDefaults() *PointUncertaintyCircle {
	this := PointUncertaintyCircle{}
	return &this
}

// GetPoint returns the Point field value
func (o *PointUncertaintyCircle) GetPoint() GeographicalCoordinates {
	if o == nil {
		var ret GeographicalCoordinates
		return ret
	}

	return o.Point
}

// GetPointOk returns a tuple with the Point field value
// and a boolean to check if the value has been set.
func (o *PointUncertaintyCircle) GetPointOk() (*GeographicalCoordinates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Point, true
}

// SetPoint sets field value
func (o *PointUncertaintyCircle) SetPoint(v GeographicalCoordinates) {
	o.Point = v
}

// GetUncertainty returns the Uncertainty field value
func (o *PointUncertaintyCircle) GetUncertainty() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Uncertainty
}

// GetUncertaintyOk returns a tuple with the Uncertainty field value
// and a boolean to check if the value has been set.
func (o *PointUncertaintyCircle) GetUncertaintyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uncertainty, true
}

// SetUncertainty sets field value
func (o *PointUncertaintyCircle) SetUncertainty(v float32) {
	o.Uncertainty = v
}

func (o PointUncertaintyCircle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PointUncertaintyCircle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedGADShape, errGADShape := json.Marshal(o.GADShape)
	if errGADShape != nil {
		return map[string]interface{}{}, errGADShape
	}
	errGADShape = json.Unmarshal([]byte(serializedGADShape), &toSerialize)
	if errGADShape != nil {
		return map[string]interface{}{}, errGADShape
	}
	toSerialize["point"] = o.Point
	toSerialize["uncertainty"] = o.Uncertainty
	return toSerialize, nil
}

func (o *PointUncertaintyCircle) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"point",
		"uncertainty",
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

	varPointUncertaintyCircle := _PointUncertaintyCircle{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPointUncertaintyCircle)

	if err != nil {
		return err
	}

	*o = PointUncertaintyCircle(varPointUncertaintyCircle)

	return err
}

type NullablePointUncertaintyCircle struct {
	value *PointUncertaintyCircle
	isSet bool
}

func (v NullablePointUncertaintyCircle) Get() *PointUncertaintyCircle {
	return v.value
}

func (v *NullablePointUncertaintyCircle) Set(val *PointUncertaintyCircle) {
	v.value = val
	v.isSet = true
}

func (v NullablePointUncertaintyCircle) IsSet() bool {
	return v.isSet
}

func (v *NullablePointUncertaintyCircle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePointUncertaintyCircle(val *PointUncertaintyCircle) *NullablePointUncertaintyCircle {
	return &NullablePointUncertaintyCircle{value: val, isSet: true}
}

func (v NullablePointUncertaintyCircle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePointUncertaintyCircle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
