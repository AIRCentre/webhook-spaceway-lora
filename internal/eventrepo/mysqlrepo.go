package eventrepo

import (
	"errors"
	"fmt"
	"strings"

	"github.com/AIRCentre/webhook-spaceway-lora/external/mysqldriver"
)

type mysqlrepo struct {
	mysqlDriver mysqldriver.I
}

func (r *mysqlrepo) Insert(deviceId string, payload EventPayload) error {
	q := buildQuery(deviceId, payload)
	_, _, err := r.mysqlDriver.Exec(q)
	if err != nil {
		return errors.New("insert failed due to mysql driver error: " + err.Error())
	}
	return nil
}

// func formatTimestamp(ts string) string {
// 	// Remove the "(Western European Standard Time)" part
// 	index := strings.Index(ts, "(")
// 	if index != -1 {
// 		ts = strings.TrimSpace(ts[:index])
// 	}

// 	// Define the input date format
// 	inputFormat := "Mon Jan 2 2006 15:04:05 GMT-0700"

// 	// Parse the input date string
// 	t, err := time.Parse(inputFormat, ts)
// 	if err != nil {
// 		return "0000-00-00 00:00:00" // Replace with a sensible default or handle the error as appropriate for your use case.

// 	}

// 	// Define the output date format
// 	outputFormat := "2006-01-02 15:04:05"

// 	// Format the date as per the SQL format
// 	formattedDate := t.Format(outputFormat)

// 	return formattedDate
// }

// func parseUnixEpochTime(epoch int64) string {
// 	t := time.Unix(epoch, 0)
// 	return t.Format("2006-01-02 15:04:05")
// }

func buildQuery(deviceId string, payload EventPayload) string {
	query := fmt.Sprintf(`
		INSERT INTO swarm_events (device_id, timestamp, latitude_deg, longitude_deg, altitude, speed, heading, gps_jamming, gps_spoofing, battery_v, temperature_c, rssi_dbm, tr, ts, td, hp, vp, tf)
		VALUES ('%s', %d, %f, %f, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d);`,
		deviceId,
		payload.Timestamp,
		payload.Latitude,
		payload.Longitude,
		payload.Altitude,
		payload.Speed,
		payload.Heading,
		payload.GPSJamming,
		payload.GPSSpoofing,
		payload.BatteryVoltage,
		payload.Temperature,
		payload.RSSI,
		payload.Tr,
		payload.Ts,
		payload.Td,
		payload.Hp,
		payload.Vp,
		payload.Tf)

	return strings.TrimSpace(query)
}

func NewMysqlRepo(driver mysqldriver.I) *mysqlrepo {
	return &mysqlrepo{
		mysqlDriver: driver,
	}
}

// func buildQuery(payload EventPayload) string {
// 	query := fmt.Sprintf(`
// 		INSERT INTO swarm_events (timestamp, latitude_deg, longitude_deg, altitude, speed, heading, gps_jamming, gps_spoofing, battery_v, temperature_c, rssi_dbm,tr, ts, td, hp, vp, tf)
// 		VALUES ('%s', %d, '%s', '%s', %d, %d, %f, %f, %d, %d, %d, %d, %d, %d, %d, %d, %d, '%s', %d);`,
// 		payload.Device,
// 		payload.PacketID,
// 		formatTimestamp(payload.Timestamp),
// 		formatTimestamp(payload.RxTime),
// 		payload.Altitude,
// 		payload.Heading,
// 		payload.Latitude,
// 		payload.Longitude,
// 		payload.GPSJamming,
// 		payload.GPSSpoofing,
// 		payload.Temperature,
// 		payload.BatteryVoltage,
// 		payload.Speed,
// 		payload.TelemetrySNR,
// 		payload.TelemetryRSSI,
// 		payload.TelemetryTime,
// 		payload.RSSIBackground,
// 		payload.TelemetryType,
// 		payload.Version)

// 	return strings.TrimSpace(query)
// }
