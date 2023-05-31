package logger

import (
	"fmt"
	"time"
)

type ErrorLog struct {
	At           time.Time
	IP           string
	Method       string
	Origin       string
	Domain       int
	Host         string
	RequestId    string
	ErrorMessage string
	ErrorDetail  string
}

func (l ErrorLog) String() string {
	return fmt.Sprintf(
		`%s %s "%s %s" %d %s %s %s %s`,
		l.At.Local().Format(time.RFC3339),
		l.IP,
		l.Method,
		l.Origin,
		l.Domain,
		l.Host,
		l.RequestId,
		l.ErrorMessage,
		l.ErrorDetail,
	)
}
