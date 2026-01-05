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
	"time"
)

// NFProfile struct for NFProfile
type NFProfile struct {
	NfInstanceId                     string                            `json:"nfInstanceId"`
	NfInstanceName                   *string                           `json:"nfInstanceName,omitempty"`
	NfType                           NFType                            `json:"nfType"`
	NfStatus                         NFStatus                          `json:"nfStatus"`
	PlmnList                         []PlmnId                          `json:"plmnList,omitempty"`
	SNssais                          []Snssai                          `json:"sNssais,omitempty"`
	PerPlmnSnssaiList                []PlmnSnssai                      `json:"perPlmnSnssaiList,omitempty"`
	NsiList                          []string                          `json:"nsiList,omitempty"`
	Fqdn                             *string                           `json:"fqdn,omitempty"`
	Ipv4Addresses                    []string                          `json:"ipv4Addresses,omitempty"`
	Ipv6Addresses                    []Ipv6Addr                        `json:"ipv6Addresses,omitempty"`
	Capacity                         *int32                            `json:"capacity,omitempty"`
	Load                             *int32                            `json:"load,omitempty"`
	Locality                         *string                           `json:"locality,omitempty"`
	Priority                         *int32                            `json:"priority,omitempty"`
	UdrInfo                          *UdrInfo                          `json:"udrInfo,omitempty"`
	UdrInfoExt                       []UdrInfo                         `json:"udrInfoExt,omitempty"`
	UdmInfo                          *UdmInfo                          `json:"udmInfo,omitempty"`
	UdmInfoExt                       []UdmInfo                         `json:"udmInfoExt,omitempty"`
	AusfInfo                         *AusfInfo                         `json:"ausfInfo,omitempty"`
	AusfInfoExt                      []AusfInfo                        `json:"ausfInfoExt,omitempty"`
	AmfInfo                          *AmfInfo                          `json:"amfInfo,omitempty"`
	AmfInfoExt                       []AmfInfo                         `json:"amfInfoExt,omitempty"`
	SmfInfo                          *SmfInfo                          `json:"smfInfo,omitempty"`
	SmfInfoExt                       []SmfInfo                         `json:"smfInfoExt,omitempty"`
	UpfInfo                          *UpfInfo                          `json:"upfInfo,omitempty"`
	UpfInfoExt                       []UpfInfo                         `json:"upfInfoExt,omitempty"`
	PcfInfo                          *PcfInfo                          `json:"pcfInfo,omitempty"`
	PcfInfoExt                       []PcfInfo                         `json:"pcfInfoExt,omitempty"`
	BsfInfo                          *BsfInfo                          `json:"bsfInfo,omitempty"`
	BsfInfoExt                       []BsfInfo                         `json:"bsfInfoExt,omitempty"`
	ChfInfo                          *ChfInfo                          `json:"chfInfo,omitempty"`
	ChfInfoExt                       []ChfInfo                         `json:"chfInfoExt,omitempty"`
	NwdafInfo                        *NwdafInfo                        `json:"nwdafInfo,omitempty"`
	CustomInfo                       map[string]interface{}            `json:"customInfo,omitempty"`
	RecoveryTime                     *time.Time                        `json:"recoveryTime,omitempty"`
	NfServicePersistence             *bool                             `json:"nfServicePersistence,omitempty"`
	NfServices                       []NFService                       `json:"nfServices,omitempty"`
	DefaultNotificationSubscriptions []DefaultNotificationSubscription `json:"defaultNotificationSubscriptions,omitempty"`
}

// NewNFProfile instantiates a new NFProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNFProfile(nfInstanceId string, nfType NFType, nfStatus NFStatus) *NFProfile {
	this := NFProfile{}
	this.NfInstanceId = nfInstanceId
	this.NfType = nfType
	this.NfStatus = nfStatus
	nfServicePersistence := false
	this.NfServicePersistence = &nfServicePersistence
	return &this
}

// NewNFProfileWithDefaults instantiates a new NFProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNFProfileWithDefaults() *NFProfile {
	this := NFProfile{}
	nfServicePersistence := false
	this.NfServicePersistence = &nfServicePersistence
	return &this
}

// GetNfInstanceId returns the NfInstanceId field value
func (o *NFProfile) GetNfInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NfInstanceId
}

// GetNfInstanceIdOk returns a tuple with the NfInstanceId field value
// and a boolean to check if the value has been set.
func (o *NFProfile) GetNfInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NfInstanceId, true
}

// SetNfInstanceId sets field value
func (o *NFProfile) SetNfInstanceId(v string) {
	o.NfInstanceId = v
}

// GetNfInstanceName returns the NfInstanceName field value if set, zero value otherwise.
func (o *NFProfile) GetNfInstanceName() string {
	if o == nil || o.NfInstanceName == nil {
		var ret string
		return ret
	}
	return *o.NfInstanceName
}

// GetNfInstanceNameOk returns a tuple with the NfInstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetNfInstanceNameOk() (*string, bool) {
	if o == nil || o.NfInstanceName == nil {
		return nil, false
	}
	return o.NfInstanceName, true
}

// HasNfInstanceName returns a boolean if a field has been set.
func (o *NFProfile) HasNfInstanceName() bool {
	if o != nil && o.NfInstanceName != nil {
		return true
	}

	return false
}

