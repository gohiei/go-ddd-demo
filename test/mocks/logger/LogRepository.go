// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	entity "cypt/internal/logger/entity"

	mock "github.com/stretchr/testify/mock"
)

// LogRepository is an autogenerated mock type for the LogRepository type
type LogRepository struct {
	mock.Mock
}

// WriteAccessLog provides a mock function with given fields: log
func (_m *LogRepository) WriteAccessLog(log *entity.AccessLog) {
	_m.Called(log)
}

// WriteErrorLog provides a mock function with given fields: log
func (_m *LogRepository) WriteErrorLog(log *entity.ErrorLog) {
	_m.Called(log)
}

// WriteHTTPRequestLog provides a mock function with given fields: log
func (_m *LogRepository) WriteHTTPRequestLog(log *entity.HTTPRequestLog) {
	_m.Called(log)
}

// WritePostLog provides a mock function with given fields: log
func (_m *LogRepository) WritePostLog(log *entity.PostLog) {
	_m.Called(log)
}

// NewLogRepository creates a new instance of LogRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLogRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *LogRepository {
	mock := &LogRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
