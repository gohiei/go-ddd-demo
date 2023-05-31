package logger

import (
	entity "cypt/internal/logger/entity"
)

type LogRepository interface {
	WriteAccessLog(entity.AccessLog)
	WritePostLog(entity.PostLog)
}
