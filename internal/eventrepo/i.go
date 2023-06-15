package eventrepo

type I interface {
	Insert(payload SwarmPayload) error
}
