package dddcore

import (
	"cypt/internal/dddcore"
	"fmt"
)

type TestEventBus struct {
}

func (b *TestEventBus) Post(e dddcore.Event) {
	fmt.Println(e.GetName())
}

func (b *TestEventBus) PostAll(ar dddcore.AggregateRoot) {
	for _, e := range ar.GetDomainEvents() {
		b.Post(e)
	}
}

func (b *TestEventBus) Register()   {}
func (b *TestEventBus) Unregister() {}

var _ dddcore.EventBus = (*TestEventBus)(nil)

func NewTestEventBus() TestEventBus {
	return TestEventBus{}
}
