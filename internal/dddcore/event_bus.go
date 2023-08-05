package dddcore

// EventBus represents an event bus that facilitates event publishing and subscription.
type EventBus interface {
	Post(ev Event) error
	PostAll(ar AggregateRoot) error
	Register(h EventHandler) error
	Unregister(h EventHandler) error
}

// EventHandler represents a handler for events.
type EventHandler interface {
	Name() string
	EventName() string
	When(name string, message []byte) error
}
