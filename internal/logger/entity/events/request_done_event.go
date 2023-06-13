package logger

import (
	"net/http"
	"time"

	"cypt/internal/dddcore"
)

const (
	RequestDoneEventName = "request.done"
)

type RequestDoneEventResponse struct {
	Latency       time.Duration
	StatusCode    int
	ContentLength int
	ResponseData  string
}

type RequestDoneEvent struct {
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

var _ dddcore.Event = (*RequestDoneEvent)(nil)

func NewRequestDoneEvent(occurredAt time.Time, clientIp string, req *http.Request, res *RequestDoneEventResponse) *RequestDoneEvent {
	return &RequestDoneEvent{
		BaseEvent: dddcore.NewEvent(RequestDoneEventName),
		At:        occurredAt,
		IP:        clientIp,

		UserAgent:   req.Header.Get("User-Agent"),
		XFF:         req.Header.Get("X-Forwarded-For"),
		RequestId:   req.Header.Get("X-Request-Id"),
		Host:        req.Host,
		Domain:      req.Header.Get("domain"),
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
