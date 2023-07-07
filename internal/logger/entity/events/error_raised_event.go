package logger

import (
	"net/http"
	"time"

	"cypt/internal/dddcore"
)

// ErrorRaisedEventName represents the event name for an error being raised
const (
	ErrorRaisedEventName = "error.raised"
)

// ErrorRaisedEventResponse represents the response data for an error raised event
type ErrorRaisedEventResponse struct {
	Latency       time.Duration
	StatusCode    int
	ContentLength int
	ResponseData  string
}

// ErrorRaisedEvent represents the event for an error being raised
type ErrorRaisedEvent struct {
	*dddcore.BaseEvent

	// Event information
	At          time.Time `json:"at"`
	IP          string    `json:"ip"`
	RequestID   string    `json:"request_id"`
	Host        string    `json:"host"`
	Domain      string    `json:"domain"`
	Method      string    `json:"method"`
	Origin      string    `json:"origin"`
	RequestBody string    `json:"request_body"`
	Error       error     `json:"error"`
}

var _ dddcore.Event = (*ErrorRaisedEvent)(nil)

// NewErrorRaisedEvent creates a new ErrorRaisedEvent
func NewErrorRaisedEvent(occurredAt time.Time, clientIP string, req *http.Request, err dddcore.Error) *ErrorRaisedEvent {
	return &ErrorRaisedEvent{
		BaseEvent:   dddcore.NewEvent(ErrorRaisedEventName),
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
