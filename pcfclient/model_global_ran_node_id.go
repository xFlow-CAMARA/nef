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

// checks if the GlobalRanNodeId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalRanNodeId{}

// GlobalRanNodeId One of the six attributes n3IwfId, gNbIdm, ngeNbId, wagfId, tngfId, eNbId shall be present.
type GlobalRanNodeId struct {
	PlmnId PlmnId
	// This represents the identifier of the N3IWF ID as specified in clause 9.3.1.57 of  3GPP TS 38.413 in hexadecimal representation. Each character in the string shall take a value  of \"0\" to \"9\", \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The most significant  character representing the 4 most significant bits of the N3IWF ID shall appear first in the  string, and the character representing the 4 least significant bit of the N3IWF ID shall  appear last in the string.
	N3IwfId *string
	GNbId   *GNbId
	// This represents the identifier of the ng-eNB ID as specified in clause 9.3.1.8 of  3GPP TS 38.413. The value of the ng-eNB ID shall be encoded in hexadecimal representation.  Each character in the string shall take a value of \"0\" to \"9\", \"a\" to \"f\" or \"A\" to \"F\" and  shall represent 4 bits. The padding 0 shall be added to make multiple nibbles, so the most  significant character representing the padding 0 if required together with the 4 most  significant bits of the ng-eNB ID shall appear first in the string, and the character  representing the 4 least significant bit of the ng-eNB ID (to form a nibble) shall appear last  in the string.
	NgeNbId *string
	// This represents the identifier of the W-AGF ID as specified in clause 9.3.1.162 of  3GPP TS 38.413 in hexadecimal representation. Each character in the string shall take a value  of \"0\" to \"9\", \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The most significant  character representing the 4 most significant bits of the W-AGF ID shall appear first in the  string, and the character representing the 4 least significant bit of the W-AGF ID shall  appear last in the string.
	WagfId *string
	// This represents the identifier of the TNGF ID as specified in clause 9.3.1.161 of  3GPP TS 38.413  in hexadecimal representation. Each character in the string shall take a value of \"0\" to \"9\", \"a\"  to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The most significant character representing the  4 most significant bits of the TNGF ID shall appear first in the string, and the character  representing the 4 least significant bit of the TNGF ID shall appear last in the string.
	TngfId *string
	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).
	Nid *string
	// This represents the identifier of the eNB ID as specified in clause 9.2.1.37 of  3GPP TS 36.413. The string shall be formatted with the following pattern  '^('MacroeNB-[A-Fa-f0-9]{5}|LMacroeNB-[A-Fa-f0-9]{6}|SMacroeNB-[A-Fa-f0-9]{5} |HomeeNB-[A-Fa-f0-9]{7})$'. The value of the eNB ID shall be encoded in hexadecimal representation. Each character in the  string shall take a value of \"0\" to \"9\", \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits.  The padding 0 shall be added to make multiple nibbles, so the most significant character  representing the padding 0 if required together with the 4 most significant bits of the eNB ID  shall appear first in the string, and the character representing the 4 least significant bit  of the eNB ID (to form a nibble) shall appear last in the string.
	ENbId *string
}

type _GlobalRanNodeId GlobalRanNodeId

// NewGlobalRanNodeId instantiates a new GlobalRanNodeId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalRanNodeId(plmnId PlmnId) *GlobalRanNodeId {
	this := GlobalRanNodeId{}
	return &this
}

// NewGlobalRanNodeIdWithDefaults instantiates a new GlobalRanNodeId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalRanNodeIdWithDefaults() *GlobalRanNodeId {
	this := GlobalRanNodeId{}
	return &this
}

// GetPlmnId returns the PlmnId field value
func (o *GlobalRanNodeId) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *GlobalRanNodeId) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *GlobalRanNodeId) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