// SetNfInstanceName gets a reference to the given string and assigns it to the NfInstanceName field.
func (o *NFProfile) SetNfInstanceName(v string) {
	o.NfInstanceName = &v
}

// GetNfType returns the NfType field value
func (o *NFProfile) GetNfType() NFType {
	if o == nil {
		var ret NFType
		return ret
	}

	return o.NfType
}

// GetNfTypeOk returns a tuple with the NfType field value
// and a boolean to check if the value has been set.
func (o *NFProfile) GetNfTypeOk() (*NFType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NfType, true
}

// SetNfType sets field value
func (o *NFProfile) SetNfType(v NFType) {
	o.NfType = v
}

// GetNfStatus returns the NfStatus field value
func (o *NFProfile) GetNfStatus() NFStatus {
	if o == nil {
		var ret NFStatus
		return ret
	}

	return o.NfStatus
}

// GetNfStatusOk returns a tuple with the NfStatus field value
// and a boolean to check if the value has been set.
func (o *NFProfile) GetNfStatusOk() (*NFStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NfStatus, true
}

// SetNfStatus sets field value
func (o *NFProfile) SetNfStatus(v NFStatus) {
	o.NfStatus = v
}

// GetPlmnList returns the PlmnList field value if set, zero value otherwise.
func (o *NFProfile) GetPlmnList() []PlmnId {
	if o == nil || o.PlmnList == nil {
		var ret []PlmnId
		return ret
	}
	return o.PlmnList
}

// GetPlmnListOk returns a tuple with the PlmnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetPlmnListOk() ([]PlmnId, bool) {
	if o == nil || o.PlmnList == nil {
		return nil, false
	}
	return o.PlmnList, true
}

// HasPlmnList returns a boolean if a field has been set.
func (o *NFProfile) HasPlmnList() bool {
	if o != nil && o.PlmnList != nil {
		return true
	}

	return false
}

// SetPlmnList gets a reference to the given []PlmnId and assigns it to the PlmnList field.
func (o *NFProfile) SetPlmnList(v []PlmnId) {
	o.PlmnList = v
}

// GetSNssais returns the SNssais field value if set, zero value otherwise.
func (o *NFProfile) GetSNssais() []Snssai {
	if o == nil || o.SNssais == nil {
		var ret []Snssai
		return ret
	}
	return o.SNssais
}

// GetSNssaisOk returns a tuple with the SNssais field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetSNssaisOk() ([]Snssai, bool) {
	if o == nil || o.SNssais == nil {
		return nil, false
	}
	return o.SNssais, true
}

// HasSNssais returns a boolean if a field has been set.
func (o *NFProfile) HasSNssais() bool {
	if o != nil && o.SNssais != nil {
		return true
	}

	return false
}

// SetSNssais gets a reference to the given []Snssai and assigns it to the SNssais field.
func (o *NFProfile) SetSNssais(v []Snssai) {
	o.SNssais = v
}

// GetPerPlmnSnssaiList returns the PerPlmnSnssaiList field value if set, zero value otherwise.
func (o *NFProfile) GetPerPlmnSnssaiList() []PlmnSnssai {
	if o == nil || o.PerPlmnSnssaiList == nil {
		var ret []PlmnSnssai
		return ret
	}
	return o.PerPlmnSnssaiList
}

// GetPerPlmnSnssaiListOk returns a tuple with the PerPlmnSnssaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetPerPlmnSnssaiListOk() ([]PlmnSnssai, bool) {
	if o == nil || o.PerPlmnSnssaiList == nil {
		return nil, false
	}
	return o.PerPlmnSnssaiList, true
}

// HasPerPlmnSnssaiList returns a boolean if a field has been set.
func (o *NFProfile) HasPerPlmnSnssaiList() bool {
	if o != nil && o.PerPlmnSnssaiList != nil {
		return true
	}

	return false
}

// SetPerPlmnSnssaiList gets a reference to the given []PlmnSnssai and assigns it to the PerPlmnSnssaiList field.
func (o *NFProfile) SetPerPlmnSnssaiList(v []PlmnSnssai) {
	o.PerPlmnSnssaiList = v
}

// GetNsiList returns the NsiList field value if set, zero value otherwise.
func (o *NFProfile) GetNsiList() []string {
	if o == nil || o.NsiList == nil {
		var ret []string
		return ret
	}
	return o.NsiList
}

// GetNsiListOk returns a tuple with the NsiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetNsiListOk() ([]string, bool) {
	if o == nil || o.NsiList == nil {
		return nil, false
	}
	return o.NsiList, true
}

// HasNsiList returns a boolean if a field has been set.
func (o *NFProfile) HasNsiList() bool {
	if o != nil && o.NsiList != nil {
		return true
	}

	return false
}

