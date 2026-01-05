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

// checks if the Local2dPointUncertaintyEllipse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Local2dPointUncertaintyEllipse{}

// Local2dPointUncertaintyEllipse Local 2D point with uncertainty ellipse
type Local2dPointUncertaintyEllipse struct {
	GADShape
	LocalOrigin        LocalOrigin               `json:"localOrigin"`
	Point              RelativeCartesianLocation `json:"point"`
	UncertaintyEllipse UncertaintyEllipse        `json:"uncertaintyEllipse"`
	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`
}

type _Local2dPointUncertaintyEllipse Local2dPointUncertaintyEllipse

// NewLocal2dPointUncertaintyEllipse instantiates a new Local2dPointUncertaintyEllipse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocal2dPointUncertaintyEllipse(localOrigin LocalOrigin, point RelativeCartesianLocation, uncertaintyEllipse UncertaintyEllipse, confidence int32, shape SupportedGADShapes) *Local2dPointUncertaintyEllipse {
	this := Local2dPointUncertaintyEllipse{}
	this.Shape = shape
	this.LocalOrigin = localOrigin
	this.Point = point
	this.UncertaintyEllipse = uncertaintyEllipse
	this.Confidence = confidence
	return &this
}

// NewLocal2dPointUncertaintyEllipseWithDefaults instantiates a new Local2dPointUncertaintyEllipse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocal2dPointUncertaintyEllipseWithDefaults() *Local2dPointUncertaintyEllipse {
	this := Local2dPointUncertaintyEllipse{}
	return &this
}

// GetLocalOrigin returns the LocalOrigin field value
func (o *Local2dPointUncertaintyEllipse) GetLocalOrigin() LocalOrigin {
	if o == nil {
		var ret LocalOrigin
		return ret
	}

	return o.LocalOrigin
}

// GetLocalOriginOk returns a tuple with the LocalOrigin field value
// and a boolean to check if the value has been set.
func (o *Local2dPointUncertaintyEllipse) GetLocalOriginOk() (*LocalOrigin, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocalOrigin, true
}

// SetLocalOrigin sets field value
func (o *Local2dPointUncertaintyEllipse) SetLocalOrigin(v LocalOrigin) {
	o.LocalOrigin = v
}

// GetPoint returns the Point field value
func (o *Local2dPointUncertaintyEllipse) GetPoint() RelativeCartesianLocation {
	if o == nil {
		var ret RelativeCartesianLocation
		return ret
	}

	return o.Point
}

// GetPointOk returns a tuple with the Point field value
// and a boolean to check if the value has been set.
func (o *Local2dPointUncertaintyEllipse) GetPointOk() (*RelativeCartesianLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Point, true
}

// SetPoint sets field value
func (o *Local2dPointUncertaintyEllipse) SetPoint(v RelativeCartesianLocation) {
	o.Point = v
}

// GetUncertaintyEllipse returns the UncertaintyEllipse field value
func (o *Local2dPointUncertaintyEllipse) GetUncertaintyEllipse() UncertaintyEllipse {
	if o == nil {
		var ret UncertaintyEllipse
		return ret
	}

	return o.UncertaintyEllipse
}

// GetUncertaintyEllipseOk returns a tuple with the UncertaintyEllipse field value
// and a boolean to check if the value has been set.
func (o *Local2dPointUncertaintyEllipse) GetUncertaintyEllipseOk() (*UncertaintyEllipse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UncertaintyEllipse, true
}

// SetUncertaintyEllipse sets field value
func (o *Local2dPointUncertaintyEllipse) SetUncertaintyEllipse(v UncertaintyEllipse) {
	o.UncertaintyEllipse = v
}

// GetConfidence returns the Confidence field value
func (o *Local2dPointUncertaintyEllipse) GetConfidence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value
// and a boolean to check if the value has been set.
func (o *Local2dPointUncertaintyEllipse) GetConfidenceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Confidence, true
}

// SetConfidence sets field value
func (o *Local2dPointUncertaintyEllipse) SetConfidence(v int32) {
	o.Confidence = v
}

func (o Local2dPointUncertaintyEllipse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Local2dPointUncertaintyEllipse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedGADShape, errGADShape := json.Marshal(o.GADShape)
	if errGADShape != nil {
		return map[string]interface{}{}, errGADShape
	}
	errGADShape = json.Unmarshal([]byte(serializedGADShape), &toSerialize)
	if errGADShape != nil {
		return map[string]interface{}{}, errGADShape
	}
	toSerialize["localOrigin"] = o.LocalOrigin
	toSerialize["point"] = o.Point
	toSerialize["uncertaintyEllipse"] = o.UncertaintyEllipse
	toSerialize["confidence"] = o.Confidence
	return toSerialize, nil
}

func (o *Local2dPointUncertaintyEllipse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"localOrigin",
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

	varLocal2dPointUncertaintyEllipse := _Local2dPointUncertaintyEllipse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLocal2dPointUncertaintyEllipse)

	if err != nil {
		return err
	}

	*o = Local2dPointUncertaintyEllipse(varLocal2dPointUncertaintyEllipse)

	return err
}

type NullableLocal2dPointUncertaintyEllipse struct {
	value *Local2dPointUncertaintyEllipse
	isSet bool
}

func (v NullableLocal2dPointUncertaintyEllipse) Get() *Local2dPointUncertaintyEllipse {
	return v.value
}

func (v *NullableLocal2dPointUncertaintyEllipse) Set(val *Local2dPointUncertaintyEllipse) {
	v.value = val
	v.isSet = true
}

func (v NullableLocal2dPointUncertaintyEllipse) IsSet() bool {
	return v.isSet
}

func (v *NullableLocal2dPointUncertaintyEllipse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocal2dPointUncertaintyEllipse(val *Local2dPointUncertaintyEllipse) *NullableLocal2dPointUncertaintyEllipse {
	return &NullableLocal2dPointUncertaintyEllipse{value: val, isSet: true}
}

func (v NullableLocal2dPointUncertaintyEllipse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocal2dPointUncertaintyEllipse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
