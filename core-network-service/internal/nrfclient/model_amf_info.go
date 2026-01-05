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

package nrfclient

import (
	"encoding/json"
)

// AmfInfo struct for AmfInfo
type AmfInfo struct {
	AmfSetId             string              `json:"amfSetId"`
	AmfRegionId          string              `json:"amfRegionId"`
	GuamiList            []Guami             `json:"guamiList"`
	TaiList              []Tai               `json:"taiList,omitempty"`
	TaiRangeList         []TaiRange          `json:"taiRangeList,omitempty"`
	BackupInfoAmfFailure []Guami             `json:"backupInfoAmfFailure,omitempty"`
	BackupInfoAmfRemoval []Guami             `json:"backupInfoAmfRemoval,omitempty"`
	N2InterfaceAmfInfo   *N2InterfaceAmfInfo `json:"n2InterfaceAmfInfo,omitempty"`
}

// NewAmfInfo instantiates a new AmfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfInfo(amfSetId string, amfRegionId string, guamiList []Guami) *AmfInfo {
	this := AmfInfo{}
	this.AmfSetId = amfSetId
	this.AmfRegionId = amfRegionId
	this.GuamiList = guamiList
	return &this
}

// NewAmfInfoWithDefaults instantiates a new AmfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfInfoWithDefaults() *AmfInfo {
	this := AmfInfo{}
	return &this
}

// GetAmfSetId returns the AmfSetId field value
func (o *AmfInfo) GetAmfSetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmfSetId
}

// GetAmfSetIdOk returns a tuple with the AmfSetId field value
// and a boolean to check if the value has been set.
func (o *AmfInfo) GetAmfSetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmfSetId, true
}

// SetAmfSetId sets field value
func (o *AmfInfo) SetAmfSetId(v string) {
	o.AmfSetId = v
}

// GetAmfRegionId returns the AmfRegionId field value
func (o *AmfInfo) GetAmfRegionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmfRegionId
}

// GetAmfRegionIdOk returns a tuple with the AmfRegionId field value
// and a boolean to check if the value has been set.
func (o *AmfInfo) GetAmfRegionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmfRegionId, true
}

// SetAmfRegionId sets field value
func (o *AmfInfo) SetAmfRegionId(v string) {
	o.AmfRegionId = v
}

// GetGuamiList returns the GuamiList field value
func (o *AmfInfo) GetGuamiList() []Guami {
	if o == nil {
		var ret []Guami
		return ret
	}

	return o.GuamiList
}

// GetGuamiListOk returns a tuple with the GuamiList field value
// and a boolean to check if the value has been set.
func (o *AmfInfo) GetGuamiListOk() ([]Guami, bool) {
	if o == nil {
		return nil, false
	}
	return o.GuamiList, true
}

// SetGuamiList sets field value
func (o *AmfInfo) SetGuamiList(v []Guami) {
	o.GuamiList = v
}

// GetTaiList returns the TaiList field value if set, zero value otherwise.
func (o *AmfInfo) GetTaiList() []Tai {
	if o == nil || o.TaiList == nil {
		var ret []Tai
		return ret
	}
	return o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfInfo) GetTaiListOk() ([]Tai, bool) {
	if o == nil || o.TaiList == nil {
		return nil, false
	}
	return o.TaiList, true
}

// HasTaiList returns a boolean if a field has been set.
func (o *AmfInfo) HasTaiList() bool {
	if o != nil && o.TaiList != nil {
		return true
	}

	return false
}

// SetTaiList gets a reference to the given []Tai and assigns it to the TaiList field.
func (o *AmfInfo) SetTaiList(v []Tai) {
	o.TaiList = v
}

// GetTaiRangeList returns the TaiRangeList field value if set, zero value otherwise.
func (o *AmfInfo) GetTaiRangeList() []TaiRange {
	if o == nil || o.TaiRangeList == nil {
		var ret []TaiRange
		return ret
	}
	return o.TaiRangeList
}

// GetTaiRangeListOk returns a tuple with the TaiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfInfo) GetTaiRangeListOk() ([]TaiRange, bool) {
	if o == nil || o.TaiRangeList == nil {
		return nil, false
	}
	return o.TaiRangeList, true
}

// HasTaiRangeList returns a boolean if a field has been set.
func (o *AmfInfo) HasTaiRangeList() bool {
	if o != nil && o.TaiRangeList != nil {
		return true
	}

	return false
}

// SetTaiRangeList gets a reference to the given []TaiRange and assigns it to the TaiRangeList field.
func (o *AmfInfo) SetTaiRangeList(v []TaiRange) {
	o.TaiRangeList = v
}

// GetBackupInfoAmfFailure returns the BackupInfoAmfFailure field value if set, zero value otherwise.
func (o *AmfInfo) GetBackupInfoAmfFailure() []Guami {
	if o == nil || o.BackupInfoAmfFailure == nil {
		var ret []Guami
		return ret
	}
	return o.BackupInfoAmfFailure
}

