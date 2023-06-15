package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/AIRCentre/webhook-spaceway-lora/external/mysqldriver"
	"github.com/stretchr/testify/assert"
)

func TestHandlePushFromSwarm(t *testing.T) {
	t.Parallel()

	t.Run("sends POST request to downlink endpoint with a swarm payload and verifies that it was stored in the database", func(t *testing.T) {
		// given
		payload := `{
			"Device": "F-0x06eb2",
			"Packet Id": 52053866,
			"Timestamp": "Thu Mar 23 2023 01:00:06 GMT+0000 (Western European Standard Time)",
			"Rx Time": "Thu Mar 23 2023 16:30:52 GMT+0000 (Western European Standard Time)",
			"Altitude": 438,
			"Heading": 338,
			"Latitude": 40.2516,
			"Longitude": -7.4872,
			"GPS Jamming": 84,
			"GPS Spoofing": 1,
			"Temperature": 19,
			"Battery Voltage": 4021,
			"Speed": 0,
			"Telemetry SNR": -9,
			"Telemetry RSSI": -114,
			"Telemetry Time": 1679526068,
			"RSSI Background": -104,
			"Telemetry Type": "ASSET_TRACKER",
			"Version": 1
		}`

		// when
		statusCode, err := makePostRequest("http://localhost:3000/uplink", payload)
		if err != nil {
			t.Fatal("request failed, check args")
		}

		// then
		assert.Equal(t, http.StatusOK, statusCode)

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
		data, err := driver.Query("SELECT * FROM swarm_events WHERE Device='F-0x06eb2'")
		if err != nil {
			t.Fatal("query failed, check query string")
		}
		assert.Equal(t, 1, len(data))
	})
}

func makePostRequest(url string, payload string) (status int, err error) {
	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return 0, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