// SetNsiList gets a reference to the given []string and assigns it to the NsiList field.
func (o *NFProfile) SetNsiList(v []string) {
	o.NsiList = v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *NFProfile) GetFqdn() string {
	if o == nil || o.Fqdn == nil {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetFqdnOk() (*string, bool) {
	if o == nil || o.Fqdn == nil {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *NFProfile) HasFqdn() bool {
	if o != nil && o.Fqdn != nil {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *NFProfile) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetIpv4Addresses returns the Ipv4Addresses field value if set, zero value otherwise.
func (o *NFProfile) GetIpv4Addresses() []string {
	if o == nil || o.Ipv4Addresses == nil {
		var ret []string
		return ret
	}
	return o.Ipv4Addresses
}

// GetIpv4AddressesOk returns a tuple with the Ipv4Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetIpv4AddressesOk() ([]string, bool) {
	if o == nil || o.Ipv4Addresses == nil {
		return nil, false
	}
	return o.Ipv4Addresses, true
}

// HasIpv4Addresses returns a boolean if a field has been set.
func (o *NFProfile) HasIpv4Addresses() bool {
	if o != nil && o.Ipv4Addresses != nil {
		return true
	}

	return false
}

// SetIpv4Addresses gets a reference to the given []string and assigns it to the Ipv4Addresses field.
func (o *NFProfile) SetIpv4Addresses(v []string) {
	o.Ipv4Addresses = v
}

// GetIpv6Addresses returns the Ipv6Addresses field value if set, zero value otherwise.
func (o *NFProfile) GetIpv6Addresses() []Ipv6Addr {
	if o == nil || o.Ipv6Addresses == nil {
		var ret []Ipv6Addr
		return ret
	}
	return o.Ipv6Addresses
}

// GetIpv6AddressesOk returns a tuple with the Ipv6Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetIpv6AddressesOk() ([]Ipv6Addr, bool) {
	if o == nil || o.Ipv6Addresses == nil {
		return nil, false
	}
	return o.Ipv6Addresses, true
}

// HasIpv6Addresses returns a boolean if a field has been set.
func (o *NFProfile) HasIpv6Addresses() bool {
	if o != nil && o.Ipv6Addresses != nil {
		return true
	}

	return false
}

// SetIpv6Addresses gets a reference to the given []Ipv6Addr and assigns it to the Ipv6Addresses field.
func (o *NFProfile) SetIpv6Addresses(v []Ipv6Addr) {
	o.Ipv6Addresses = v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *NFProfile) GetCapacity() int32 {
	if o == nil || o.Capacity == nil {
		var ret int32
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetCapacityOk() (*int32, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *NFProfile) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int32 and assigns it to the Capacity field.
func (o *NFProfile) SetCapacity(v int32) {
	o.Capacity = &v
}

// GetLoad returns the Load field value if set, zero value otherwise.
func (o *NFProfile) GetLoad() int32 {
	if o == nil || o.Load == nil {
		var ret int32
		return ret
	}
	return *o.Load
}

// GetLoadOk returns a tuple with the Load field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetLoadOk() (*int32, bool) {
	if o == nil || o.Load == nil {
		return nil, false
	}
	return o.Load, true
}

// HasLoad returns a boolean if a field has been set.
func (o *NFProfile) HasLoad() bool {
	if o != nil && o.Load != nil {
		return true
	}

	return false
}

// SetLoad gets a reference to the given int32 and assigns it to the Load field.
func (o *NFProfile) SetLoad(v int32) {
	o.Load = &v
}

// GetLocality returns the Locality field value if set, zero value otherwise.
func (o *NFProfile) GetLocality() string {
	if o == nil || o.Locality == nil {
		var ret string
		return ret
	}
	return *o.Locality
}

// GetLocalityOk returns a tuple with the Locality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetLocalityOk() (*string, bool) {
	if o == nil || o.Locality == nil {
		return nil, false
	}
	return o.Locality, true
}

// HasLocality returns a boolean if a field has been set.
func (o *NFProfile) HasLocality() bool {
	if o != nil && o.Locality != nil {
		return true
	}

	return false
}

// SetLocality gets a reference to the given string and assigns it to the Locality field.
func (o *NFProfile) SetLocality(v string) {
	o.Locality = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *NFProfile) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *NFProfile) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *NFProfile) SetPriority(v int32) {
	o.Priority = &v
}

// GetUdrInfo returns the UdrInfo field value if set, zero value otherwise.
func (o *NFProfile) GetUdrInfo() UdrInfo {
	if o == nil || o.UdrInfo == nil {
		var ret UdrInfo
		return ret
	}
	return *o.UdrInfo
}

// GetUdrInfoOk returns a tuple with the UdrInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetUdrInfoOk() (*UdrInfo, bool) {
	if o == nil || o.UdrInfo == nil {
		return nil, false
	}
	return o.UdrInfo, true
}

// HasUdrInfo returns a boolean if a field has been set.
func (o *NFProfile) HasUdrInfo() bool {
	if o != nil && o.UdrInfo != nil {
		return true
	}

	return false
}

// SetUdrInfo gets a reference to the given UdrInfo and assigns it to the UdrInfo field.
func (o *NFProfile) SetUdrInfo(v UdrInfo) {
	o.UdrInfo = &v
}

// GetUdrInfoExt returns the UdrInfoExt field value if set, zero value otherwise.
func (o *NFProfile) GetUdrInfoExt() []UdrInfo {
	if o == nil || o.UdrInfoExt == nil {
		var ret []UdrInfo
		return ret
	}
	return o.UdrInfoExt
}

// GetUdrInfoExtOk returns a tuple with the UdrInfoExt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetUdrInfoExtOk() ([]UdrInfo, bool) {
	if o == nil || o.UdrInfoExt == nil {
		return nil, false
	}
	return o.UdrInfoExt, true
}

