# Webhook For Space-based LoRa Communication of Vessel Tracking Data
*Developed by Spaceway*
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
  + Underscore rathern than spaces
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

