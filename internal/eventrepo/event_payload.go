package eventrepo

type EventPayload struct {
	Device         string  `json:"device"`
	PacketID       int     `json:"packet_id"`
	Timestamp      string  `json:"timestamp"`
	RxTime         string  `json:"rx_time"`
	Altitude       int     `json:"altitude"`
	Heading        int     `json:"heading"`
	Latitude       float64 `json:"latitude_deg"`
	Longitude      float64 `json:"longitude_deg"`
	GPSJamming     int     `json:"gps_jamming"`
	GPSSpoofing    int     `json:"gps_spoofing"`
	Temperature    int     `json:"temperature_c"`
	BatteryVoltage int     `json:"battery_v"`
	Speed          int     `json:"speed"`
	TelemetrySNR   int     `json:"telemetry_snr_db"`
	TelemetryRSSI  int     `json:"telemetry_rssi_dbm"`
	TelemetryTime  int     `json:"telemetry_time"`
	RSSIBackground int     `json:"rssi_background_dbm"`
	TelemetryType  string  `json:"telemetry_type"`
	Version        int     `json:"version"`
}
