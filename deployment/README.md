# open-exposure NEF — Quickstart & Deployment Guide
This document explains how to deploy the open-exposure NEF provided in this repository. It focuses on the two common ways to run the NEF:

- Standalone NEF deployment (useful for integrating with an existing core network)
- NEF + Core Simulator deployment (fully self-contained environment for testing and development)

Both deployments are provided as Docker Compose manifests under the `artifacts/` folder.

## Supported Features
- CAPIF Integration
- Monitoring Events (3GPP)
- Traffic Influence (3GPP)
- Application Sessions with QoS (3GPP)
- Ue Identifier (3GPP)
- Ue Address (3GPP)
- Dataset Exporter Service
- Ue Profile Service
- PCF Discovery via NRF
- AMF Discovery via NRF
- SMF Discovery via NRF
- Docker Compose Manifests

## Supported APIs
- 3GPP Traffic Influence (rel 18.4)
- 3GPP Monitoring Events (rel 18.4)
- 3GPP AsSessionWithQoS (rel 18.4)
- 3GPP Ue Identifier (rel 18.4)
- 3GPP Ue Address (rel 18.4)

## What you'll find in `artifacts/`

This folder includes ready-to-run compose manifests and example configuration files:

- `nef-compose/` — compose manifest and `config/` for running NEF only
- `nef-coresim-compose/` — compose manifest and `config/` for running NEF together with a Core Simulator, Prometheus and Grafana
- Example configuration YAMLs for the supported 3GPP services in `config/`. 

## Building the images
The latest images are available in our dockerhub repository under the `openexposure/<service-name>:<tag>` convention. You will find the latest release and develop images thanks to our CI/CD pipeline.

If you want to build locally the different services, all you have to do is:

``` bash
git clone https://gitlab.eurecom.fr/open-exposure/nef/<service-name>
cd <service-name>

docker build -t openexposure/<service-name>:<tag> -f docker/Dockerfile .
```
The procedure is identical for all the services. You can choose your custom tag, but please remind to update it in your docker compose file when deploying.

## Deploying NEF (standalone)

Use this mode when you want to run only the NEF and connect it to your existing core network components.

1. Open a terminal and change into the artifacts NEF compose folder:

```bash
cd artifacts/nef-compose
```

2. Inspect `docker-compose.yaml` and the `config/` directory to review service settings and ports. The `config/` files include example YAMLs for the services the NEF exposes. Tune the configuration to adapt to your setup. You can find examples on how to configure the single components in their respective repos.

3. Start the stack in the background:

```bash
docker compose up -d
```

4. Verify containers are running and healthy:

```bash
docker compose ps
```

5. Follow logs for the different microservices:

```bash
docker compose logs -f <service_name>
```

6. Stop and remove the stack when finished:

```bash
docker compose down
```

## Deploying NEF + Core Simulator (integrated)

This mode runs the NEF together wit the open-exposure core-simulator so you can test 3GPP flows end-to-end.

1. Change to the combined compose folder:

```bash
cd artifacts/nef-coresim-compose
```

2. Review `docker-compose.yaml` and the `config/` directory. The `grafana/` and `prometheus/` subfolders contains the dashboard to observe the core simualator metrics.

3. Start the full stack:

```bash
docker compose up -d
```

4. Check running services:

```bash
docker compose ps
```

5. Useful logs:

```bash
docker compose logs -f <service_name>
```

6. Once started, please follow the core-simulator documentation in order to start the simulator. 

7. When finished, tear down the environment:

```bash
docker compose down
```

## Quick checks and a basic smoke test

- Container health: `docker compose ps` and the `STATUS`/`HEALTH` columns
- Logs: `docker compose logs -f <service>`
- Simple HTTP check (replace HOST:PORT and path according to your compose ports):


## Troubleshooting and tips

- Port conflicts: If a compose service fails to start due to a port being in use, change the published port in the corresponding `docker-compose.yaml`.
- Logs: `docker compose logs --tail 200 <service>` is useful for debugging start failures.
- Resource limits: If containers are OOM-killed, increase available RAM or reduce services launched.
- Compose file differences: If your Docker Compose CLI is v1, use `docker-compose` instead of `docker compose`.

## File map and purpose

- `artifacts/nef-compose/` — NEF-only compose manifest and example configs
- `artifacts/nef-coresim-compose/` — NEF + Core Simulator compose manifest; contains `grafana/` and `prometheus/` example configs
- `config/*.yaml` — example service configuration files (asSessionWithQos.yaml, monitoringEvent.yaml, trafficInfluence.yaml, ueAddress.yaml, ueId.yaml, coreSimulator.yaml)

## Next steps

- Customize `config/` YAMLs to match your environment (IP addresses, upstream core components, certificates)
- Add secrets and TLS configuration before deploying in a production environment
- Import Grafana dashboards from `grafana/dashboards` for quick observability
