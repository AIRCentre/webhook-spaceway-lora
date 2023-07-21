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
