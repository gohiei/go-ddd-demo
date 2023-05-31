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

	// request
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

	// response
	StatusCode    int    `json:"status_code"`
	ContentLength int    `json:"content_log"`
	Latency       int64  `json:"latency"`
	ResponseData  string `json:"response_data"`
}

var _ dddcore.Event = (*ErrorRaisedEvent)(nil)

func NewErrorRaisedEvent(occurredAt time.Time, clientIp string, req *http.Request, res *ErrorRaisedEventResponse) *ErrorRaisedEvent {
	return &ErrorRaisedEvent{
		BaseEvent:   dddcore.NewEvent(ErrorRaisedEventName),
		At:          occurredAt,
		UserAgent:   req.Header.Get("User-Agent"),
		XFF:         req.Header.Get("X-Forwarded-For"),
		RequestId:   req.Header.Get("X-Request-Id"),
		Host:        req.Host,
		Domain:      req.Header.Get("domain"),
		IP:          clientIp,
		Method:      req.Method,
		Origin:      req.RequestURI,
		HttpVersion: req.Proto,
		Refer:       req.Header.Get("Referer"),
		RequestBody: req.PostForm.Encode(),

		StatusCode:    res.StatusCode,
		ContentLength: res.ContentLength,
		Latency:       res.Latency.Milliseconds(),
		ResponseData:  res.ResponseData,
	}
}
