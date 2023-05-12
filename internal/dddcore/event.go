package dddcore

import "time"

type BaseEvent struct {
	Id         string    `json:"id"`
	OccurredOn time.Time `json:"occurred_on"`
	Name       string    `json:"name"`
}

func (e *BaseEvent) GetId() string {
	return e.Id
}

func (e *BaseEvent) GetName() string {
	return e.Name
}

func (e *BaseEvent) GetOccurredOn() time.Time {
	return e.OccurredOn
}

func NewEvent(name string) *BaseEvent {
	return &BaseEvent{
		Id:         NewUUID().String(),
		Name:       name,
		OccurredOn: time.Now(),
	}
}

type Event interface {
	GetId() string
	GetName() string
	GetOccurredOn() time.Time
}
