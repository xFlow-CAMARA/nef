# Core Network Service

## Overview

The **Core Network Service** acts as an interface between the 5G Core Network and the internal data layer.  
It subscribes to the relevant 3GPP APIs exposed by the **Access and Mobility Management Function (AMF)** and the **Session Management Function (SMF)**, collects notifications, processes them, and stores the results into the database.  

All processed information is persisted in a normalized JSON format as defined in [`models/ue_info.go`](https://gitlab.eurecom.fr/open-exposure/nef/ue-profile-service/-/blob/develop/internal/models/ue_info.go?ref_type=heads).

## Responsibilities

- Subscribe to AMF and SMF notification endpoints.
- Process and normalize incoming 3GPP event data.
- Persist UE information into the Redis database using the defined JSON schema.
- Serve as the data provider for other northbound services (e.g., Monitoring Event, UE Identifier, UE Address).

## Adaptability

The service can be adapted to integrate with **custom core network solutions** that may not expose all the required 3GPP interfaces.  
As long as the service maintains the JSON format defined in `models/ue_info.go` when storing data in Redis, downstream services will remain fully compatible.

## Dependencies

- **AMF and SMF APIs** (3GPP TS 29-series interfaces)  
- **Redis** (primary datastore for normalized UE information)  


## Configuration


| Variable        | Description                     | Default         |
|----------------|---------------------------------|-----------------|
| `REDIS_ADDR`   | Redis service url            | `redis:6379`     |
| `USE_NRF`   | Discover AMF and SMF via NRF            | `False`          |
| `CN_HTTP_VERSION`| Switch between HTTP1.1 and HTTP2 | `1` |
| `AMF_IP_ADDR` | AMF SBI url, used when USE_NRF=false         | `""` |
| `SMF_IP_ADDR`| SMF SBI url, used when USE_NRF=false| `""` |
| `EVENT_NOTIFY_URI` | The callback url to receive event notifications         | `http://core-network-service:9090`       |
| `SERVER_ADDR`| The callback server configuration | `core-network-service:9090` |



## Data Model

The JSON schema used for persisting UE information is defined in [`models/ue_info.go`](https://gitlab.eurecom.fr/open-exposure/nef/ue-profile-service/-/blob/develop/internal/models/ue_info.go?ref_type=heads). 
This schema **must remain consistent** across all integrations to ensure interoperability with other services in the system.