package dddcore

// AggregateRoot is an interface for entities that persist state changes
type AggregateRoot interface {
	AddDomainEvent(Event)
	GetDomainEvents() []Event
}

// BaseAggregateRoot is a struct that provides a base implementation of the AggregateRoot interface
type BaseAggregateRoot struct {
	events []Event
}

// AddDomainEvent adds a domain event to the aggregate root
func (ar *BaseAggregateRoot) AddDomainEvent(event Event) {
	if ar.events == nil {
		ar.events = make([]Event, 0, 10)
	}

	ar.events = append(ar.events, event)
}

// GetDomainEvents returns all domain events collected by the aggregate root
func (ar *BaseAggregateRoot) GetDomainEvents() []Event {
	return ar.events
}

// NewAggregateRoot creates a new instance of BaseAggregateRoot
func NewAggregateRoot() *BaseAggregateRoot {
	return &BaseAggregateRoot{}
}
