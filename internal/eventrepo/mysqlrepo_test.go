package eventrepo_test

import (
	"errors"
	"testing"

	"github.com/AIRCentre/webhook-spaceway-lora/external/mysqldriver"
	"github.com/AIRCentre/webhook-spaceway-lora/internal/eventrepo"
	"github.com/AIRCentre/webhook-spaceway-lora/util"

	"github.com/stretchr/testify/assert"
)

var ValidPayload eventrepo.EventPayload = eventrepo.EventPayload{
	Timestamp:            1686920651,
	Latitude:             38.6534,
	Longitude:            -27.2188,
	Altitude:             10,
	Speed:                1,
	Heading:              0,
	Battery:              4000,
	CPUTemperature:       28,
	RSSI:                 -112,
	SNR:                  0,
	TimestampAtReception: 1686920654,
	RSSIBackground:       -99}

func TestMysqlRepo(t *testing.T) {
	t.Parallel()

	t.Run("calling the insert method should exec a query once", func(t *testing.T) {
		t.Parallel()
		// given
		fakePayload := eventrepo.EventPayload{}
		drivermock := mysqldriver.NewMock()
		repo := eventrepo.NewMysqlRepo(drivermock)

		// when
		repo.Insert("", fakePayload)

		// then
		assert.Equal(t, 1, drivermock.ExecCallCount)
	})

	t.Run("calling the insert method with a valid payload should exec the correct sql insert query #1", func(t *testing.T) {
		t.Parallel()
		// given
		deviceId := "123456"
		fakePayload := eventrepo.EventPayload{
			Timestamp:            1686920652,
			Latitude:             38.6532,
			Longitude:            -27.2182,
			Altitude:             12,
			Speed:                92,
			Heading:              2,
			Battery:              4002,
			CPUTemperature:       22,
			RSSI:                 -104,
			SNR:                  2,
			TimestampAtReception: 1686920789,
			RSSIBackground:       -92,
		}
		drivermock := mysqldriver.NewMock()
		repo := eventrepo.NewMysqlRepo(drivermock)

		// when
		repo.Insert(deviceId, fakePayload)

		// then
		expectedQuery := `
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
			VALUES (
				'123456', 
				1686920652, 
				38.653200, 
				-27.218200, 
				12, 
				92, 
				2, 
				4002, 
				22,
				-104,
				2,
				1686920789,
				-92
			);`

		util.SQLEq(t, expectedQuery, drivermock.GetLastExec())
	})

	t.Run("calling the insert method with a valid payload should exec the correct sql insert query #2", func(t *testing.T) {
		t.Parallel()
		// given
		deviceId := "123567"
		fakePayload := eventrepo.EventPayload{
			Timestamp:            1686920123,
			Latitude:             38.6345,
			Longitude:            -27.2567,
			Altitude:             13,
			Speed:                3,
			Heading:              4,
			Battery:              3002,
			CPUTemperature:       12,
			RSSI:                 -104,
			SNR:                  4,
			TimestampAtReception: 1686920123,
			RSSIBackground:       -82,
		}
		drivermock := mysqldriver.NewMock()
		repo := eventrepo.NewMysqlRepo(drivermock)

		// when
		repo.Insert(deviceId, fakePayload)

		// then
		expectedQuery := `
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
			VALUES (
				'123567', 
				1686920123, 
				38.634500, 
				-27.256700, 
				13, 
				3, 
				4, 
				3002, 
				12,
				-104,
				4,
				1686920123,
				-82
			);`

		util.SQLEq(t, expectedQuery, drivermock.GetLastExec())
	})

	t.Run("handles an error from mysql driver #1", func(t *testing.T) {
		t.Parallel()
		// given
		drivermock := mysqldriver.NewMock()
		drivermock.SetError(errors.New("fake db error"))
		repo := eventrepo.NewMysqlRepo(drivermock)

		// when
		err := repo.Insert("123", ValidPayload)

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
		err := repo.Insert("123", ValidPayload)

		// then
		assert.EqualError(t, err, "insert failed due to mysql driver error: different err msg")
	})
}