// HasUdrInfoExt returns a boolean if a field has been set.
func (o *NFProfile) HasUdrInfoExt() bool {
	if o != nil && o.UdrInfoExt != nil {
		return true
	}

	return false
}

// SetUdrInfoExt gets a reference to the given []UdrInfo and assigns it to the UdrInfoExt field.
func (o *NFProfile) SetUdrInfoExt(v []UdrInfo) {
	o.UdrInfoExt = v
}

// GetUdmInfo returns the UdmInfo field value if set, zero value otherwise.
func (o *NFProfile) GetUdmInfo() UdmInfo {
	if o == nil || o.UdmInfo == nil {
		var ret UdmInfo
		return ret
	}
	return *o.UdmInfo
}

// GetUdmInfoOk returns a tuple with the UdmInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetUdmInfoOk() (*UdmInfo, bool) {
	if o == nil || o.UdmInfo == nil {
		return nil, false
	}
	return o.UdmInfo, true
}

// HasUdmInfo returns a boolean if a field has been set.
func (o *NFProfile) HasUdmInfo() bool {
	if o != nil && o.UdmInfo != nil {
		return true
	}

	return false
}

// SetUdmInfo gets a reference to the given UdmInfo and assigns it to the UdmInfo field.
func (o *NFProfile) SetUdmInfo(v UdmInfo) {
	o.UdmInfo = &v
}

// GetUdmInfoExt returns the UdmInfoExt field value if set, zero value otherwise.
func (o *NFProfile) GetUdmInfoExt() []UdmInfo {
	if o == nil || o.UdmInfoExt == nil {
		var ret []UdmInfo
		return ret
	}
	return o.UdmInfoExt
}

// GetUdmInfoExtOk returns a tuple with the UdmInfoExt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetUdmInfoExtOk() ([]UdmInfo, bool) {
	if o == nil || o.UdmInfoExt == nil {
		return nil, false
	}
	return o.UdmInfoExt, true
}

// HasUdmInfoExt returns a boolean if a field has been set.
func (o *NFProfile) HasUdmInfoExt() bool {
	if o != nil && o.UdmInfoExt != nil {
		return true
	}

	return false
}

// SetUdmInfoExt gets a reference to the given []UdmInfo and assigns it to the UdmInfoExt field.
func (o *NFProfile) SetUdmInfoExt(v []UdmInfo) {
	o.UdmInfoExt = v
}

// GetAusfInfo returns the AusfInfo field value if set, zero value otherwise.
func (o *NFProfile) GetAusfInfo() AusfInfo {
	if o == nil || o.AusfInfo == nil {
		var ret AusfInfo
		return ret
	}
	return *o.AusfInfo
}

// GetAusfInfoOk returns a tuple with the AusfInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetAusfInfoOk() (*AusfInfo, bool) {
	if o == nil || o.AusfInfo == nil {
		return nil, false
	}
	return o.AusfInfo, true
}

// HasAusfInfo returns a boolean if a field has been set.
func (o *NFProfile) HasAusfInfo() bool {
	if o != nil && o.AusfInfo != nil {
		return true
	}

	return false
}

// SetAusfInfo gets a reference to the given AusfInfo and assigns it to the AusfInfo field.
func (o *NFProfile) SetAusfInfo(v AusfInfo) {
	o.AusfInfo = &v
}

// GetAusfInfoExt returns the AusfInfoExt field value if set, zero value otherwise.
func (o *NFProfile) GetAusfInfoExt() []AusfInfo {
	if o == nil || o.AusfInfoExt == nil {
		var ret []AusfInfo
		return ret
	}
	return o.AusfInfoExt
}

// GetAusfInfoExtOk returns a tuple with the AusfInfoExt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetAusfInfoExtOk() ([]AusfInfo, bool) {
	if o == nil || o.AusfInfoExt == nil {
		return nil, false
	}
	return o.AusfInfoExt, true
}

// HasAusfInfoExt returns a boolean if a field has been set.
func (o *NFProfile) HasAusfInfoExt() bool {
	if o != nil && o.AusfInfoExt != nil {
		return true
	}

	return false
}

// SetAusfInfoExt gets a reference to the given []AusfInfo and assigns it to the AusfInfoExt field.
func (o *NFProfile) SetAusfInfoExt(v []AusfInfo) {
	o.AusfInfoExt = v
}

// GetAmfInfo returns the AmfInfo field value if set, zero value otherwise.
func (o *NFProfile) GetAmfInfo() AmfInfo {
	if o == nil || o.AmfInfo == nil {
		var ret AmfInfo
		return ret
	}
	return *o.AmfInfo
}

// GetAmfInfoOk returns a tuple with the AmfInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetAmfInfoOk() (*AmfInfo, bool) {
	if o == nil || o.AmfInfo == nil {
		return nil, false
	}
	return o.AmfInfo, true
}

// HasAmfInfo returns a boolean if a field has been set.
func (o *NFProfile) HasAmfInfo() bool {
	if o != nil && o.AmfInfo != nil {
		return true
	}

	return false
}

// SetAmfInfo gets a reference to the given AmfInfo and assigns it to the AmfInfo field.
func (o *NFProfile) SetAmfInfo(v AmfInfo) {
	o.AmfInfo = &v
}

