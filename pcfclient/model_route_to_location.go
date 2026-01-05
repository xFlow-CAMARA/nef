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
)

// checks if the RouteToLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteToLocation{}

// RouteToLocation At least one of the \"routeInfo\" attribute and the \"routeProfId\" attribute shall be included in the \"RouteToLocation\" data type.
type RouteToLocation struct {
	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	Dnai      string
	RouteInfo NullableRouteInformation
	// Identifies the routing profile Id.
	RouteProfId NullableString
}

type _RouteToLocation RouteToLocation

// NewRouteToLocation instantiates a new RouteToLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteToLocation(dnai string) *RouteToLocation {
	this := RouteToLocation{}
	return &this
}

// NewRouteToLocationWithDefaults instantiates a new RouteToLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteToLocationWithDefaults() *RouteToLocation {
	this := RouteToLocation{}
	return &this
}

// GetDnai returns the Dnai field value
func (o *RouteToLocation) GetDnai() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dnai
}

// GetDnaiOk returns a tuple with the Dnai field value
// and a boolean to check if the value has been set.
func (o *RouteToLocation) GetDnaiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dnai, true
}

// SetDnai sets field value
func (o *RouteToLocation) SetDnai(v string) {
	o.Dnai = v
}

// GetRouteInfo returns the RouteInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RouteToLocation) GetRouteInfo() RouteInformation {
	if o == nil || IsNil(o.RouteInfo.Get()) {
		var ret RouteInformation
		return ret
	}
	return *o.RouteInfo.Get()
}

// GetRouteInfoOk returns a tuple with the RouteInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RouteToLocation) GetRouteInfoOk() (*RouteInformation, bool) {
	if o == nil {
		return nil, false
	}
	return o.RouteInfo.Get(), o.RouteInfo.IsSet()
}

// HasRouteInfo returns a boolean if a field has been set.
func (o *RouteToLocation) HasRouteInfo() bool {
	if o != nil && o.RouteInfo.IsSet() {
		return true
	}

	return false
}

// SetRouteInfo gets a reference to the given NullableRouteInformation and assigns it to the RouteInfo field.
func (o *RouteToLocation) SetRouteInfo(v RouteInformation) {
	o.RouteInfo.Set(&v)
}

// SetRouteInfoNil sets the value for RouteInfo to be an explicit nil
func (o *RouteToLocation) SetRouteInfoNil() {
	o.RouteInfo.Set(nil)
}

// UnsetRouteInfo ensures that no value is present for RouteInfo, not even an explicit nil
func (o *RouteToLocation) UnsetRouteInfo() {
	o.RouteInfo.Unset()
}

// GetRouteProfId returns the RouteProfId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RouteToLocation) GetRouteProfId() string {
	if o == nil || IsNil(o.RouteProfId.Get()) {
		var ret string
		return ret
	}
	return *o.RouteProfId.Get()
}

// GetRouteProfIdOk returns a tuple with the RouteProfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RouteToLocation) GetRouteProfIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RouteProfId.Get(), o.RouteProfId.IsSet()
}

// HasRouteProfId returns a boolean if a field has been set.
func (o *RouteToLocation) HasRouteProfId() bool {
	if o != nil && o.RouteProfId.IsSet() {
		return true
	}

	return false
}

// SetRouteProfId gets a reference to the given NullableString and assigns it to the RouteProfId field.
func (o *RouteToLocation) SetRouteProfId(v string) {
	o.RouteProfId.Set(&v)
}

// SetRouteProfIdNil sets the value for RouteProfId to be an explicit nil
func (o *RouteToLocation) SetRouteProfIdNil() {
	o.RouteProfId.Set(nil)
}

// UnsetRouteProfId ensures that no value is present for RouteProfId, not even an explicit nil
func (o *RouteToLocation) UnsetRouteProfId() {
	o.RouteProfId.Unset()
}

func (o RouteToLocation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RouteToLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dnai"] = o.Dnai
	if o.RouteInfo.IsSet() {
		toSerialize["routeInfo"] = o.RouteInfo.Get()
	}
	if o.RouteProfId.IsSet() {
		toSerialize["routeProfId"] = o.RouteProfId.Get()
	}
	return toSerialize, nil
}

func (o *RouteToLocation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dnai",
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

	varRouteToLocation := _RouteToLocation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRouteToLocation)

	if err != nil {
		return err
	}

	*o = RouteToLocation(varRouteToLocation)

	return err
}

type NullableRouteToLocation struct {
	value *RouteToLocation
	isSet bool
}

func (v NullableRouteToLocation) Get() *RouteToLocation {
	return v.value
}

func (v *NullableRouteToLocation) Set(val *RouteToLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteToLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteToLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteToLocation(val *RouteToLocation) *NullableRouteToLocation {
	return &NullableRouteToLocation{value: val, isSet: true}
}

func (v NullableRouteToLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteToLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
