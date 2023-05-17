package dddcore

import "time"

type BaseEvent struct {
	ID         string    `json:"id"`
	OccurredOn time.Time `json:"occurred_on"`
	Name       string    `json:"name"`
}

func (e *BaseEvent) GetID() string {
	return e.ID
}

func (e *BaseEvent) GetName() string {
	return e.Name
}

func (e *BaseEvent) GetOccurredOn() time.Time {
	return e.OccurredOn
}

func NewEvent(name string) *BaseEvent {
	return &BaseEvent{
		ID:         NewUUID().String(),
		Name:       name,
		OccurredOn: time.Now(),
	}
}

type Event interface {
	GetID() string
	GetName() string
	GetOccurredOn() time.Time
}