// GetAmfInfoExt returns the AmfInfoExt field value if set, zero value otherwise.
func (o *NFProfile) GetAmfInfoExt() []AmfInfo {
	if o == nil || o.AmfInfoExt == nil {
		var ret []AmfInfo
		return ret
	}
	return o.AmfInfoExt
}

// GetAmfInfoExtOk returns a tuple with the AmfInfoExt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetAmfInfoExtOk() ([]AmfInfo, bool) {
	if o == nil || o.AmfInfoExt == nil {
		return nil, false
	}
	return o.AmfInfoExt, true
}

// HasAmfInfoExt returns a boolean if a field has been set.
func (o *NFProfile) HasAmfInfoExt() bool {
	if o != nil && o.AmfInfoExt != nil {
		return true
	}

	return false
}

// SetAmfInfoExt gets a reference to the given []AmfInfo and assigns it to the AmfInfoExt field.
func (o *NFProfile) SetAmfInfoExt(v []AmfInfo) {
	o.AmfInfoExt = v
}

// GetSmfInfo returns the SmfInfo field value if set, zero value otherwise.
func (o *NFProfile) GetSmfInfo() SmfInfo {
	if o == nil || o.SmfInfo == nil {
		var ret SmfInfo
		return ret
	}
	return *o.SmfInfo
}

// GetSmfInfoOk returns a tuple with the SmfInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetSmfInfoOk() (*SmfInfo, bool) {
	if o == nil || o.SmfInfo == nil {
		return nil, false
	}
	return o.SmfInfo, true
}

// HasSmfInfo returns a boolean if a field has been set.
func (o *NFProfile) HasSmfInfo() bool {
	if o != nil && o.SmfInfo != nil {
		return true
	}

	return false
}

// SetSmfInfo gets a reference to the given SmfInfo and assigns it to the SmfInfo field.
func (o *NFProfile) SetSmfInfo(v SmfInfo) {
	o.SmfInfo = &v
}

// GetSmfInfoExt returns the SmfInfoExt field value if set, zero value otherwise.
func (o *NFProfile) GetSmfInfoExt() []SmfInfo {
	if o == nil || o.SmfInfoExt == nil {
		var ret []SmfInfo
		return ret
	}
	return o.SmfInfoExt
}

// GetSmfInfoExtOk returns a tuple with the SmfInfoExt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetSmfInfoExtOk() ([]SmfInfo, bool) {
	if o == nil || o.SmfInfoExt == nil {
		return nil, false
	}
	return o.SmfInfoExt, true
}

// HasSmfInfoExt returns a boolean if a field has been set.
func (o *NFProfile) HasSmfInfoExt() bool {
	if o != nil && o.SmfInfoExt != nil {
		return true
	}

	return false
}

// SetSmfInfoExt gets a reference to the given []SmfInfo and assigns it to the SmfInfoExt field.
func (o *NFProfile) SetSmfInfoExt(v []SmfInfo) {
	o.SmfInfoExt = v
}

// GetUpfInfo returns the UpfInfo field value if set, zero value otherwise.
func (o *NFProfile) GetUpfInfo() UpfInfo {
	if o == nil || o.UpfInfo == nil {
		var ret UpfInfo
		return ret
	}
	return *o.UpfInfo
}

// GetUpfInfoOk returns a tuple with the UpfInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetUpfInfoOk() (*UpfInfo, bool) {
	if o == nil || o.UpfInfo == nil {
		return nil, false
	}
	return o.UpfInfo, true
}

// HasUpfInfo returns a boolean if a field has been set.
func (o *NFProfile) HasUpfInfo() bool {
	if o != nil && o.UpfInfo != nil {
		return true
	}

	return false
}

// SetUpfInfo gets a reference to the given UpfInfo and assigns it to the UpfInfo field.
func (o *NFProfile) SetUpfInfo(v UpfInfo) {
	o.UpfInfo = &v
}

// GetUpfInfoExt returns the UpfInfoExt field value if set, zero value otherwise.
func (o *NFProfile) GetUpfInfoExt() []UpfInfo {
	if o == nil || o.UpfInfoExt == nil {
		var ret []UpfInfo
		return ret
	}
	return o.UpfInfoExt
}

// GetUpfInfoExtOk returns a tuple with the UpfInfoExt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetUpfInfoExtOk() ([]UpfInfo, bool) {
	if o == nil || o.UpfInfoExt == nil {
		return nil, false
	}
	return o.UpfInfoExt, true
}

// HasUpfInfoExt returns a boolean if a field has been set.
func (o *NFProfile) HasUpfInfoExt() bool {
	if o != nil && o.UpfInfoExt != nil {
		return true
	}

	return false
}

// SetUpfInfoExt gets a reference to the given []UpfInfo and assigns it to the UpfInfoExt field.
func (o *NFProfile) SetUpfInfoExt(v []UpfInfo) {
	o.UpfInfoExt = v
}

// GetPcfInfo returns the PcfInfo field value if set, zero value otherwise.
func (o *NFProfile) GetPcfInfo() PcfInfo {
	if o == nil || o.PcfInfo == nil {
		var ret PcfInfo
		return ret
	}
	return *o.PcfInfo
}

