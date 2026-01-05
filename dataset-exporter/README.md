# Dataset Exporter Service

The Dataset Exporter Service is part of the Open-Exposure NEF stack. It enables exporting event exposure data into a downloadable CSV format for analysis or archival purposes. The service listens to a Redis broadcast channel and records event messages during a defined capture window.

## Overview

This service allows external consumers to initiate and retrieve anonymized event datasets. It masks user identities to ensure privacy compliance and supports time-bounded data capture.

## Endpoints

### 1. Start Capture

- **Endpoint:** `/dataset-exporter/v1/captures`
- **Method:** `POST`
- **Request Body:**

```json
{
  "duration": "60s"
}
```

- **Response:**

```json
{
  "status": "IN_PROGRESS",
  "remainingTime": 60,
  "downloadPath":  "/dataset-exporter/v1/captures/download?id={sessionId}",
}
```

- **Description:** Starts a capture session for the specified duration. The service subscribes to the Redis broadcast channel and logs events into a CSV file.

---

### 2. Capture Status

- **Endpoint:** `/dataset-exporter/v1/captures/status?id={sessionId}`
- **Method:** `GET`
- **Response:**

```json
{
  "status": "IN_PROGRESS",
  "remainingTime": 25,
  "downloadPath":  "/dataset-exporter/v1/captures/download?id={sessionId}",
}
```

- **Description:** Retrieves the current status of the capture session (either `IN_PROGRESS` or `COMPLETED`), the remaining time in seconds and the URL to download the CSV file.

---

### 3. Download CSV

- **Endpoint:** `/dataset-exporter/v1/captures/download?id={sessionId}`
- **Method:** `GET`
- **Description:** Returns the anonymized event dataset as a CSV file. Only available when the capture session has completed.

## Privacy

All user identities are masked in the exported dataset to prevent the exposure of sensitive subscriber information.

## Dependencies

- **Redis**: Required to subscribe to the event broadcast channel and retrieve events in real time.

## Configuration

| Variable        | Description                | Default     |
|----------------|----------------------------|-------------|
| `REDIS_SVC`   | Redis service url            | `redis:6379`     |
| `NBI_PORT`   | Northbound server port              | `8080`          |
| `HTTPS`| Use https for northbound server | `false` |
| `HTTP_VER` | Switch between HTTP1.1 and HTTP2         | `1`          |

## Running Locally

```bash
go run cmd/dataset-exporter/main.go
```

## License

GPLv2