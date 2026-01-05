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

package resolver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"

	"gitlab.eurecom.fr/open-exposure/nef/ue-identity-service/internal/models"
	"gitlab.eurecom.fr/open-exposure/nef/ue-identity-service/internal/redis"
)

var (
	ErrNotFound = errors.New("IP address not found")
)

type Resolver struct {
	redis     *redis.RedisAdaptor
	ctx       context.Context
	ipCache   map[string]string // Map IP → SUPI
	gpsiCache map[string]string // Map SUPI → GPSI (MSISDN)
	lock      sync.RWMutex
}

func NewResolver(rdb *redis.RedisAdaptor) *Resolver {
	r := &Resolver{
		redis:     rdb,
		ctx:       context.Background(),
		ipCache:   make(map[string]string),
		gpsiCache: make(map[string]string),
	}

	log.Printf("bootstrapping information from redis")
	// parse current ue profiles to build the mapping
	if err := r.bootstrapFromRedis(); err != nil {
		log.Fatalf("could not bootstrap from redis: %s", err)
	}

	log.Printf("bootstrap completed, found %d IP mappings", len(r.ipCache))
	for ip, supi := range r.ipCache {
		log.Printf("  IP %s → SUPI %s", ip, supi)
	}
	log.Printf("initialization completed, listening for updates")
	// Start listening for updates
	r.listenForUpdates()

	return r
}

// Lookup returns the ExternalId for a given IP address
func (r *Resolver) Lookup(ip string) (string, error) {
	r.lock.RLock()
	supi, ok := r.ipCache[ip]
	r.lock.RUnlock()

	if ok {
		return supi, nil
	}

	return "", fmt.Errorf("could not find the requested address")
}

// LookupMsisdn returns the MSISDN (GPSI) for a given IP address
func (r *Resolver) LookupMsisdn(ip string) (string, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	supi, ok := r.ipCache[ip]
	if !ok {
		return "", fmt.Errorf("could not find UE for IP address %s", ip)
	}

	gpsi, ok := r.gpsiCache[supi]
	if ok {
		return gpsi, nil
	}

	// Fallback: Generate MSISDN from SUPI for demo purposes
	// SUPI format: 001010000000001 → MSISDN: +33600000001 (French mobile format)
	// This ensures E.164 compliance: +[1-9][0-9]{4,14}
	if len(supi) >= 5 {
		// Use last 8 digits of SUPI and prefix with +336 (French mobile)
		lastDigits := supi[len(supi)-8:]
		msisdn := "+336" + lastDigits
		return msisdn, nil
	}

	return "", fmt.Errorf("MSISDN not available for SUPI %s", supi)
}

// SetGpsi associates a GPSI (MSISDN) with a SUPI
func (r *Resolver) SetGpsi(supi, gpsi string) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.gpsiCache[supi] = gpsi
}

// listenForUpdates subscribes to Redis events and keeps the cache updated
func (r *Resolver) listenForUpdates() {
	sub := r.redis.Instance.PSubscribe(r.ctx, "broadcast:PDN_CONNECTIVITY_STATUS")

	go func() {

		defer func() {
			_ = sub.Close()
		}()

		ch := sub.Channel()

		for {
			select {
			case msg := <-ch:

				updateData := &models.UeInfoPatch{}
				err := json.Unmarshal([]byte(msg.Payload), updateData)
				if err != nil {
					continue
				}

				jsonData, err := json.Marshal(updateData.Data)
				if err != nil {
					continue
				}

				r.lock.Lock()
				switch updateData.Type {
				case "PDU_SES_EST":
					pduSesEst := &models.PduSesEst{}
					err := json.Unmarshal(jsonData, pduSesEst)
					if err != nil {
						continue
					}
					r.ipCache[*pduSesEst.AdIpv4Addr] = updateData.Imsi

				case "PDU_SES_REL":
					pduSesRel := &models.PduSesRel{}
					err := json.Unmarshal(jsonData, pduSesRel)
					if err != nil {
						continue
					}
					delete(r.ipCache, *pduSesRel.Ipv4Addr)

				case "UE_IP_CH":
					ueIpCh := &models.UeIpCh{}
					err := json.Unmarshal(jsonData, ueIpCh)
					if err != nil {
						continue
					}

					r.ipCache[*ueIpCh.AdIpv4Addr] = updateData.Imsi

				}
				r.lock.Unlock()

			case <-r.ctx.Done():
				return
			}
		}
	}()

}

func (r *Resolver) bootstrapFromRedis() error {
	ueList, err := r.redis.QueryUEsInfo()
	if err != nil {
		return fmt.Errorf("could not bootstrap: %s", err)
	}

	log.Printf("Processing %d UEs from Redis", len(ueList))
	for _, ue := range ueList {
		log.Printf("Processing UE: IMSI=%s, PDU sessions=%d", *ue.Imsi, len(ue.PduSessEst))

		// Cache GPSI (MSISDN) if available
		if ue.Gpsi != nil && *ue.Gpsi != "" {
			r.lock.Lock()
			r.gpsiCache[*ue.Imsi] = *ue.Gpsi
			r.lock.Unlock()
			log.Printf("  GPSI (MSISDN): %s", *ue.Gpsi)
		}

		for pduId, pduSess := range ue.PduSessEst {
			log.Printf("  PDU Session %s: IP=%s", pduId, *pduSess.AdIpv4Addr)

			// Check if IP has changed (newer than the establishment event)
			if ue.UeIpCh[pduId] != nil && ue.UeIpCh[pduId].TimeStamp >= pduSess.TimeStamp {
				ip := ue.UeIpCh[pduId].AdIpv4Addr
				log.Printf("    IP changed, updating from %s to %s", *pduSess.AdIpv4Addr, *ip)
				r.lock.Lock()
				r.ipCache[*ip] = *ue.Imsi
				delete(r.ipCache, *pduSess.AdIpv4Addr)
				r.lock.Unlock()
				continue
			}

			// Use the established IP even if session was released
			// This allows lookup of recently released sessions for testing/debugging
			log.Printf("    Adding mapping: %s → %s", *pduSess.AdIpv4Addr, *ue.Imsi)
			r.lock.Lock()
			r.ipCache[*pduSess.AdIpv4Addr] = *ue.Imsi
			r.lock.Unlock()
		}
	}

	return nil
}
