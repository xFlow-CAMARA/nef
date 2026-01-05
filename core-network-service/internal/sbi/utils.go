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

package sbi

import (
	"context"
	"log"
	"net/url"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"

	amf_client "gitlab.eurecom.fr/open-exposure/nef/core-network-service/internal/amfclient"
	nrf_client "gitlab.eurecom.fr/open-exposure/nef/core-network-service/internal/nrfclient"
	smf_client "gitlab.eurecom.fr/open-exposure/nef/core-network-service/internal/smfclient"
)

// ------------------------------------------------------------------------------
// InitConfig - Initialize global variables (cfg and mongoClient) and subscribe to AMF and SMF
func InitConfig() {
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err.Error())
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr: config.Database.Uri,
	})

	amfs := []string{}
	smfs := []string{}

	if config.Nrf.Use {
		amfs = nrfDiscoverAmfInstances()
		smfs = nrfDiscoverSmfInstances()
	} else {
		amfs = append(amfs, config.Amf.IpAddr+"/"+string(nrf_client.NAMF_EVTS)+"/v1")
		smfs = append(smfs, config.Smf.IpAddr+"/"+string(nrf_client.NSMF_EVENT_EXPOSURE)+"/v1")
	}

	for _, amf := range amfs {
		if amf != "" {
			// Subscribe to all event notifications from AMF
			amfEventSubscription(
				amf,
				config.Server.NotifUri+config.Amf.ApiRoute,
				config.Amf.NotifCorrId,
				uuid.NewString(),
			)
		}
	}

	for _, smf := range smfs {
		// Subscribe to all event notifications from SMF
		if smf != "" {
			smfEventSubscription(
				smf,
				config.Server.NotifUri+config.Smf.ApiRoute,
				uuid.NewString(),
			)
		}
	}

}

// ------------------------------------------------------------------------------
func amfEventSubscription(
	amfUrl string,
	amfEventNotifyUri string,
	amfNotifyCorrelationId string,
	amfNfId string,
) {
	// Store all AMF event types
	var amfEvents []amf_client.AmfEvent
	for _, amfEventTypeAnyOf := range amf_client.AllowedAmfEventTypeAnyOfEnumValues {
		amfEvents = append(amfEvents, *amf_client.NewAmfEvent(amfEventTypeAnyOf))
	}
	// Subscribe to all AMF event types
	amfCreateEventSubscription := *amf_client.NewAmfCreateEventSubscription(
		*amf_client.NewAmfEventSubscription(
			amfEvents,
			amfEventNotifyUri,
			amfNotifyCorrelationId,
			amfNfId,
		),
	)
	configuration := amf_client.NewConfiguration(amfUrl)

	amfApiClient := amf_client.NewAPIClient(configuration)
	resp, r, err := amfApiClient.SubscriptionsCollectionCollectionApi.CreateSubscription(
		context.Background()).AmfCreateEventSubscription(amfCreateEventSubscription).Execute()
	if err != nil {
		log.Printf(
			"Error when calling `SubscriptionsCollectionCollectionApi.CreateSubscription``: %v\n",
			err,
		)
		log.Printf("Full HTTP response: %v\n", r)
	}

	log.Printf(
		"namf-evts subscription created for:")
	for _, e := range resp.Subscription.EventList {
		log.Printf("\t%s\n", string(e.GetType()))
	}
}

// ------------------------------------------------------------------------------
func smfEventSubscription(smfUrl string, smfEventNotifyUri string, smfNfId string) {

	// Store all SMF event types
	var smfEventSubs []smf_client.EventSubscription
	smfEventSubs = append(smfEventSubs,
		*smf_client.NewEventSubscription(smf_client.SMFEVENTANYOF_PDU_SES_EST),
	)
	smfEventSubs = append(smfEventSubs,
		*smf_client.NewEventSubscription(smf_client.SMFEVENTANYOF_UE_IP_CH),
	)
	smfEventSubs = append(smfEventSubs,
		*smf_client.NewEventSubscription(smf_client.SMFEVENTANYOF_PLMN_CH),
	)
	smfEventSubs = append(smfEventSubs,
		*smf_client.NewEventSubscription(smf_client.SMFEVENTANYOF_DDDS),
	)
	smfEventSubs = append(smfEventSubs,
		*smf_client.NewEventSubscription(smf_client.SMFEVENTANYOF_PDU_SES_REL),
	)
	smfEventSubs = append(smfEventSubs,
		*smf_client.NewEventSubscription(smf_client.SMFEVENTANYOF_QOS_MON),
	)
	// Subscribe to all SMF event types
	nsmfEventExposure := *smf_client.NewNsmfEventExposure(
		smfNfId,
		smfEventNotifyUri,
		smfEventSubs,
	)
	configuration := smf_client.NewConfiguration(smfUrl)
	smfApiClient := smf_client.NewAPIClient(configuration)
	resp, r, err := smfApiClient.SubscriptionsCollectionApi.CreateIndividualSubcription(
		context.Background()).NsmfEventExposure(nsmfEventExposure).Execute()
	if err != nil {
		log.Printf(
			"Error when calling `SubscriptionsCollectionApi.CreateIndividualSubcription``: %v\n",
			err,
		)
		log.Printf("Full HTTP response: %v\n", r)
	}
	// response from `CreateIndividualSubcription`: NsmfEventExposure

	log.Printf(
		"nsmf-event-exposure subscription created for: ")
	for _, e := range resp.EventSubs {
		log.Printf("\t%s\n", string(e.GetEvent()))
	}
}

