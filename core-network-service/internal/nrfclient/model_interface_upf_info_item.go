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

// InterfaceUpfInfoItem struct for InterfaceUpfInfoItem
type InterfaceUpfInfoItem struct {
	InterfaceType         UPInterfaceType `json:"interfaceType"`
	Ipv4EndpointAddresses []string        `json:"ipv4EndpointAddresses,omitempty"`
	Ipv6EndpointAddresses []Ipv6Addr      `json:"ipv6EndpointAddresses,omitempty"`
	EndpointFqdn          *string         `json:"endpointFqdn,omitempty"`
	NetworkInstance       *string         `json:"networkInstance,omitempty"`
}

// NewInterfaceUpfInfoItem instantiates a new InterfaceUpfInfoItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceUpfInfoItem(interfaceType UPInterfaceType) *InterfaceUpfInfoItem {
	this := InterfaceUpfInfoItem{}
	this.InterfaceType = interfaceType
	return &this
}

// NewInterfaceUpfInfoItemWithDefaults instantiates a new InterfaceUpfInfoItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceUpfInfoItemWithDefaults() *InterfaceUpfInfoItem {
	this := InterfaceUpfInfoItem{}
	return &this
}

// GetInterfaceType returns the InterfaceType field value
func (o *InterfaceUpfInfoItem) GetInterfaceType() UPInterfaceType {
	if o == nil {
		var ret UPInterfaceType
		return ret
	}

	return o.InterfaceType
}

// GetInterfaceTypeOk returns a tuple with the InterfaceType field value
// and a boolean to check if the value has been set.
func (o *InterfaceUpfInfoItem) GetInterfaceTypeOk() (*UPInterfaceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InterfaceType, true
}

// SetInterfaceType sets field value
func (o *InterfaceUpfInfoItem) SetInterfaceType(v UPInterfaceType) {
	o.InterfaceType = v
}

// GetIpv4EndpointAddresses returns the Ipv4EndpointAddresses field value if set, zero value otherwise.
func (o *InterfaceUpfInfoItem) GetIpv4EndpointAddresses() []string {
	if o == nil || o.Ipv4EndpointAddresses == nil {
		var ret []string
		return ret
	}
	return o.Ipv4EndpointAddresses
}

// GetIpv4EndpointAddressesOk returns a tuple with the Ipv4EndpointAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceUpfInfoItem) GetIpv4EndpointAddressesOk() ([]string, bool) {
	if o == nil || o.Ipv4EndpointAddresses == nil {
		return nil, false
	}
	return o.Ipv4EndpointAddresses, true
}

// HasIpv4EndpointAddresses returns a boolean if a field has been set.
func (o *InterfaceUpfInfoItem) HasIpv4EndpointAddresses() bool {
	if o != nil && o.Ipv4EndpointAddresses != nil {
		return true
	}

	return false
}

// SetIpv4EndpointAddresses gets a reference to the given []string and assigns it to the Ipv4EndpointAddresses field.
func (o *InterfaceUpfInfoItem) SetIpv4EndpointAddresses(v []string) {
	o.Ipv4EndpointAddresses = v
}

// GetIpv6EndpointAddresses returns the Ipv6EndpointAddresses field value if set, zero value otherwise.
func (o *InterfaceUpfInfoItem) GetIpv6EndpointAddresses() []Ipv6Addr {
	if o == nil || o.Ipv6EndpointAddresses == nil {
		var ret []Ipv6Addr
		return ret
	}
	return o.Ipv6EndpointAddresses
}

// GetIpv6EndpointAddressesOk returns a tuple with the Ipv6EndpointAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceUpfInfoItem) GetIpv6EndpointAddressesOk() ([]Ipv6Addr, bool) {
	if o == nil || o.Ipv6EndpointAddresses == nil {
		return nil, false
	}
	return o.Ipv6EndpointAddresses, true
}

