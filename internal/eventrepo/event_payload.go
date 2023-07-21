package eventrepo

type EventPayload struct {
	Timestamp      int64   `json:"dt"`
	Latitude       float64 `json:"lt"`
	Longitude      float64 `json:"ln"`
	Altitude       int     `json:"al"`
	Speed          int     `json:"sp"`
	Heading        int     `json:"hd"`
	GPSJamming     int     `json:"gj"`
	GPSSpoofing    int     `json:"gs"`
	BatteryVoltage int     `json:"bv"`
	Temperature    int     `json:"tp"`
	RSSI           int     `json:"rs"`
	Tr             int     `json:"tr"`
	Ts             int     `json:"ts"`
	Td             int     `json:"td"`
	Hp             int     `json:"hp"`
	Vp             int     `json:"vp"`
	Tf             int     `json:"tf"`
}
