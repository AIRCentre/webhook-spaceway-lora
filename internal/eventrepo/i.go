package eventrepo

type I interface {
	Insert(deviceId string, payload EventPayload) error
}
