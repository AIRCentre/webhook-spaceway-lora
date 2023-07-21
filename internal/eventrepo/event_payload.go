package eventrepo

type EventPayload struct {
	Timestamp            int64   `json:"dt"`
	Latitude             float64 `json:"lt"`
	Longitude            float64 `json:"ln"`
	Altitude             int     `json:"al"`
	Speed                int     `json:"sp"`
	Heading              int     `json:"hd"`
	BatteryVoltage       int     `json:"bv"`
	Temperature          int     `json:"tp"`
	RSSIBackground       int     `json:"rs"`
	RSSI                 int     `json:"tr"`
	SNR                  int     `json:"ts"`
	TimestampAtReception int     `json:"td"`
}

// [SQL Fields]
// timestamp INT,
// latitude_deg FLOAT,
// longitude_deg FLOAT,
// altitude INT,
// speed_mps INT,
// heading_deg INT,
// battery_v INT,
// cpu_temperature_c INT,
// rssi_dbm INT,
// snr_db INT,
// timestamp_at_reception INT,
// rssi_background_dbm INT,
