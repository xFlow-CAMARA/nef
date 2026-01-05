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
NRF NFDiscovery Service

NRF NFDiscovery Service. © 2019, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.0.alpha-1
*/

package nrfclient

import (
	"encoding/json"
)

// SmfInfo struct for SmfInfo
type SmfInfo struct {
	SNssaiSmfInfoList []SnssaiSmfInfoItem `json:"sNssaiSmfInfoList"`
	TaiList           []Tai               `json:"taiList,omitempty"`
	TaiRangeList      []TaiRange          `json:"taiRangeList,omitempty"`
	PgwFqdn           *string             `json:"pgwFqdn,omitempty"`
	AccessType        []AccessType        `json:"accessType,omitempty"`
}

// NewSmfInfo instantiates a new SmfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmfInfo(sNssaiSmfInfoList []SnssaiSmfInfoItem) *SmfInfo {
	this := SmfInfo{}
	this.SNssaiSmfInfoList = sNssaiSmfInfoList
	return &this
}

// NewSmfInfoWithDefaults instantiates a new SmfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmfInfoWithDefaults() *SmfInfo {
	this := SmfInfo{}
	return &this
}

// GetSNssaiSmfInfoList returns the SNssaiSmfInfoList field value
func (o *SmfInfo) GetSNssaiSmfInfoList() []SnssaiSmfInfoItem {
	if o == nil {
		var ret []SnssaiSmfInfoItem
		return ret
	}

	return o.SNssaiSmfInfoList
}

// GetSNssaiSmfInfoListOk returns a tuple with the SNssaiSmfInfoList field value
// and a boolean to check if the value has been set.
func (o *SmfInfo) GetSNssaiSmfInfoListOk() ([]SnssaiSmfInfoItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.SNssaiSmfInfoList, true
}

// SetSNssaiSmfInfoList sets field value
func (o *SmfInfo) SetSNssaiSmfInfoList(v []SnssaiSmfInfoItem) {
	o.SNssaiSmfInfoList = v
}

// GetTaiList returns the TaiList field value if set, zero value otherwise.
func (o *SmfInfo) GetTaiList() []Tai {
	if o == nil || o.TaiList == nil {
		var ret []Tai
		return ret
	}
	return o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfInfo) GetTaiListOk() ([]Tai, bool) {
	if o == nil || o.TaiList == nil {
		return nil, false
	}
	return o.TaiList, true
}

// HasTaiList returns a boolean if a field has been set.
func (o *SmfInfo) HasTaiList() bool {
	if o != nil && o.TaiList != nil {
		return true
	}

	return false
}

// SetTaiList gets a reference to the given []Tai and assigns it to the TaiList field.
func (o *SmfInfo) SetTaiList(v []Tai) {
	o.TaiList = v
}

// GetTaiRangeList returns the TaiRangeList field value if set, zero value otherwise.
func (o *SmfInfo) GetTaiRangeList() []TaiRange {
	if o == nil || o.TaiRangeList == nil {
		var ret []TaiRange
		return ret
	}
	return o.TaiRangeList
}

// GetTaiRangeListOk returns a tuple with the TaiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfInfo) GetTaiRangeListOk() ([]TaiRange, bool) {
	if o == nil || o.TaiRangeList == nil {
		return nil, false
	}
	return o.TaiRangeList, true
}

// HasTaiRangeList returns a boolean if a field has been set.
func (o *SmfInfo) HasTaiRangeList() bool {
	if o != nil && o.TaiRangeList != nil {
		return true
	}

	return false
}

// SetTaiRangeList gets a reference to the given []TaiRange and assigns it to the TaiRangeList field.
func (o *SmfInfo) SetTaiRangeList(v []TaiRange) {
	o.TaiRangeList = v
}

// GetPgwFqdn returns the PgwFqdn field value if set, zero value otherwise.
func (o *SmfInfo) GetPgwFqdn() string {
	if o == nil || o.PgwFqdn == nil {
		var ret string
		return ret
	}
	return *o.PgwFqdn
}

// GetPgwFqdnOk returns a tuple with the PgwFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfInfo) GetPgwFqdnOk() (*string, bool) {
	if o == nil || o.PgwFqdn == nil {
		return nil, false
	}
	return o.PgwFqdn, true
}

// HasPgwFqdn returns a boolean if a field has been set.
func (o *SmfInfo) HasPgwFqdn() bool {
	if o != nil && o.PgwFqdn != nil {
		return true
	}

	return false
}

// SetPgwFqdn gets a reference to the given string and assigns it to the PgwFqdn field.
func (o *SmfInfo) SetPgwFqdn(v string) {
	o.PgwFqdn = &v
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *SmfInfo) GetAccessType() []AccessType {
	if o == nil || o.AccessType == nil {
		var ret []AccessType
		return ret
	}
	return o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfInfo) GetAccessTypeOk() ([]AccessType, bool) {
	if o == nil || o.AccessType == nil {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *SmfInfo) HasAccessType() bool {
	if o != nil && o.AccessType != nil {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given []AccessType and assigns it to the AccessType field.
func (o *SmfInfo) SetAccessType(v []AccessType) {
	o.AccessType = v
}

func (o SmfInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sNssaiSmfInfoList"] = o.SNssaiSmfInfoList
	}
	if o.TaiList != nil {
		toSerialize["taiList"] = o.TaiList
	}
	if o.TaiRangeList != nil {
		toSerialize["taiRangeList"] = o.TaiRangeList
	}
	if o.PgwFqdn != nil {
		toSerialize["pgwFqdn"] = o.PgwFqdn
	}
	if o.AccessType != nil {
		toSerialize["accessType"] = o.AccessType
	}
	return json.Marshal(toSerialize)
}

type NullableSmfInfo struct {
	value *SmfInfo
	isSet bool
}

func (v NullableSmfInfo) Get() *SmfInfo {
	return v.value
}

func (v *NullableSmfInfo) Set(val *SmfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSmfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSmfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmfInfo(val *SmfInfo) *NullableSmfInfo {
	return &NullableSmfInfo{value: val, isSet: true}
}

func (v NullableSmfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
