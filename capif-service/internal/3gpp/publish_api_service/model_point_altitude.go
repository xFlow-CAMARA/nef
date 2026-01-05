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

// checks if the PointAltitude type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PointAltitude{}

// PointAltitude Ellipsoid point with altitude.
type PointAltitude struct {
	GADShape
	Point GeographicalCoordinates `json:"point"`
	// Indicates value of altitude.
	Altitude float64 `json:"altitude"`
}

type _PointAltitude PointAltitude

// NewPointAltitude instantiates a new PointAltitude object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPointAltitude(point GeographicalCoordinates, altitude float64, shape SupportedGADShapes) *PointAltitude {
	this := PointAltitude{}
	this.Shape = shape
	this.Point = point
	this.Altitude = altitude
	return &this
}

// NewPointAltitudeWithDefaults instantiates a new PointAltitude object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPointAltitudeWithDefaults() *PointAltitude {
	this := PointAltitude{}
	return &this
}

// GetPoint returns the Point field value
func (o *PointAltitude) GetPoint() GeographicalCoordinates {
	if o == nil {
		var ret GeographicalCoordinates
		return ret
	}

	return o.Point
}

// GetPointOk returns a tuple with the Point field value
// and a boolean to check if the value has been set.
func (o *PointAltitude) GetPointOk() (*GeographicalCoordinates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Point, true
}

// SetPoint sets field value
func (o *PointAltitude) SetPoint(v GeographicalCoordinates) {
	o.Point = v
}

// GetAltitude returns the Altitude field value
func (o *PointAltitude) GetAltitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Altitude
}

// GetAltitudeOk returns a tuple with the Altitude field value
// and a boolean to check if the value has been set.
func (o *PointAltitude) GetAltitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Altitude, true
}

// SetAltitude sets field value
func (o *PointAltitude) SetAltitude(v float64) {
	o.Altitude = v
}

func (o PointAltitude) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PointAltitude) ToMap() (map[string]interface{}, error) {
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
	toSerialize["altitude"] = o.Altitude
	return toSerialize, nil
}

func (o *PointAltitude) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"point",
		"altitude",
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

	varPointAltitude := _PointAltitude{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPointAltitude)

	if err != nil {
		return err
	}

	*o = PointAltitude(varPointAltitude)

	return err
}

type NullablePointAltitude struct {
	value *PointAltitude
	isSet bool
}

func (v NullablePointAltitude) Get() *PointAltitude {
	return v.value
}

func (v *NullablePointAltitude) Set(val *PointAltitude) {
	v.value = val
	v.isSet = true
}

func (v NullablePointAltitude) IsSet() bool {
	return v.isSet
}

func (v *NullablePointAltitude) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePointAltitude(val *PointAltitude) *NullablePointAltitude {
	return &NullablePointAltitude{value: val, isSet: true}
}

func (v NullablePointAltitude) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePointAltitude) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
