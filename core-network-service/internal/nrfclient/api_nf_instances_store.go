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
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// NFInstancesStoreApiService NFInstancesStoreApi service
type NFInstancesStoreApiService service

type ApiSearchNFInstancesRequest struct {
	ctx                     context.Context
	ApiService              *NFInstancesStoreApiService
	targetNfType            *NFType
	requesterNfType         *NFType
	requesterNfInstanceId   *string
	serviceNames            *[]ServiceName
	requesterNfInstanceFqdn *string
	targetPlmnList          *[]PlmnId
	requesterPlmnList       *[]PlmnId
	targetNfInstanceId      *string
	targetNfFqdn            *string
	hnrfUri                 *string
	snssais                 *[]Snssai
	requesterSnssais        *[]Snssai
	plmnSpecificSnssaiList  *[]PlmnSnssai
	dnn                     *string
	nsiList                 *[]string
	smfServingArea          *string
	tai                     *Tai
	amfRegionId             *string
	amfSetId                *string
	guami                   *Guami
	supi                    *string
	ueIpv4Address           *string
	ipDomain                *string
	ueIpv6Prefix            *Ipv6Prefix
	pgwInd                  *bool
	pgw                     *string
	gpsi                    *string
	externalGroupIdentity   *string
	dataSet                 *DataSetId
	routingIndicator        *string
	groupIdList             *[]string
	dnaiList                *[]string
	pduSessionTypes         *[]PduSessionType
	eventIdList             *[]EventId
	nwdafEventList          *[]NwdafEvent
	supportedFeatures       *string
	upfIwkEpsInd            *bool
	chfSupportedPlmn        *PlmnId
	preferredLocality       *string
	accessType              *AccessType
	limit                   *int32
	requiredFeatures        *[]string
	complexQuery            *ComplexQuery
	maxPayloadSize          *int32
	atsssCapability         *AtsssCapability
	upfUeIpAddrInd          *bool
	ifNoneMatch             *string
}

// Type of the target NF
func (r ApiSearchNFInstancesRequest) TargetNfType(targetNfType NFType) ApiSearchNFInstancesRequest {
	r.targetNfType = &targetNfType
	return r
}

// Type of the requester NF
func (r ApiSearchNFInstancesRequest) RequesterNfType(requesterNfType NFType) ApiSearchNFInstancesRequest {
	r.requesterNfType = &requesterNfType
	return r
}

// NfInstanceId of the requester NF
func (r ApiSearchNFInstancesRequest) RequesterNfInstanceId(requesterNfInstanceId string) ApiSearchNFInstancesRequest {
	r.requesterNfInstanceId = &requesterNfInstanceId
	return r
}

// Names of the services offered by the NF
func (r ApiSearchNFInstancesRequest) ServiceNames(serviceNames []ServiceName) ApiSearchNFInstancesRequest {
	r.serviceNames = &serviceNames
	return r
}

// FQDN of the requester NF
func (r ApiSearchNFInstancesRequest) RequesterNfInstanceFqdn(requesterNfInstanceFqdn string) ApiSearchNFInstancesRequest {
	r.requesterNfInstanceFqdn = &requesterNfInstanceFqdn
	return r
}

// Id of the PLMN of the target NF
func (r ApiSearchNFInstancesRequest) TargetPlmnList(targetPlmnList []PlmnId) ApiSearchNFInstancesRequest {
	r.targetPlmnList = &targetPlmnList
	return r
}

// Id of the PLMN where the NF issuing the Discovery request is located
func (r ApiSearchNFInstancesRequest) RequesterPlmnList(requesterPlmnList []PlmnId) ApiSearchNFInstancesRequest {
	r.requesterPlmnList = &requesterPlmnList
	return r
}

// Identity of the NF instance being discovered
func (r ApiSearchNFInstancesRequest) TargetNfInstanceId(targetNfInstanceId string) ApiSearchNFInstancesRequest {
	r.targetNfInstanceId = &targetNfInstanceId
	return r
}

// FQDN of the NF instance being discovered
func (r ApiSearchNFInstancesRequest) TargetNfFqdn(targetNfFqdn string) ApiSearchNFInstancesRequest {
	r.targetNfFqdn = &targetNfFqdn
	return r
}

