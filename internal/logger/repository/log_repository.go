package logger

import (
	entity "cypt/internal/logger/entity"
)

// LogRepository defines the interface for a log repository.
type LogRepository interface {
	WriteAccessLog(log *entity.AccessLog)           // Write access log entry
	WritePostLog(log *entity.PostLog)               // Write post log entry
	WriteErrorLog(log *entity.ErrorLog)             // Write error log entry
	WriteHTTPRequestLog(log *entity.HTTPRequestLog) // Write http-request log entry
}
