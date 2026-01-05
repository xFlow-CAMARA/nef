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

// checks if the PointUncertaintyEllipse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PointUncertaintyEllipse{}

// PointUncertaintyEllipse Ellipsoid point with uncertainty ellipse.
type PointUncertaintyEllipse struct {
	GADShape
	Point              GeographicalCoordinates `json:"point"`
	UncertaintyEllipse UncertaintyEllipse      `json:"uncertaintyEllipse"`
	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`
}

type _PointUncertaintyEllipse PointUncertaintyEllipse

// NewPointUncertaintyEllipse instantiates a new PointUncertaintyEllipse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPointUncertaintyEllipse(point GeographicalCoordinates, uncertaintyEllipse UncertaintyEllipse, confidence int32, shape SupportedGADShapes) *PointUncertaintyEllipse {
	this := PointUncertaintyEllipse{}
	this.Shape = shape
	this.Point = point
	this.UncertaintyEllipse = uncertaintyEllipse
	this.Confidence = confidence
	return &this
}

// NewPointUncertaintyEllipseWithDefaults instantiates a new PointUncertaintyEllipse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPointUncertaintyEllipseWithDefaults() *PointUncertaintyEllipse {
	this := PointUncertaintyEllipse{}
	return &this
}

// GetPoint returns the Point field value
func (o *PointUncertaintyEllipse) GetPoint() GeographicalCoordinates {
	if o == nil {
		var ret GeographicalCoordinates
		return ret
	}

	return o.Point
}

// GetPointOk returns a tuple with the Point field value
// and a boolean to check if the value has been set.
func (o *PointUncertaintyEllipse) GetPointOk() (*GeographicalCoordinates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Point, true
}

// SetPoint sets field value
func (o *PointUncertaintyEllipse) SetPoint(v GeographicalCoordinates) {
	o.Point = v
}

// GetUncertaintyEllipse returns the UncertaintyEllipse field value
func (o *PointUncertaintyEllipse) GetUncertaintyEllipse() UncertaintyEllipse {
	if o == nil {
		var ret UncertaintyEllipse
		return ret
	}

	return o.UncertaintyEllipse
}

// GetUncertaintyEllipseOk returns a tuple with the UncertaintyEllipse field value
// and a boolean to check if the value has been set.
func (o *PointUncertaintyEllipse) GetUncertaintyEllipseOk() (*UncertaintyEllipse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UncertaintyEllipse, true
}

// SetUncertaintyEllipse sets field value
func (o *PointUncertaintyEllipse) SetUncertaintyEllipse(v UncertaintyEllipse) {
	o.UncertaintyEllipse = v
}

// GetConfidence returns the Confidence field value
func (o *PointUncertaintyEllipse) GetConfidence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value
// and a boolean to check if the value has been set.
func (o *PointUncertaintyEllipse) GetConfidenceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Confidence, true
}

// SetConfidence sets field value
func (o *PointUncertaintyEllipse) SetConfidence(v int32) {
	o.Confidence = v
}

func (o PointUncertaintyEllipse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PointUncertaintyEllipse) ToMap() (map[string]interface{}, error) {
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
	toSerialize["uncertaintyEllipse"] = o.UncertaintyEllipse
	toSerialize["confidence"] = o.Confidence
	return toSerialize, nil
}

func (o *PointUncertaintyEllipse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"point",
		"uncertaintyEllipse",
		"confidence",
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

	varPointUncertaintyEllipse := _PointUncertaintyEllipse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPointUncertaintyEllipse)

	if err != nil {
		return err
	}

	*o = PointUncertaintyEllipse(varPointUncertaintyEllipse)

	return err
}

type NullablePointUncertaintyEllipse struct {
	value *PointUncertaintyEllipse
	isSet bool
}

func (v NullablePointUncertaintyEllipse) Get() *PointUncertaintyEllipse {
	return v.value
}

func (v *NullablePointUncertaintyEllipse) Set(val *PointUncertaintyEllipse) {
	v.value = val
	v.isSet = true
}

func (v NullablePointUncertaintyEllipse) IsSet() bool {
	return v.isSet
}

func (v *NullablePointUncertaintyEllipse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePointUncertaintyEllipse(val *PointUncertaintyEllipse) *NullablePointUncertaintyEllipse {
	return &NullablePointUncertaintyEllipse{value: val, isSet: true}
}

func (v NullablePointUncertaintyEllipse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePointUncertaintyEllipse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
