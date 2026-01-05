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

// NFService struct for NFService
type NFService struct {
	ServiceInstanceId                string                            `json:"serviceInstanceId"`
	ServiceName                      ServiceName                       `json:"serviceName"`
	Versions                         []NFServiceVersion                `json:"versions"`
	Scheme                           UriScheme                         `json:"scheme"`
	NfServiceStatus                  NFServiceStatus                   `json:"nfServiceStatus"`
	Fqdn                             *string                           `json:"fqdn,omitempty"`
	IpEndPoints                      []IpEndPoint                      `json:"ipEndPoints,omitempty"`
	ApiPrefix                        *string                           `json:"apiPrefix,omitempty"`
	DefaultNotificationSubscriptions []DefaultNotificationSubscription `json:"defaultNotificationSubscriptions,omitempty"`
	Capacity                         *int32                            `json:"capacity,omitempty"`
	Load                             *int32                            `json:"load,omitempty"`
	Priority                         *int32                            `json:"priority,omitempty"`
	RecoveryTime                     *time.Time                        `json:"recoveryTime,omitempty"`
	ChfServiceInfo                   *ChfServiceInfo                   `json:"chfServiceInfo,omitempty"`
	SupportedFeatures                *string                           `json:"supportedFeatures,omitempty"`
}

// NewNFService instantiates a new NFService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNFService(serviceInstanceId string, serviceName ServiceName, versions []NFServiceVersion, scheme UriScheme, nfServiceStatus NFServiceStatus) *NFService {
	this := NFService{}
	this.ServiceInstanceId = serviceInstanceId
	this.ServiceName = serviceName
	this.Versions = versions
	this.Scheme = scheme
	this.NfServiceStatus = nfServiceStatus
	return &this
}

// NewNFServiceWithDefaults instantiates a new NFService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNFServiceWithDefaults() *NFService {
	this := NFService{}
	return &this
}

// GetServiceInstanceId returns the ServiceInstanceId field value
func (o *NFService) GetServiceInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceInstanceId
}

// GetServiceInstanceIdOk returns a tuple with the ServiceInstanceId field value
// and a boolean to check if the value has been set.
func (o *NFService) GetServiceInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceInstanceId, true
}

// SetServiceInstanceId sets field value
func (o *NFService) SetServiceInstanceId(v string) {
	o.ServiceInstanceId = v
}