// Uri of the home NRF
func (r ApiSearchNFInstancesRequest) HnrfUri(hnrfUri string) ApiSearchNFInstancesRequest {
	r.hnrfUri = &hnrfUri
	return r
}

// Slice info of the target NF
func (r ApiSearchNFInstancesRequest) Snssais(snssais []Snssai) ApiSearchNFInstancesRequest {
	r.snssais = &snssais
	return r
}

// Slice info of the requester NF
func (r ApiSearchNFInstancesRequest) RequesterSnssais(requesterSnssais []Snssai) ApiSearchNFInstancesRequest {
	r.requesterSnssais = &requesterSnssais
	return r
}

// PLMN specific Slice info of the target NF
func (r ApiSearchNFInstancesRequest) PlmnSpecificSnssaiList(plmnSpecificSnssaiList []PlmnSnssai) ApiSearchNFInstancesRequest {
	r.plmnSpecificSnssaiList = &plmnSpecificSnssaiList
	return r
}

// Dnn supported by the BSF, SMF or UPF
func (r ApiSearchNFInstancesRequest) Dnn(dnn string) ApiSearchNFInstancesRequest {
	r.dnn = &dnn
	return r
}

// NSI IDs that are served by the services being discovered
func (r ApiSearchNFInstancesRequest) NsiList(nsiList []string) ApiSearchNFInstancesRequest {
	r.nsiList = &nsiList
	return r
}

func (r ApiSearchNFInstancesRequest) SmfServingArea(smfServingArea string) ApiSearchNFInstancesRequest {
	r.smfServingArea = &smfServingArea
	return r
}

// Tracking Area Identity
func (r ApiSearchNFInstancesRequest) Tai(tai Tai) ApiSearchNFInstancesRequest {
	r.tai = &tai
	return r
}

// AMF Region Identity
func (r ApiSearchNFInstancesRequest) AmfRegionId(amfRegionId string) ApiSearchNFInstancesRequest {
	r.amfRegionId = &amfRegionId
	return r
}

// AMF Set Identity
func (r ApiSearchNFInstancesRequest) AmfSetId(amfSetId string) ApiSearchNFInstancesRequest {
	r.amfSetId = &amfSetId
	return r
}

// Guami used to search for an appropriate AMF
func (r ApiSearchNFInstancesRequest) Guami(guami Guami) ApiSearchNFInstancesRequest {
	r.guami = &guami
	return r
}

// SUPI of the user
func (r ApiSearchNFInstancesRequest) Supi(supi string) ApiSearchNFInstancesRequest {
	r.supi = &supi
	return r
}

// IPv4 address of the UE
func (r ApiSearchNFInstancesRequest) UeIpv4Address(ueIpv4Address string) ApiSearchNFInstancesRequest {
	r.ueIpv4Address = &ueIpv4Address
	return r
}

// IP domain of the UE, which supported by BSF
func (r ApiSearchNFInstancesRequest) IpDomain(ipDomain string) ApiSearchNFInstancesRequest {
	r.ipDomain = &ipDomain
	return r
}

// IPv6 prefix of the UE
func (r ApiSearchNFInstancesRequest) UeIpv6Prefix(ueIpv6Prefix Ipv6Prefix) ApiSearchNFInstancesRequest {
	r.ueIpv6Prefix = &ueIpv6Prefix
	return r
}

// Combined PGW-C and SMF or a standalone SMF
func (r ApiSearchNFInstancesRequest) PgwInd(pgwInd bool) ApiSearchNFInstancesRequest {
	r.pgwInd = &pgwInd
	return r
}

// PGW FQDN of a combined PGW-C and SMF
func (r ApiSearchNFInstancesRequest) Pgw(pgw string) ApiSearchNFInstancesRequest {
	r.pgw = &pgw
	return r
}

// GPSI of the user
func (r ApiSearchNFInstancesRequest) Gpsi(gpsi string) ApiSearchNFInstancesRequest {
	r.gpsi = &gpsi
	return r
}

