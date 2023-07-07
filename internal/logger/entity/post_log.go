package logger

import (
	"fmt"
	"time"
)

// PostLog represents a log entry for a POST request.
type PostLog struct {
	At            time.Time `json:"time"`        // Timestamp of the log entry
	IP            string    `json:"ip"`          // IP address of the client
	Method        string    `json:"method"`      // HTTP method
	Origin        string    `json:"origin"`      // Origin of the request
	StatusCode    int       `json:"status_code"` // HTTP status code
	ContentLength int       `json:"length"`      // Length of the response content
	Domain        string    `json:"domain"`      // Domain of the request
	Host          string    `json:"host"`        // Host of the request
	RequestID     string    `json:"request_id"`  // Request ID
	RequestBody   string    `json:"request"`     // Request body
	ResponseData  string    `json:"response"`    // Response data
}

// String returns a formatted string representation of the PostLog.
func (l PostLog) String() string {
	return fmt.Sprintf(
		`%s %s "%s %s" %d %d %s "%s" "%s" "%s" %s`,
		l.At.Local().Format(time.RFC3339),
		l.IP,
		l.Method,
		l.Origin,
		l.StatusCode,
		l.ContentLength,
		l.Domain,
		l.Host,
		l.RequestID,
		l.RequestBody,
		l.ResponseData,
	)
}
