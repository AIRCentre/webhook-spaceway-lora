  ```json
   {
    1- "device": "",
    2- "packet_id": 0,
    3- "timestamp": "RFC 822/1123 date-time format",
    4- "rx_time": "RFC 822/1123 date-time format",
    5- "altitude": 0.0,
    6- "heading": 0,
    7- "latitude_deg": 0.0,
    8- "longitude_deg": 0.0,
    9- "gps_jamming": 0,
    10- "gps_spoofing": 0,
    11- "temperature_c": 0.0,
    12- "battery_v": 0,
    13- "speed": 0.0,
    14- "telemetry_snr_db": 0.0,
    15- "telemetry_rssi_dbm": 0,
    16- "telemetry_time": 0,
    17- "rssi_background": 0,
    18- "telemetry_type": "",
    19- "version": 0
  }
  ```

{"dt":1686920651,"lt":38.6534,"ln":-27.2188,"al":10,"sp":1,"hd":0,"gj":90,"gs":1,"bv":4000,"tp":28,"rs":-112,"tr":0,"ts":0,"td":0,"hp":331,"vp":420,"tf":161914}

{
3- "dt":1686920651, (Unix epoch time GMT+0)
7- "lt":38.6534,
8- "ln":-27.2188,
5- "al":10,
13- "sp":1,
6- "hd":0,
9- "gj":90,
10- "gs":1,
12- "bv":4000,
11- "tp":28,
?15- "rs":-112,
"tr":0,
"ts":0,
"td":0,
"hp":331,
"vp":420,
"tf":161914
}
count = 17

{"dt":1686920651,
"lt":38.6534,
"ln":-27.2188,
"al":10,
"sp":1,
"hd":0,
"gj":90,
"gs":1,
"bv":4000,
"tp":28,
"rs":-112,
"tr":0,
"ts":0,
"td":0,
"hp":331,
"vp":420,
"tf":161914
}
type EventPayload struct {
	Dt      string  `json:"timestamp_unix_epoch_time"`
	Lt      float64 `json:"latitude_deg"`
	Ln      float64 `json:"longitude_deg"`
	Al      int     `json:"altitude"`
	Sp      int     `json:"speed"`
	Hd      int     `json:"heading"`
	Gj      int     `json:"gps_jamming"`
	Gs      int     `json:"gps_spoofing"`
	Bv      int     `json:"battery_v"`
	Tp      int     `json:"temperature_c"`
	Rs      int     `json:"rssi_dbm"`
    Tr      int     `json:"tr"`
    Ts      int     `json:"ts"`
    Td      int     `json:"td"`
    Hp      int     `json:"hp"`
    Vp      int     `json:"vp"`
    Tf      int     `json:"tf"`
}


(dt, lt, ln, al, sp, hd, gj, gs, bv, tp, rs, tr, ts, td, hp, vp, tf)


(timestamp_unix_epoch_time, latitude_deg, longitude_deg, altitude, speed, heading, gps_jamming, gps_spoofing, battery_v, temperature_c, rssi_dbm,tr, ts, td, hp, vp, tf)






{
"dt":1686994251,
"lt":0.0000,
"ln":0.0000,
"al":0,
"sp":0,
"hd":0,
"gj":90,
"gs":1,
"bv":3985,
"tp":26,
"rs":-99,
"tr":-118,
"ts":-8,
"td":1686986993,
"hp":9999,
"vp":9999,
"tf":0
}

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
   snr_db INT,
   tr INT,
   ts INT,
   td INT,
   hp INT,
   vp INT,
   tf INT,
   signal_quality VARCHAR(20),
   PRIMARY KEY (id),
   UNIQUE KEY id_UNIQUE (id)
) AUTO_INCREMENT=1;

-- Create a trigger to automatically calculate signal_quality
DELIMITER //
CREATE TRIGGER calculate_signal_quality BEFORE INSERT ON Vessel_location.swarm_events
FOR EACH ROW
BEGIN
  SET NEW.signal_quality =
    CASE
      WHEN NEW.rssi_dbm < -120 OR NEW.snr_db < -20 THEN 'Very Poor Signal Quality'
      WHEN (NEW.rssi_dbm >= -120 AND NEW.rssi_dbm <= -110) OR (NEW.snr_db >= -20 AND NEW.snr_db <= -10) THEN 'Poor Signal Quality'
      WHEN (NEW.rssi_dbm >= -110 AND NEW.rssi_dbm <= -100) OR (NEW.snr_db >= -10 AND NEW.snr_db <= 0) THEN 'Fair Signal Quality'
      WHEN (NEW.rssi_dbm >= -100 AND NEW.rssi_dbm <= -90) OR (NEW.snr_db >= 0 AND NEW.snr_db <= 10) THEN 'Good Signal Quality'
      WHEN NEW.rssi_dbm > -90 OR NEW.snr_db > 10 THEN 'Excellent Signal Quality'
      ELSE 'Unknown'
    END;
END //
DELIMITER ;
