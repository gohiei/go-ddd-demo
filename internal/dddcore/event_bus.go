package dddcore

// EventBus represents an event bus that facilitates event publishing and subscription.
type EventBus interface {
	Post(ev Event)
	PostAll(ar AggregateRoot)
	Register(h EventHandler)
	Unregister(h EventHandler)
}

// EventHandler represents a handler for events.
type EventHandler interface {
	Name() string
	EventName() string
	When(name string, message []byte) error
}
