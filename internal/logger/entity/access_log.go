package logger

import (
	"fmt"
	"time"
)

// AccessLog represents an access log entry.
type AccessLog struct {
	At            time.Time `json:"time"`        // Timestamp of the log entry
	Method        string    `json:"method"`      // HTTP method
	Origin        string    `json:"origin"`      // Origin of the request
	HttpVersion   string    `json:"version"`     // HTTP version
	UserAgent     string    `json:"user_agent"`  // User agent string
	XFF           string    `json:"xff"`         // X-Forwarded-For header value
	StatusCode    int       `json:"status_code"` // HTTP status code
	ContentLength int       `json:"length"`      // Content length of the response
	Latency       int64     `json:"latency"`     // Request latency in nanoseconds
	Domain        string    `json:"domain"`      // Domain of the request
	Host          string    `json:"host"`        // Host of the request
	RequestId     string    `json:"request_id"`  // Request ID
	IP            string    `json:"ip"`          // IP address of the client
}

// String formats the AccessLog as a string.
func (l AccessLog) String() string {
	return fmt.Sprintf(
		`"%s" %s "%s %s %s" "%s" "%s" %d %d %d %s "%s" "%s"`,
		l.At.Local().Format(time.RFC3339),
		l.IP,
		l.Method,
		l.Origin,
		l.HttpVersion,
		l.UserAgent,
		l.XFF,
		l.StatusCode,
		l.ContentLength,
		l.Latency,
		l.Domain,
		l.Host,
		l.RequestId,
	)
}
