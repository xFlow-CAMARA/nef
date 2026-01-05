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
	"encoding/json"
)

// checks if the InterfaceDescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterfaceDescription{}

// InterfaceDescription Represents the description of an API's interface.
type InterfaceDescription struct {
	// string identifying a Ipv4 address formatted in the \"dotted decimal\" notation as defined in IETF RFC 1166.
	Ipv4Addr *string
	// string identifying a Ipv6 address formatted according to clause 4 in IETF RFC 5952. The mixed Ipv4 Ipv6 notation according to clause 5 of IETF RFC 5952 shall not be used.
	Ipv6Addr *string
	// Fully Qualified Domain Name
	Fqdn *string
	// Unsigned integer with valid values between 0 and 65535.
	Port *int32
	// A string representing a sequence of path segments that starts with the slash character.
	ApiPrefix *string
	// Security methods supported by the interface, it take precedence over the security methods provided in AefProfile, for this specific interface.
	SecurityMethods []string
}

// NewInterfaceDescription instantiates a new InterfaceDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceDescription() *InterfaceDescription {
	this := InterfaceDescription{}
	return &this
}

// NewInterfaceDescriptionWithDefaults instantiates a new InterfaceDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceDescriptionWithDefaults() *InterfaceDescription {
	this := InterfaceDescription{}
	return &this
}

// GetIpv4Addr returns the Ipv4Addr field value if set, zero value otherwise.
func (o *InterfaceDescription) GetIpv4Addr() string {
	if o == nil || IsNil(o.Ipv4Addr) {
		var ret string
		return ret
	}
	return *o.Ipv4Addr
}

// GetIpv4AddrOk returns a tuple with the Ipv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceDescription) GetIpv4AddrOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv4Addr) {
		return nil, false
	}
	return o.Ipv4Addr, true
}

// HasIpv4Addr returns a boolean if a field has been set.
func (o *InterfaceDescription) HasIpv4Addr() bool {
	if o != nil && !IsNil(o.Ipv4Addr) {
		return true
	}

	return false
}

// SetIpv4Addr gets a reference to the given string and assigns it to the Ipv4Addr field.
func (o *InterfaceDescription) SetIpv4Addr(v string) {
	o.Ipv4Addr = &v
}

// GetIpv6Addr returns the Ipv6Addr field value if set, zero value otherwise.
func (o *InterfaceDescription) GetIpv6Addr() string {
	if o == nil || IsNil(o.Ipv6Addr) {
		var ret string
		return ret
	}
	return *o.Ipv6Addr
}

// GetIpv6AddrOk returns a tuple with the Ipv6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceDescription) GetIpv6AddrOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv6Addr) {
		return nil, false
	}
	return o.Ipv6Addr, true
}

// HasIpv6Addr returns a boolean if a field has been set.
func (o *InterfaceDescription) HasIpv6Addr() bool {
	if o != nil && !IsNil(o.Ipv6Addr) {
		return true
	}

	return false
}

// SetIpv6Addr gets a reference to the given string and assigns it to the Ipv6Addr field.
func (o *InterfaceDescription) SetIpv6Addr(v string) {
	o.Ipv6Addr = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *InterfaceDescription) GetFqdn() string {
	if o == nil || IsNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceDescription) GetFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.Fqdn) {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *InterfaceDescription) HasFqdn() bool {
	if o != nil && !IsNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *InterfaceDescription) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *InterfaceDescription) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceDescription) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *InterfaceDescription) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *InterfaceDescription) SetPort(v int32) {
	o.Port = &v
}

// GetApiPrefix returns the ApiPrefix field value if set, zero value otherwise.
func (o *InterfaceDescription) GetApiPrefix() string {
	if o == nil || IsNil(o.ApiPrefix) {
		var ret string
		return ret
	}
	return *o.ApiPrefix
}

// GetApiPrefixOk returns a tuple with the ApiPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceDescription) GetApiPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.ApiPrefix) {
		return nil, false
	}
	return o.ApiPrefix, true
}

// HasApiPrefix returns a boolean if a field has been set.
func (o *InterfaceDescription) HasApiPrefix() bool {
	if o != nil && !IsNil(o.ApiPrefix) {
		return true
	}

	return false
}

// SetApiPrefix gets a reference to the given string and assigns it to the ApiPrefix field.
func (o *InterfaceDescription) SetApiPrefix(v string) {
	o.ApiPrefix = &v
}

// GetSecurityMethods returns the SecurityMethods field value if set, zero value otherwise.
func (o *InterfaceDescription) GetSecurityMethods() []string {
	if o == nil || IsNil(o.SecurityMethods) {
		var ret []string
		return ret
	}
	return o.SecurityMethods
}

// GetSecurityMethodsOk returns a tuple with the SecurityMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceDescription) GetSecurityMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.SecurityMethods) {
		return nil, false
	}
	return o.SecurityMethods, true
}

// HasSecurityMethods returns a boolean if a field has been set.
func (o *InterfaceDescription) HasSecurityMethods() bool {
	if o != nil && !IsNil(o.SecurityMethods) {
		return true
	}

	return false
}

// SetSecurityMethods gets a reference to the given []string and assigns it to the SecurityMethods field.
func (o *InterfaceDescription) SetSecurityMethods(v []string) {
	o.SecurityMethods = v
}

func (o InterfaceDescription) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterfaceDescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ipv4Addr) {
		toSerialize["ipv4Addr"] = o.Ipv4Addr
	}
	if !IsNil(o.Ipv6Addr) {
		toSerialize["ipv6Addr"] = o.Ipv6Addr
	}
	if !IsNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.ApiPrefix) {
		toSerialize["apiPrefix"] = o.ApiPrefix
	}
	if !IsNil(o.SecurityMethods) {
		toSerialize["securityMethods"] = o.SecurityMethods
	}
	return toSerialize, nil
}

type NullableInterfaceDescription struct {
	value *InterfaceDescription
	isSet bool
}

func (v NullableInterfaceDescription) Get() *InterfaceDescription {
	return v.value
}

func (v *NullableInterfaceDescription) Set(val *InterfaceDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceDescription(val *InterfaceDescription) *NullableInterfaceDescription {
	return &NullableInterfaceDescription{value: val, isSet: true}
}

func (v NullableInterfaceDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
