## NEF Integration Guide

This document describes how to integrate the NEF stack with either a real 5G Core Network or with the included open-exposure CoreSim. It explains which 3GPP interfaces and capabilities are required by the northbound services (Monitoring Event, UE Address, UE ID, Traffic Influence, AS Session with QoS) and how the `core-network-service` adapts AMF/SMF notifications for the rest of the system.

### Summary (short)
- You can integrate using a real core (OAI, Open5GS, free5GC, etc.) or using the open-exposure CoreSim included in this repository.
- Monitoring Event, UE Address and UE ID require AMF and SMF event exposure (3GPP notification/subscription APIs).
- `core-network-service` acts as an adapter that subscribes to AMF/SMF events, normalizes them and stores the resulting UE state in Redis for northbound services.
- Traffic Influence and AS Session with QoS require a Core with a PCF implementing the relevant policy control features and the NPCF Policy Authorization API.
- The system has been tested with: OAI (monitoring-event, ue-id, ue-address), Open5GS (as-session-with-qos) and free5GC (traffic influence).

### Integration options

1) Real 5G Core Network

- Use a compliant core that exposes the AMF and SMF event-notification APIs. Configure the core to send notifications (or accept subscriptions) to the `core-network-service` callback endpoint.
- If the core supports NRF discovery you can enable service discovery (see `USE_NRF` below). Otherwise provide explicit `AMF_IP_ADDR` / `SMF_IP_ADDR` values.

2) open-exposure CoreSim

- The open-exposure CoreSim can be used when no real core is available. It emulates AMF/SMF event exposure, PCF APIs and is useful for development, testing and CI.
- When using the open-exposure CoreSim, point the `core-network-service` at the simulator endpoints or use NRF discovery depending on your composition.

### Required interfaces and services

- AMF event exposure: required. The monitoring-event, ue-address and ue-id services rely on AMF notifications (e.g., UE registration, mobility, reachability).
- SMF event exposure: required. SMF notifications (e.g., PDU session events) are used to enrich UE state for the above services.
- Core Network Service: this service subscribes to AMF/SMF notifications, normalizes and persists UE information into Redis (JSON schema defined in `internal/models/ue_info.go`). All northbound services read from Redis via this schema.

Note: The `core-network-service` is designed to be compatible with any Core Network that supports the standard AMF and SMF event exposure interfaces.

### Adapting non-standard interfaces

If your core exposes non-standard or vendor-specific interfaces, you can adapt the integration by implementing a thin adapter in the `core-network-service` component. See the `core-network-service` README for developer guidance and examples on how to add custom adapters and to map non-standard events into the normalized UE JSON model.

Files to check:

- `core-network-service/README.md` — integration and adapter guidance
- `core-network-service/internal/` — implementation entry points for subscriptions and event processing

### Traffic Influence and AS Session with QoS

- These features depend on policy control capabilities in the core. Specifically, a PCF that supports the relevant policy features and the NPCF Policy Authorization API is required.
- AS Session with QoS: tested with Open5GS. Requires the core to support session-level QoS control and to expose the necessary SMF/PCF interactions.
- Traffic Influence: tested with free5GC. Requires PCF support and the NPCF policy authorization API.

### Tested cores (quick matrix)

The table below shows which northbound APIs were validated with each core during integration testing. An "X" indicates the feature was tested and confirmed with the listed core.

| Core / Northbound API | Monitoring Event | UE Address | UE ID | Traffic Influence | AS Session with QoS |
|---|:---:|:---:|:---:|:---:|:---:|
| OAI | X | X | X |  |  |
| Open5GS |  |  |  |  | X |
| free5GC |  |  |  | X |  |

These are the combinations used during integration testing. Other cores that implement the same 3GPP interfaces should also work when configured correctly.

### Configuration pointers

Key environment variables used by `core-network-service` (defaults shown where available):

- `REDIS_ADDR` — Redis service address (default: `redis:6379`)
- `USE_NRF` — Discover AMF and SMF via NRF (default: `False`)
- `CN_HTTP_VERSION` — HTTP version to use when contacting Core Network APIs (`1` or `2`, default: `1`)
- `AMF_IP_ADDR` — AMF SBI URL used when `USE_NRF=false` (default: `""`)
- `SMF_IP_ADDR` — SMF SBI URL used when `USE_NRF=false` (default: `""`)
- `EVENT_NOTIFY_URI` — Callback URL that AMF/SMF will use to send event notifications (default: `http://core-network-service:9090`)
- `SERVER_ADDR` — Address where `core-network-service` listens (default: `core-network-service:9090`)

Typical integration steps

1. Ensure Redis is running and reachable using `REDIS_ADDR`.
2. Deploy `core-network-service` and configure its environment variables according to your environment (NRF vs static AMF/SMF addresses).
3. Configure your Core (or open-exposure CoreSim) to send AMF and SMF events to `EVENT_NOTIFY_URI` or allow `core-network-service` to subscribe to the core's event APIs.
4. Start `core-network-service`. It will subscribe to / accept notifications, normalize events and persist them in Redis using the agreed JSON model.
5. Start northbound services (Monitoring Event, UE ID, UE Address, Traffic Influence, AS Session with QoS). They will read normalized UE data from Redis.


### Quick checklist

- [ ] Redis reachable and configured (`REDIS_ADDR`)
- [ ] `core-network-service` deployed and configured (`USE_NRF` / `AMF_IP_ADDR` / `SMF_IP_ADDR` / `EVENT_NOTIFY_URI`)
- [ ] Core (real or simulator) configured to expose AMF/SMF event APIs and to call/accept subscriptions from `core-network-service`
- [ ] PCF in the core supports NPCF policy authorization API when using Traffic Influence or AS Session with QoS
- [ ] Northbound services pointed to the running `core-network-service` / Redis

### Troubleshooting tips

- If Monitoring Event / UE Address / UE ID show missing data: verify AMF/SMF notifications are being sent and that `core-network-service` logs show successful subscription/notifications. Check `EVENT_NOTIFY_URI` and network connectivity.
- If Redis is empty: confirm the `core-network-service` has processed events and that no errors appear on startup (missing schema, DB permissions, etc.).
- When integrating a non-standard core: implement an adapter in `core-network-service` that translates vendor events into the normalized UE JSON model and follow the `core-network-service` README examples.

### Where to read more

- `core-network-service/README.md` — How to run and extend the adapter, plus developer notes for custom event mappings.
- `internal/models/ue_info.go` — The normalized UE JSON schema used by all northbound services.

---

