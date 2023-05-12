package dddcore

type AggregateRoot interface {
	AddDomainEvent(Event)
	GetDomainEvents() []Event
}
type BaseAggregateRoot struct {
	events []Event
}

func (ar *BaseAggregateRoot) AddDomainEvent(event Event) {
	if ar.events == nil {
		ar.events = make([]Event, 0, 10)
	}

	ar.events = append(ar.events, event)
}

func (ar *BaseAggregateRoot) GetDomainEvents() []Event {
	return ar.events
}

func NewAggregateRoot() *BaseAggregateRoot {
	return &BaseAggregateRoot{}
}
