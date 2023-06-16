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
	Timestamp:      1686920651,
	Latitude:       38.6534,
	Longitude:      -27.2188,
	Altitude:       10,
	Speed:          1,
	Heading:        0,
	GPSJamming:     90,
	GPSSpoofing:    1,
	BatteryVoltage: 4000,
	Temperature:    28,
	RSSI:           -112,
	Tr:             0,
	Ts:             0,
	Td:             0,
	Hp:             331,
	Vp:             420,
	Tf:             161914,
}

func TestMysqlRepo(t *testing.T) {
	t.Parallel()

	t.Run("calling the insert method should exec a query once", func(t *testing.T) {
		t.Parallel()
		// given
		fakePayload := eventrepo.EventPayload{}
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
		fakePayload := eventrepo.EventPayload{
			Timestamp:      1686920652,
			Latitude:       38.6532,
			Longitude:      -27.2182,
			Altitude:       12,
			Speed:          2,
			Heading:        2,
			GPSJamming:     92,
			GPSSpoofing:    2,
			BatteryVoltage: 4002,
			Temperature:    22,
			RSSI:           -114,
			Tr:             2,
			Ts:             2,
			Td:             2,
			Hp:             332,
			Vp:             422,
			Tf:             161912,
		}
		drivermock := mysqldriver.NewMock()
		repo := eventrepo.NewMysqlRepo(drivermock)

		// when
		repo.Insert(fakePayload)

		// then
		expectedQuery := `INSERT INTO swarm_events (timestamp, latitude_deg, longitude_deg, altitude, speed, heading, gps_jamming, gps_spoofing, battery_v, temperature_c, rssi_dbm,tr, ts, td, hp, vp, tf)
			VALUES (1686920652, 38.653200, -27.218200, 12, 2, 2, 92, 2, 4002, 22, -114, 2, 2, 2, 332, 422, 161912);`

		util.SQLEq(t, expectedQuery, drivermock.GetLastExec())
	})

	t.Run("calling the insert method with a valid payload should exec the correct sql insert query #2", func(t *testing.T) {
		t.Parallel()
		// given
		fakePayload := eventrepo.EventPayload{
			Timestamp:      1686920653,
			Latitude:       38.6533,
			Longitude:      -27.2183,
			Altitude:       13,
			Speed:          3,
			Heading:        3,
			GPSJamming:     93,
			GPSSpoofing:    3,
			BatteryVoltage: 4003,
			Temperature:    23,
			RSSI:           -113,
			Tr:             3,
			Ts:             3,
			Td:             3,
			Hp:             333,
			Vp:             423,
			Tf:             161913,
		}
		drivermock := mysqldriver.NewMock()
		repo := eventrepo.NewMysqlRepo(drivermock)

		// when
		repo.Insert(fakePayload)

		// then
		expectedQuery := `INSERT INTO swarm_events (timestamp, latitude_deg, longitude_deg, altitude, speed, heading, gps_jamming, gps_spoofing, battery_v, temperature_c, rssi_dbm,tr, ts, td, hp, vp, tf)
			VALUES (1686920653, 38.653300, -27.218300, 13, 3, 3, 93, 3, 4003, 23, -113, 3, 3, 3, 333, 423, 161913);`

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
