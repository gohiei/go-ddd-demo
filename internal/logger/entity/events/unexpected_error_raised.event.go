package events

import (
	"net/http"
	"time"

	"cypt/internal/dddcore"
)

// UnexpectedErrorRaisedEventName represents the event name for an unexpected error being raised
const (
	UnexpectedErrorRaisedEventName = "unexpected_error.raised"
)

// UnexpectedErrorRaisedEventResponse represents the response data for an unexpected error raised event
type UnexpectedErrorRaisedEventResponse struct {
	Latency       time.Duration
	StatusCode    int
	ContentLength int
	ResponseData  string
}

// UnexpectedErrorRaisedEvent represents the event for an error being raised
type UnexpectedErrorRaisedEvent struct {
	*dddcore.BaseEvent

	// Event information
	At          time.Time `json:"time"`
	IP          string    `json:"ip"`
	RequestID   string    `json:"request_id"`
	Host        string    `json:"host"`
	Domain      string    `json:"domain"`
	Method      string    `json:"method"`
	Origin      string    `json:"origin"`
	RequestBody string    `json:"req_body"`
	Error       error     `json:"error"`
}

var _ dddcore.Event = (*UnexpectedErrorRaisedEvent)(nil)

// NewUnexpectedErrorRaisedEvent creates a new ErrorRaisedEvent
func NewUnexpectedErrorRaisedEvent(occurredAt time.Time, clientIP string, req *http.Request, err dddcore.Error) *UnexpectedErrorRaisedEvent {
	return &UnexpectedErrorRaisedEvent{
		BaseEvent:   dddcore.NewEvent(UnexpectedErrorRaisedEventName),
		At:          occurredAt,
		IP:          clientIP,
		RequestID:   req.Header.Get("X-Request-Id"),
		Host:        req.Host,
		Domain:      req.Header.Get("domain"),
		Method:      req.Method,
		Origin:      req.RequestURI,
		RequestBody: req.PostForm.Encode(),

		Error: err,
	}
}
