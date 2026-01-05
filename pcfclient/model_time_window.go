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
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.6
*/

package pcfclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the TimeWindow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeWindow{}

// TimeWindow Represents a time window identified by a start time and a stop time.
type TimeWindow struct {
	// string with format \"date-time\" as defined in OpenAPI.
	StartTime time.Time `json:"startTime"`
	// string with format \"date-time\" as defined in OpenAPI.
	StopTime time.Time `json:"stopTime"`
}

type _TimeWindow TimeWindow

// NewTimeWindow instantiates a new TimeWindow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeWindow(startTime time.Time, stopTime time.Time) *TimeWindow {
	this := TimeWindow{}
	this.StartTime = startTime
	this.StopTime = stopTime
	return &this
}

// NewTimeWindowWithDefaults instantiates a new TimeWindow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeWindowWithDefaults() *TimeWindow {
	this := TimeWindow{}
	return &this
}

// GetStartTime returns the StartTime field value
func (o *TimeWindow) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *TimeWindow) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *TimeWindow) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetStopTime returns the StopTime field value
func (o *TimeWindow) GetStopTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StopTime
}

// GetStopTimeOk returns a tuple with the StopTime field value
// and a boolean to check if the value has been set.
func (o *TimeWindow) GetStopTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StopTime, true
}

// SetStopTime sets field value
func (o *TimeWindow) SetStopTime(v time.Time) {
	o.StopTime = v
}

func (o TimeWindow) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeWindow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startTime"] = o.StartTime
	toSerialize["stopTime"] = o.StopTime
	return toSerialize, nil
}

func (o *TimeWindow) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"startTime",
		"stopTime",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTimeWindow := _TimeWindow{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTimeWindow)

	if err != nil {
		return err
	}

	*o = TimeWindow(varTimeWindow)

	return err
}

type NullableTimeWindow struct {
	value *TimeWindow
	isSet bool
}

func (v NullableTimeWindow) Get() *TimeWindow {
	return v.value
}

func (v *NullableTimeWindow) Set(val *TimeWindow) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeWindow) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeWindow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeWindow(val *TimeWindow) *NullableTimeWindow {
	return &NullableTimeWindow{value: val, isSet: true}
}

func (v NullableTimeWindow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeWindow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