// GetPcfInfoOk returns a tuple with the PcfInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetPcfInfoOk() (*PcfInfo, bool) {
	if o == nil || o.PcfInfo == nil {
		return nil, false
	}
	return o.PcfInfo, true
}

// HasPcfInfo returns a boolean if a field has been set.
func (o *NFProfile) HasPcfInfo() bool {
	if o != nil && o.PcfInfo != nil {
		return true
	}

	return false
}

// SetPcfInfo gets a reference to the given PcfInfo and assigns it to the PcfInfo field.
func (o *NFProfile) SetPcfInfo(v PcfInfo) {
	o.PcfInfo = &v
}

// GetPcfInfoExt returns the PcfInfoExt field value if set, zero value otherwise.
func (o *NFProfile) GetPcfInfoExt() []PcfInfo {
	if o == nil || o.PcfInfoExt == nil {
		var ret []PcfInfo
		return ret
	}
	return o.PcfInfoExt
}

// GetPcfInfoExtOk returns a tuple with the PcfInfoExt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetPcfInfoExtOk() ([]PcfInfo, bool) {
	if o == nil || o.PcfInfoExt == nil {
		return nil, false
	}
	return o.PcfInfoExt, true
}

// HasPcfInfoExt returns a boolean if a field has been set.
func (o *NFProfile) HasPcfInfoExt() bool {
	if o != nil && o.PcfInfoExt != nil {
		return true
	}

	return false
}

// SetPcfInfoExt gets a reference to the given []PcfInfo and assigns it to the PcfInfoExt field.
func (o *NFProfile) SetPcfInfoExt(v []PcfInfo) {
	o.PcfInfoExt = v
}

// GetBsfInfo returns the BsfInfo field value if set, zero value otherwise.
func (o *NFProfile) GetBsfInfo() BsfInfo {
	if o == nil || o.BsfInfo == nil {
		var ret BsfInfo
		return ret
	}
	return *o.BsfInfo
}

// GetBsfInfoOk returns a tuple with the BsfInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetBsfInfoOk() (*BsfInfo, bool) {
	if o == nil || o.BsfInfo == nil {
		return nil, false
	}
	return o.BsfInfo, true
}

// HasBsfInfo returns a boolean if a field has been set.
func (o *NFProfile) HasBsfInfo() bool {
	if o != nil && o.BsfInfo != nil {
		return true
	}

	return false
}

// SetBsfInfo gets a reference to the given BsfInfo and assigns it to the BsfInfo field.
func (o *NFProfile) SetBsfInfo(v BsfInfo) {
	o.BsfInfo = &v
}

// GetBsfInfoExt returns the BsfInfoExt field value if set, zero value otherwise.
func (o *NFProfile) GetBsfInfoExt() []BsfInfo {
	if o == nil || o.BsfInfoExt == nil {
		var ret []BsfInfo
		return ret
	}
	return o.BsfInfoExt
}

// GetBsfInfoExtOk returns a tuple with the BsfInfoExt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetBsfInfoExtOk() ([]BsfInfo, bool) {
	if o == nil || o.BsfInfoExt == nil {
		return nil, false
	}
	return o.BsfInfoExt, true
}

// HasBsfInfoExt returns a boolean if a field has been set.
func (o *NFProfile) HasBsfInfoExt() bool {
	if o != nil && o.BsfInfoExt != nil {
		return true
	}

	return false
}

// SetBsfInfoExt gets a reference to the given []BsfInfo and assigns it to the BsfInfoExt field.
func (o *NFProfile) SetBsfInfoExt(v []BsfInfo) {
	o.BsfInfoExt = v
}

// GetChfInfo returns the ChfInfo field value if set, zero value otherwise.
func (o *NFProfile) GetChfInfo() ChfInfo {
	if o == nil || o.ChfInfo == nil {
		var ret ChfInfo
		return ret
	}
	return *o.ChfInfo
}

// GetChfInfoOk returns a tuple with the ChfInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetChfInfoOk() (*ChfInfo, bool) {
	if o == nil || o.ChfInfo == nil {
		return nil, false
	}
	return o.ChfInfo, true
}

// HasChfInfo returns a boolean if a field has been set.
func (o *NFProfile) HasChfInfo() bool {
	if o != nil && o.ChfInfo != nil {
		return true
	}

	return false
}

// SetChfInfo gets a reference to the given ChfInfo and assigns it to the ChfInfo field.
func (o *NFProfile) SetChfInfo(v ChfInfo) {
	o.ChfInfo = &v
}

// GetChfInfoExt returns the ChfInfoExt field value if set, zero value otherwise.
func (o *NFProfile) GetChfInfoExt() []ChfInfo {
	if o == nil || o.ChfInfoExt == nil {
		var ret []ChfInfo
		return ret
	}
	return o.ChfInfoExt
}

// GetChfInfoExtOk returns a tuple with the ChfInfoExt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetChfInfoExtOk() ([]ChfInfo, bool) {
	if o == nil || o.ChfInfoExt == nil {
		return nil, false
	}
	return o.ChfInfoExt, true
}

// HasChfInfoExt returns a boolean if a field has been set.
func (o *NFProfile) HasChfInfoExt() bool {
	if o != nil && o.ChfInfoExt != nil {
		return true
	}

	return false
}

