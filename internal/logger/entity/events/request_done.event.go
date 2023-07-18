// Package events provides functionality for logging and event handling related to logging.
package events

import (
	"fmt"
	"net/http"
	"strconv"
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
	At          time.Time `json:"time"`
	UserAgent   string    `json:"user_agent"`
	XFF         string    `json:"xff"`
	RequestID   string    `json:"request_id"`
	Host        string    `json:"host"`
	Domain      int       `json:"domain"`
	IP          string    `json:"ip"`
	Method      string    `json:"method"`
	Origin      string    `json:"origin"`
	HTTPVersion string    `json:"version"`
	RequestBody string    `json:"req_body"`
	Refer       string    `json:"refer"`
	FullPath    string    `json:"full_path"`
	SessionID   string    `json:"session_id"`
	Agent       string    `json:"agent"`

	// Response information
	StatusCode    int    `json:"status_code"`
	ContentLength int    `json:"content_length"`
	Latency       int64  `json:"latency"`
	ResponseBody  string `json:"res_body"`
}

var _ dddcore.Event = (*RequestDoneEvent)(nil)

// NewRequestDoneEvent creates a new RequestDoneEvent
func NewRequestDoneEvent(
	occurredAt time.Time,
	clientIP string,
	fullPath string,
	req *http.Request,
	res *RequestDoneEventResponse,
) *RequestDoneEvent {

	domain, err := strconv.Atoi(req.Header.Get("Domain"))

	if err != nil {
		domain = 0
	}

	return &RequestDoneEvent{
		BaseEvent:     dddcore.NewEvent(RequestDoneEventName),
		At:            occurredAt,
		IP:            clientIP,
		UserAgent:     req.Header.Get("User-Agent"),
		XFF:           req.Header.Get("X-Forwarded-For"),
		RequestID:     req.Header.Get("X-Request-Id"),
		Host:          req.Host,
		Domain:        domain,
		Method:        req.Method,
		Origin:        req.RequestURI,
		HTTPVersion:   req.Proto,
		Refer:         req.Header.Get("Referer"),
		RequestBody:   req.PostForm.Encode(),
		StatusCode:    res.StatusCode,
		ContentLength: res.ContentLength,
		Latency:       res.Latency.Milliseconds(),
		FullPath:      fmt.Sprintf("%s %s", req.Method, fullPath),
		ResponseBody:  res.ResponseData,
		SessionID:     req.Header.Get("Session-Id"),
		Agent:         req.Header.Get("X-Agent"),
	}
}
