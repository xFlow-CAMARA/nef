# Capif Service 

## Overview

Implements the **3GPP CAPIF API Provider Function** interacting with the CAPIF Core Function (CCF) [3GPP TS 29.222](https://www.etsi.org/deliver/etsi_ts/129200_129299/129222/16.11.00_60/ts_129222v161100p.pdf). This component allows to interact with the CAPIF Core Function in order to publish open-exposure NEF APIs to CAPIF.

## API Specification

Reference Technical Specification:
[3GPP TS 29.222 Common API Framework for 3GPP Northbound APIs](https://www.etsi.org/deliver/etsi_ts/129200_129299/129222/16.11.00_60/ts_129222v161100p.pdf)

## Functionalities
It acts as a middleware between the different open-exposure microservices and the Capif Core Function. The communication between the microservice and the open-exposure capif-service is done via the libcapif library. 
The communication between the capif service and the CCF is done accordingly to [3GPP TS 29.222](https://www.etsi.org/deliver/etsi_ts/129200_129299/129222/16.11.00_60/ts_129222v161100p.pdf)
## Configuration

| Variable        | Description                | Default     |
|----------------|----------------------------|-------------|
| `CAPIF_FQDN`   | Capif Core Function FQDN      | `capif.etsi.org`     |
| `CAPIF_AUTH_PORT`   | CCF port reserved for authenitaction              | `8080`          |
| `CAPIF_USERNAME`   | Capif Username of the NEF instance              | ``          |
| `CAPIF_PASSWORD`   | Capif Password of the NEF instance             | ``          |
| `INGRESS_FQDN`   | Specifies the fqdn of the API Gateway of the NEF              | ``          |


## Running Locally

```bash
go run cmd/main.go
```