// SetChfInfoExt gets a reference to the given []ChfInfo and assigns it to the ChfInfoExt field.
func (o *NFProfile) SetChfInfoExt(v []ChfInfo) {
	o.ChfInfoExt = v
}

// GetNwdafInfo returns the NwdafInfo field value if set, zero value otherwise.
func (o *NFProfile) GetNwdafInfo() NwdafInfo {
	if o == nil || o.NwdafInfo == nil {
		var ret NwdafInfo
		return ret
	}
	return *o.NwdafInfo
}

// GetNwdafInfoOk returns a tuple with the NwdafInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetNwdafInfoOk() (*NwdafInfo, bool) {
	if o == nil || o.NwdafInfo == nil {
		return nil, false
	}
	return o.NwdafInfo, true
}

// HasNwdafInfo returns a boolean if a field has been set.
func (o *NFProfile) HasNwdafInfo() bool {
	if o != nil && o.NwdafInfo != nil {
		return true
	}

	return false
}

// SetNwdafInfo gets a reference to the given NwdafInfo and assigns it to the NwdafInfo field.
func (o *NFProfile) SetNwdafInfo(v NwdafInfo) {
	o.NwdafInfo = &v
}

// GetCustomInfo returns the CustomInfo field value if set, zero value otherwise.
func (o *NFProfile) GetCustomInfo() map[string]interface{} {
	if o == nil || o.CustomInfo == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomInfo
}

// GetCustomInfoOk returns a tuple with the CustomInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetCustomInfoOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomInfo == nil {
		return nil, false
	}
	return o.CustomInfo, true
}

// HasCustomInfo returns a boolean if a field has been set.
func (o *NFProfile) HasCustomInfo() bool {
	if o != nil && o.CustomInfo != nil {
		return true
	}

	return false
}

// SetCustomInfo gets a reference to the given map[string]interface{} and assigns it to the CustomInfo field.
func (o *NFProfile) SetCustomInfo(v map[string]interface{}) {
	o.CustomInfo = v
}

// GetRecoveryTime returns the RecoveryTime field value if set, zero value otherwise.
func (o *NFProfile) GetRecoveryTime() time.Time {
	if o == nil || o.RecoveryTime == nil {
		var ret time.Time
		return ret
	}
	return *o.RecoveryTime
}

// GetRecoveryTimeOk returns a tuple with the RecoveryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetRecoveryTimeOk() (*time.Time, bool) {
	if o == nil || o.RecoveryTime == nil {
		return nil, false
	}
	return o.RecoveryTime, true
}

// HasRecoveryTime returns a boolean if a field has been set.
func (o *NFProfile) HasRecoveryTime() bool {
	if o != nil && o.RecoveryTime != nil {
		return true
	}

	return false
}

// SetRecoveryTime gets a reference to the given time.Time and assigns it to the RecoveryTime field.
func (o *NFProfile) SetRecoveryTime(v time.Time) {
	o.RecoveryTime = &v
}

// GetNfServicePersistence returns the NfServicePersistence field value if set, zero value otherwise.
func (o *NFProfile) GetNfServicePersistence() bool {
	if o == nil || o.NfServicePersistence == nil {
		var ret bool
		return ret
	}
	return *o.NfServicePersistence
}

// GetNfServicePersistenceOk returns a tuple with the NfServicePersistence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetNfServicePersistenceOk() (*bool, bool) {
	if o == nil || o.NfServicePersistence == nil {
		return nil, false
	}
	return o.NfServicePersistence, true
}

// HasNfServicePersistence returns a boolean if a field has been set.
func (o *NFProfile) HasNfServicePersistence() bool {
	if o != nil && o.NfServicePersistence != nil {
		return true
	}

	return false
}

// SetNfServicePersistence gets a reference to the given bool and assigns it to the NfServicePersistence field.
func (o *NFProfile) SetNfServicePersistence(v bool) {
	o.NfServicePersistence = &v
}

// GetNfServices returns the NfServices field value if set, zero value otherwise.
func (o *NFProfile) GetNfServices() []NFService {
	if o == nil || o.NfServices == nil {
		var ret []NFService
		return ret
	}
	return o.NfServices
}

// GetNfServicesOk returns a tuple with the NfServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetNfServicesOk() ([]NFService, bool) {
	if o == nil || o.NfServices == nil {
		return nil, false
	}
	return o.NfServices, true
}

// HasNfServices returns a boolean if a field has been set.
func (o *NFProfile) HasNfServices() bool {
	if o != nil && o.NfServices != nil {
		return true
	}

	return false
}

// SetNfServices gets a reference to the given []NFService and assigns it to the NfServices field.
func (o *NFProfile) SetNfServices(v []NFService) {
	o.NfServices = v
}

// GetDefaultNotificationSubscriptions returns the DefaultNotificationSubscriptions field value if set, zero value otherwise.
func (o *NFProfile) GetDefaultNotificationSubscriptions() []DefaultNotificationSubscription {
	if o == nil || o.DefaultNotificationSubscriptions == nil {
		var ret []DefaultNotificationSubscription
		return ret
	}
	return o.DefaultNotificationSubscriptions
}

// GetDefaultNotificationSubscriptionsOk returns a tuple with the DefaultNotificationSubscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetDefaultNotificationSubscriptionsOk() ([]DefaultNotificationSubscription, bool) {
	if o == nil || o.DefaultNotificationSubscriptions == nil {
		return nil, false
	}
	return o.DefaultNotificationSubscriptions, true
}

