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
	"net/http"

	pcfclient "gitlab.eurecom.fr/open-exposure/nef/pcfclient"
	"gitlab.eurecom.fr/open-exposure/nef/traffic-influence/internal/connector"
	contexts "gitlab.eurecom.fr/open-exposure/nef/traffic-influence/internal/context"
	models "gitlab.eurecom.fr/open-exposure/nef/traffic-influence/internal/models"
	"gitlab.eurecom.fr/open-exposure/nef/traffic-influence/pkg/config"
)

type app interface {
	Cfg() *config.AppConfig
	Connector() *connector.Connector
	Ctx() *contexts.TraffInflAppCtx
}

type Service struct {
	app
}

func NewTraffInflService(app app) *Service {
	svc := &Service{app: app}
	return svc
}

// ------------------------------------------------------------------------------
func (s *Service) CreateTrafficInfluenceSub(afId string, trafficInfluSub *models.TrafficInfluSub) (string, int, error) {
	var loc string

	af := s.Ctx().GetAf(afId)
	if af == nil {
		af = s.Ctx().AddAf(afId)
	}

	af.Mu.Lock()
	defer af.Mu.Unlock()

	if len(trafficInfluSub.Gpsi) > 0 || len(trafficInfluSub.Ipv4Addr) > 0 || len(trafficInfluSub.Ipv6Addr) > 0 {
		// Single UE, sent to PCF
		pa_ctx, err := s.trafficInfluence2PolicyAuthorization(trafficInfluSub)
		if err == nil {
			loc, err = s.Connector().CreatePolicyAuthzSubscription(*pa_ctx)
			if err != nil {
				return "", http.StatusInternalServerError, fmt.Errorf("could not create PCF Policy Authorization context")
			}
			tiLoc, subCtx := af.NewAfSubscription(trafficInfluSub)
			subCtx.AppSessId = loc
			return tiLoc, http.StatusCreated, nil
		} else {
			return "", http.StatusBadRequest, err
		}
	} else if len(trafficInfluSub.ExternalGroupId) > 0 || trafficInfluSub.AnyUeInd {
		return "", http.StatusNotImplemented, fmt.Errorf("udr policies not supported yet")
	} else {
		return "", http.StatusNotImplemented, fmt.Errorf("not in the single or multiple UE cases")
	}

}

// ------------------------------------------------------------------------------
func (s *Service) UpdateTrafficInfluenceSub(afId string, subId string, trafficInfluSub *models.TrafficInfluSub) (int, error) {

	af := s.Ctx().GetAf(afId)

	if af != nil {
		af.Mu.Lock()
		defer af.Mu.Unlock()

		sub := af.GetAfSubscription(subId)
		if sub != nil {
			if len(sub.AppSessId) > 0 {
				// here we update the pcf subscription
				if len(trafficInfluSub.Gpsi) > 0 || len(trafficInfluSub.Ipv4Addr) > 0 || len(trafficInfluSub.Ipv6Addr) > 0 {
					// Single UE, sent to PCF
					pa_ctx, err := s.trafficInfluence2PolicyAuthorization(trafficInfluSub)
					if err == nil {
						newPcfLoc, err := s.Connector().UpdatePolicyAuthzSubscription(sub.AppSessId, *pa_ctx)
						if err != nil {
							return http.StatusInternalServerError, fmt.Errorf("could not update PCF Policy Authorization context")
						}
						sub.AppSessId = newPcfLoc
						sub.Data = trafficInfluSub
						return http.StatusCreated, nil
					} else {
						return http.StatusBadRequest, err
					}
				}
			}
		}
	}
	return http.StatusNotFound, fmt.Errorf("could not find af/subId")

}

// ------------------------------------------------------------------------------
func (s *Service) DeleteTrafficInfluenceSub(afId string, subId string) error {

	af := s.Ctx().GetAf(afId)

	if af != nil {
		af.Mu.Lock()
		defer af.Mu.Unlock()

		sub := af.GetAfSubscription(subId)
		if sub != nil {
			if len(sub.AppSessId) > 0 {
				err := s.Connector().RemovePolicyAuthzSubscription(sub.AppSessId)
				if err == nil {
					err = af.DeleteAfscription(subId)
				}
				return err
			}
		}
	}
	return fmt.Errorf("could not find af/subId")
}

