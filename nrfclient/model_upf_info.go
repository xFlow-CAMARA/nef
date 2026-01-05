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

// UpfInfo struct for UpfInfo
type UpfInfo struct {
	SNssaiUpfInfoList    []SnssaiUpfInfoItem    `json:"sNssaiUpfInfoList"`
	SmfServingArea       []string               `json:"smfServingArea,omitempty"`
	InterfaceUpfInfoList []InterfaceUpfInfoItem `json:"interfaceUpfInfoList,omitempty"`
	IwkEpsInd            *bool                  `json:"iwkEpsInd,omitempty"`
	PduSessionTypes      []PduSessionType       `json:"pduSessionTypes,omitempty"`
	AtsssCapability      *AtsssCapability       `json:"atsssCapability,omitempty"`
	UeIpAddrInd          *bool                  `json:"ueIpAddrInd,omitempty"`
}

// NewUpfInfo instantiates a new UpfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpfInfo(sNssaiUpfInfoList []SnssaiUpfInfoItem) *UpfInfo {
	this := UpfInfo{}
	this.SNssaiUpfInfoList = sNssaiUpfInfoList
	var iwkEpsInd bool = false
	this.IwkEpsInd = &iwkEpsInd
	var ueIpAddrInd bool = false
	this.UeIpAddrInd = &ueIpAddrInd
	return &this
}

// NewUpfInfoWithDefaults instantiates a new UpfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpfInfoWithDefaults() *UpfInfo {
	this := UpfInfo{}
	var iwkEpsInd bool = false
	this.IwkEpsInd = &iwkEpsInd
	var ueIpAddrInd bool = false
	this.UeIpAddrInd = &ueIpAddrInd
	return &this
}

// GetSNssaiUpfInfoList returns the SNssaiUpfInfoList field value
func (o *UpfInfo) GetSNssaiUpfInfoList() []SnssaiUpfInfoItem {
	if o == nil {
		var ret []SnssaiUpfInfoItem
		return ret
	}

	return o.SNssaiUpfInfoList
}

// GetSNssaiUpfInfoListOk returns a tuple with the SNssaiUpfInfoList field value
// and a boolean to check if the value has been set.
func (o *UpfInfo) GetSNssaiUpfInfoListOk() ([]SnssaiUpfInfoItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.SNssaiUpfInfoList, true
}

// SetSNssaiUpfInfoList sets field value
func (o *UpfInfo) SetSNssaiUpfInfoList(v []SnssaiUpfInfoItem) {
	o.SNssaiUpfInfoList = v
}

// GetSmfServingArea returns the SmfServingArea field value if set, zero value otherwise.
func (o *UpfInfo) GetSmfServingArea() []string {
	if o == nil || o.SmfServingArea == nil {
		var ret []string
		return ret
	}
	return o.SmfServingArea
}

// GetSmfServingAreaOk returns a tuple with the SmfServingArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfInfo) GetSmfServingAreaOk() ([]string, bool) {
	if o == nil || o.SmfServingArea == nil {
		return nil, false
	}
	return o.SmfServingArea, true
}

// HasSmfServingArea returns a boolean if a field has been set.
func (o *UpfInfo) HasSmfServingArea() bool {
	if o != nil && o.SmfServingArea != nil {
		return true
	}

	return false
}

// SetSmfServingArea gets a reference to the given []string and assigns it to the SmfServingArea field.
func (o *UpfInfo) SetSmfServingArea(v []string) {
	o.SmfServingArea = v
}

// GetInterfaceUpfInfoList returns the InterfaceUpfInfoList field value if set, zero value otherwise.
func (o *UpfInfo) GetInterfaceUpfInfoList() []InterfaceUpfInfoItem {
	if o == nil || o.InterfaceUpfInfoList == nil {
		var ret []InterfaceUpfInfoItem
		return ret
	}
	return o.InterfaceUpfInfoList
}

// GetInterfaceUpfInfoListOk returns a tuple with the InterfaceUpfInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfInfo) GetInterfaceUpfInfoListOk() ([]InterfaceUpfInfoItem, bool) {
	if o == nil || o.InterfaceUpfInfoList == nil {
		return nil, false
	}
	return o.InterfaceUpfInfoList, true
}

// HasInterfaceUpfInfoList returns a boolean if a field has been set.
func (o *UpfInfo) HasInterfaceUpfInfoList() bool {
	if o != nil && o.InterfaceUpfInfoList != nil {
		return true
	}

	return false
}

// SetInterfaceUpfInfoList gets a reference to the given []InterfaceUpfInfoItem and assigns it to the InterfaceUpfInfoList field.
func (o *UpfInfo) SetInterfaceUpfInfoList(v []InterfaceUpfInfoItem) {
	o.InterfaceUpfInfoList = v
}

// GetIwkEpsInd returns the IwkEpsInd field value if set, zero value otherwise.
func (o *UpfInfo) GetIwkEpsInd() bool {
	if o == nil || o.IwkEpsInd == nil {
		var ret bool
		return ret
	}
	return *o.IwkEpsInd
}

