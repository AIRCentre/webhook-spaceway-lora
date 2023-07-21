package tests

import (
	"testing"

	"github.com/AIRCentre/webhook-spaceway-lora/external/mysqldriver"
)

func TestMysqlDriver(t *testing.T) {
	t.Parallel()
	driver, err := mysqldriver.New(
		"root",
		"mysecretpassword",
		"localhost",
		"3306",
		"Vessel_location",
	)
	if err != nil {
		t.Fatal("failed to connect to test_db, check connection args")
	}
	_, err = driver.Query("SELECT * FROM swarm_events")
	if err != nil {
		t.Fatal("query failed")
	}
}
