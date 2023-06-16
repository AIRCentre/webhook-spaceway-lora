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
	Dt: 1686920651,
	Lt: 38.6534,
	Ln: -27.2188,
	Al: 10,
	Sp: 1,
	Hd: 0,
	Gj: 90,
	Gs: 1,
	Bv: 4000,
	Tp: 28,
	Rs: -112,
	Tr: 0,
	Ts: 0,
	Td: 0,
	Hp: 331,
	Vp: 420,
	Tf: 161914,
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
			Dt: 1686920652,
			Lt: 38.6532,
			Ln: -27.2182,
			Al: 12,
			Sp: 2,
			Hd: 2,
			Gj: 92,
			Gs: 2,
			Bv: 4002,
			Tp: 22,
			Rs: -114,
			Tr: 2,
			Ts: 2,
			Td: 2,
			Hp: 332,
			Vp: 422,
			Tf: 161912,
		}
		drivermock := mysqldriver.NewMock()
		repo := eventrepo.NewMysqlRepo(drivermock)

		// when
		repo.Insert(fakePayload)

		// then
		expectedQuery := `INSERT INTO swarm_events (timestamp, latitude_deg, longitude_deg, altitude, speed, heading, gps_jamming, gps_spoofing, battery_v, temperature_c, rssi_dbm,tr, ts, td, hp, vp, tf)
			VALUES (1686920652, 38.6532, -27.2182, 12, 2, 2, 92, 2, 4002, 22, -114, 2, 2, 2, 332, 422, 161912);`

		util.SQLEq(t, expectedQuery, drivermock.GetLastExec())
	})

	t.Run("calling the insert method with a valid payload should exec the correct sql insert query #2", func(t *testing.T) {
		t.Parallel()
		// given
		fakePayload := eventrepo.EventPayload{
			Dt: 1686920653,
			Lt: 38.6533,
			Ln: -27.2183,
			Al: 13,
			Sp: 3,
			Hd: 3,
			Gj: 93,
			Gs: 3,
			Bv: 4003,
			Tp: 23,
			Rs: -113,
			Tr: 3,
			Ts: 3,
			Td: 3,
			Hp: 333,
			Vp: 423,
			Tf: 161913,
		}
		drivermock := mysqldriver.NewMock()
		repo := eventrepo.NewMysqlRepo(drivermock)

		// when
		repo.Insert(fakePayload)

		// then
		expectedQuery := `INSERT INTO swarm_events (timestamp, latitude_deg, longitude_deg, altitude, speed, heading, gps_jamming, gps_spoofing, battery_v, temperature_c, rssi_dbm,tr, ts, td, hp, vp, tf)
			VALUES (1686920653, 38.6533, -27.2183, 13, 3, 3, 93, 3, 4003, 23, -113, 3, 3, 3, 333, 423, 161913);`

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
