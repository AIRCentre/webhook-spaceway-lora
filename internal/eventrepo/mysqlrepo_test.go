package eventrepo_test

import (
	"errors"
	"testing"

	"github.com/AIRCentre/webhook-spaceway-lora/external/mysqldriver"
	"github.com/AIRCentre/webhook-spaceway-lora/internal/eventrepo"

	"github.com/AIRCentre/webhook-spaceway-lora/util"
	"github.com/stretchr/testify/assert"
)

var ValidPayload eventrepo.SwarmPayload = eventrepo.SwarmPayload{
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
		fakePayload := eventrepo.SwarmPayload{}
		drivermock := mysqldriver.NewMock()
		repo := eventrepo.NewMysqlRepo(drivermock)

		// when
		repo.Insert(fakePayload)

		// then
		assert.Equal(t, 1, drivermock.ExecCallCount)
	})

	t.Run("calling the insert method with a valid payload should exec the correct sql insert query #1", func(t *testing.T) {
		t.Parallel()
		// given
		fakePayload := eventrepo.SwarmPayload{
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
		drivermock := mysqldriver.NewMock()
		repo := eventrepo.NewMysqlRepo(drivermock)

		// when
		repo.Insert(fakePayload)

		// then
		expectedQuery := `INSERT INTO swarm_events (device, packet_id, timestamp, rx_time, altitude, heading, latitude_deg, longitude_deg, gps_jamming, gps_spoofing, temperature_c, battery_voltage, speed, telemetry_snr_db, telemetry_rssi_dbm, telemetry_time, rssi_background_dbm, telemetry_type, version)
			VALUES ('F-0x06eb2', 52053866, '2023-03-23 01:00:06', '2023-03-23 16:30:52', 438, 338, 40.251600, -7.487200, 84, 1, 19, 4021, 0, -9, -114, 1679526068, -104, 'ASSET_TRACKER', 1);`

		util.SQLEq(t, expectedQuery, drivermock.GetLastExec())
	})

	t.Run("calling the insert method with a valid payload should exec the correct sql insert query #2", func(t *testing.T) {
		t.Parallel()
		// given
		fakePayload := eventrepo.SwarmPayload{
			Device:         "H-0x98ab1",
			PacketID:       98765432,
			Timestamp:      "Tue Jun 13 2023 10:30:15 GMT+0000 (Western European Standard Time)",
			RxTime:         "Tue Jun 13 2023 18:45:22 GMT+0000 (Western European Standard Time)",
			Altitude:       762,
			Heading:        221,
			Latitude:       41.8993,
			Longitude:      -8.7325,
			GPSJamming:     62,
			GPSSpoofing:    2,
			Temperature:    27,
			BatteryVoltage: 4037,
			Speed:          15,
			TelemetrySNR:   -7,
			TelemetryRSSI:  -112,
			TelemetryTime:  1679584409,
			RSSIBackground: -98,
			TelemetryType:  "LOCATION_TRACKER",
			Version:        3,
		}
		drivermock := mysqldriver.NewMock()
		repo := eventrepo.NewMysqlRepo(drivermock)

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
		repo := eventrepo.NewMysqlRepo(drivermock)

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
		repo := eventrepo.NewMysqlRepo(drivermock)

		// when
		err := repo.Insert(ValidPayload)

		// then
		assert.EqualError(t, err, "insert failed due to mysql driver error: different err msg")
	})
}
