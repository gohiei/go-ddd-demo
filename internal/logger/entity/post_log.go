package logger

import (
	"fmt"
	"time"
)

type PostLog struct {
	At            time.Time `json:"time"`
	IP            string    `json:"ip"`
	Method        string    `json:"method"`
	Origin        string    `json:"origin"`
	StatusCode    int       `json:"status_code"`
	ContentLength int       `json:"length"`
	Domain        string    `json:"domain"`
	Host          string    `json:"host"`
	RequestId     string    `json:"request_id"`
	RequestBody   string    `json:"request"`
	ResponseData  string    `json:"response"`
}

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
		l.RequestId,
		l.RequestBody,
		l.ResponseData,
	)
}