// external group identifier of the user
func (r ApiSearchNFInstancesRequest) ExternalGroupIdentity(externalGroupIdentity string) ApiSearchNFInstancesRequest {
	r.externalGroupIdentity = &externalGroupIdentity
	return r
}

// data set supported by the NF
func (r ApiSearchNFInstancesRequest) DataSet(dataSet DataSetId) ApiSearchNFInstancesRequest {
	r.dataSet = &dataSet
	return r
}

// routing indicator in SUCI
func (r ApiSearchNFInstancesRequest) RoutingIndicator(routingIndicator string) ApiSearchNFInstancesRequest {
	r.routingIndicator = &routingIndicator
	return r
}

// Group IDs of the NFs being discovered
func (r ApiSearchNFInstancesRequest) GroupIdList(groupIdList []string) ApiSearchNFInstancesRequest {
	r.groupIdList = &groupIdList
	return r
}

// Data network access identifiers of the NFs being discovered
func (r ApiSearchNFInstancesRequest) DnaiList(dnaiList []string) ApiSearchNFInstancesRequest {
	r.dnaiList = &dnaiList
	return r
}

// list of PDU Session Type required to be supported by the target NF
func (r ApiSearchNFInstancesRequest) PduSessionTypes(pduSessionTypes []PduSessionType) ApiSearchNFInstancesRequest {
	r.pduSessionTypes = &pduSessionTypes
	return r
}

// Analytics event(s) requested to be supported by the Nnwdaf_AnalyticsInfo service
func (r ApiSearchNFInstancesRequest) EventIdList(eventIdList []EventId) ApiSearchNFInstancesRequest {
	r.eventIdList = &eventIdList
	return r
}

// Analytics event(s) requested to be supported by the Nnwdaf_EventsSubscription service.
func (r ApiSearchNFInstancesRequest) NwdafEventList(nwdafEventList []NwdafEvent) ApiSearchNFInstancesRequest {
	r.nwdafEventList = &nwdafEventList
	return r
}

// Features required to be supported by the target NF
func (r ApiSearchNFInstancesRequest) SupportedFeatures(supportedFeatures string) ApiSearchNFInstancesRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

// UPF supporting interworking with EPS or not
func (r ApiSearchNFInstancesRequest) UpfIwkEpsInd(upfIwkEpsInd bool) ApiSearchNFInstancesRequest {
	r.upfIwkEpsInd = &upfIwkEpsInd
	return r
}

// PLMN ID supported by a CHF
func (r ApiSearchNFInstancesRequest) ChfSupportedPlmn(chfSupportedPlmn PlmnId) ApiSearchNFInstancesRequest {
	r.chfSupportedPlmn = &chfSupportedPlmn
	return r
}

// preferred target NF location
func (r ApiSearchNFInstancesRequest) PreferredLocality(preferredLocality string) ApiSearchNFInstancesRequest {
	r.preferredLocality = &preferredLocality
	return r
}

// AccessType supported by the target NF
func (r ApiSearchNFInstancesRequest) AccessType(accessType AccessType) ApiSearchNFInstancesRequest {
	r.accessType = &accessType
	return r
}

// Maximum number of NFProfiles to return in the response
func (r ApiSearchNFInstancesRequest) Limit(limit int32) ApiSearchNFInstancesRequest {
	r.limit = &limit
	return r
}

// Features required to be supported by the target NF
func (r ApiSearchNFInstancesRequest) RequiredFeatures(requiredFeatures []string) ApiSearchNFInstancesRequest {
	r.requiredFeatures = &requiredFeatures
	return r
}

// the complex query condition expression
func (r ApiSearchNFInstancesRequest) ComplexQuery(complexQuery ComplexQuery) ApiSearchNFInstancesRequest {
	r.complexQuery = &complexQuery
	return r
}

// Maximum payload size of the response expressed in kilo octets
func (r ApiSearchNFInstancesRequest) MaxPayloadSize(maxPayloadSize int32) ApiSearchNFInstancesRequest {
	r.maxPayloadSize = &maxPayloadSize
	return r
}

// ATSSS Capability
func (r ApiSearchNFInstancesRequest) AtsssCapability(atsssCapability AtsssCapability) ApiSearchNFInstancesRequest {
	r.atsssCapability = &atsssCapability
	return r
}

