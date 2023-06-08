package logger

import (
	"cypt/internal/dddcore"
	"time"
)

type ErrorLog struct {
	At          time.Time     `json:"time"`
	Method      string        `json:"method"`
	Origin      string        `json:"origin"`
	Domain      string        `json:"domain"`
	Host        string        `json:"host"`
	RequestId   string        `json:"request_id"`
	IP          string        `json:"ip"`
	RequestBody string        `json:"request"`
	Error       dddcore.Error `json:"error"`
}
