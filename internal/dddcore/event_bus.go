package dddcore

// @todo
type EventBus interface {
	Post(ev Event)
	PostAll(ar AggregateRoot)
	Register(h Handler)
	Unregister(h Handler)
}

type Handler interface {
	Name() string
	EventName() string
	Handle(name string, message []byte)
}