// ------------------------------------------------------------------------------
func (s *Service) GetIndividualTrafficInfluenceSub(afId string, subId string) (*models.TrafficInfluSub, int, error) {

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
func (s *Service) GetAllTrafficInfluenceSub(afId string) ([]*models.TrafficInfluSub, int, error) {

	af := s.Ctx().GetAf(afId)

	if af != nil {
		af.Mu.Lock()
		defer af.Mu.Unlock()

		subs := af.GetAfSubscriptions()
		var result []*models.TrafficInfluSub
		for _, sub := range subs {
			result = append(result, sub.Data)
		}
		return result, http.StatusOK, nil
	}
	return nil, http.StatusNotFound, fmt.Errorf("could not find afId")

}

// ------------------------------------------------------------------------------
func (s *Service) trafficInfluence2PolicyAuthorization(trafficInfluSub *models.TrafficInfluSub) (*pcfclient.AppSessionContext, error) {

	/*Convert NEF model to PCF models*/
	err := validateSubscriptionData(trafficInfluSub)
	if err != nil {
		return nil, err
	}

	req := pcfclient.AppSessionContextReqData{
		AfAppId: &trafficInfluSub.AfAppId,
		AfRoutReq: &pcfclient.AfRoutingRequirement{
			AppReloc: &trafficInfluSub.AppReloInd,
		},
		Supi:     pcfclient.PtrString(trafficInfluSub.Gpsi),
		UeIpv4:   pcfclient.PtrString(trafficInfluSub.Ipv4Addr),
		UeMac:    pcfclient.PtrString(trafficInfluSub.MacAddr),
		NotifUri: trafficInfluSub.NotificationDestination,
		SuppFeat: s.Cfg().SupportedFeat,
		Dnn:      pcfclient.PtrString(trafficInfluSub.Dnn),
		SliceInfo: &pcfclient.Snssai{
			Sst: trafficInfluSub.Snssai.Sst,
			Sd:  pcfclient.PtrString(trafficInfluSub.Snssai.Sd),
		},
	}

	/* RouteToLocs */
	for _, routeToLoc := range trafficInfluSub.TrafficRoutes {
		rToL := pcfclient.RouteToLocation{}
		rToL.SetDnai(routeToLoc.Dnai)
		if routeToLoc.RouteProfId != nil {
			rToL.RouteProfId = *pcfclient.NewNullableString(routeToLoc.RouteProfId)
		}
		if routeToLoc.RouteInfo != nil {
			routeInfo := pcfclient.RouteInformation{
				Ipv4Addr:   &routeToLoc.RouteInfo.Ipv4Addr,
				PortNumber: routeToLoc.RouteInfo.PortNumber,
			}
			rToL.SetRouteInfo(routeInfo)
		}
		req.AfRoutReq.RouteToLocs = append(req.AfRoutReq.RouteToLocs, rToL)
	}

	for _, flow := range trafficInfluSub.TrafficFilters {
		tfcInfo := &pcfclient.TrafficCorrelationInfo{}
		tfcInfo.TfcCorrId = pcfclient.PtrString(flow.FlowDescriptions[0])
		req.AfRoutReq.TfcCorreInfo.Set(tfcInfo)
		break
	}

	ctx := &pcfclient.AppSessionContext{}
	ctx.SetAscReqData(req)
	return ctx, nil
}

func validateSubscriptionData(data *models.TrafficInfluSub) error {

	var err error
	err = AssertStringNotEmpty(data.Ipv4Addr)
	if err != nil {
		return fmt.Errorf("field Ipv4Addr not provided")
	}
	err = AssertStringNotEmpty(data.NotificationDestination)
	if err != nil {
		return fmt.Errorf("field NotificationDestination not provided")
	}
	err = AssertStringNotEmpty(data.Dnn)
	if err != nil {
		return fmt.Errorf("field Dnn not provided")
	}

	if len(data.TrafficFilters) == 0 {
		return fmt.Errorf("field TrafficFilters not provided")
	}

	for _, flow := range data.TrafficFilters {
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

	if len(data.TrafficRoutes) == 0 {
		return fmt.Errorf("field TrafficRoutes not provided")
	}

	for _, route := range data.TrafficRoutes {
		err = AssertStringNotEmpty(route.Dnai)
		if err != nil {
			return fmt.Errorf("invalid route info")
		}
	}
	return nil
}

func AssertStringNotEmpty(s string) error {
	if len(s) == 0 {
		return fmt.Errorf("mandatory field not present")
	}
	return nil
}
