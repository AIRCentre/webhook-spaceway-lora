package mysqlrepo

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/AIRCentre/webhook-spaceway-lora/external/mysqldriver"
)

// type SwarmPayload struct {
// 	Device         string  `json:"Device"`
// 	PacketID       int     `json:"Packet Id"`
// 	Timestamp      string  `json:"Timestamp"`
// 	RxTime         string  `json:"Rx Time"`
// 	Altitude       int     `json:"Altitude"`
// 	Heading        int     `json:"Heading"`
// 	Latitude       float64 `json:"Latitude"`
// 	Longitude      float64 `json:"Longitude"`
// 	GPSJamming     int     `json:"GPS Jamming"`
// 	GPSSpoofing    int     `json:"GPS Spoofing"`
// 	Temperature    int     `json:"Temperature"`
// 	BatteryVoltage int     `json:"Battery Voltage"`
// 	Speed          int     `json:"Speed"`
// 	TelemetrySNR   int     `json:"Telemetry SNR"`
// 	TelemetryRSSI  int     `json:"Telemetry RSSI"`
// 	TelemetryTime  int     `json:"Telemetry Time"`
// 	RSSIBackground int     `json:"RSSI Background"`
// 	TelemetryType  string  `json:"Telemetry Type"`
// 	Version        int     `json:"Version"`
// }


type SwarmPayload struct {
	device              string  `json:"Device"`
	packet_id           int     `json:"Packet Id"`
	timestamp           string  `json:"Timestamp"`
	rx_time             string  `json:"Rx Time"`
	altitude            int     `json:"Altitude"`
	heading             int     `json:"Heading"`
	latitude_deg        float64 `json:"Latitude"`
	longitude_deg       float64 `json:"Longitude"`
	gps_jamming         int     `json:"GPS Jamming"`
	gps_spoofing        int     `json:"GPS Spoofing"`
	temperature         int     `json:"Temperature"`
	battery_v           int     `json:"Battery Voltage"`
	speed               int     `json:"Speed"`
	telemetry_snr_db    int     `json:"Telemetry SNR"`
	telemetry_rssi_dbm  int     `json:"Telemetry RSSI"`
	telemetry_time      int     `json:"Telemetry Time"`
	rssi_background_dbm int     `json:"RSSI Background"`
	telemetry_type      string  `json:"Telemetry Type"`
	version             int     `json:"Version"`
}


type mysqlrepo struct {
	mysqlDriver mysqldriver.I
}

func (r *mysqlrepo) Insert(payload SwarmPayload) error {
	q := insertQuery(payload)
	_, _, err := r.mysqlDriver.Exec(q)
	if err != nil {
		return errors.New("insert failed due to mysql driver error: " + err.Error())
	}
	return nil
}

func formatTimestamp(ts string) string {
	// Remove the "(Western European Standard Time)" part
	index := strings.Index(ts, "(")
	if index != -1 {
		ts = strings.TrimSpace(ts[:index])
	}

	// Define the input date format
	inputFormat := "Mon Jan 2 2006 15:04:05 GMT-0700"

	// Parse the input date string
	t, err := time.Parse(inputFormat, ts)
	if err != nil {
		return "0000-00-00 00:00:00" // Replace with a sensible default or handle the error as appropriate for your use case.

	}

	// Define the output date format
	outputFormat := "2006-01-02 15:04:05"

	// Format the date as per the SQL format
	formattedDate := t.Format(outputFormat)

	return formattedDate
}

func insertQuery(payload SwarmPayload) string {
	query := fmt.Sprintf(`
		INSERT INTO swarm_events (device, packet_id, timestamp, rx_time, altitude, heading, latitude_deg, longitude_deg, gps_jamming, gps_spoofing, temperature, battery_voltage, speed, telemetry_snr_db, telemetry_rssi_dbm, telemetry_time, rssi_background_dbm, telemetry_type, version)
		VALUES ('%s', %d, '%s', '%s', %d, %d, %f, %f, %d, %d, %d, %d, %d, %d, %d, %d, %d, '%s', %d);`,
		payload.Device,
		payload.PacketID,
		formatTimestamp(payload.Timestamp),
		formatTimestamp(payload.RxTime),
		payload.Altitude,
		payload.Heading,
		payload.Latitude,
		payload.Longitude,
		payload.GPSJamming,
		payload.GPSSpoofing,
		payload.Temperature,
		payload.BatteryVoltage,
		payload.Speed,
		payload.TelemetrySNR,
		payload.TelemetryRSSI,
		payload.TelemetryTime,
		payload.RSSIBackground,
		payload.TelemetryType,
		payload.Version)

	return strings.TrimSpace(query)
}

func New(driver mysqldriver.I) *mysqlrepo {
	return &mysqlrepo{
		mysqlDriver: driver,
	}
}
