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

package capif

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"gitlab.eurecom.fr/open-exposure/nef/capif-service/internal/3gpp/publish_api_service"
	"gitlab.eurecom.fr/open-exposure/nef/capif-service/internal/utils"
)

var providerInstance *ApiProvider

func InitializeDefaultProvider() {
	/*Create a provider instance*/
	providerInstance = NewApiProvider()
	providerInstance.InstanceName = "open-exposure-nef" + uuid.New().String()
}

func RegisterDefaultProvider() error {

	/*retrieve auth details*/
	auth_info := AuthProviderResponse{}
	err := retrieveAuthorizationDetails(&auth_info)
	if err != nil {
		return fmt.Errorf("cannot reach authentication provider due to: %s", err.Error())
	}

	log.Printf("access token refreshesd")

	/*Store CA ROOT certificate for CAPIF ingress*/
	caroot := auth_info.CaRoot
	err = os.WriteFile(utils.Config.CertsPath+"ca.crt", []byte(caroot), 0755)
	if err != nil {
		return fmt.Errorf("cannot store CA ROOT certficate due to: %s", err.Error())
	}
	/*Parse CORE Endpoints*/
	log.Printf(`Core Instance Details:
	ApiProviderManagement: %s
	ServiceAPIs: %s
	ApiInvokerManagement: %s
	PublishedAPIs: %s
	CapifSecurity: %s`,
		auth_info.CcfApiOnboardingUrl,
		auth_info.CcfPublishUrl,
		auth_info.CcfOnboardingUrl,
		auth_info.CcfDiscoverUrl,
		auth_info.CcfSecurityUrl)

	/*Attempt registration to CAPIF Core*/
	err = RegisterProvider(providerInstance, auth_info.AccessToken)
	if err != nil {
		return fmt.Errorf("cannot register API Provider to Capif Core due to: %s", err.Error())
	}

	log.Default().Printf("API Provider instance has been registered (uuid=%s)", providerInstance.InstanceId)
	return nil
}

func DestroyDefaultProvider() error {
	token, err := retrieveAccessToken()
	if err != nil {
		return err
	}

	err = UnregisterProvider(providerInstance, token)
	if err != nil {
		return fmt.Errorf("cannot unregister API Provider from Capif Core")

	}

	log.Default().Printf("API Provider instance has been unregistered (uuid=%s)", providerInstance.InstanceId)
	return nil
}

func PublishNewApi(api *publish_api_service.ServiceAPIDescription) (string, error) {
	token, err := retrieveAccessToken()
	if err != nil {
		return "", err
	}

	if api == nil {
		return "", fmt.Errorf("malformed request: api definition is not present")
	}

	/* Add AEF Id to API model */
	api.AefProfiles[0].AefId = providerInstance.FuncId[AEF]
	/* Instantiate new service profile */
	service := NewServiceApiProfile(api.ApiName)
	service.ServiceModel = api

	/* append igress controller fqdn to the service endpoint */
	if utils.Config.IngressFqdn != "" {
		api.AefProfiles[0].InterfaceDescriptions = append(api.AefProfiles[0].InterfaceDescriptions,
			publish_api_service.InterfaceDescription{
				Fqdn: publish_api_service.PtrString(utils.Config.IngressFqdn),
			})
	}

	sid, err := PublishService(providerInstance, service, token)
	if err != nil {
		return "", err
	}
	providerInstance.Services[sid] = service
	return sid, nil
}

func DeletePublishedApi(serviceId string) error {
	token, err := retrieveAccessToken()
	if err != nil {
		return err
	}
	err = DeleteService(providerInstance, serviceId, token)
	if err != nil {
		return err
	}
	return nil
}

func NewLogEntry(serviceId string, logData *LogInfo) error {
	/*token, err := retrieveAccessToken()
	if err != nil {
		return err
	}*/
	err := PushLoggingEntry(serviceId, logData)
	if err != nil {
		return err
	}

	return nil
}

func ValidateToken(serviceId string, authData *OauthInfo) error {

	/* retrieve service related api name */
	apiName := providerInstance.Services[serviceId].ServiceModel.ApiName
	aefId := providerInstance.FuncId[ApiProviderFunctionType("AEF")]

	err := ValidateInvokerToken(aefId, apiName, authData)
	if err != nil {
		return err
	}
	return nil
}
