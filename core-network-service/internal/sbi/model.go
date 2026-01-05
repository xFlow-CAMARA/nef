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
	"github.com/go-redis/redis"

	amf_client "gitlab.eurecom.fr/open-exposure/nef/core-network-service/internal/amfclient"
	smf_client "gitlab.eurecom.fr/open-exposure/nef/core-network-service/internal/smfclient"
)

// Global variable
var redisClient *redis.Client
var config SbiConfig

// ------------------------------------------------------------------------------
// Type of EngineConfig structure
type SbiConfig struct {
	Nrf struct {
		NrfName string `envconfig:"NRF_NAME"`
		Mcc     string `envconfig:"MCC"`
		Mnc     string `envconfig:"MNC"`
		Use     bool   `envconfig:"USE_NRF"`
	}
	Amf struct {
		IpAddr            string `envconfig:"AMF_IP_ADDR"`
		SubRoute          string `envconfig:"AMF_SUBSCR_ROUTE"`
		ApiRoute          string `envconfig:"AMF_API_ROUTE"`
		NotifCorrId       string `envconfig:"AMF_NOTIFY_CORRELATION_ID"`
		NotifId           string `envconfig:"AMF_NOTIFICATION_ID"`
		NorifForwardRoute string `envconfig:"AMF_NOTIFICATION_FORWARD_ROUTE"`
	}
	Smf struct {
		IpAddr            string `envconfig:"SMF_IP_ADDR"`
		SubRoute          string `envconfig:"SMF_SUBSCR_ROUTE"`
		ApiRoute          string `envconfig:"SMF_API_ROUTE"`
		NotifCorrId       string `envconfig:"SMF_NOTIFY_CORRELATION_ID"`
		NotifId           string `envconfig:"SMF_NOTIFICATION_ID"`
		NorifForwardRoute string `envconfig:"SMF_NOTIFICATION_FORWARD_ROUTE"`
	}
	Database struct {
		Uri string `envconfig:"REDIS_ADDR"`
	}
	Server struct {
		NotifUri string `envconfig:"EVENT_NOTIFY_URI"`
		Uri      string `envconfig:"SERVER_ADDR"`
	}
}
type pduSesEst struct {
	AdIpv4Addr  *string
	Dnn         *string
	PduSeId     *int32
	PduSessType *smf_client.PduSessionType
	Snssai      *smf_client.Snssai
	TimeStamp   int64
}

type pduSesRel struct {
	Ipv4Addr    *string
	Dnn         *string
	PduSeId     *int32
	PduSessType *smf_client.PduSessionType
	Snssai      *smf_client.Snssai
	TimeStamp   int64
}

type ueIpCh struct {
	AdIpv4Addr *string
	PduSeId    *int32
	TimeStamp  int64
}

type ddds struct {
	DddStatus  *smf_client.DlDataDeliveryStatus
	PduSeId    *int32
	UeIpv4Addr *string
	Dnn        *string
	Snssai     *smf_client.Snssai
	TimeStamp  int64
}

type qosMon struct {
	Customized_data *smf_client.CustomizedData
	PduSeId         *int32
	TimeStamp       int64
}

type rmInfo struct {
	RmInfo    amf_client.RmInfo
	TimeStamp int64
}

type cmInfo struct {
	CmInfo    amf_client.CmInfo
	TimeStamp int64
}

type location struct {
	UserLocation amf_client.UserLocation
	TimeStamp    int64
}

type lossOfConnectReason struct {
	LossOfConnectReason amf_client.LossOfConnectivityReasonAnyOf
	TimeStamp           int64
}
