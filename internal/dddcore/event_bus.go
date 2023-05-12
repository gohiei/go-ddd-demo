package dddcore

// @todo
type EventBus interface {
	Post(Event)
	PostAll(AggregateRoot)
	Register()
	Unregister()
}
