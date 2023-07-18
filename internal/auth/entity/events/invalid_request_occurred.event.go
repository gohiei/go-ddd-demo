// Package events provides event structures and functions related to authentication functionality.
package events

import "cypt/internal/dddcore"

const (
	InvalidRequestOccurredEventName = "invalid_request.occurred"
)

// InvalidRequestOccurredEvent represents an event that occurs when an invalid request is detected.
type InvalidRequestOccurredEvent struct {
	dddcore.Event
	IP     string `json:"ip"`
	XFF    string `json:"xff"`
	Token  string `json:"token"`
	Method string `json:"method"`
	URL    string `json:"url"`
	Error  string `json:"error"`
}

var _ dddcore.Event = (*InvalidRequestOccurredEvent)(nil)

// NewInvalidRequestOccurredEvent creates a new instance of the InvalidRequestOccurredEvent.
func NewInvalidRequestOccurredEvent(token, method, url, ip, xff string, err error) *InvalidRequestOccurredEvent {
	errStr := "-"

	if err != nil {
		errStr = err.Error()
	}

	return &InvalidRequestOccurredEvent{
		Event:  dddcore.NewEvent(InvalidRequestOccurredEventName),
		IP:     ip,
		XFF:    xff,
		Token:  token,
		Method: method,
		URL:    url,
		Error:  errStr,
	}
}