// GetBackupInfoAmfFailureOk returns a tuple with the BackupInfoAmfFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfInfo) GetBackupInfoAmfFailureOk() ([]Guami, bool) {
	if o == nil || o.BackupInfoAmfFailure == nil {
		return nil, false
	}
	return o.BackupInfoAmfFailure, true
}

// HasBackupInfoAmfFailure returns a boolean if a field has been set.
func (o *AmfInfo) HasBackupInfoAmfFailure() bool {
	if o != nil && o.BackupInfoAmfFailure != nil {
		return true
	}

	return false
}

// SetBackupInfoAmfFailure gets a reference to the given []Guami and assigns it to the BackupInfoAmfFailure field.
func (o *AmfInfo) SetBackupInfoAmfFailure(v []Guami) {
	o.BackupInfoAmfFailure = v
}

// GetBackupInfoAmfRemoval returns the BackupInfoAmfRemoval field value if set, zero value otherwise.
func (o *AmfInfo) GetBackupInfoAmfRemoval() []Guami {
	if o == nil || o.BackupInfoAmfRemoval == nil {
		var ret []Guami
		return ret
	}
	return o.BackupInfoAmfRemoval
}

// GetBackupInfoAmfRemovalOk returns a tuple with the BackupInfoAmfRemoval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfInfo) GetBackupInfoAmfRemovalOk() ([]Guami, bool) {
	if o == nil || o.BackupInfoAmfRemoval == nil {
		return nil, false
	}
	return o.BackupInfoAmfRemoval, true
}

// HasBackupInfoAmfRemoval returns a boolean if a field has been set.
func (o *AmfInfo) HasBackupInfoAmfRemoval() bool {
	if o != nil && o.BackupInfoAmfRemoval != nil {
		return true
	}

	return false
}

// SetBackupInfoAmfRemoval gets a reference to the given []Guami and assigns it to the BackupInfoAmfRemoval field.
func (o *AmfInfo) SetBackupInfoAmfRemoval(v []Guami) {
	o.BackupInfoAmfRemoval = v
}

// GetN2InterfaceAmfInfo returns the N2InterfaceAmfInfo field value if set, zero value otherwise.
func (o *AmfInfo) GetN2InterfaceAmfInfo() N2InterfaceAmfInfo {
	if o == nil || o.N2InterfaceAmfInfo == nil {
		var ret N2InterfaceAmfInfo
		return ret
	}
	return *o.N2InterfaceAmfInfo
}

// GetN2InterfaceAmfInfoOk returns a tuple with the N2InterfaceAmfInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfInfo) GetN2InterfaceAmfInfoOk() (*N2InterfaceAmfInfo, bool) {
	if o == nil || o.N2InterfaceAmfInfo == nil {
		return nil, false
	}
	return o.N2InterfaceAmfInfo, true
}

// HasN2InterfaceAmfInfo returns a boolean if a field has been set.
func (o *AmfInfo) HasN2InterfaceAmfInfo() bool {
	if o != nil && o.N2InterfaceAmfInfo != nil {
		return true
	}

	return false
}

// SetN2InterfaceAmfInfo gets a reference to the given N2InterfaceAmfInfo and assigns it to the N2InterfaceAmfInfo field.
func (o *AmfInfo) SetN2InterfaceAmfInfo(v N2InterfaceAmfInfo) {
	o.N2InterfaceAmfInfo = &v
}

func (o AmfInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amfSetId"] = o.AmfSetId
	}
	if true {
		toSerialize["amfRegionId"] = o.AmfRegionId
	}
	if true {
		toSerialize["guamiList"] = o.GuamiList
	}
	if o.TaiList != nil {
		toSerialize["taiList"] = o.TaiList
	}
	if o.TaiRangeList != nil {
		toSerialize["taiRangeList"] = o.TaiRangeList
	}
	if o.BackupInfoAmfFailure != nil {
		toSerialize["backupInfoAmfFailure"] = o.BackupInfoAmfFailure
	}
	if o.BackupInfoAmfRemoval != nil {
		toSerialize["backupInfoAmfRemoval"] = o.BackupInfoAmfRemoval
	}
	if o.N2InterfaceAmfInfo != nil {
		toSerialize["n2InterfaceAmfInfo"] = o.N2InterfaceAmfInfo
	}
	return json.Marshal(toSerialize)
}

type NullableAmfInfo struct {
	value *AmfInfo
	isSet bool
}

func (v NullableAmfInfo) Get() *AmfInfo {
	return v.value
}

func (v *NullableAmfInfo) Set(val *AmfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfInfo(val *AmfInfo) *NullableAmfInfo {
	return &NullableAmfInfo{value: val, isSet: true}
}

func (v NullableAmfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
