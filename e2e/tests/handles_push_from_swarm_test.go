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
			"dt": 1686920651,
			"lt": 38.6534,
			"ln": -27.2188,
			"al": 10,
			"sp": 1,
			"hd": 0,
			"gj": 90,
			"gs": 1,
			"bv": 4000,
			"tp": 28,
			"rs": -112,
			"tr": 0,
			"ts": 0,
			"td": 0,
			"hp": 331,
			"vp": 420,
			"tf": 161914
		  }`

		// when
		statusCode, err := makePostRequest("http://localhost:3000/uplink?access_key=test", payload)
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
		data, err := driver.Query("SELECT * FROM swarm_events")
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
