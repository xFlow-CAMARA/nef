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

// checks if the Ipv4AddressRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ipv4AddressRange{}

// Ipv4AddressRange Range of IPv4 addresses
type Ipv4AddressRange struct {
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	Start string `json:"start" validate:"regexp=^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	End string `json:"end" validate:"regexp=^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
}

type _Ipv4AddressRange Ipv4AddressRange

// NewIpv4AddressRange instantiates a new Ipv4AddressRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpv4AddressRange(start string, end string) *Ipv4AddressRange {
	this := Ipv4AddressRange{}
	this.Start = start
	this.End = end
	return &this
}

// NewIpv4AddressRangeWithDefaults instantiates a new Ipv4AddressRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpv4AddressRangeWithDefaults() *Ipv4AddressRange {
	this := Ipv4AddressRange{}
	return &this
}

// GetStart returns the Start field value
func (o *Ipv4AddressRange) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *Ipv4AddressRange) GetStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *Ipv4AddressRange) SetStart(v string) {
	o.Start = v
}

// GetEnd returns the End field value
func (o *Ipv4AddressRange) GetEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *Ipv4AddressRange) GetEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *Ipv4AddressRange) SetEnd(v string) {
	o.End = v
}

func (o Ipv4AddressRange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ipv4AddressRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start"] = o.Start
	toSerialize["end"] = o.End
	return toSerialize, nil
}

func (o *Ipv4AddressRange) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"start",
		"end",
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

	varIpv4AddressRange := _Ipv4AddressRange{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIpv4AddressRange)

	if err != nil {
		return err
	}

	*o = Ipv4AddressRange(varIpv4AddressRange)

	return err
}

type NullableIpv4AddressRange struct {
	value *Ipv4AddressRange
	isSet bool
}

func (v NullableIpv4AddressRange) Get() *Ipv4AddressRange {
	return v.value
}

func (v *NullableIpv4AddressRange) Set(val *Ipv4AddressRange) {
	v.value = val
	v.isSet = true
}

func (v NullableIpv4AddressRange) IsSet() bool {
	return v.isSet
}

func (v *NullableIpv4AddressRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpv4AddressRange(val *Ipv4AddressRange) *NullableIpv4AddressRange {
	return &NullableIpv4AddressRange{value: val, isSet: true}
}

func (v NullableIpv4AddressRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpv4AddressRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
