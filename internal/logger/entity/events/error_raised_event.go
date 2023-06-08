package logger

import (
	"cypt/internal/dddcore"
	"net/http"
	"time"
)

const (
	ErrorRaisedEventName = "error.raised"
)

type ErrorRaisedEventResponse struct {
	Latency       time.Duration
	StatusCode    int
	ContentLength int
	ResponseData  string
}

type ErrorRaisedEvent struct {
	*dddcore.BaseEvent

	At          time.Time `json:"at"`
	IP          string    `json:"ip"`
	RequestId   string    `json:"request_id"`
	Host        string    `json:"host"`
	Domain      string    `json:"domain"`
	Method      string    `json:"method"`
	Origin      string    `json:"origin"`
	RequestBody string    `json:"request_body"`
	Error       error     `json:"error"`
}

var _ dddcore.Event = (*ErrorRaisedEvent)(nil)

func NewErrorRaisedEvent(
	occurredAt time.Time,
	clientIp string,
	req *http.Request,
	err dddcore.Error,
) *ErrorRaisedEvent {
	return &ErrorRaisedEvent{
		BaseEvent: dddcore.NewEvent(ErrorRaisedEventName),
		At:        occurredAt,
		IP:        clientIp,

		RequestId:   req.Header.Get("X-Request-Id"),
		Host:        req.Host,
		Domain:      req.Header.Get("domain"),
		Method:      req.Method,
		Origin:      req.RequestURI,
		RequestBody: req.PostForm.Encode(),

		Error: err,
	}
}