// GetServiceName returns the ServiceName field value
func (o *NFService) GetServiceName() ServiceName {
	if o == nil {
		var ret ServiceName
		return ret
	}

	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *NFService) GetServiceNameOk() (*ServiceName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value
func (o *NFService) SetServiceName(v ServiceName) {
	o.ServiceName = v
}

// GetVersions returns the Versions field value
func (o *NFService) GetVersions() []NFServiceVersion {
	if o == nil {
		var ret []NFServiceVersion
		return ret
	}

	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value
// and a boolean to check if the value has been set.
func (o *NFService) GetVersionsOk() ([]NFServiceVersion, bool) {
	if o == nil {
		return nil, false
	}
	return o.Versions, true
}

// SetVersions sets field value
func (o *NFService) SetVersions(v []NFServiceVersion) {
	o.Versions = v
}

// GetScheme returns the Scheme field value
func (o *NFService) GetScheme() UriScheme {
	if o == nil {
		var ret UriScheme
		return ret
	}

	return o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value
// and a boolean to check if the value has been set.
func (o *NFService) GetSchemeOk() (*UriScheme, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scheme, true
}

// SetScheme sets field value
func (o *NFService) SetScheme(v UriScheme) {
	o.Scheme = v
}

// GetNfServiceStatus returns the NfServiceStatus field value
func (o *NFService) GetNfServiceStatus() NFServiceStatus {
	if o == nil {
		var ret NFServiceStatus
		return ret
	}

	return o.NfServiceStatus
}

// GetNfServiceStatusOk returns a tuple with the NfServiceStatus field value
// and a boolean to check if the value has been set.
func (o *NFService) GetNfServiceStatusOk() (*NFServiceStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NfServiceStatus, true
}

// SetNfServiceStatus sets field value
func (o *NFService) SetNfServiceStatus(v NFServiceStatus) {
	o.NfServiceStatus = v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *NFService) GetFqdn() string {
	if o == nil || o.Fqdn == nil {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetFqdnOk() (*string, bool) {
	if o == nil || o.Fqdn == nil {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *NFService) HasFqdn() bool {
	if o != nil && o.Fqdn != nil {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *NFService) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetIpEndPoints returns the IpEndPoints field value if set, zero value otherwise.
func (o *NFService) GetIpEndPoints() []IpEndPoint {
	if o == nil || o.IpEndPoints == nil {
		var ret []IpEndPoint
		return ret
	}
	return o.IpEndPoints
}

// GetIpEndPointsOk returns a tuple with the IpEndPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetIpEndPointsOk() ([]IpEndPoint, bool) {
	if o == nil || o.IpEndPoints == nil {
		return nil, false
	}
	return o.IpEndPoints, true
}

// HasIpEndPoints returns a boolean if a field has been set.
func (o *NFService) HasIpEndPoints() bool {
	if o != nil && o.IpEndPoints != nil {
		return true
	}

	return false
}

// SetIpEndPoints gets a reference to the given []IpEndPoint and assigns it to the IpEndPoints field.
func (o *NFService) SetIpEndPoints(v []IpEndPoint) {
	o.IpEndPoints = v
}

// GetApiPrefix returns the ApiPrefix field value if set, zero value otherwise.
func (o *NFService) GetApiPrefix() string {
	if o == nil || o.ApiPrefix == nil {
		var ret string
		return ret
	}
	return *o.ApiPrefix
}

// GetApiPrefixOk returns a tuple with the ApiPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetApiPrefixOk() (*string, bool) {
	if o == nil || o.ApiPrefix == nil {
		return nil, false
	}
	return o.ApiPrefix, true
}

// HasApiPrefix returns a boolean if a field has been set.
func (o *NFService) HasApiPrefix() bool {
	if o != nil && o.ApiPrefix != nil {
		return true
	}

	return false
}

// SetApiPrefix gets a reference to the given string and assigns it to the ApiPrefix field.
func (o *NFService) SetApiPrefix(v string) {
	o.ApiPrefix = &v
}

// GetDefaultNotificationSubscriptions returns the DefaultNotificationSubscriptions field value if set, zero value otherwise.
func (o *NFService) GetDefaultNotificationSubscriptions() []DefaultNotificationSubscription {
	if o == nil || o.DefaultNotificationSubscriptions == nil {
		var ret []DefaultNotificationSubscription
		return ret
	}
	return o.DefaultNotificationSubscriptions
}

// GetDefaultNotificationSubscriptionsOk returns a tuple with the DefaultNotificationSubscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetDefaultNotificationSubscriptionsOk() ([]DefaultNotificationSubscription, bool) {
	if o == nil || o.DefaultNotificationSubscriptions == nil {
		return nil, false
	}
	return o.DefaultNotificationSubscriptions, true
}

// HasDefaultNotificationSubscriptions returns a boolean if a field has been set.
func (o *NFService) HasDefaultNotificationSubscriptions() bool {
	if o != nil && o.DefaultNotificationSubscriptions != nil {
		return true
	}

	return false
}

// SetDefaultNotificationSubscriptions gets a reference to the given []DefaultNotificationSubscription and assigns it to the DefaultNotificationSubscriptions field.
func (o *NFService) SetDefaultNotificationSubscriptions(v []DefaultNotificationSubscription) {
	o.DefaultNotificationSubscriptions = v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *NFService) GetCapacity() int32 {
	if o == nil || o.Capacity == nil {
		var ret int32
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetCapacityOk() (*int32, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *NFService) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int32 and assigns it to the Capacity field.
func (o *NFService) SetCapacity(v int32) {
	o.Capacity = &v
}

// GetLoad returns the Load field value if set, zero value otherwise.
func (o *NFService) GetLoad() int32 {
	if o == nil || o.Load == nil {
		var ret int32
		return ret
	}
	return *o.Load
}

// GetLoadOk returns a tuple with the Load field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetLoadOk() (*int32, bool) {
	if o == nil || o.Load == nil {
		return nil, false
	}
	return o.Load, true
}

// HasLoad returns a boolean if a field has been set.
func (o *NFService) HasLoad() bool {
	if o != nil && o.Load != nil {
		return true
	}

	return false
}

// SetLoad gets a reference to the given int32 and assigns it to the Load field.
func (o *NFService) SetLoad(v int32) {
	o.Load = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *NFService) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *NFService) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *NFService) SetPriority(v int32) {
	o.Priority = &v
}

// GetRecoveryTime returns the RecoveryTime field value if set, zero value otherwise.
func (o *NFService) GetRecoveryTime() time.Time {
	if o == nil || o.RecoveryTime == nil {
		var ret time.Time
		return ret
	}
	return *o.RecoveryTime
}

// GetRecoveryTimeOk returns a tuple with the RecoveryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetRecoveryTimeOk() (*time.Time, bool) {
	if o == nil || o.RecoveryTime == nil {
		return nil, false
	}
	return o.RecoveryTime, true
}

// HasRecoveryTime returns a boolean if a field has been set.
func (o *NFService) HasRecoveryTime() bool {
	if o != nil && o.RecoveryTime != nil {
		return true
	}

	return false
}

// SetRecoveryTime gets a reference to the given time.Time and assigns it to the RecoveryTime field.
func (o *NFService) SetRecoveryTime(v time.Time) {
	o.RecoveryTime = &v
}

// GetChfServiceInfo returns the ChfServiceInfo field value if set, zero value otherwise.
func (o *NFService) GetChfServiceInfo() ChfServiceInfo {
	if o == nil || o.ChfServiceInfo == nil {
		var ret ChfServiceInfo
		return ret
	}
	return *o.ChfServiceInfo
}

// GetChfServiceInfoOk returns a tuple with the ChfServiceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetChfServiceInfoOk() (*ChfServiceInfo, bool) {
	if o == nil || o.ChfServiceInfo == nil {
		return nil, false
	}
	return o.ChfServiceInfo, true
}

// HasChfServiceInfo returns a boolean if a field has been set.
func (o *NFService) HasChfServiceInfo() bool {
	if o != nil && o.ChfServiceInfo != nil {
		return true
	}

	return false
}

// SetChfServiceInfo gets a reference to the given ChfServiceInfo and assigns it to the ChfServiceInfo field.
func (o *NFService) SetChfServiceInfo(v ChfServiceInfo) {
	o.ChfServiceInfo = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *NFService) GetSupportedFeatures() string {
	if o == nil || o.SupportedFeatures == nil {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || o.SupportedFeatures == nil {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *NFService) HasSupportedFeatures() bool {
	if o != nil && o.SupportedFeatures != nil {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *NFService) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o NFService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serviceInstanceId"] = o.ServiceInstanceId
	}
	if true {
		toSerialize["serviceName"] = o.ServiceName
	}
	if true {
		toSerialize["versions"] = o.Versions
	}
	if true {
		toSerialize["scheme"] = o.Scheme
	}
	if true {
		toSerialize["nfServiceStatus"] = o.NfServiceStatus
	}
	if o.Fqdn != nil {
		toSerialize["fqdn"] = o.Fqdn
	}
	if o.IpEndPoints != nil {
		toSerialize["ipEndPoints"] = o.IpEndPoints
	}
	if o.ApiPrefix != nil {
		toSerialize["apiPrefix"] = o.ApiPrefix
	}
	if o.DefaultNotificationSubscriptions != nil {
		toSerialize["defaultNotificationSubscriptions"] = o.DefaultNotificationSubscriptions
	}
	if o.Capacity != nil {
		toSerialize["capacity"] = o.Capacity
	}
	if o.Load != nil {
		toSerialize["load"] = o.Load
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.RecoveryTime != nil {
		toSerialize["recoveryTime"] = o.RecoveryTime
	}
	if o.ChfServiceInfo != nil {
		toSerialize["chfServiceInfo"] = o.ChfServiceInfo
	}
	if o.SupportedFeatures != nil {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return json.Marshal(toSerialize)
}

type NullableNFService struct {
	value *NFService
	isSet bool
}

func (v NullableNFService) Get() *NFService {
	return v.value
}

func (v *NullableNFService) Set(val *NFService) {
	v.value = val
	v.isSet = true
}

func (v NullableNFService) IsSet() bool {
	return v.isSet
}

func (v *NullableNFService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFService(val *NFService) *NullableNFService {
	return &NullableNFService{value: val, isSet: true}
}

func (v NullableNFService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
