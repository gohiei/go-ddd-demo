package events

import (
	"net/http"
	"time"

	"cypt/internal/dddcore"
	"cypt/internal/logger/entity"
)

const (
	HTTPRequestDoneEventName = "http_request.done"
)

// HTTPRequestDoneEvent represents an event for a completed HTTP request.
type HTTPRequestDoneEvent struct {
	*dddcore.BaseEvent
	At         time.Time     `json:"time"`        // Timestamp of the request
	Method     string        `json:"method"`      // HTTP method (e.g., GET, POST)
	Origin     string        `json:"origin"`      // Request URL
	Host       string        `json:"host"`        // Host of the request
	ReqHeader  http.Header   `json:"req_header"`  // Request headers
	ReqBody    string        `json:"req_body"`    // Request body
	StatusCode int           `json:"status_code"` // Response status code
	Latency    time.Duration `json:"latency"`     // Request latency
	Error      error         `json:"error"`       // Error, if any, encountered during the request
	ResHeader  http.Header   `json:"res_header"`  // Response headers
	ResBody    string        `json:"res_body"`    // Response body
}

// Ensure HTTPRequestDoneEvent implements the dddcore.Event interface
var _ dddcore.Event = (*HTTPRequestDoneEvent)(nil)

// NewHTTPRequestDoneEvent creates a new HTTPRequestDoneEvent based on the given HTTPRequestLog.
func NewHTTPRequestDoneEvent(log *entity.HTTPRequestLog) *HTTPRequestDoneEvent {
	return &HTTPRequestDoneEvent{
		BaseEvent:  dddcore.NewEvent(HTTPRequestDoneEventName), // Create a new base event with the given event name
		At:         log.At,
		Method:     log.Method,
		Origin:     log.Origin,
		Host:       log.Host,
		ReqHeader:  log.ReqHeader,
		ReqBody:    log.ReqBody,
		StatusCode: log.StatusCode,
		Latency:    log.Latency,
		Error:      log.Error,
		ResHeader:  log.ResHeader,
		ResBody:    log.ResBody,
	}
}
