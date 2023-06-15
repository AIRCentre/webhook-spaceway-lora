package events_repository

type SwarmPayload struct {
	Device         string  `json:"Device"`
	PacketID       int     `json:"Packet Id"`
	Timestamp      string  `json:"Timestamp"`
	RxTime         string  `json:"Rx Time"`
	Altitude       int     `json:"Altitude"`
	Heading        int     `json:"Heading"`
	Latitude       float64 `json:"Latitude"`
	Longitude      float64 `json:"Longitude"`
	GPSJamming     int     `json:"GPS Jamming"`
	GPSSpoofing    int     `json:"GPS Spoofing"`
	Temperature    int     `json:"Temperature"`
	BatteryVoltage int     `json:"Battery Voltage"`
	Speed          int     `json:"Speed"`
	TelemetrySNR   int     `json:"Telemetry SNR"`
	TelemetryRSSI  int     `json:"Telemetry RSSI"`
	TelemetryTime  int     `json:"Telemetry Time"`
	RSSIBackground int     `json:"RSSI Background"`
	TelemetryType  string  `json:"Telemetry Type"`
	Version        int     `json:"Version"`
}
