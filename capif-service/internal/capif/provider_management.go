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
	"context"
	"fmt"
	"log"
	"os"

	apiProvManagement "gitlab.eurecom.fr/open-exposure/nef/capif-service/internal/3gpp/api_provider_management"
	"gitlab.eurecom.fr/open-exposure/nef/capif-service/internal/utils"
)

func RegisterProvider(profile *ApiProvider, token string) error {
	log.Printf("Registering API Provider instance (uuid=%s)", profile.InstanceId)

	// Register 3 entities (AEF, APF, AMF)
	provDetail := &apiProvManagement.APIProviderEnrolmentDetails{}
	fillProviderEnrollementDetails(profile, token, provDetail)
	configuration := apiProvManagement.NewConfiguration()
	configuration.AddDefaultHeader("Authorization", "Bearer "+token)
	apiProvManagementClient := apiProvManagement.NewAPIClient(configuration)
	resp, r, err := apiProvManagementClient.DefaultAPI.RegistrationsPost(
		context.Background()).APIProviderEnrolmentDetails(*provDetail).Execute()
	if err != nil {
		log.Printf(
			"Error when calling `DefaultAPI.RegistrationsPost`: %v\n",
			err,
		)
		log.Printf("Full HTTP response: %v\n", r)
	}

	if len(resp.ApiProvFuncs) != 3 {
		return fmt.Errorf("enrollment was not successful")
	}

	/* Get enrollment results and store signed certificates */
	for _, entity := range resp.ApiProvFuncs {
		regInfo := entity.GetRegInfo()
		role := entity.GetApiProvFuncRole()
		if role != "" {
			crtBytes := []byte(regInfo.GetApiProvCert())
			profile.FuncId[ApiProviderFunctionType(role)] = *entity.ApiProvFuncId
			/* try to store certificate on file as well */
			err = os.WriteFile(utils.Config.CertsPath+""+role+".crt", crtBytes, 0755)
			if err != nil {
				return fmt.Errorf("could not store certificate on file: %v", err)
			}
		} else {
			return fmt.Errorf("enrollment was not successful")
		}
	}

	profile.InstanceId = *resp.ApiProvDomId
	return nil
}

func fillProviderEnrollementDetails(profile *ApiProvider, access_token string, provDetail *apiProvManagement.APIProviderEnrolmentDetails) {
	entities := []ApiProviderFunctionType{AMF, APF, AEF}
	provDetail.ApiProvDomInfo = apiProvManagement.PtrString(profile.InstanceName)
	provDetail.ApiProvName = apiProvManagement.PtrString(profile.InstanceName)
	provDetail.RegSec = access_token
	provDetail.ApiProvFuncs = make([]apiProvManagement.APIProviderFunctionDetails, 0)
	for _, entity := range entities {
		csr, _, err := CreateCsr(string(entity))
		if err != nil {
			continue
		}
		apiProvFunc := apiProvManagement.APIProviderFunctionDetails{
			ApiProvFuncInfo: apiProvManagement.PtrString(profile.InstanceName + "_" + string(entity)),
			ApiProvFuncRole: string(entity),
			RegInfo: apiProvManagement.RegistrationInformation{
				ApiProvPubKey: string(csr),
			},
		}
		provDetail.SuppFeat = apiProvManagement.PtrString("0")
		provDetail.ApiProvFuncs = append(provDetail.ApiProvFuncs, apiProvFunc)
	}

}

func UnregisterProvider(profile *ApiProvider, token string) error {
	log.Printf("Unregistering API Provider instance (uuid=%s)", profile.InstanceId)

	configuration := apiProvManagement.NewConfiguration()
	configuration.HTTPClient = httpClientAMFCerts()
	apiProvManagementClient := apiProvManagement.NewAPIClient(configuration)
	r, err := apiProvManagementClient.DefaultAPI.RegistrationsRegistrationIdDelete(context.Background(), profile.InstanceId).Execute()
	if err != nil {
		log.Printf(
			"Error when calling `DefaultAPI.RegistrationsPost`: %v\n",
			err,
		)
		log.Printf("Full HTTP response: %v\n", r)
	}
	return nil
}
