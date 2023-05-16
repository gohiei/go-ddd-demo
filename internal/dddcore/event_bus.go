package dddcore

// @todo
type EventBus interface {
	Post(ev Event)
	PostAll(ar AggregateRoot)
	Register(h EventHandler)
	Unregister(h EventHandler)
}

type EventHandler interface {
	Name() string
	EventName() string
	When(name string, message []byte)
}