// GetN3IwfId returns the N3IwfId field value if set, zero value otherwise.
func (o *GlobalRanNodeId) GetN3IwfId() string {
	if o == nil || IsNil(o.N3IwfId) {
		var ret string
		return ret
	}
	return *o.N3IwfId
}

// GetN3IwfIdOk returns a tuple with the N3IwfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalRanNodeId) GetN3IwfIdOk() (*string, bool) {
	if o == nil || IsNil(o.N3IwfId) {
		return nil, false
	}
	return o.N3IwfId, true
}

// HasN3IwfId returns a boolean if a field has been set.
func (o *GlobalRanNodeId) HasN3IwfId() bool {
	if o != nil && !IsNil(o.N3IwfId) {
		return true
	}

	return false
}

// SetN3IwfId gets a reference to the given string and assigns it to the N3IwfId field.
func (o *GlobalRanNodeId) SetN3IwfId(v string) {
	o.N3IwfId = &v
}

// GetGNbId returns the GNbId field value if set, zero value otherwise.
func (o *GlobalRanNodeId) GetGNbId() GNbId {
	if o == nil || IsNil(o.GNbId) {
		var ret GNbId
		return ret
	}
	return *o.GNbId
}

// GetGNbIdOk returns a tuple with the GNbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalRanNodeId) GetGNbIdOk() (*GNbId, bool) {
	if o == nil || IsNil(o.GNbId) {
		return nil, false
	}
	return o.GNbId, true
}

// HasGNbId returns a boolean if a field has been set.
func (o *GlobalRanNodeId) HasGNbId() bool {
	if o != nil && !IsNil(o.GNbId) {
		return true
	}

	return false
}

// SetGNbId gets a reference to the given GNbId and assigns it to the GNbId field.
func (o *GlobalRanNodeId) SetGNbId(v GNbId) {
	o.GNbId = &v
}

// GetNgeNbId returns the NgeNbId field value if set, zero value otherwise.
func (o *GlobalRanNodeId) GetNgeNbId() string {
	if o == nil || IsNil(o.NgeNbId) {
		var ret string
		return ret
	}
	return *o.NgeNbId
}

// GetNgeNbIdOk returns a tuple with the NgeNbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalRanNodeId) GetNgeNbIdOk() (*string, bool) {
	if o == nil || IsNil(o.NgeNbId) {
		return nil, false
	}
	return o.NgeNbId, true
}

// HasNgeNbId returns a boolean if a field has been set.
func (o *GlobalRanNodeId) HasNgeNbId() bool {
	if o != nil && !IsNil(o.NgeNbId) {
		return true
	}

	return false
}

// SetNgeNbId gets a reference to the given string and assigns it to the NgeNbId field.
func (o *GlobalRanNodeId) SetNgeNbId(v string) {
	o.NgeNbId = &v
}

// GetWagfId returns the WagfId field value if set, zero value otherwise.
func (o *GlobalRanNodeId) GetWagfId() string {
	if o == nil || IsNil(o.WagfId) {
		var ret string
		return ret
	}
	return *o.WagfId
}

// GetWagfIdOk returns a tuple with the WagfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalRanNodeId) GetWagfIdOk() (*string, bool) {
	if o == nil || IsNil(o.WagfId) {
		return nil, false
	}
	return o.WagfId, true
}

// HasWagfId returns a boolean if a field has been set.
func (o *GlobalRanNodeId) HasWagfId() bool {
	if o != nil && !IsNil(o.WagfId) {
		return true
	}

	return false
}

// SetWagfId gets a reference to the given string and assigns it to the WagfId field.
func (o *GlobalRanNodeId) SetWagfId(v string) {
	o.WagfId = &v
}

// GetTngfId returns the TngfId field value if set, zero value otherwise.
func (o *GlobalRanNodeId) GetTngfId() string {
	if o == nil || IsNil(o.TngfId) {
		var ret string
		return ret
	}
	return *o.TngfId
}

