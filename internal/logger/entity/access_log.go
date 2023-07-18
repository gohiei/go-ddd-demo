// Package logger provides functionality for logging and event handling related to logging.
package logger

import (
	"cypt/internal/dddcore"
	"encoding/json"
	"fmt"
	"time"
)

// AccessLog represents an access log entry.
type AccessLog struct {
	At            time.Time     `json:"time"`        // Timestamp of the log entry
	Method        string        `json:"method"`      // HTTP method
	Origin        string        `json:"origin"`      // Origin of the request
	HTTPVersion   string        `json:"version"`     // HTTP version
	UserAgent     string        `json:"user_agent"`  // User agent string
	XFF           string        `json:"xff"`         // X-Forwarded-For header value
	StatusCode    int           `json:"status_code"` // HTTP status code
	ContentLength int           `json:"length"`      // Content length of the response
	Latency       time.Duration `json:"latency"`     // Request latency in nanoseconds
	Host          string        `json:"host"`        // Host of the request
	IP            string        `json:"ip"`          // IP address of the client
	Domain        int           `json:"domain"`      // Domain of the request
	RequestID     string        `json:"request_id"`  // Request ID
	SessionID     string        `json:"session_id"`  // Session ID
	FullPath      string        `json:"full_path"`   // API FullPath
	Agent         string        `json:"agent"`       // Agent
}

// String formats the AccessLog as a string.
func (l *AccessLog) String() string {
	return fmt.Sprintf(
		`"%s" %s "%s %s %s" "%s" "%s" %d %d %d %d "%s" "%s" "%s"`,
		l.At.Local().Format(time.RFC3339),
		l.IP,
		l.Method,
		l.Origin,
		l.HTTPVersion,
		l.UserAgent,
		l.XFF,
		l.StatusCode,
		l.ContentLength,
		l.Latency,
		l.Domain,
		l.Host,
		l.RequestID,
		l.FullPath,
	)
}

func (l *AccessLog) JSON() ([]byte, error) {
	b, err := json.Marshal(l)

	if err != nil {
		return nil, dddcore.NewErrorBy(err)
	}

	return b, nil
}
