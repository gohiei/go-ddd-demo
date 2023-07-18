// Package dddcore provides core functionality for Domain-Driven Design (DDD) concepts.
package dddcore

import "time"

// BaseEvent is a base struct for representing events.
type BaseEvent struct {
	ID         string    `json:"id"`
	OccurredOn time.Time `json:"occurred_on"`
	Name       string    `json:"name"`
}

// GetID returns the ID of the event.
func (e *BaseEvent) GetID() string {
	return e.ID
}

// GetName returns the name of the event.
func (e *BaseEvent) GetName() string {
	return e.Name
}

// GetOccurredOn returns the timestamp when the event occurred.
func (e *BaseEvent) GetOccurredOn() time.Time {
	return e.OccurredOn
}

// NewEvent creates a new instance of the BaseEvent with the given name and generates a unique ID and timestamp.
func NewEvent(name string) *BaseEvent {
	return &BaseEvent{
		ID:         NewUUID().String(),
		Name:       name,
		OccurredOn: time.Now(),
	}
}

// Event is an interface that represents an event.
type Event interface {
	GetID() string
	GetName() string
	GetOccurredOn() time.Time
}