// GetTngfIdOk returns a tuple with the TngfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalRanNodeId) GetTngfIdOk() (*string, bool) {
	if o == nil || IsNil(o.TngfId) {
		return nil, false
	}
	return o.TngfId, true
}

// HasTngfId returns a boolean if a field has been set.
func (o *GlobalRanNodeId) HasTngfId() bool {
	if o != nil && !IsNil(o.TngfId) {
		return true
	}

	return false
}

// SetTngfId gets a reference to the given string and assigns it to the TngfId field.
func (o *GlobalRanNodeId) SetTngfId(v string) {
	o.TngfId = &v
}

// GetNid returns the Nid field value if set, zero value otherwise.
func (o *GlobalRanNodeId) GetNid() string {
	if o == nil || IsNil(o.Nid) {
		var ret string
		return ret
	}
	return *o.Nid
}

// GetNidOk returns a tuple with the Nid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalRanNodeId) GetNidOk() (*string, bool) {
	if o == nil || IsNil(o.Nid) {
		return nil, false
	}
	return o.Nid, true
}

// HasNid returns a boolean if a field has been set.
func (o *GlobalRanNodeId) HasNid() bool {
	if o != nil && !IsNil(o.Nid) {
		return true
	}

	return false
}

// SetNid gets a reference to the given string and assigns it to the Nid field.
func (o *GlobalRanNodeId) SetNid(v string) {
	o.Nid = &v
}

// GetENbId returns the ENbId field value if set, zero value otherwise.
func (o *GlobalRanNodeId) GetENbId() string {
	if o == nil || IsNil(o.ENbId) {
		var ret string
		return ret
	}
	return *o.ENbId
}

// GetENbIdOk returns a tuple with the ENbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalRanNodeId) GetENbIdOk() (*string, bool) {
	if o == nil || IsNil(o.ENbId) {
		return nil, false
	}
	return o.ENbId, true
}

// HasENbId returns a boolean if a field has been set.
func (o *GlobalRanNodeId) HasENbId() bool {
	if o != nil && !IsNil(o.ENbId) {
		return true
	}

	return false
}

// SetENbId gets a reference to the given string and assigns it to the ENbId field.
func (o *GlobalRanNodeId) SetENbId(v string) {
	o.ENbId = &v
}

func (o GlobalRanNodeId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlobalRanNodeId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plmnId"] = o.PlmnId
	if !IsNil(o.N3IwfId) {
		toSerialize["n3IwfId"] = o.N3IwfId
	}
	if !IsNil(o.GNbId) {
		toSerialize["gNbId"] = o.GNbId
	}
	if !IsNil(o.NgeNbId) {
		toSerialize["ngeNbId"] = o.NgeNbId
	}
	if !IsNil(o.WagfId) {
		toSerialize["wagfId"] = o.WagfId
	}
	if !IsNil(o.TngfId) {
		toSerialize["tngfId"] = o.TngfId
	}
	if !IsNil(o.Nid) {
		toSerialize["nid"] = o.Nid
	}
	if !IsNil(o.ENbId) {
		toSerialize["eNbId"] = o.ENbId
	}
	return toSerialize, nil
}

func (o *GlobalRanNodeId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"plmnId",
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

	varGlobalRanNodeId := _GlobalRanNodeId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGlobalRanNodeId)

	if err != nil {
		return err
	}

	*o = GlobalRanNodeId(varGlobalRanNodeId)

	return err
}

type NullableGlobalRanNodeId struct {
	value *GlobalRanNodeId
	isSet bool
}

func (v NullableGlobalRanNodeId) Get() *GlobalRanNodeId {
	return v.value
}

func (v *NullableGlobalRanNodeId) Set(val *GlobalRanNodeId) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalRanNodeId) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalRanNodeId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalRanNodeId(val *GlobalRanNodeId) *NullableGlobalRanNodeId {
	return &NullableGlobalRanNodeId{value: val, isSet: true}
}

func (v NullableGlobalRanNodeId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalRanNodeId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
