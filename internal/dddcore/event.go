package dddcore

import "time"

type IEvent interface {
	GetId() string
	GetOccurredOn() time.Time
	GetName() string
}

type Event struct {
	Id         string    `json:"id"`
	OccurredOn time.Time `json:"occurred_on"`
	Name       string    `json:"name"`
}

func NewEvent(name string) *Event {
	return &Event{
		Id:         NewUUID().String(),
		OccurredOn: time.Now(),
		Name:       name,
	}
}

func (e *Event) GetId() string {
	return e.Id
}

func (e *Event) GetOccurredOn() time.Time {
	return e.OccurredOn
}

func (e *Event) GetName() string {
	return e.Name
}
