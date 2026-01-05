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

// checks if the TwapId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TwapId{}

// TwapId Contain the TWAP Identifier as defined in clause 4.2.8.5.3 of 3GPP TS 23.501 or the WLAN location information as defined in clause 4.5.7.2.8 of 3GPP TS 23.402.
type TwapId struct {
	// This IE shall contain the SSID of the access point to which the UE is attached, that is received over NGAP, see IEEE Std 802.11-2012.
	SsId string `json:"ssId"`
	// When present, it shall contain the BSSID of the access point to which the UE is attached, for trusted WLAN access, see IEEE Std 802.11-2012.
	BssId *string `json:"bssId,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	CivicAddress *string `json:"civicAddress,omitempty"`
}

type _TwapId TwapId

// NewTwapId instantiates a new TwapId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTwapId(ssId string) *TwapId {
	this := TwapId{}
	this.SsId = ssId
	return &this
}

// NewTwapIdWithDefaults instantiates a new TwapId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTwapIdWithDefaults() *TwapId {
	this := TwapId{}
	return &this
}

// GetSsId returns the SsId field value
func (o *TwapId) GetSsId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SsId
}

// GetSsIdOk returns a tuple with the SsId field value
// and a boolean to check if the value has been set.
func (o *TwapId) GetSsIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SsId, true
}

// SetSsId sets field value
func (o *TwapId) SetSsId(v string) {
	o.SsId = v
}

// GetBssId returns the BssId field value if set, zero value otherwise.
func (o *TwapId) GetBssId() string {
	if o == nil || IsNil(o.BssId) {
		var ret string
		return ret
	}
	return *o.BssId
}

// GetBssIdOk returns a tuple with the BssId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwapId) GetBssIdOk() (*string, bool) {
	if o == nil || IsNil(o.BssId) {
		return nil, false
	}
	return o.BssId, true
}

// HasBssId returns a boolean if a field has been set.
func (o *TwapId) HasBssId() bool {
	if o != nil && !IsNil(o.BssId) {
		return true
	}

	return false
}

// SetBssId gets a reference to the given string and assigns it to the BssId field.
func (o *TwapId) SetBssId(v string) {
	o.BssId = &v
}

// GetCivicAddress returns the CivicAddress field value if set, zero value otherwise.
func (o *TwapId) GetCivicAddress() string {
	if o == nil || IsNil(o.CivicAddress) {
		var ret string
		return ret
	}
	return *o.CivicAddress
}

// GetCivicAddressOk returns a tuple with the CivicAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwapId) GetCivicAddressOk() (*string, bool) {
	if o == nil || IsNil(o.CivicAddress) {
		return nil, false
	}
	return o.CivicAddress, true
}

// HasCivicAddress returns a boolean if a field has been set.
func (o *TwapId) HasCivicAddress() bool {
	if o != nil && !IsNil(o.CivicAddress) {
		return true
	}

	return false
}

// SetCivicAddress gets a reference to the given string and assigns it to the CivicAddress field.
func (o *TwapId) SetCivicAddress(v string) {
	o.CivicAddress = &v
}

func (o TwapId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TwapId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ssId"] = o.SsId
	if !IsNil(o.BssId) {
		toSerialize["bssId"] = o.BssId
	}
	if !IsNil(o.CivicAddress) {
		toSerialize["civicAddress"] = o.CivicAddress
	}
	return toSerialize, nil
}

func (o *TwapId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ssId",
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

	varTwapId := _TwapId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTwapId)

	if err != nil {
		return err
	}

	*o = TwapId(varTwapId)

	return err
}

type NullableTwapId struct {
	value *TwapId
	isSet bool
}

func (v NullableTwapId) Get() *TwapId {
	return v.value
}

func (v *NullableTwapId) Set(val *TwapId) {
	v.value = val
	v.isSet = true
}

func (v NullableTwapId) IsSet() bool {
	return v.isSet
}

func (v *NullableTwapId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTwapId(val *TwapId) *NullableTwapId {
	return &NullableTwapId{value: val, isSet: true}
}

func (v NullableTwapId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTwapId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
