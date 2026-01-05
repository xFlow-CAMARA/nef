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

// checks if the Polygon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Polygon{}

// Polygon Polygon.
type Polygon struct {
	GADShape
	// List of points.
	PointList []GeographicalCoordinates `json:"pointList"`
}

type _Polygon Polygon

// NewPolygon instantiates a new Polygon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolygon(pointList []GeographicalCoordinates, shape SupportedGADShapes) *Polygon {
	this := Polygon{}
	this.Shape = shape
	this.PointList = pointList
	return &this
}

// NewPolygonWithDefaults instantiates a new Polygon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolygonWithDefaults() *Polygon {
	this := Polygon{}
	return &this
}

// GetPointList returns the PointList field value
func (o *Polygon) GetPointList() []GeographicalCoordinates {
	if o == nil {
		var ret []GeographicalCoordinates
		return ret
	}

	return o.PointList
}

// GetPointListOk returns a tuple with the PointList field value
// and a boolean to check if the value has been set.
func (o *Polygon) GetPointListOk() ([]GeographicalCoordinates, bool) {
	if o == nil {
		return nil, false
	}
	return o.PointList, true
}

// SetPointList sets field value
func (o *Polygon) SetPointList(v []GeographicalCoordinates) {
	o.PointList = v
}

func (o Polygon) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Polygon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedGADShape, errGADShape := json.Marshal(o.GADShape)
	if errGADShape != nil {
		return map[string]interface{}{}, errGADShape
	}
	errGADShape = json.Unmarshal([]byte(serializedGADShape), &toSerialize)
	if errGADShape != nil {
		return map[string]interface{}{}, errGADShape
	}
	toSerialize["pointList"] = o.PointList
	return toSerialize, nil
}

func (o *Polygon) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pointList",
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

	varPolygon := _Polygon{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPolygon)

	if err != nil {
		return err
	}

	*o = Polygon(varPolygon)

	return err
}

type NullablePolygon struct {
	value *Polygon
	isSet bool
}

func (v NullablePolygon) Get() *Polygon {
	return v.value
}

func (v *NullablePolygon) Set(val *Polygon) {
	v.value = val
	v.isSet = true
}

func (v NullablePolygon) IsSet() bool {
	return v.isSet
}

func (v *NullablePolygon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolygon(val *Polygon) *NullablePolygon {
	return &NullablePolygon{value: val, isSet: true}
}

func (v NullablePolygon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolygon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
