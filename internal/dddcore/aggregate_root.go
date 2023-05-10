package dddcore

type AggregateRoot struct {
	events []IEvent
}

func (ar *AggregateRoot) AddDomainEvent(event IEvent) {
	if ar.events == nil {
		ar.events = make([]IEvent, 10)
	}

	ar.events = append(ar.events, event)
}
