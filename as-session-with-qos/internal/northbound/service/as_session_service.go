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

package service

import (
	"fmt"
	"log"
	"net/http"

	"gitlab.eurecom.fr/open-exposure/nef/as-session-with-qos/internal/connector"
	contexts "gitlab.eurecom.fr/open-exposure/nef/as-session-with-qos/internal/context"
	"gitlab.eurecom.fr/open-exposure/nef/as-session-with-qos/internal/models"
	"gitlab.eurecom.fr/open-exposure/nef/as-session-with-qos/pkg/config"
	"gitlab.eurecom.fr/open-exposure/nef/pcfclient"
)

type app interface {
	Cfg() *config.AppConfig
	Connector() *connector.Connector
	Ctx() *contexts.AsSessionAppCtx
}

type Service struct {
	app
}

func NewAsSessionWithQoSService(app app) *Service {
	svc := &Service{app: app}
	return svc
}

// ------------------------------------------------------------------------------
func (s *Service) PostSessionWithQoSSubscription(afId string, data *models.AsSessionWithQoSSubscription) (string, int, error) {
	var loc string

	af := s.Ctx().GetAf(afId)
	if af == nil {
		af = s.Ctx().AddAf(afId)
	}

	af.Mu.Lock()
	defer af.Mu.Unlock()

	if len(data.Gpsi) > 0 || len(data.UeIpv4Addr) > 0 || len(data.UeIpv6Addr) > 0 {
		// Single UE, sent to PCF
		pa_ctx, err := s.sessionWithQoS2PolicyAuthz(afId, data)
		if err == nil {
			loc, err = s.Connector().CreatePolicyAuthzSubscription(*pa_ctx)
			if err != nil {
				return "", http.StatusInternalServerError, fmt.Errorf("could not create PCF Policy Authorization context")
			}
			tiLoc, subCtx := af.NewAfSubscription(data)
			subCtx.AppSessId = loc
			return tiLoc, http.StatusCreated, nil
		} else {
			return "", http.StatusBadRequest, err
		}
	} else if len(data.ExtGroupId) > 0 {
		return "", http.StatusNotImplemented, fmt.Errorf("udr policies not supported yet")
	} else {
		return "", http.StatusNotImplemented, fmt.Errorf("not in the single or multiple UE cases")
	}

}

// ------------------------------------------------------------------------------
func (s *Service) DeleteSessionWithQoSSubscription(afId string, subId string) (int, error) {

	af := s.Ctx().GetAf(afId)

	if af != nil {
		af.Mu.Lock()
		defer af.Mu.Unlock()

		sub := af.GetAfSubscription(subId)
		if sub != nil {
			if len(sub.AppSessId) > 0 {
				err := s.Connector().RemovePolicyAuthzSubscription(sub.AppSessId)
				if err != nil {
					return http.StatusInternalServerError, err
				} else {
					err = af.DeleteAfscription(subId)
					if err != nil {
						return http.StatusInternalServerError, err
					}
					return http.StatusNoContent, nil
				}
			}
		}
	}
	return http.StatusNotFound, fmt.Errorf("could not find af/subId")
}

// ------------------------------------------------------------------------------
func (s *Service) FetchSessionWithQoSSubscription(afId string, subId string) (*models.AsSessionWithQoSSubscription, int, error) {

	af := s.Ctx().GetAf(afId)

	if af != nil {
		af.Mu.Lock()
		defer af.Mu.Unlock()

		sub := af.GetAfSubscription(subId)
		if sub != nil {
			return sub.Data, http.StatusOK, nil
		}

	}
	return nil, http.StatusNotFound, fmt.Errorf("could not find af/subId")
}

// ------------------------------------------------------------------------------
func (s *Service) sessionWithQoS2PolicyAuthz(afId string, data *models.AsSessionWithQoSSubscription) (*pcfclient.AppSessionContext, error) {
	/*Convert NEF model to PCF models*/

	err := validateSubscriptionData(data)
	if err != nil {
		return nil, err
	}

	req := pcfclient.AppSessionContextReqData{
		AfAppId: &afId,
		UeIpv4:  pcfclient.PtrString(data.UeIpv4Addr),
		// TODO add support for ipv6
		NotifUri: data.NotificationDestination,
		SuppFeat: s.Cfg().SupportedFeat,
		Dnn:      pcfclient.PtrString(data.Dnn),
		SliceInfo: &pcfclient.Snssai{
			Sst: data.Snssai.Sst,
			Sd:  pcfclient.PtrString(data.Snssai.Sd),
		},
		QosDuration: pcfclient.PtrInt32(data.QosDuration),
	}

	medComponents := make(map[string]pcfclient.MediaComponent)
	req.MedComponents = &medComponents

	medComponent := pcfclient.MediaComponent{}
	medComponent.SetMedCompN(1)
	medComponent.SetFStatus("ENABLED")
	medComponent.SetAfAppId(afId)
	/*Based on the QoS reference retrieve the BwInfo*/
	qosInfo, ok := s.Cfg().QosConf[data.QosReference]
	if !ok {
		log.Default().Printf("requested QoS reference do not exists")
		return nil, fmt.Errorf("invalid QoS reference")
	}
	medComponent.SetMarBwDl(qosInfo.MarBwDl)
	medComponent.SetMarBwUl(qosInfo.MarBwUl)
	medComponent.SetMedType(qosInfo.MediaType)

	/*subComponent containing target flows information*/
	medSubComponents := make(map[string]pcfclient.MediaSubComponent)
	medComponent.MedSubComps = &medSubComponents
	for _, flow := range data.FlowInfo {
		medSubComponent := pcfclient.MediaSubComponent{}
		medSubComponent.FNum = flow.FlowId
		medSubComponent.FlowUsage = pcfclient.PtrString("NO_INFO")

		medSubComponent.FDescs = append(medSubComponent.FDescs, flow.FlowDescriptions...)

		medSubComponents[string(flow.FlowId)] = medSubComponent
	}
	medComponents["1"] = medComponent

	log.Printf("%+v", *req.MedComponents)
	ctx := &pcfclient.AppSessionContext{}
	ctx.SetAscReqData(req)
	return ctx, nil

}

// ------------------------------------------------------------------------------
func validateSubscriptionData(data *models.AsSessionWithQoSSubscription) error {

	var err error
	err = AssertStringNotEmpty(data.UeIpv4Addr)
	if err != nil {
		return fmt.Errorf("field UeIpv4Addr not provided")
	}
	err = AssertStringNotEmpty(data.NotificationDestination)
	if err != nil {
		return fmt.Errorf("field NotificationDestination not provided")
	}
	err = AssertStringNotEmpty(data.Dnn)
	if err != nil {
		return fmt.Errorf("field Dnn not provided")
	}
	err = AssertStringNotEmpty(data.QosReference)
	if err != nil {
		return fmt.Errorf("field QosReference not provided")
	}

	if len(data.FlowInfo) == 0 {
		return fmt.Errorf("field FlowInfo not provided")
	}

	for _, flow := range data.FlowInfo {
		if len(flow.FlowDescriptions) == 0 {
			return fmt.Errorf("no FlowDescription provided")
		}
		for _, flowD := range flow.FlowDescriptions {
			err = AssertStringNotEmpty(flowD)
			if err != nil {
				return fmt.Errorf("invalid flow descriptor")
			}
		}
	}
	return nil
}

// ------------------------------------------------------------------------------
func AssertStringNotEmpty(s string) error {
	if len(s) == 0 {
		return fmt.Errorf("mandatory field not present")
	}
	return nil
}
