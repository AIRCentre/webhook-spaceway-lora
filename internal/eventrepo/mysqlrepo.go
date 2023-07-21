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
		INSERT INTO swarm_events (
			device_id,
			timestamp,
			latitude_deg,
			longitude_deg,
			altitude,
			speed_mps,
			heading_deg,
			battery_v,
			cpu_temperature_c,
			rssi_dbm,
			snr_db,
			timestamp_at_reception,
			rssi_background_dbm
		)
		VALUES ('%s', %d, %f, %f, %d, %d, %d, %d, %d, %d, %d, %d, %d);`,
		deviceId,
		payload.Timestamp,
		payload.Latitude,
		payload.Longitude,
		payload.Altitude,
		payload.Speed,
		payload.Heading,
		payload.Battery,
		payload.CPUTemperature,
		payload.RSSI,
		payload.SNR,
		payload.TimestampAtReception,
		payload.RSSIBackground,
	)

	return strings.TrimSpace(query)
}

func NewMysqlRepo(driver mysqldriver.I) *mysqlrepo {
	return &mysqlrepo{
		mysqlDriver: driver,
	}
}
