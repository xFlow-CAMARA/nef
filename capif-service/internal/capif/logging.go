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
	"time"

	"gitlab.eurecom.fr/open-exposure/nef/capif-service/internal/3gpp/logging_api_invocation"
)

func PushLoggingEntry(serviceId string, logData *LogInfo) error {
	/* find service profile */
	service := providerInstance.Services[serviceId]

	logginEntry := logging_api_invocation.NewInvocationLogWithDefaults()
	logginEntry.AefId = providerInstance.FuncId[AEF]
	logginEntry.ApiInvokerId = logData.InvokerId

	apiLog := logging_api_invocation.NewLogWithDefaults()
	apiLog.ApiId = serviceId
	apiLog.ApiName = service.ServiceName
	apiLog.ApiVersion = service.ServiceModel.AefProfiles[0].Versions[0].ApiVersion

	time := time.Unix(logData.TimeStamp, 0)
	src := &logging_api_invocation.InterfaceDescription{
		Ipv4Addr: &logData.SourceIp,
		Port:     logging_api_invocation.PtrInt32(logData.SourcePort),
	}
	dst := &logging_api_invocation.InterfaceDescription{
		Ipv4Addr: &logData.DestinationIp,
		Port:     logging_api_invocation.PtrInt32(logData.DestinationPort),
	}

	apiLog.InvocationTime = &time
	apiLog.SrcInterface = *logging_api_invocation.NewNullableInterfaceDescription(src)
	apiLog.DestInterface = *logging_api_invocation.NewNullableInterfaceDescription(dst)

	apiLog.Operation = logData.Operation
	apiLog.ResourceName = logData.Endpoint
	apiLog.Uri = logging_api_invocation.PtrString(logData.Endpoint)
	apiLog.Result = logData.Result
	apiLog.Protocol = logData.Protocol

	logginEntry.Logs = []logging_api_invocation.Log{*apiLog}

	configuration := logging_api_invocation.NewConfiguration()
	LoggingApiClient := logging_api_invocation.NewAPIClient(configuration)
	_, r, err := LoggingApiClient.DefaultAPI.AefIdLogsPost(context.Background(), logginEntry.AefId).InvocationLog(*logginEntry).Execute()

	if err != nil {
		log.Printf(
			"Error when calling `DefaultAPI.AefIdLogsPost`: %v\n",
			err,
		)
		log.Printf("Full HTTP response: %v\n", r)
		return err
	}

	return nil
}

func PushLoggingEntryMock(serviceId string, logData *LogInfo, token string) error {
	/* find service profile */

	logginEntry := logging_api_invocation.NewInvocationLogWithDefaults()
	logginEntry.AefId = providerInstance.FuncId[AEF]
	logginEntry.ApiInvokerId = "INVbeb515b0e11edf0547c114ff86f20d"

	apiLog := logging_api_invocation.NewLogWithDefaults()
	apiLog.ApiId = serviceId
	apiLog.ApiName = "sample-api"
	apiLog.ApiVersion = "v1"

	time := time.Unix(logData.TimeStamp, 0)
	src := &logging_api_invocation.InterfaceDescription{
		Ipv4Addr: &logData.SourceIp,
		Port:     logging_api_invocation.PtrInt32(logData.SourcePort),
	}
	dst := &logging_api_invocation.InterfaceDescription{
		Ipv4Addr: &logData.DestinationIp,
		Port:     logging_api_invocation.PtrInt32(logData.DestinationPort),
	}

	apiLog.InvocationTime = &time
	apiLog.SrcInterface = *logging_api_invocation.NewNullableInterfaceDescription(src)
	apiLog.DestInterface = *logging_api_invocation.NewNullableInterfaceDescription(dst)

	apiLog.Operation = logData.Operation
	apiLog.Uri = logging_api_invocation.PtrString(logData.Endpoint)
	apiLog.Result = logData.Result
	apiLog.Protocol = logData.Protocol

	logginEntry.Logs = []logging_api_invocation.Log{*apiLog}

	configuration := logging_api_invocation.NewConfiguration()
	LoggingApiClient := logging_api_invocation.NewAPIClient(configuration)
	_, r, err := LoggingApiClient.DefaultAPI.AefIdLogsPost(context.Background(), logginEntry.AefId).InvocationLog(*logginEntry).Execute()

	if err != nil {
		log.Printf(
			"Error when calling `DefaultAPI.AefIdLogsPost`: %v\n",
			err,
		)
		log.Printf("Full HTTP response: %v\n", r)
		return err
	}

	return nil
}
