package auth

import "cypt/internal/dddcore"

const (
	InvalidRequestOccurredEventName = "invalid_request.occurred"
)

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
