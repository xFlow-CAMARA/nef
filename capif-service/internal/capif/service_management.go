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
	"log"

	serviceApi "gitlab.eurecom.fr/open-exposure/nef/capif-service/internal/3gpp/publish_api_service"
)

func PublishService(profile *ApiProvider, service *ServiceApiProfile, token string) (string, error) {
	log.Printf("Publishing new service using provider (uuid=%s)", profile.InstanceId)

	apfId := profile.FuncId[APF]
	configuration := serviceApi.NewConfiguration()
	ServiceApiClient := serviceApi.NewAPIClient(configuration)
	resp, r, err := ServiceApiClient.DefaultAPI.ApfIdServiceApisPost(context.Background(), apfId).ServiceAPIDescription(*service.ServiceModel).Execute()

	if err != nil {
		log.Printf(
			"Error when calling `DefaultAPI.RegistrationsPost`: %v\n",
			err,
		)
		log.Printf("Full HTTP response: %v\n", r)
		return "", err
	}

	if resp.ApiId != nil {
		service.ServiceId = *resp.ApiId
	}

	log.Printf("Published new service %s (uuid=%s)", service.ServiceName, service.ServiceId)
	return service.ServiceId, nil
}

func DeleteService(profile *ApiProvider, apiId string, token string) error {
	log.Printf("Deleting service (uuid=%s) using provider (uuid=%s)", apiId, profile.InstanceId)

	apfId := profile.FuncId[APF]
	configuration := serviceApi.NewConfiguration()
	ServiceApiClient := serviceApi.NewAPIClient(configuration)
	r, err := ServiceApiClient.DefaultAPI.ApfIdServiceApisServiceApiIdDelete(context.Background(), apiId, apfId).Execute()

	if err != nil {
		log.Printf(
			"Error when calling `DefaultAPI.RegistrationsPost`: %v\n",
			err,
		)
		log.Printf("Full HTTP response: %v\n", r)
		return err
	}
	log.Printf("Deleted service (uuid=%s)", apiId)
	return nil
}