// ------------------------------------------------------------------------------
func nrfDiscoverInstances(targetNfType nrf_client.NFType) *nrf_client.SearchResult {

	requesterNfType := nrf_client.NFType_NEF
	configuration := nrf_client.NewConfiguration()
	nrfApiClient := nrf_client.NewAPIClient(configuration)
	resp, r, err := nrfApiClient.NFInstancesStoreApi.SearchNFInstances(
		context.Background()).TargetNfType(targetNfType).RequesterNfType(requesterNfType).Execute()

	if err != nil {
		log.Printf(
			"Error when calling `nrfApiClient.NFInstancesStoreApi`: %v\n",
			err,
		)
		log.Printf("Full HTTP response: %v\n", r)
	}
	return resp
}

// ------------------------------------------------------------------------------
func nrfDiscoverSmfInstances() []string {
	targetNfType := nrf_client.NFType_SMF
	res := nrfDiscoverInstances(targetNfType)
	if res != nil {
		smfs := []string{}
		for _, smf := range res.NfInstances {
			smfs = append(smfs, smfProfileGetEventExposureInfo(smf))
		}
		return smfs
	} else {
		return []string{}
	}
}

// ------------------------------------------------------------------------------
func nrfDiscoverAmfInstances() []string {
	targetNfType := nrf_client.NFType_AMF
	res := nrfDiscoverInstances(targetNfType)
	if res != nil {
		amfs := []string{}
		for _, amf := range res.NfInstances {
			amfs = append(amfs, amfProfileGetEventExposureInfo(amf))
		}
		return amfs
	} else {
		return []string{}
	}

}

// ------------------------------------------------------------------------------
func smfProfileGetEventExposureInfo(smf nrf_client.NFProfile) string {

	for _, service := range smf.GetNfServices() {

		if service.GetServiceName().ServiceNameAnyOf.Ptr() != nil &&
			*service.GetServiceName().ServiceNameAnyOf.Ptr() == nrf_client.NSMF_EVENT_EXPOSURE {

			ip := service.IpEndPoints[0].GetIpv4Address()
			port := strconv.FormatInt(int64(service.GetIpEndPoints()[0].GetPort()), 10)
			ver := service.GetVersions()[0].GetApiVersionInUri()

			url := url.URL{
				Scheme: "http",
				Host:   ip + ":" + port,
				Path:   string(nrf_client.NSMF_EVENT_EXPOSURE) + "/" + ver,
			}

			log.Printf("Found SMF(%s) supporting %s\n ip: %s\n port: %s\n apiVersion: %s\n", smf.GetNfInstanceId(), string(nrf_client.NSMF_EVENT_EXPOSURE), ip, port, ver)
			return url.String()
		}
	}
	return ""
}

// ------------------------------------------------------------------------------
func amfProfileGetEventExposureInfo(amf nrf_client.NFProfile) string {

	for _, service := range amf.GetNfServices() {
		if service.GetServiceName().ServiceNameAnyOf.Ptr() != nil &&
			*service.GetServiceName().ServiceNameAnyOf.Ptr() == nrf_client.NAMF_EVTS {

			ip := service.IpEndPoints[0].GetIpv4Address()
			port := strconv.FormatInt(int64(service.GetIpEndPoints()[0].GetPort()), 10)
			ver := service.GetVersions()[0].GetApiVersionInUri()

			url := url.URL{
				Scheme: "http",
				Host:   ip + ":" + port,
				Path:   string(nrf_client.NAMF_EVTS) + "/" + ver,
			}

			log.Printf("Found AMF(%s) supporting %s\n ip: %s\n port: %s\n apiVersion: %s\n", amf.GetNfInstanceId(), string(nrf_client.NAMF_EVTS), ip, port, ver)

			return url.String()
		}
	}
	return ""
}
