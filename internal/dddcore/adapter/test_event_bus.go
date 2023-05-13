package dddcore

import (
	"cypt/internal/dddcore"
	"encoding/json"
)

type TestEventBus struct {
	handlers map[string][]dddcore.Handler
}

func (b *TestEventBus) Post(e dddcore.Event) {
	name := e.GetName()

	handlers, ok := b.handlers[name]

	if !ok || len(handlers) == 0 {
		return
	}

	jsonData, err := json.Marshal(e)

	if err != nil {
		return
	}

	for _, handler := range handlers {
		handler.Handle(e.GetName(), jsonData)
	}
}

func (b *TestEventBus) PostAll(ar dddcore.AggregateRoot) {
	for _, e := range ar.GetDomainEvents() {
		b.Post(e)
	}
}

func (b *TestEventBus) Register(h dddcore.Handler) {
	name := h.EventName()
	_, ok := b.handlers[name]

	if !ok {
		b.handlers[name] = make([]dddcore.Handler, 0, 10)
	}

	b.handlers[name] = append(b.handlers[name], h)
}

func (b *TestEventBus) Unregister(h dddcore.Handler) {
	name := h.EventName()
	handlers, ok := b.handlers[name]

	if !ok {
		return
	}

	var index = 0
	for i, handler := range handlers {
		if handler == h {
			index = i
			break
		}
	}

	b.handlers[name] = append(b.handlers[name][:index], b.handlers[name][index+1:]...)
}

var _ dddcore.EventBus = (*TestEventBus)(nil)

func NewTestEventBus() TestEventBus {
	return TestEventBus{
		handlers: make(map[string][]dddcore.Handler),
	}
}
