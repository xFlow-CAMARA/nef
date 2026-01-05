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

// checks if the AefProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AefProfile{}

// AefProfile Represents the AEF profile data.
type AefProfile struct {
	// Identifier of the API exposing function
	AefId string
	// API version
	Versions   []Version
	Protocol   Protocol
	DataFormat *string
	// Security methods supported by the AEF
	SecurityMethods []string
	// Domain to which API belongs to
	DomainName *string
	// Interface details
	InterfaceDescriptions []InterfaceDescription
	AefLocation           *AefLocation
	ServiceKpis           *ServiceKpis
	UeIpRange             NullableIpAddrRange
}

type _AefProfile AefProfile

// NewAefProfile instantiates a new AefProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAefProfile(aefId string, versions []Version) *AefProfile {
	this := AefProfile{}
	return &this
}

// NewAefProfileWithDefaults instantiates a new AefProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAefProfileWithDefaults() *AefProfile {
	this := AefProfile{}
	return &this
}

// GetAefId returns the AefId field value
func (o *AefProfile) GetAefId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AefId
}

// GetAefIdOk returns a tuple with the AefId field value
// and a boolean to check if the value has been set.
func (o *AefProfile) GetAefIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AefId, true
}

// SetAefId sets field value
func (o *AefProfile) SetAefId(v string) {
	o.AefId = v
}

// GetVersions returns the Versions field value
func (o *AefProfile) GetVersions() []Version {
	if o == nil {
		var ret []Version
		return ret
	}

	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value
// and a boolean to check if the value has been set.
func (o *AefProfile) GetVersionsOk() ([]Version, bool) {
	if o == nil {
		return nil, false
	}
	return o.Versions, true
}

// SetVersions sets field value
func (o *AefProfile) SetVersions(v []Version) {
	o.Versions = v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *AefProfile) GetProtocol() Protocol {
	if o == nil || IsNil(o.Protocol) {
		var ret Protocol
		return ret
	}
	return o.Protocol
}

// HasProtocol returns a boolean if a field has been set.
func (o *AefProfile) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given Protocol and assigns it to the Protocol field.
func (o *AefProfile) SetProtocol(v Protocol) {
	o.Protocol = v
}

// GetDataFormat returns the DataFormat field value if set, zero value otherwise.
func (o *AefProfile) GetDataFormat() string {
	if o == nil || IsNil(o.DataFormat) {
		var ret string
		return ret
	}
	return *o.DataFormat
}

// GetDataFormatOk returns a tuple with the DataFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AefProfile) GetDataFormatOk() (*string, bool) {
	if o == nil || IsNil(o.DataFormat) {
		return nil, false
	}
	return o.DataFormat, true
}

// HasDataFormat returns a boolean if a field has been set.
func (o *AefProfile) HasDataFormat() bool {
	if o != nil && !IsNil(o.DataFormat) {
		return true
	}

	return false
}

// SetDataFormat gets a reference to the given DataFormat and assigns it to the DataFormat field.
func (o *AefProfile) SetDataFormat(v string) {
	o.DataFormat = &v
}

// GetSecurityMethods returns the SecurityMethods field value if set, zero value otherwise.
func (o *AefProfile) GetSecurityMethods() []string {
	if o == nil || IsNil(o.SecurityMethods) {
		var ret []string
		return ret
	}
	return o.SecurityMethods
}

// GetSecurityMethodsOk returns a tuple with the SecurityMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AefProfile) GetSecurityMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.SecurityMethods) {
		return nil, false
	}
	return o.SecurityMethods, true
}

// HasSecurityMethods returns a boolean if a field has been set.
func (o *AefProfile) HasSecurityMethods() bool {
	if o != nil && !IsNil(o.SecurityMethods) {
		return true
	}

	return false
}

