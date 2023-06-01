# Webhook For Space-based LoRa Communication of Vessel Tracking Data
_____

Webhook for the real-time event-driven delivery of location data to database. The webhook pushes space-based IoT location data to database.   
The repo contains the code of the webhook and all supporting documentation, including the Dockerfile for its deployment.

### Database  
+ MySQL 8.
+ Credentials as Github Secrets.
+ Naming tables:
  + English
  + Lowercase
+ Naming columns:
  + English
  + Lowecase
  + Underscore rather than spaces
  + Include the unit

Exemples:
+ station_id           (key)
+ battery_v
+ rssi_dbm
+ snr_db
+ latitude_deg         (WGS84)
+ longitude_deg        (WGS84)
+ gw_eui
+ timestamp_utc_iso_string 
+ wind_speed_kmh
+ temperature_c
+ radiation_kjm2
+ rel_humidity_pctg
+ wind_direction_bin

### Contents  
1. Documentation  
 1.1 Technical description  
 1.1 Description of resources (inlcuding the harware)  
 1.1 Description of the webhook solution  
2. Code  
2.1 Anotated code  
2.1 Dockerfile  
2.1 Notes for the deployment of the Dockerfile  
3. Webhook monitoring and control services  


## Development Setup

#### Prerequisites
- Go 1.20 - available [here](https://golang.org/dl/).
- Git - available [here](https://git-scm.com/downloads).
- Docker - available [here](https://www.docker.com/products/docker-desktop).

#### Dependencies
Run `go mod download` to install dependencies

#### Testing
- Run `go test ./modules/...` to execute unit tests
- Run `sh ./e2e/run_tests.sh` to execute end-to-end tests


## CI Workflow
This project uses GitHub Actions to enable continous integration (CI).

On every commit, the pipeline builds, tests and pushes a new version of a Docker image called `webhook-spaceway-lora` to GitHub's container Registry (ghcr.io). It becomes avaliable in [AIR Centre's GitHub Packages page](https://github.com/orgs/AIRCentre/packages). 

Include `no-ci` in the commit message to prevent the workflow from running. This is useful when updating just the readme or other files that dont affect system behavior.
