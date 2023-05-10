package dddcore

// @todo
type EventBus interface {
	Post(Event)
	PostAll([]Event)
	Register()
	Unregister()
}
