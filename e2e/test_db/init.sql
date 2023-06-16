-- -- create databese: meteorology
-- CREATE DATABASE IF NOT EXISTS meteorology;

-- -- create table: stations
-- CREATE TABLE IF NOT EXISTS meteorology.stations(
--     station_id VARCHAR(45) NOT NULL,
--     place VARCHAR(50) DEFAULT NULL,
--     latitude_deg DECIMAL(8,6) NOT NULL,
--     longitude_deg DECIMAL(9,6) NOT NULL,
--     source VARCHAR(45) DEFAULT NULL,
--     PRIMARY KEY (station_id),
--     UNIQUE KEY station_id_UNIQUE (station_id)
-- );

-- -- create table: meteo_1hr
-- CREATE TABLE IF NOT EXISTS meteorology.meteo_1hr( 
--     id INT UNSIGNED NOT NULL AUTO_INCREMENT,  
--     station_id VARCHAR(45) NOT NULL,  
--     timestamp TIMESTAMP NOT NULL,
--     wind_speed_kmh FLOAT UNSIGNED DEFAULT NULL,
--     temperature_c FLOAT DEFAULT NULL,
--     radiation_kjm2 FLOAT UNSIGNED DEFAULT NULL,
--     wind_direction_bin FLOAT DEFAULT NULL,
--     precipitation_accum_mm FLOAT UNSIGNED DEFAULT NULL,
--     rel_humidity_pctg FLOAT UNSIGNED DEFAULT NULL,
--     pressure_hpa FLOAT DEFAULT NULL,
--     PRIMARY KEY (id), 
--     UNIQUE KEY id_UNIQUE (id),  
--     KEY fk_idx (station_id),  
--     CONSTRAINT fk FOREIGN KEY (station_id) REFERENCES meteorology.stations (station_id)
-- );

-- -- add data to stations
-- INSERT INTO meteorology.stations (station_id, place, latitude_deg, longitude_deg, source) VALUES
--   ('st1', 'a place', 1.1, -1.1, 'SOME_SRC'),
--   ('st2', 'other location', 2.2, -2.2, 'SOME_SRC'),
--   ('st3', 'yet another location', 3.3, -3.3, 'OTHR_SRC');

-- -- add data to meto_1hr
-- INSERT INTO meteorology.meteo_1hr (
--     station_id, 
--     timestamp, 
--     wind_speed_kmh, 
--     temperature_c, 
--     radiation_kjm2, 
--     wind_direction_bin, 
--     precipitation_accum_mm,
--     rel_humidity_pctg,
--     pressure_hpa
--   ) 
--   VALUES 
--   ( "st1", "2023-01-01 00:00:00+00:00", 1.1, 20.1, 1000, 180.0, 10.1, 80.2, 1200.3),
--   ( "st1", "2023-01-02 00:00:00+00:00", 2.1, 21.1, 2000, 280.0, 20.1, 70.2, 1000.3),
--   ( "st2", "2023-01-01 00:00:00+00:00", 3.1, 23.1, 1040, 100.0, 30.1, 60.2, 1700.3),
--   ( "st2", "2023-01-02 00:00:00+00:00", 4.1, 24.1, 2000, 150.0, 23.1, 73.2, 1020.4),
--   ( "st3", DATE_FORMAT(NOW(), '%Y-%m-%d 00:00:00+00:00'), NULL, 25.2, NULL, NULL, NULL, 96.7, 1021.4),
--   ( "st3", "2023-01-01 00:00:00+00:00", 20.1, NULL, NULL, 180.0, NULL, NULL, NULL);


-- -- create databese: vessel_location
-- CREATE DATABASE IF NOT EXISTS vessel_location;

-- -- create table: gateways
-- CREATE TABLE IF NOT EXISTS vessel_location.gateways(
-- gw_eui VARCHAR(45) NOT NULL,
-- gw_lon_deg_wgs84 DECIMAL(9,6) NOT NULL,
-- gw_lat_deg_wgs84 DECIMAL(8,6) NOT NULL,
-- PRIMARY KEY (gw_eui),
-- UNIQUE KEY gw_eui_UNIQUE (gw_eui)
-- );

-- -- create table: events
-- CREATE TABLE IF NOT EXISTS vessel_location.events(
-- id int unsigned NOT NULL AUTO_INCREMENT,
-- timestamp_utc_iso_string DATETIME NOT NULL,
-- dev_eui VARCHAR(45) NOT NULL,
-- dev_lon_deg_wgs84 DECIMAL(9,6) NOT NULL,
-- dev_lat_deg_wgs84 DECIMAL(8,6) NOT NULL,
-- rssi_dbm INT NOT NULL,
-- snr_db DECIMAL(2,1) NOT NULL,
-- distance_to_gw_km FLOAT NOT NULL,
-- gw_eui VARCHAR(45) NOT NULL,
-- PRIMARY KEY (id),
-- UNIQUE KEY id_UNIQUE (id),
-- KEY fk_idx (gw_eui),
-- CONSTRAINT fk FOREIGN KEY (gw_eui) REFERENCES vessel_location.gateways (gw_eui)
-- );

-- create databese: meteorology
CREATE DATABASE IF NOT EXISTS Vessel_location;

-- Cretate table for SWARM data
CREATE TABLE IF NOT EXISTS Vessel_location.swarm_events(
   id INT UNSIGNED NOT NULL AUTO_INCREMENT,  
   timestamp INT,
   latitude_deg FLOAT,
   longitude_deg FLOAT,
   altitude INT,
   speed INT,
   heading INT,
   gps_jamming INT,
   gps_spoofing INT,
   battery_v INT,
   temperature_c INT,
   rssi_dbm INT,
   tr INT,
   ts INT,
   td INT,
   hp INT,
   vp INT,
   tf INT,
   PRIMARY KEY (id),
   UNIQUE KEY id_UNIQUE (id)
);



-- -- Cretate table for SWARM data
-- CREATE TABLE IF NOT EXISTS Vessel_location.swarm_events(
--    id INT UNSIGNED NOT NULL AUTO_INCREMENT,  
--    device VARCHAR(15),
--    packet_id INT,
--    timestamp DATETIME,
--    rx_time DATETIME,
--    altitude INT,
--    heading INT,
--    latitude_deg FLOAT,
--    longitude_deg FLOAT,
--    gps_jamming INT,
--    gps_spoofing INT,
--    temperature_c INT,
--    battery_v INT,
--    speed INT,
--    telemetry_snr_db INT,
--    telemetry_rssi_dbm INT,
--    telemetry_time INT,
--    rssi_background_dbm INT,
--    telemetry_type VARCHAR(25),
--    version INT,
--    PRIMARY KEY (id),
--    UNIQUE KEY id_UNIQUE (id)
-- );
