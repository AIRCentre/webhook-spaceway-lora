package eventrepo

type I interface {
	Insert(EventPayload) error
}
