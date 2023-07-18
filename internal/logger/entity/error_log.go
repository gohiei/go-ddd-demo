package entity

import (
	"time"

	"cypt/internal/dddcore"
)

// ErrorLog represents an error log entry.
type ErrorLog struct {
	At          time.Time     `json:"time"`       // Timestamp of the log entry
	Method      string        `json:"method"`     // HTTP method
	Origin      string        `json:"origin"`     // Origin of the request
	Domain      int           `json:"domain"`     // Domain of the request
	Host        string        `json:"host"`       // Host of the request
	RequestID   string        `json:"request_id"` // Request ID
	IP          string        `json:"ip"`         // IP address of the client
	RequestBody string        `json:"req_body"`   // Request body
	Error       dddcore.Error `json:"error"`      // Error information
}
