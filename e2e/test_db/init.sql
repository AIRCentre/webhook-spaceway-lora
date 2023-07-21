
-- create databese: meteorology
CREATE DATABASE IF NOT EXISTS Vessel_location;

-- Cretate table for SWARM data
CREATE TABLE IF NOT EXISTS Vessel_location.swarm_events(
   id INT UNSIGNED NOT NULL AUTO_INCREMENT,  
   device_id VARCHAR(45),
   timestamp INT,
   latitude_deg FLOAT,
   longitude_deg FLOAT,
   altitude INT,
   speed_mps INT,
   heading_deg INT,
   battery_v INT,
   cpu_temperature_c INT,
   rssi_dbm INT,
   snr_db INT,
   timestamp_at_reception INT,
   rssi_background_dbm INT,
   signal_strength VARCHAR(20),
   PRIMARY KEY (id),
   UNIQUE KEY id_UNIQUE (id)
);


-- Create a trigger to automatically calculate signal_quality
CREATE TRIGGER calculate_signal_quality BEFORE INSERT ON Vessel_location.swarm_events
FOR EACH ROW
BEGIN
  SET NEW.signal_quality =
    CASE
      WHEN NEW.rssi_dbm < -120 OR NEW.snr_db < -20 THEN 'Very Poor'
      WHEN (NEW.rssi_dbm >= -120 AND NEW.rssi_dbm <= -110) OR (NEW.snr_db >= -20 AND NEW.snr_db <= -10) THEN 'Poor'
      WHEN (NEW.rssi_dbm >= -110 AND NEW.rssi_dbm <= -100) OR (NEW.snr_db >= -10 AND NEW.snr_db <= 0) THEN 'Fair'
      WHEN (NEW.rssi_dbm >= -100 AND NEW.rssi_dbm <= -90) OR (NEW.snr_db >= 0 AND NEW.snr_db <= 10) THEN 'Good'
      WHEN NEW.rssi_dbm > -90 OR NEW.snr_db > 10 THEN 'Excellent'
      ELSE 'Unknown'
    END;
END;

