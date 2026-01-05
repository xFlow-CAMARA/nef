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

package connector

import (
	"context"
	"log"
	"net/url"
	"strconv"

	nrf_client "gitlab.eurecom.fr/open-exposure/nef/nrfclient"
)

// ------------------------------------------------------------------------------
func (c *Connector) DiscoverPCFEndpoint() string {
	if c.Cfg().Sbi.UseNrf {
		pcfs := c.nrfDiscoverPCFInstances()
		if len(pcfs) > 0 {
			//TODO implement support for multiple PCFs - for the moment just
			//return the first and only one
			if len(pcfs) > 1 {
				log.Printf("Warning: Multiple PCFs are not supported yet")
			}
			server_addr := c.pcfProfileGetPolicyAuthorizationInfo(pcfs[0])
			if server_addr == "" {
				log.Fatal("npcf_policyAuthorization api is not supported")
			}
			return server_addr
		}
		return ""
	} else {
		if c.Cfg().Sbi.PcfSvc == "" {
			log.Default().Fatalf("NRF is disabled and PCF svc is not set")
		}
		return c.Cfg().Sbi.PcfSvc + "/npcf-policyauthorization/v1"
	}
}

// ------------------------------------------------------------------------------
func (c *Connector) pcfProfileGetPolicyAuthorizationInfo(pcf nrf_client.NFProfile) string {

	for _, service := range pcf.GetNfServices() {

		if service.GetServiceName().ServiceNameAnyOf.Ptr() != nil &&
			*service.GetServiceName().ServiceNameAnyOf.Ptr() == nrf_client.NPCF_POLICYAUTHORIZATION {

			ip := service.IpEndPoints[0].GetIpv4Address()
			port := strconv.FormatInt(int64(service.GetIpEndPoints()[0].GetPort()), 10)
			ver := service.GetVersions()[0].GetApiVersionInUri()

			url := url.URL{
				Scheme: "http",
				Host:   ip + ":" + port,
				Path:   string(nrf_client.NPCF_POLICYAUTHORIZATION) + "/" + ver,
			}

			log.Printf("Found PCF(%s) supporting %s\n ip: %s\n port: %s\n apiVersion: %s\n", pcf.GetNfInstanceId(), string(nrf_client.NPCF_POLICYAUTHORIZATION), ip, port, ver)
			return url.String()
		}
	}
	return ""
}

// ------------------------------------------------------------------------------
func (c *Connector) nrfDiscoverPCFInstances() []nrf_client.NFProfile {
	targetNfType := nrf_client.NFType_PCF
	requesterNfType := nrf_client.NFType_NEF
	configuration := nrf_client.NewConfiguration(c.Cfg().Sbi.NrfSvc, c.Cfg().Sbi.Httpversion)
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

	if resp != nil {
		return resp.NfInstances

	} else {
		return []nrf_client.NFProfile{}
	}

}