// HasIpv6EndpointAddresses returns a boolean if a field has been set.
func (o *InterfaceUpfInfoItem) HasIpv6EndpointAddresses() bool {
	if o != nil && o.Ipv6EndpointAddresses != nil {
		return true
	}

	return false
}

// SetIpv6EndpointAddresses gets a reference to the given []Ipv6Addr and assigns it to the Ipv6EndpointAddresses field.
func (o *InterfaceUpfInfoItem) SetIpv6EndpointAddresses(v []Ipv6Addr) {
	o.Ipv6EndpointAddresses = v
}

// GetEndpointFqdn returns the EndpointFqdn field value if set, zero value otherwise.
func (o *InterfaceUpfInfoItem) GetEndpointFqdn() string {
	if o == nil || o.EndpointFqdn == nil {
		var ret string
		return ret
	}
	return *o.EndpointFqdn
}

// GetEndpointFqdnOk returns a tuple with the EndpointFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceUpfInfoItem) GetEndpointFqdnOk() (*string, bool) {
	if o == nil || o.EndpointFqdn == nil {
		return nil, false
	}
	return o.EndpointFqdn, true
}

// HasEndpointFqdn returns a boolean if a field has been set.
func (o *InterfaceUpfInfoItem) HasEndpointFqdn() bool {
	if o != nil && o.EndpointFqdn != nil {
		return true
	}

	return false
}

// SetEndpointFqdn gets a reference to the given string and assigns it to the EndpointFqdn field.
func (o *InterfaceUpfInfoItem) SetEndpointFqdn(v string) {
	o.EndpointFqdn = &v
}

// GetNetworkInstance returns the NetworkInstance field value if set, zero value otherwise.
func (o *InterfaceUpfInfoItem) GetNetworkInstance() string {
	if o == nil || o.NetworkInstance == nil {
		var ret string
		return ret
	}
	return *o.NetworkInstance
}

// GetNetworkInstanceOk returns a tuple with the NetworkInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceUpfInfoItem) GetNetworkInstanceOk() (*string, bool) {
	if o == nil || o.NetworkInstance == nil {
		return nil, false
	}
	return o.NetworkInstance, true
}

// HasNetworkInstance returns a boolean if a field has been set.
func (o *InterfaceUpfInfoItem) HasNetworkInstance() bool {
	if o != nil && o.NetworkInstance != nil {
		return true
	}

	return false
}

// SetNetworkInstance gets a reference to the given string and assigns it to the NetworkInstance field.
func (o *InterfaceUpfInfoItem) SetNetworkInstance(v string) {
	o.NetworkInstance = &v
}

func (o InterfaceUpfInfoItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["interfaceType"] = o.InterfaceType
	}
	if o.Ipv4EndpointAddresses != nil {
		toSerialize["ipv4EndpointAddresses"] = o.Ipv4EndpointAddresses
	}
	if o.Ipv6EndpointAddresses != nil {
		toSerialize["ipv6EndpointAddresses"] = o.Ipv6EndpointAddresses
	}
	if o.EndpointFqdn != nil {
		toSerialize["endpointFqdn"] = o.EndpointFqdn
	}
	if o.NetworkInstance != nil {
		toSerialize["networkInstance"] = o.NetworkInstance
	}
	return json.Marshal(toSerialize)
}

type NullableInterfaceUpfInfoItem struct {
	value *InterfaceUpfInfoItem
	isSet bool
}

func (v NullableInterfaceUpfInfoItem) Get() *InterfaceUpfInfoItem {
	return v.value
}

func (v *NullableInterfaceUpfInfoItem) Set(val *InterfaceUpfInfoItem) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceUpfInfoItem) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceUpfInfoItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceUpfInfoItem(val *InterfaceUpfInfoItem) *NullableInterfaceUpfInfoItem {
	return &NullableInterfaceUpfInfoItem{value: val, isSet: true}
}

func (v NullableInterfaceUpfInfoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceUpfInfoItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