// UPF supporting allocating UE IP addresses/prefixes
func (r ApiSearchNFInstancesRequest) UpfUeIpAddrInd(upfUeIpAddrInd bool) ApiSearchNFInstancesRequest {
	r.upfUeIpAddrInd = &upfUeIpAddrInd
	return r
}

// Validator for conditional requests, as described in IETF RFC 7232, 3.2
func (r ApiSearchNFInstancesRequest) IfNoneMatch(ifNoneMatch string) ApiSearchNFInstancesRequest {
	r.ifNoneMatch = &ifNoneMatch
	return r
}

func (r ApiSearchNFInstancesRequest) Execute() (*SearchResult, *http.Response, error) {
	return r.ApiService.SearchNFInstancesExecute(r)
}

/*
SearchNFInstances Search a collection of NF Instances

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSearchNFInstancesRequest
*/
func (a *NFInstancesStoreApiService) SearchNFInstances(ctx context.Context) ApiSearchNFInstancesRequest {
	return ApiSearchNFInstancesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SearchResult
func (a *NFInstancesStoreApiService) SearchNFInstancesExecute(r ApiSearchNFInstancesRequest) (*SearchResult, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SearchResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NFInstancesStoreApiService.SearchNFInstances")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/nf-instances"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.targetNfType == nil {
		return localVarReturnValue, nil, reportError("targetNfType is required and must be specified")
	}
	if r.requesterNfType == nil {
		return localVarReturnValue, nil, reportError("requesterNfType is required and must be specified")
	}

	localVarQueryParams.Add("target-nf-type", parameterToString(*r.targetNfType, ""))
	localVarQueryParams.Add("requester-nf-type", parameterToString(*r.requesterNfType, ""))
	if r.requesterNfInstanceId != nil {
		localVarQueryParams.Add("requester-nf-instance-id", parameterToString(*r.requesterNfInstanceId, ""))
	}
	if r.serviceNames != nil {
		localVarQueryParams.Add("service-names", parameterToString(*r.serviceNames, "csv"))
	}
	if r.requesterNfInstanceFqdn != nil {
		localVarQueryParams.Add("requester-nf-instance-fqdn", parameterToString(*r.requesterNfInstanceFqdn, ""))
	}
	if r.targetPlmnList != nil {
		localVarQueryParams.Add("target-plmn-list", parameterToString(*r.targetPlmnList, "csv"))
	}
	if r.requesterPlmnList != nil {
		localVarQueryParams.Add("requester-plmn-list", parameterToString(*r.requesterPlmnList, "csv"))
	}
	if r.targetNfInstanceId != nil {
		localVarQueryParams.Add("target-nf-instance-id", parameterToString(*r.targetNfInstanceId, ""))
	}
	if r.targetNfFqdn != nil {
		localVarQueryParams.Add("target-nf-fqdn", parameterToString(*r.targetNfFqdn, ""))
	}
	if r.hnrfUri != nil {
		localVarQueryParams.Add("hnrf-uri", parameterToString(*r.hnrfUri, ""))
	}
	if r.snssais != nil {
		localVarQueryParams.Add("snssais", parameterToString(*r.snssais, "csv"))
	}
	if r.requesterSnssais != nil {
		localVarQueryParams.Add("requester-snssais", parameterToString(*r.requesterSnssais, "csv"))
	}
	if r.plmnSpecificSnssaiList != nil {
		localVarQueryParams.Add("plmn-specific-snssai-list", parameterToString(*r.plmnSpecificSnssaiList, "csv"))
	}
	if r.dnn != nil {
		localVarQueryParams.Add("dnn", parameterToString(*r.dnn, ""))
	}
	if r.nsiList != nil {
		localVarQueryParams.Add("nsi-list", parameterToString(*r.nsiList, "csv"))
	}
	if r.smfServingArea != nil {
		localVarQueryParams.Add("smf-serving-area", parameterToString(*r.smfServingArea, ""))
	}
	if r.tai != nil {
		localVarQueryParams.Add("tai", parameterToString(*r.tai, ""))
	}
	if r.amfRegionId != nil {
		localVarQueryParams.Add("amf-region-id", parameterToString(*r.amfRegionId, ""))
	}
	if r.amfSetId != nil {
		localVarQueryParams.Add("amf-set-id", parameterToString(*r.amfSetId, ""))
	}
	if r.guami != nil {
		localVarQueryParams.Add("guami", parameterToString(*r.guami, ""))
	}
	if r.supi != nil {
		localVarQueryParams.Add("supi", parameterToString(*r.supi, ""))
	}
	if r.ueIpv4Address != nil {
		localVarQueryParams.Add("ue-ipv4-address", parameterToString(*r.ueIpv4Address, ""))
	}
	if r.ipDomain != nil {
		localVarQueryParams.Add("ip-domain", parameterToString(*r.ipDomain, ""))
	}
	if r.ueIpv6Prefix != nil {
		localVarQueryParams.Add("ue-ipv6-prefix", parameterToString(*r.ueIpv6Prefix, ""))
	}
	if r.pgwInd != nil {
		localVarQueryParams.Add("pgw-ind", parameterToString(*r.pgwInd, ""))
	}
	if r.pgw != nil {
		localVarQueryParams.Add("pgw", parameterToString(*r.pgw, ""))
	}
	if r.gpsi != nil {
		localVarQueryParams.Add("gpsi", parameterToString(*r.gpsi, ""))
	}
	if r.externalGroupIdentity != nil {
		localVarQueryParams.Add("external-group-identity", parameterToString(*r.externalGroupIdentity, ""))
	}
	if r.dataSet != nil {
		localVarQueryParams.Add("data-set", parameterToString(*r.dataSet, ""))
	}
	if r.routingIndicator != nil {
		localVarQueryParams.Add("routing-indicator", parameterToString(*r.routingIndicator, ""))
	}
	if r.groupIdList != nil {
		localVarQueryParams.Add("group-id-list", parameterToString(*r.groupIdList, "csv"))
	}
	if r.dnaiList != nil {
		localVarQueryParams.Add("dnai-list", parameterToString(*r.dnaiList, "csv"))
	}
	if r.pduSessionTypes != nil {
		localVarQueryParams.Add("pdu-session-types", parameterToString(*r.pduSessionTypes, "csv"))
	}
	if r.eventIdList != nil {
		localVarQueryParams.Add("event-id-list", parameterToString(*r.eventIdList, "csv"))
	}
	if r.nwdafEventList != nil {
		localVarQueryParams.Add("nwdaf-event-list", parameterToString(*r.nwdafEventList, "csv"))
	}
	if r.supportedFeatures != nil {
		localVarQueryParams.Add("supported-features", parameterToString(*r.supportedFeatures, ""))
	}
	if r.upfIwkEpsInd != nil {
		localVarQueryParams.Add("upf-iwk-eps-ind", parameterToString(*r.upfIwkEpsInd, ""))
	}
	if r.chfSupportedPlmn != nil {
		localVarQueryParams.Add("chf-supported-plmn", parameterToString(*r.chfSupportedPlmn, ""))
	}
	if r.preferredLocality != nil {
		localVarQueryParams.Add("preferred-locality", parameterToString(*r.preferredLocality, ""))
	}
	if r.accessType != nil {
		localVarQueryParams.Add("access-type", parameterToString(*r.accessType, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.requiredFeatures != nil {
		localVarQueryParams.Add("required-features", parameterToString(*r.requiredFeatures, "csv"))
	}
	if r.complexQuery != nil {
		localVarQueryParams.Add("complex-query", parameterToString(*r.complexQuery, ""))
	}
	if r.maxPayloadSize != nil {
		localVarQueryParams.Add("max-payload-size", parameterToString(*r.maxPayloadSize, ""))
	}
	if r.atsssCapability != nil {
		localVarQueryParams.Add("atsss-capability", parameterToString(*r.atsssCapability, ""))
	}
	if r.upfUeIpAddrInd != nil {
		localVarQueryParams.Add("upf-ue-ip-addr-ind", parameterToString(*r.upfUeIpAddrInd, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ifNoneMatch != nil {
		localVarHeaderParams["If-None-Match"] = parameterToString(*r.ifNoneMatch, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 411 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 413 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 415 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 501 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
