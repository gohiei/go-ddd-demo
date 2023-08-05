package adapter

import (
	"encoding/json"

	"cypt/internal/dddcore"
)

// TestEventBus is a simple test implementation of the EventBus interface.
type TestEventBus struct {
	handlers map[string][]dddcore.EventHandler
}

var _ dddcore.EventBus = (*TestEventBus)(nil)

// Post publishes an event to the event bus.
func (b *TestEventBus) Post(e dddcore.Event) error {
	name := e.GetName()

	handlers, ok := b.handlers[name]

	if !ok || len(handlers) == 0 {
		return nil
	}

	jsonData, err := json.Marshal(e)

	if err != nil {
		return err
	}

	for _, handler := range handlers {
		_ = handler.When(e.GetName(), jsonData)
	}

	return nil
}

// PostAll publishes all the domain events of an aggregate root to the event bus.
func (b *TestEventBus) PostAll(ar dddcore.AggregateRoot) error {
	for _, e := range ar.GetDomainEvents() {
		b.Post(e)
	}

	return nil
}

// Register registers an event handler with the event bus.
func (b *TestEventBus) Register(h dddcore.EventHandler) error {
	name := h.EventName()
	_, ok := b.handlers[name]

	if !ok {
		b.handlers[name] = make([]dddcore.EventHandler, 0, 10)
	}

	b.handlers[name] = append(b.handlers[name], h)

	return nil
}

// Unregister unregisters an event handler from the event bus.
func (b *TestEventBus) Unregister(h dddcore.EventHandler) error {
	name := h.EventName()
	handlers, ok := b.handlers[name]
	if !ok {
		return nil
	}

	var index = -1
	for i, handler := range handlers {
		if handler == h {
			index = i
			break
		}
	}

	if index >= 0 {
		b.handlers[name] = append(handlers[:index], handlers[index+1:]...)
	}

	return nil
}

// NewTestEventBus creates a new instance of TestEventBus.
func NewTestEventBus() TestEventBus {
	return TestEventBus{
		handlers: make(map[string][]dddcore.EventHandler),
	}
}
