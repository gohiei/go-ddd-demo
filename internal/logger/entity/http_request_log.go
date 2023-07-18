package logger

import (
	"fmt"
	"net/http"
	"time"
)

type HTTPRequestLog struct {
	At         time.Time     `json:"time"`
	Host       string        `json:"host"`
	Method     string        `json:"method"`
	Origin     string        `json:"origin"`
	ReqHeader  http.Header   `json:"req_header"`
	ReqBody    string        `json:"req_body"`
	StatusCode int           `json:"status_code"`
	Latency    time.Duration `json:"latency"`
	Error      error         `json:"error"`
	ResHeader  http.Header   `json:"res_header"`
	ResBody    string        `json:"res_body"`
}

// String formats the HTTPRequestLog as a string.
func (l *HTTPRequestLog) String() string {
	return fmt.Sprintf(
		`%s "%s %s" %s %s %s %d %d %s %s`,
		l.At.Local().Format(time.RFC3339),
		l.Method,
		l.Origin,
		l.Host,
		l.ReqHeader,
		l.ReqBody,
		l.StatusCode,
		l.Latency,
		l.ResHeader,
		l.ResBody,
	)
}
