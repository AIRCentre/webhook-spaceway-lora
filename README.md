# Webhook For SWARM Data
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

## Examles of actual device payloads
```json
// payload 1
{"packetId":63464221,"deviceType":1,"deviceId":28338,"userApplicationId":65002,"organizationId":66429,"data":"eyJkdCI6MTY4NzE5OTQ1MCwibHQiOjM4LjY1NTIsImxuIjotMjcuMjE5MywiYWwiOjMzLCJzcCI6OSwiaGQiOjY1LCJnaiI6ODksImdzIjoxLCJidiI6MzkyMCwidHAiOjI2LCJycyI6LTgyLCJ0ciI6LTExNSwidHMiOi0yLCJ0ZCI6MTY4NzE5NjUwOSwiaHAiOjEyNjYsInZwIjoxMDAsInRmIjoyMDU0NDB9","len":174,"status":0,"hiveRxTime":"2023-06-20T11:33:56"}

// payload 1 data after base64 decode
{"dt":1687199450,"lt":38.6552,"ln":-27.2193,"al":33,"sp":9,"hd":65,"gj":89,"gs":1,"bv":3920,"tp":26,"rs":-82,"tr":-115,"ts":-2,"td":1687196509,"hp":1266,"vp":100,"tf":205440}

// payload 2
{"packetId":63476528,"deviceType":1,"deviceId":28338,"userApplicationId":65002,"organizationId":66429,"data":"eyJkdCI6MTY4NzI2NzY4MiwibHQiOjM4LjY3NDUsImxuIjotMjcuMjUwNywiYWwiOjEzMywic3AiOjYsImhkIjoxMDAsImdqIjo4NSwiZ3MiOjEsImJ2IjozODk4LCJ0cCI6MzMsInJzIjotMTExLCJ0ciI6LTExNSwidHMiOi02LCJ0ZCI6MTY4NzI2NzY3MywiaHAiOjEzMiwidnAiOjIzNCwidGYiOjIxMTYwfQ==","len":175,"status":0,"hiveRxTime":"2023-06-20T13:31:26"}

// payload 2 data after base64 decode
{"dt":1687267682,"lt":38.6745,"ln":-27.2507,"al":133,"sp":6,"hd":100,"gj":85,"gs":1,"bv":3898,"tp":33,"rs":-111,"tr":-115,"ts":-6,"td":1687267673,"hp":132,"vp":234,"tf":21160}
```


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
  TODO
  ```


### Example Request:
+ Method: `POST`
+ URL: `https://loraspaceway.aircentre.io/uplink?access_key=your_access_key_here`
+ Headers: `Content-Type: application/json`
+ Body:
  ```json
  TODO
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
This project leverages GitHub Actions for a CI/CD process. Each commit triggers a workflow which builds, tests, and pushes a new Docker image called `webhook-spaceway-lora` to [GitHub's container Registry (ghcr.io)](https://github.com/AIRCentre/webhook-spaceway-lora/pkgs/container/webhook-spaceway-lora). Ansible, within the GitHub Actions workflow, deploys this updated image automatically.

To skip the workflow, add `no-ci` in the commit message - helpful when updating non-impactful files like the readme.
