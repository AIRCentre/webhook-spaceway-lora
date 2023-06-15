package events_repository

type I interface {
	Insert(payload SwarmPayload) error
}
