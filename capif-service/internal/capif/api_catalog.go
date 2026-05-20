package capif

import (
	"log"

	serviceApi "gitlab.eurecom.fr/open-exposure/nef/capif-service/internal/3gpp/publish_api_service"
)

// camaraAPIs maps Kong route path prefixes to CAMARA API names and versions.
var camaraAPIs = []struct {
	PathPrefix string
	ApiName    string
	Version    string
}{
	{"/quality-on-demand", "quality-on-demand", "v1"},
	{"/location-retrieval", "location-retrieval", "v0"},
	{"/traffic-influence", "traffic-influence", "vwip"},
	{"/device-status", "device-status", "v0"},
	{"/device-reachability-status", "device-reachability-status", "v1"},
	{"/number-verification", "number-verification", "vwip"},
	{"/sim-swap", "sim-swap", "v0"},
}

// ServiceByPath returns the service profile whose name matches the given
// request path prefix. Returns nil if no match.
func ServiceByPath(apiPath string) *ServiceApiProfile {
	for _, def := range camaraAPIs {
		if len(apiPath) >= len(def.PathPrefix) && apiPath[:len(def.PathPrefix)] == def.PathPrefix {
			for _, svc := range providerInstance.Services {
				if svc.ServiceName == def.ApiName {
					return svc
				}
			}
		}
	}
	return nil
}

// CatalogEntry is a single entry in the published API catalog.
type CatalogEntry struct {
	ApiName   string `json:"api_name"`
	ServiceId string `json:"service_id"`
	AefId     string `json:"aef_id"`
	Scope     string `json:"scope"`
}

// GetPublishedCatalog returns the list of currently published CAMARA APIs with their AEF IDs and scopes.
func GetPublishedCatalog() []CatalogEntry {
	aefId := providerInstance.FuncId[AEF]
	entries := make([]CatalogEntry, 0, len(providerInstance.Services))
	for sid, svc := range providerInstance.Services {
		entries = append(entries, CatalogEntry{
			ApiName:   svc.ServiceName,
			ServiceId: sid,
			AefId:     aefId,
			Scope:     "3gpp#" + aefId + ":" + svc.ServiceName,
		})
	}
	return entries
}

// PublishDefaultAPIs publishes all CAMARA APIs to the CAPIF Core at startup.
func PublishDefaultAPIs() {
	token, err := retrieveAccessToken()
	if err != nil {
		log.Printf("PublishDefaultAPIs: cannot retrieve access token: %v", err)
		return
	}

	aefId := providerInstance.FuncId[AEF]

	for _, def := range camaraAPIs {
		// Skip if already published in this session
		alreadyPublished := false
		for _, svc := range providerInstance.Services {
			if svc.ServiceName == def.ApiName {
				alreadyPublished = true
				break
			}
		}
		if alreadyPublished {
			log.Printf("PublishDefaultAPIs: %s already published, skipping", def.ApiName)
			continue
		}

		resource := serviceApi.NewResource(def.ApiName, "REQUEST_RESPONSE", def.PathPrefix)

		domain := "camara.local"

		supportedFeatures := "0"
		description := serviceApi.NewServiceAPIDescription(def.ApiName)
		description.SupportedFeatures = &supportedFeatures
		description.AefProfiles = []serviceApi.AefProfile{
			{
				AefId:           aefId,
				Protocol:        serviceApi.HTTP_2,
				DomainName:      &domain,
				SecurityMethods: []string{"PSK"},
				Versions: []serviceApi.Version{
					{
						ApiVersion: def.Version,
						Resources:  []serviceApi.Resource{*resource},
					},
				},
			},
		}

		profile := NewServiceApiProfile(def.ApiName)
		profile.ServiceModel = description

		sid, err := PublishService(providerInstance, profile, token)
		if err != nil {
			log.Printf("PublishDefaultAPIs: failed to publish %s: %v", def.ApiName, err)
			continue
		}
		providerInstance.Services[sid] = profile
		log.Printf("PublishDefaultAPIs: published %s (id=%s)", def.ApiName, sid)
	}
}
