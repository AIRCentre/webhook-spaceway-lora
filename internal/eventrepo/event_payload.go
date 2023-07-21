package eventrepo

type EventPayload struct {
	Timestamp            int64   `json:"dt"`
	Latitude             float64 `json:"lt"`
	Longitude            float64 `json:"ln"`
	Altitude             int     `json:"al"`
	Speed                int     `json:"sp"`
	Heading              int     `json:"hd"`
	Battery              int     `json:"bv"`
	CPUTemperature       int     `json:"tp"`
	RSSI                 int     `json:"tr"`
	SNR                  int     `json:"ts"`
	TimestampAtReception int     `json:"td"`
	RSSIBackground       int     `json:"rs"`
}
