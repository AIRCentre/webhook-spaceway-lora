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
		payload := `{"packetId":63464221,"deviceType":1,"deviceId":28338,"userApplicationId":65002,"organizationId":66429,"data":"eyJkdCI6MTY4NzE5OTQ1MCwibHQiOjM4LjY1NTIsImxuIjotMjcuMjE5MywiYWwiOjMzLCJzcCI6OSwiaGQiOjY1LCJnaiI6ODksImdzIjoxLCJidiI6MzkyMCwidHAiOjI2LCJycyI6LTgyLCJ0ciI6LTExNSwidHMiOi0yLCJ0ZCI6MTY4NzE5NjUwOSwiaHAiOjEyNjYsInZwIjoxMDAsInRmIjoyMDU0NDB9","len":174,"status":0,"hiveRxTime":"2023-06-20T11:33:56"}`
		// decoded base64 payload data
		// {"dt":1687199450,"lt":38.6552,"ln":-27.2193,"al":33,"sp":9,"hd":65,"gj":89,"gs":1,"bv":3920,"tp":26,"rs":-82,"tr":-115,"ts":-2,"td":1687196509,"hp":1266,"vp":100,"tf":205440}

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
		data, err := driver.Query("SELECT * FROM swarm_events WHERE device_id='28338' AND altitude=33")
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