// GetIwkEpsIndOk returns a tuple with the IwkEpsInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfInfo) GetIwkEpsIndOk() (*bool, bool) {
	if o == nil || o.IwkEpsInd == nil {
		return nil, false
	}
	return o.IwkEpsInd, true
}

// HasIwkEpsInd returns a boolean if a field has been set.
func (o *UpfInfo) HasIwkEpsInd() bool {
	if o != nil && o.IwkEpsInd != nil {
		return true
	}

	return false
}

// SetIwkEpsInd gets a reference to the given bool and assigns it to the IwkEpsInd field.
func (o *UpfInfo) SetIwkEpsInd(v bool) {
	o.IwkEpsInd = &v
}

// GetPduSessionTypes returns the PduSessionTypes field value if set, zero value otherwise.
func (o *UpfInfo) GetPduSessionTypes() []PduSessionType {
	if o == nil || o.PduSessionTypes == nil {
		var ret []PduSessionType
		return ret
	}
	return o.PduSessionTypes
}

// GetPduSessionTypesOk returns a tuple with the PduSessionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfInfo) GetPduSessionTypesOk() ([]PduSessionType, bool) {
	if o == nil || o.PduSessionTypes == nil {
		return nil, false
	}
	return o.PduSessionTypes, true
}

// HasPduSessionTypes returns a boolean if a field has been set.
func (o *UpfInfo) HasPduSessionTypes() bool {
	if o != nil && o.PduSessionTypes != nil {
		return true
	}

	return false
}

// SetPduSessionTypes gets a reference to the given []PduSessionType and assigns it to the PduSessionTypes field.
func (o *UpfInfo) SetPduSessionTypes(v []PduSessionType) {
	o.PduSessionTypes = v
}

// GetAtsssCapability returns the AtsssCapability field value if set, zero value otherwise.
func (o *UpfInfo) GetAtsssCapability() AtsssCapability {
	if o == nil || o.AtsssCapability == nil {
		var ret AtsssCapability
		return ret
	}
	return *o.AtsssCapability
}

// GetAtsssCapabilityOk returns a tuple with the AtsssCapability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfInfo) GetAtsssCapabilityOk() (*AtsssCapability, bool) {
	if o == nil || o.AtsssCapability == nil {
		return nil, false
	}
	return o.AtsssCapability, true
}

// HasAtsssCapability returns a boolean if a field has been set.
func (o *UpfInfo) HasAtsssCapability() bool {
	if o != nil && o.AtsssCapability != nil {
		return true
	}

	return false
}

// SetAtsssCapability gets a reference to the given AtsssCapability and assigns it to the AtsssCapability field.
func (o *UpfInfo) SetAtsssCapability(v AtsssCapability) {
	o.AtsssCapability = &v
}

// GetUeIpAddrInd returns the UeIpAddrInd field value if set, zero value otherwise.
func (o *UpfInfo) GetUeIpAddrInd() bool {
	if o == nil || o.UeIpAddrInd == nil {
		var ret bool
		return ret
	}
	return *o.UeIpAddrInd
}

// GetUeIpAddrIndOk returns a tuple with the UeIpAddrInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfInfo) GetUeIpAddrIndOk() (*bool, bool) {
	if o == nil || o.UeIpAddrInd == nil {
		return nil, false
	}
	return o.UeIpAddrInd, true
}

// HasUeIpAddrInd returns a boolean if a field has been set.
func (o *UpfInfo) HasUeIpAddrInd() bool {
	if o != nil && o.UeIpAddrInd != nil {
		return true
	}

	return false
}

// SetUeIpAddrInd gets a reference to the given bool and assigns it to the UeIpAddrInd field.
func (o *UpfInfo) SetUeIpAddrInd(v bool) {
	o.UeIpAddrInd = &v
}

func (o UpfInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sNssaiUpfInfoList"] = o.SNssaiUpfInfoList
	}
	if o.SmfServingArea != nil {
		toSerialize["smfServingArea"] = o.SmfServingArea
	}
	if o.InterfaceUpfInfoList != nil {
		toSerialize["interfaceUpfInfoList"] = o.InterfaceUpfInfoList
	}
	if o.IwkEpsInd != nil {
		toSerialize["iwkEpsInd"] = o.IwkEpsInd
	}
	if o.PduSessionTypes != nil {
		toSerialize["pduSessionTypes"] = o.PduSessionTypes
	}
	if o.AtsssCapability != nil {
		toSerialize["atsssCapability"] = o.AtsssCapability
	}
	if o.UeIpAddrInd != nil {
		toSerialize["ueIpAddrInd"] = o.UeIpAddrInd
	}
	return json.Marshal(toSerialize)
}

type NullableUpfInfo struct {
	value *UpfInfo
	isSet bool
}

func (v NullableUpfInfo) Get() *UpfInfo {
	return v.value
}

func (v *NullableUpfInfo) Set(val *UpfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUpfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUpfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpfInfo(val *UpfInfo) *NullableUpfInfo {
	return &NullableUpfInfo{value: val, isSet: true}
}

func (v NullableUpfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
