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

// checks if the PlmnIdNid type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlmnIdNid{}

// PlmnIdNid Contains the serving core network operator PLMN ID and, for an SNPN, the NID that together with the PLMN ID identifies the SNPN.
type PlmnIdNid struct {
	// Mobile Country Code part of the PLMN, comprising 3 digits, as defined in clause 9.3.3.5 of 3GPP TS 38.413.
	Mcc string `json:"mcc" validate:"regexp=^\\\\d{3}$"`
	// Mobile Network Code part of the PLMN, comprising 2 or 3 digits, as defined in  clause 9.3.3.5 of 3GPP TS 38.413.
	Mnc string `json:"mnc" validate:"regexp=^\\\\d{2,3}$"`
	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).
	Nid *string `json:"nid,omitempty" validate:"regexp=^[A-Fa-f0-9]{11}$"`
}

type _PlmnIdNid PlmnIdNid

// NewPlmnIdNid instantiates a new PlmnIdNid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlmnIdNid(mcc string, mnc string) *PlmnIdNid {
	this := PlmnIdNid{}
	this.Mcc = mcc
	this.Mnc = mnc
	return &this
}

// NewPlmnIdNidWithDefaults instantiates a new PlmnIdNid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlmnIdNidWithDefaults() *PlmnIdNid {
	this := PlmnIdNid{}
	return &this
}

// GetMcc returns the Mcc field value
func (o *PlmnIdNid) GetMcc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value
// and a boolean to check if the value has been set.
func (o *PlmnIdNid) GetMccOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mcc, true
}

// SetMcc sets field value
func (o *PlmnIdNid) SetMcc(v string) {
	o.Mcc = v
}

// GetMnc returns the Mnc field value
func (o *PlmnIdNid) GetMnc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mnc
}

// GetMncOk returns a tuple with the Mnc field value
// and a boolean to check if the value has been set.
func (o *PlmnIdNid) GetMncOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mnc, true
}

// SetMnc sets field value
func (o *PlmnIdNid) SetMnc(v string) {
	o.Mnc = v
}

// GetNid returns the Nid field value if set, zero value otherwise.
func (o *PlmnIdNid) GetNid() string {
	if o == nil || IsNil(o.Nid) {
		var ret string
		return ret
	}
	return *o.Nid
}

// GetNidOk returns a tuple with the Nid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnIdNid) GetNidOk() (*string, bool) {
	if o == nil || IsNil(o.Nid) {
		return nil, false
	}
	return o.Nid, true
}

// HasNid returns a boolean if a field has been set.
func (o *PlmnIdNid) HasNid() bool {
	if o != nil && !IsNil(o.Nid) {
		return true
	}

	return false
}

// SetNid gets a reference to the given string and assigns it to the Nid field.
func (o *PlmnIdNid) SetNid(v string) {
	o.Nid = &v
}

func (o PlmnIdNid) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlmnIdNid) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mcc"] = o.Mcc
	toSerialize["mnc"] = o.Mnc
	if !IsNil(o.Nid) {
		toSerialize["nid"] = o.Nid
	}
	return toSerialize, nil
}

func (o *PlmnIdNid) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mcc",
		"mnc",
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

	varPlmnIdNid := _PlmnIdNid{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPlmnIdNid)

	if err != nil {
		return err
	}

	*o = PlmnIdNid(varPlmnIdNid)

	return err
}

type NullablePlmnIdNid struct {
	value *PlmnIdNid
	isSet bool
}

func (v NullablePlmnIdNid) Get() *PlmnIdNid {
	return v.value
}

func (v *NullablePlmnIdNid) Set(val *PlmnIdNid) {
	v.value = val
	v.isSet = true
}

func (v NullablePlmnIdNid) IsSet() bool {
	return v.isSet
}

func (v *NullablePlmnIdNid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlmnIdNid(val *PlmnIdNid) *NullablePlmnIdNid {
	return &NullablePlmnIdNid{value: val, isSet: true}
}

func (v NullablePlmnIdNid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlmnIdNid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