// HasDefaultNotificationSubscriptions returns a boolean if a field has been set.
func (o *NFProfile) HasDefaultNotificationSubscriptions() bool {
	if o != nil && o.DefaultNotificationSubscriptions != nil {
		return true
	}

	return false
}

// SetDefaultNotificationSubscriptions gets a reference to the given []DefaultNotificationSubscription and assigns it to the DefaultNotificationSubscriptions field.
func (o *NFProfile) SetDefaultNotificationSubscriptions(v []DefaultNotificationSubscription) {
	o.DefaultNotificationSubscriptions = v
}

func (o NFProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nfInstanceId"] = o.NfInstanceId
	}
	if o.NfInstanceName != nil {
		toSerialize["nfInstanceName"] = o.NfInstanceName
	}
	if true {
		toSerialize["nfType"] = o.NfType
	}
	if true {
		toSerialize["nfStatus"] = o.NfStatus
	}
	if o.PlmnList != nil {
		toSerialize["plmnList"] = o.PlmnList
	}
	if o.SNssais != nil {
		toSerialize["sNssais"] = o.SNssais
	}
	if o.PerPlmnSnssaiList != nil {
		toSerialize["perPlmnSnssaiList"] = o.PerPlmnSnssaiList
	}
	if o.NsiList != nil {
		toSerialize["nsiList"] = o.NsiList
	}
	if o.Fqdn != nil {
		toSerialize["fqdn"] = o.Fqdn
	}
	if o.Ipv4Addresses != nil {
		toSerialize["ipv4Addresses"] = o.Ipv4Addresses
	}
	if o.Ipv6Addresses != nil {
		toSerialize["ipv6Addresses"] = o.Ipv6Addresses
	}
	if o.Capacity != nil {
		toSerialize["capacity"] = o.Capacity
	}
	if o.Load != nil {
		toSerialize["load"] = o.Load
	}
	if o.Locality != nil {
		toSerialize["locality"] = o.Locality
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.UdrInfo != nil {
		toSerialize["udrInfo"] = o.UdrInfo
	}
	if o.UdrInfoExt != nil {
		toSerialize["udrInfoExt"] = o.UdrInfoExt
	}
	if o.UdmInfo != nil {
		toSerialize["udmInfo"] = o.UdmInfo
	}
	if o.UdmInfoExt != nil {
		toSerialize["udmInfoExt"] = o.UdmInfoExt
	}
	if o.AusfInfo != nil {
		toSerialize["ausfInfo"] = o.AusfInfo
	}
	if o.AusfInfoExt != nil {
		toSerialize["ausfInfoExt"] = o.AusfInfoExt
	}
	if o.AmfInfo != nil {
		toSerialize["amfInfo"] = o.AmfInfo
	}
	if o.AmfInfoExt != nil {
		toSerialize["amfInfoExt"] = o.AmfInfoExt
	}
	if o.SmfInfo != nil {
		toSerialize["smfInfo"] = o.SmfInfo
	}
	if o.SmfInfoExt != nil {
		toSerialize["smfInfoExt"] = o.SmfInfoExt
	}
	if o.UpfInfo != nil {
		toSerialize["upfInfo"] = o.UpfInfo
	}
	if o.UpfInfoExt != nil {
		toSerialize["upfInfoExt"] = o.UpfInfoExt
	}
	if o.PcfInfo != nil {
		toSerialize["pcfInfo"] = o.PcfInfo
	}
	if o.PcfInfoExt != nil {
		toSerialize["pcfInfoExt"] = o.PcfInfoExt
	}
	if o.BsfInfo != nil {
		toSerialize["bsfInfo"] = o.BsfInfo
	}
	if o.BsfInfoExt != nil {
		toSerialize["bsfInfoExt"] = o.BsfInfoExt
	}
	if o.ChfInfo != nil {
		toSerialize["chfInfo"] = o.ChfInfo
	}
	if o.ChfInfoExt != nil {
		toSerialize["chfInfoExt"] = o.ChfInfoExt
	}
	if o.NwdafInfo != nil {
		toSerialize["nwdafInfo"] = o.NwdafInfo
	}
	if o.CustomInfo != nil {
		toSerialize["customInfo"] = o.CustomInfo
	}
	if o.RecoveryTime != nil {
		toSerialize["recoveryTime"] = o.RecoveryTime
	}
	if o.NfServicePersistence != nil {
		toSerialize["nfServicePersistence"] = o.NfServicePersistence
	}
	if o.NfServices != nil {
		toSerialize["nfServices"] = o.NfServices
	}
	if o.DefaultNotificationSubscriptions != nil {
		toSerialize["defaultNotificationSubscriptions"] = o.DefaultNotificationSubscriptions
	}
	return json.Marshal(toSerialize)
}

type NullableNFProfile struct {
	value *NFProfile
	isSet bool
}

func (v NullableNFProfile) Get() *NFProfile {
	return v.value
}

func (v *NullableNFProfile) Set(val *NFProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableNFProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableNFProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFProfile(val *NFProfile) *NullableNFProfile {
	return &NullableNFProfile{value: val, isSet: true}
}

func (v NullableNFProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
