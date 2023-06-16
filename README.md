# Webhook For SWARM data
### Space-based LoRa Communication of Vessel Tracking Data

---

Webhook for the real-time event-driven delivery of location data to database. The webhook pushes space-based IoT location data to database.   
The repo contains the code of the webhook and all supporting documentation, including the Dockerfile for its deployment.

### Database  
+ MySQL 8.
+ Credentials as Github Secrets:
  + MYSQL_DO_CERTIFICATE
  + MYSQL_DO_DATABASE
  + MYSQL_DO_HOST
  + MYSQL_DO_PORT
  + MYSQL_DO_PW
  + MYSQL_DO_SSLMODE
  + MYSQL_DO_USERNAME
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

### Base URL
loraspaceway.aircentre.io

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


## Usage

Send a `POST` request to the following endpoint to send device data to the webhook:
```
POST https://loraspaceway.aircentre.io/uplink
```

### Parameters:

- **access_key (mandatory)**: String that grants access to the webhook.

### Headers:

- **Content-Type**: `application/json`

### Body:

- The body of the POST request should be a JSON object containing the device data.
- JSON object format:
  ```json
   {
    "device": "",
    "packet_id": 0,
    "timestamp": "RFC 822/1123 date-time format",
    "rx_time": "RFC 822/1123 date-time format",
    "altitude": 0.0,
    "heading": 0,
    "latitude_deg": 0.0,
    "longitude_deg": 0.0,
    "gps_jamming": 0,
    "gps_spoofing": 0,
    "temperature_c": 0.0,
    "battery_v": 0,
    "speed": 0.0,
    "telemetry_snr_db": 0.0,
    "telemetry_rssi_dbm": 0,
    "telemetry_time": 0,
    "rssi_background": 0,
    "telemetry_type": "",
    "version": 0
  }
  ```


### Example Request:
+ Method: `POST`
+ URL: `https://loraspaceway.aircentre.io/uplink?access_key=your_access_key_here`
+ Headers: `Content-Type: application/json`
+ Body:
  ```json
  {
    "device": "F-0x06eb2",
    "packet_id": 52053866,
    "timestamp": "Thu Mar 23 2023 01:00:06 GMT+0000 (Western European Standard Time)",
    "rx_time": "Thu Mar 23 2023 16:30:52 GMT+0000 (Western European Standard Time)",
    "altitude": 438,
    "heading": 338,
    "latitude_deg": 40.2516,
    "longitude_deg": -7.4872,
    "gps_jamming": 84,
    "gps_spoofing": 1,
    "temperature_c": 19,
    "battery_v": 4021,
    "speed": 0,
    "telemetry_snr_db": -9,
    "telemetry_rssi_dbm": -114,
    "telemetry_time": 1679526068,
    "rssi_background": -104,
    "telemetry_type": "ASSET_TRACKER",
    "version": 1
  }
  ```



## Development Setup

#### Prerequisites
- [Go 1.20](https://golang.org/dl/)
- [Git](https://git-scm.com/downloads)
- [Docker Desktop](https://www.docker.com/products/docker-desktop)

#### Dependencies
Run `go mod download` to install dependencies

#### Testing
- Run `go test ./internal/...` to execute unit tests
- Run `sh ./e2e/run_tests.sh` to execute end-to-end tests


## CI/CD
This project leverages GitHub Actions for a CI/CD process. Each commit triggers a pipeline which builds, tests, and pushes a new Docker image called `webhook-spaceway-lora` to [GitHub's container Registry (ghcr.io)](https://github.com/AIRCentre/webhook-spaceway-lora/pkgs/container/webhook-spaceway-lora). Ansible, within the GitHub Actions workflow, deploys this updated image automatically.

To skip the workflow, add `no-ci` in the commit message - helpful when updating non-impactful files like the readme.
