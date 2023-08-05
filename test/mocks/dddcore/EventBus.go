// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	dddcore "cypt/internal/dddcore"

	mock "github.com/stretchr/testify/mock"
)

// EventBus is an autogenerated mock type for the EventBus type
type EventBus struct {
	mock.Mock
}

// Post provides a mock function with given fields: ev
func (_m *EventBus) Post(ev dddcore.Event) error {
	ret := _m.Called(ev)

	var r0 error
	if rf, ok := ret.Get(0).(func(dddcore.Event) error); ok {
		r0 = rf(ev)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PostAll provides a mock function with given fields: ar
func (_m *EventBus) PostAll(ar dddcore.AggregateRoot) error {
	ret := _m.Called(ar)

	var r0 error
	if rf, ok := ret.Get(0).(func(dddcore.AggregateRoot) error); ok {
		r0 = rf(ar)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Register provides a mock function with given fields: h
func (_m *EventBus) Register(h dddcore.EventHandler) error {
	ret := _m.Called(h)

	var r0 error
	if rf, ok := ret.Get(0).(func(dddcore.EventHandler) error); ok {
		r0 = rf(h)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unregister provides a mock function with given fields: h
func (_m *EventBus) Unregister(h dddcore.EventHandler) error {
	ret := _m.Called(h)

	var r0 error
	if rf, ok := ret.Get(0).(func(dddcore.EventHandler) error); ok {
		r0 = rf(h)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewEventBus creates a new instance of EventBus. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEventBus(t interface {
	mock.TestingT
	Cleanup(func())
}) *EventBus {
	mock := &EventBus{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
