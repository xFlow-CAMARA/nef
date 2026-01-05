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
	"fmt"
	"log"
	"strings"

	pcfclient "gitlab.eurecom.fr/open-exposure/nef/pcfclient"
)

// ------------------------------------------------------------------------------
func (c *Connector) CreatePolicyAuthzSubscription(pa_ctx pcfclient.AppSessionContext) (string, error) {
	//1.Discover PCF
	url := c.DiscoverPCFEndpoint()
	if url == "" {
		log.Printf("Could not find any availble PCF Instance")
		return "", fmt.Errorf("no PCF available, ignoring")
	}
	//2. Setup API Client and perform registration
	configuration := pcfclient.NewConfiguration(url, c.Cfg().Sbi.Httpversion)
	pcfPolicyAuthClient := pcfclient.NewAPIClient(configuration)
	_, r, err := pcfPolicyAuthClient.ApplicationSessionsCollectionAPI.PostAppSessions(
		context.Background()).AppSessionContext(pa_ctx).Execute()

	if err != nil {
		log.Printf("cannot create policy authorization subscription")
		log.Printf("Full HTTP response: %v\n", r)
		return "", err
	}
	loc := r.Header.Get("Location")
	appSessId := strings.Split(loc, "/app-sessions/")[1]
	log.Printf("activated policy authorization subscription at %s", appSessId)

	return appSessId, nil
}

// ------------------------------------------------------------------------------
func (c *Connector) UpdatePolicyAuthzSubscription(AppSessId string, pa_ctx pcfclient.AppSessionContext) (string, error) {
	//1.Discover PCF
	url := c.DiscoverPCFEndpoint()
	if url == "" {
		log.Printf("Could not find any availble PCF Instance")
		return "", fmt.Errorf("no PCF available, ignoring")
	}

	//2. Setup API Client
	configuration := pcfclient.NewConfiguration(url, c.Cfg().Sbi.Httpversion)
	pcfPolicyAuthClient := pcfclient.NewAPIClient(configuration)

	//3. Delete current PCF subscription
	_, r, err := pcfPolicyAuthClient.IndividualApplicationSessionContextDocumentAPI.DeleteAppSession(
		context.Background(), AppSessId).Execute()

	//4. Create new one if no errors
	if err == nil {
		_, r, err = pcfPolicyAuthClient.ApplicationSessionsCollectionAPI.PostAppSessions(
			context.Background()).AppSessionContext(pa_ctx).Execute()
		if err == nil {
			loc := r.Header.Get("Location")
			appSessId := strings.Split(loc, "/app-sessions/")[1]
			log.Printf("updated policy authorization subscription, new id is %s", appSessId)
			return appSessId, nil
		}
	}

	log.Printf("could not update policy authorization subscription \n")
	log.Printf("Full HTTP response: %v\n", r)
	return "", err
}

// ------------------------------------------------------------------------------
func (c *Connector) RemovePolicyAuthzSubscription(AppSessId string) error {
	//1.Discover PCF
	url := c.DiscoverPCFEndpoint()
	if url == "" {
		log.Printf("Could not find any availble PCF Instance")
		return fmt.Errorf("no PCF available, ignoring")
	}

	//2. Setup API Client and perform registration
	configuration := pcfclient.NewConfiguration(url, c.Cfg().Sbi.Httpversion)
	pcfPolicyAuthClient := pcfclient.NewAPIClient(configuration)
	_, r, err := pcfPolicyAuthClient.IndividualApplicationSessionContextDocumentAPI.DeleteAppSession(
		context.Background(), AppSessId).Execute()
	if err != nil {
		log.Printf(
			"Error when calling `SubscriptionsCollectionCollectionApi.CreateSubscription``: %v\n",
			err,
		)
		log.Printf("Full HTTP response: %v\n", r)
		return err
	}

	return nil
}
