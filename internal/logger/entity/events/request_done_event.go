package logger

import (
	"net/http"
	"time"

	"cypt/internal/dddcore"
)

// RequestDoneEventName represents the event name for a request being done
const (
	RequestDoneEventName = "request.done"
)

// RequestDoneEventResponse represents the response data for a request done event
type RequestDoneEventResponse struct {
	Latency       time.Duration
	StatusCode    int
	ContentLength int
	ResponseData  string
}

// RequestDoneEvent represents the event for a request being done
type RequestDoneEvent struct {
	*dddcore.BaseEvent

	// Request information
	At          time.Time `json:"at"`
	UserAgent   string    `json:"user_agent"`
	XFF         string    `json:"x_forwarded_for"`
	RequestId   string    `json:"request_id"`
	Host        string    `json:"host"`
	Domain      string    `json:"domain"`
	IP          string    `json:"ip"`
	Method      string    `json:"method"`
	Origin      string    `json:"origin"`
	HttpVersion string    `json:"http_version"`
	RequestBody string    `json:"request_body"`
	Refer       string    `json:"refer"`

	// Response information
	StatusCode    int    `json:"status_code"`
	ContentLength int    `json:"content_length"`
	Latency       int64  `json:"latency"`
	ResponseData  string `json:"response_data"`
}

var _ dddcore.Event = (*RequestDoneEvent)(nil)

// NewRequestDoneEvent creates a new RequestDoneEvent
func NewRequestDoneEvent(occurredAt time.Time, clientIP string, req *http.Request, res *RequestDoneEventResponse) *RequestDoneEvent {
	return &RequestDoneEvent{
		BaseEvent:     dddcore.NewEvent(RequestDoneEventName),
		At:            occurredAt,
		IP:            clientIP,
		UserAgent:     req.Header.Get("User-Agent"),
		XFF:           req.Header.Get("X-Forwarded-For"),
		RequestId:     req.Header.Get("X-Request-Id"),
		Host:          req.Host,
		Domain:        req.Header.Get("domain"),
		Method:        req.Method,
		Origin:        req.RequestURI,
		HttpVersion:   req.Proto,
		Refer:         req.Header.Get("Referer"),
		RequestBody:   req.PostForm.Encode(),
		StatusCode:    res.StatusCode,
		ContentLength: res.ContentLength,
		Latency:       res.Latency.Milliseconds(),
		ResponseData:  res.ResponseData,
	}
}
