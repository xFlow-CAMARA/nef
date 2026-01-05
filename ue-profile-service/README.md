# UE Profile Service
The UE Profile Service is a middleware component of the Open-Exposure NEF stack. It provides a consolidated view of the current status and configuration of a User Equipment (UE) by querying the redis database and by creating a in-memory cache entry.

### Overview
This service implements a simple REST API for retrieving UE profile information based on the UE's SUPI (Subscriber Permanent Identifier). It is designed to facilitate internal NEF workflows and support northbound API responses.
It will implement different endpoints to:

- retrieve all the ues (/ue-profile/v1/profiles)
- retrieve a specific UE (/ue-profile/v1/profiles/{supi})

## Logic

1. when a UE is retrieved for the first time, it is stored and cached.
2. a subscription to the UE specific channel in redis is created and allows to keep the ue profile up-to-date
3. everytime that the ue profile is retrieved, the TTL is resetted. When the TTL expires, the UE profile is removed and the subscription deleted

## API

- **Endpoint:** `/ue-profile/v1/profiles/{supi}`
- **Method:** `GET`
- **Example:**

```bash
curl -X GET http://172.26.0.4:8080/ue-profile/v1/profiles/00106000000001
```

**Example Response:**

```json
{
  "Imsi": "00106000000001",
  "RegistrationStatus": "REGISTERED",
  "ConnectionStatus": "CONNECTED",
  "Plmn": {"mcc": "001", "mnc": "06"},
  "Tac": "1",
  "NrCellId": "1784",
  "PduSessions": {
    "1": {
      "Id": 1,
      "Ipv4": "12.1.0.2",
      "Snssai": {"Sst": 1, "Sd": "010203"},
      "DlStatus": "TRANSMITTED",
      "Dnn": "intenet"
    }
  }
}
```

## Dependencies

- **Redis**: Used for retrieving real-time UE context data
- **Core Network APIs**:
  - `namf-evts`: Access to UE registration and connectivity status
  - `nsmf-event-exposure`: Access to session management and PDU session details

Ensure that the core network supports these interfaces for this service to function correctly.

## Configuration


| Variable        | Description                     | Default         |
|----------------|---------------------------------|-----------------|
| `REDIS_SVC`   | Redis service url            | `redis:6379`     |
| `NBI_PORT`   | Northbound server port              | `8080`          |
| `HTTPS`| Use https for northbound server | `false` |
| `HTTP_VER` | Switch between HTTP1.1 and HTTP2         | `1`          |

## Running Locally

```bash
cd ue-profile-service
go run cmd/ue-profile-service/main.go
```

## License

GPLv3