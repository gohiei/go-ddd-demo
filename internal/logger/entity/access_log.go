package logger

import (
	"fmt"
	"time"
)

type AccessLog struct {
	At            time.Time `json:"time"`
	Method        string    `json:"method"`
	Origin        string    `json:"origin"`
	HttpVersion   string    `json:"version"`
	UserAgent     string    `json:"user_agent"`
	XFF           string    `json:"xff"`
	StatusCode    int       `json:"status_code"`
	ContentLength int       `json:"length"`
	Latency       int64     `json:"latency"`
	Domain        string    `json:"domain"`
	Host          string    `json:"host"`
	RequestId     string    `json:"request_id"`
	IP            string    `json:"ip"`
}

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
