package mysqlrepo_test

import (
	"errors"
	"testing"

	"github.com/AIRCentre/webhook-spaceway-lora/external/mysqldriver"
	"github.com/AIRCentre/webhook-spaceway-lora/internal/mysqlrepo"
	"github.com/AIRCentre/webhook-spaceway-lora/util"
	"github.com/stretchr/testify/assert"
)

var ValidPayload mysqlrepo.SwarmPayload = mysqlrepo.SwarmPayload{
	Device:         "F-0x06eb2",
	PacketID:       52053866,
	Timestamp:      "Thu Mar 23 2023 01:00:06 GMT+0000 (Western European Standard Time)",
	RxTime:         "Thu Mar 23 2023 16:30:52 GMT+0000 (Western European Standard Time)",
	Altitude:       438,
	Heading:        338,
	Latitude:       40.2516,
	Longitude:      -7.4872,
	GPSJamming:     84,
	GPSSpoofing:    1,
	Temperature:    19,
	BatteryVoltage: 4021,
	Speed:          0,
	TelemetrySNR:   -9,
	TelemetryRSSI:  -114,
	TelemetryTime:  1679526068,
	RSSIBackground: -104,
	TelemetryType:  "ASSET_TRACKER",
	Version:        1,
}

func TestMysqlRepo(t *testing.T) {
	t.Parallel()

	t.Run("calling the insert method should exec a query once", func(t *testing.T) {
		t.Parallel()
		// given
		fakePayload := mysqlrepo.SwarmPayload{}
		drivermock := mysqldriver.NewMock()
		repo := mysqlrepo.New(drivermock)

		// when
		repo.Insert(fakePayload)

		// then
		assert.Equal(t, 1, drivermock.ExecCallCount)
	})

	t.Run("calling the insert method with a valid payload should exec the correct sql insert query #1", func(t *testing.T) {
		t.Parallel()
		// given
		fakePayload := mysqlrepo.SwarmPayload{
			device:              "F-0x06eb2",
			packet_id:           52053866,
			timestamp:           "Thu Mar 23 2023 01:00:06 GMT+0000 (Western European Standard Time)",
			rx_time:             "Thu Mar 23 2023 16:30:52 GMT+0000 (Western European Standard Time)",
			altitude:            438,
			heading:             338,
			latitude_deg:        40.2516,
			longitude_deg:       -7.4872,
			gps_jamming:         84,
			gps_spoofing:        1,
			temperature_c:       19,
			battery_v:           4021,
			speed:               0,
			telemetry_snr_db:    -9,
			telemetry_rssi_dbm:  -114,
			telemetry_time:      1679526068,
			rssi_background_dbm: -104,
			telemetry_type:      "ASSET_TRACKER",
			Version:             1,
		}
		drivermock := mysqldriver.NewMock()
		repo := mysqlrepo.New(drivermock)

		// when
		repo.Insert(fakePayload)

		// then
		expectedQuery := `INSERT INTO swarm_events (device, packet_id, timestamp, rx_time, altitude, heading, latitude_deg, longitude_deg, gps_jamming, gps_spoofing, temperature, battery_voltage, speed, telemetry_snr_db, telemetry_rssi_dbm, telemetry_time, rssi_background_dbm, telemetry_type, version)
			VALUES ('F-0x06eb2', 52053866, '2023-03-23 01:00:06', '2023-03-23 16:30:52', 438, 338, 40.251600, -7.487200, 84, 1, 19, 4021, 0, -9, -114, 1679526068, -104, 'ASSET_TRACKER', 1);`

		util.SQLEq(t, expectedQuery, drivermock.GetLastExec())
	})

	t.Run("calling the insert method with a valid payload should exec the correct sql insert query #2", func(t *testing.T) {
		t.Parallel()
		// given
		fakePayload := mysqlrepo.SwarmPayload{
			device:             "H-0x98ab1",
			packet_id:           98765432,
			timestamp:           "Tue Jun 13 2023 10:30:15 GMT+0000 (Western European Standard Time)",
			rx_time:             "Tue Jun 13 2023 18:45:22 GMT+0000 (Western European Standard Time)",
			altitude:            762,
			heading:             221,
			latitude_deg:        41.8993,
			longitude_deg:       -8.7325,
			gps_jamming:         62,
			gps_spoofing:        2,
			temperature_c:       27,
			battery_v:           4037,
			speed:               15,
			telemetry_snr_db:    -7,
			telemetry_rssi_dbm:  -112,
			telemetry_time:      1679584409,
			rssi_background_dbm: -98,
			telemetry_type:      "LOCATION_TRACKER",
			version:             3,
		}
		drivermock := mysqldriver.NewMock()
		repo := mysqlrepo.New(drivermock)

		// when
		repo.Insert(fakePayload)

		// then
		expectedQuery := `INSERT INTO swarm_events (device, packet_id, timestamp, rx_time, altitude, heading, latitude_deg, longitude_deg, gps_jamming, gps_spoofing, temperature, battery_voltage, speed, telemetry_snr_db, telemetry_rssi_dbm, telemetry_time, rssi_background_dbm, telemetry_type, version)
			VALUES ('H-0x98ab1', 98765432, '2023-06-13 10:30:15', '2023-06-13 18:45:22', 762, 221, 41.899300, -8.732500, 62, 2, 27, 4037, 15, -7, -112, 1679584409, -98, 'LOCATION_TRACKER', 3);`

		util.SQLEq(t, expectedQuery, drivermock.GetLastExec())
	})

	t.Run("handles an error from mysql driver #1", func(t *testing.T) {
		t.Parallel()
		// given
		drivermock := mysqldriver.NewMock()
		drivermock.SetError(errors.New("fake db error"))
		repo := mysqlrepo.New(drivermock)

		// when
		err := repo.Insert(ValidPayload)

		// then
		assert.EqualError(t, err, "insert failed due to mysql driver error: fake db error")
	})

	t.Run("handles an error from mysql driver #2", func(t *testing.T) {
		t.Parallel()
		// given
		drivermock := mysqldriver.NewMock()
		drivermock.SetError(errors.New("different err msg"))
		repo := mysqlrepo.New(drivermock)

		// when
		err := repo.Insert(ValidPayload)

		// then
		assert.EqualError(t, err, "insert failed due to mysql driver error: different err msg")
	})
}