// SetSecurityMethods gets a reference to the given []string and assigns it to the SecurityMethods field.
func (o *AefProfile) SetSecurityMethods(v []string) {
	o.SecurityMethods = v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *AefProfile) GetDomainName() string {
	if o == nil || IsNil(o.DomainName) {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AefProfile) GetDomainNameOk() (*string, bool) {
	if o == nil || IsNil(o.DomainName) {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *AefProfile) HasDomainName() bool {
	if o != nil && !IsNil(o.DomainName) {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *AefProfile) SetDomainName(v string) {
	o.DomainName = &v
}

// GetInterfaceDescriptions returns the InterfaceDescriptions field value if set, zero value otherwise.
func (o *AefProfile) GetInterfaceDescriptions() []InterfaceDescription {
	if o == nil || IsNil(o.InterfaceDescriptions) {
		var ret []InterfaceDescription
		return ret
	}
	return o.InterfaceDescriptions
}

// GetInterfaceDescriptionsOk returns a tuple with the InterfaceDescriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AefProfile) GetInterfaceDescriptionsOk() ([]InterfaceDescription, bool) {
	if o == nil || IsNil(o.InterfaceDescriptions) {
		return nil, false
	}
	return o.InterfaceDescriptions, true
}

// HasInterfaceDescriptions returns a boolean if a field has been set.
func (o *AefProfile) HasInterfaceDescriptions() bool {
	if o != nil && !IsNil(o.InterfaceDescriptions) {
		return true
	}

	return false
}

// SetInterfaceDescriptions gets a reference to the given []InterfaceDescription and assigns it to the InterfaceDescriptions field.
func (o *AefProfile) SetInterfaceDescriptions(v []InterfaceDescription) {
	o.InterfaceDescriptions = v
}

// GetAefLocation returns the AefLocation field value if set, zero value otherwise.
func (o *AefProfile) GetAefLocation() AefLocation {
	if o == nil || IsNil(o.AefLocation) {
		var ret AefLocation
		return ret
	}
	return *o.AefLocation
}

// GetAefLocationOk returns a tuple with the AefLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AefProfile) GetAefLocationOk() (*AefLocation, bool) {
	if o == nil || IsNil(o.AefLocation) {
		return nil, false
	}
	return o.AefLocation, true
}

// HasAefLocation returns a boolean if a field has been set.
func (o *AefProfile) HasAefLocation() bool {
	if o != nil && !IsNil(o.AefLocation) {
		return true
	}

	return false
}

// SetAefLocation gets a reference to the given AefLocation and assigns it to the AefLocation field.
func (o *AefProfile) SetAefLocation(v AefLocation) {
	o.AefLocation = &v
}

// GetServiceKpis returns the ServiceKpis field value if set, zero value otherwise.
func (o *AefProfile) GetServiceKpis() ServiceKpis {
	if o == nil || IsNil(o.ServiceKpis) {
		var ret ServiceKpis
		return ret
	}
	return *o.ServiceKpis
}

// GetServiceKpisOk returns a tuple with the ServiceKpis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AefProfile) GetServiceKpisOk() (*ServiceKpis, bool) {
	if o == nil || IsNil(o.ServiceKpis) {
		return nil, false
	}
	return o.ServiceKpis, true
}

// HasServiceKpis returns a boolean if a field has been set.
func (o *AefProfile) HasServiceKpis() bool {
	if o != nil && !IsNil(o.ServiceKpis) {
		return true
	}

	return false
}

// SetServiceKpis gets a reference to the given ServiceKpis and assigns it to the ServiceKpis field.
func (o *AefProfile) SetServiceKpis(v ServiceKpis) {
	o.ServiceKpis = &v
}

// GetUeIpRange returns the UeIpRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AefProfile) GetUeIpRange() IpAddrRange {
	if o == nil || IsNil(o.UeIpRange.Get()) {
		var ret IpAddrRange
		return ret
	}
	return *o.UeIpRange.Get()
}

// GetUeIpRangeOk returns a tuple with the UeIpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AefProfile) GetUeIpRangeOk() (*IpAddrRange, bool) {
	if o == nil {
		return nil, false
	}
	return o.UeIpRange.Get(), o.UeIpRange.IsSet()
}

// HasUeIpRange returns a boolean if a field has been set.
func (o *AefProfile) HasUeIpRange() bool {
	if o != nil && o.UeIpRange.IsSet() {
		return true
	}

	return false
}

// SetUeIpRange gets a reference to the given NullableIpAddrRange and assigns it to the UeIpRange field.
func (o *AefProfile) SetUeIpRange(v IpAddrRange) {
	o.UeIpRange.Set(&v)
}

// SetUeIpRangeNil sets the value for UeIpRange to be an explicit nil
func (o *AefProfile) SetUeIpRangeNil() {
	o.UeIpRange.Set(nil)
}

// UnsetUeIpRange ensures that no value is present for UeIpRange, not even an explicit nil
func (o *AefProfile) UnsetUeIpRange() {
	o.UeIpRange.Unset()
}

func (o AefProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AefProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["aefId"] = o.AefId
	toSerialize["versions"] = o.Versions
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.DataFormat) {
		toSerialize["dataFormat"] = o.DataFormat
	}
	if !IsNil(o.SecurityMethods) {
		toSerialize["securityMethods"] = o.SecurityMethods
	}
	if !IsNil(o.DomainName) {
		toSerialize["domainName"] = o.DomainName
	}
	if !IsNil(o.InterfaceDescriptions) {
		toSerialize["interfaceDescriptions"] = o.InterfaceDescriptions
	}
	if !IsNil(o.AefLocation) {
		toSerialize["aefLocation"] = o.AefLocation
	}
	if !IsNil(o.ServiceKpis) {
		toSerialize["serviceKpis"] = o.ServiceKpis
	}
	if o.UeIpRange.IsSet() {
		toSerialize["ueIpRange"] = o.UeIpRange.Get()
	}
	return toSerialize, nil
}

func (o *AefProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"aefId",
		"versions",
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

	varAefProfile := _AefProfile{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAefProfile)

	if err != nil {
		return err
	}

	*o = AefProfile(varAefProfile)

	return err
}

type NullableAefProfile struct {
	value *AefProfile
	isSet bool
}

func (v NullableAefProfile) Get() *AefProfile {
	return v.value
}

func (v *NullableAefProfile) Set(val *AefProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableAefProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableAefProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAefProfile(val *AefProfile) *NullableAefProfile {
	return &NullableAefProfile{value: val, isSet: true}
}

func (v NullableAefProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAefProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
