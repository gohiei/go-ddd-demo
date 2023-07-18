package logger

import (
	"encoding/json"
	"io"
	"os"
	"path"

	entity "cypt/internal/logger/entity"
	repo "cypt/internal/logger/repository"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

// ZerologLogRepository is an implementation of LogRepository using the Zerolog library.
type ZerologLogRepository struct {
	accessLogger      *zerolog.Logger
	postLogger        *zerolog.Logger
	errorLogger       *zerolog.Logger
	httpRequestLogger *zerolog.Logger
}

var _ repo.LogRepository = (*ZerologLogRepository)(nil)

// NewZerologLogRepository creates a new instance of ZerologLogRepository.
func NewZerologLogRepository(logDir string) *ZerologLogRepository {
	aWriters := io.MultiWriter(
		zerolog.ConsoleWriter{Out: os.Stdout},
		NewRollingLog(logDir, "/access.log"),
	)

	pWriters := io.MultiWriter(
		zerolog.ConsoleWriter{Out: os.Stdout},
		NewRollingLog(logDir, "/post.log"),
	)

	eWriters := io.MultiWriter(
		zerolog.ConsoleWriter{Out: os.Stdout},
		NewRollingLog(logDir, "/error.log"),
	)

	hWriters := io.MultiWriter(
		zerolog.ConsoleWriter{Out: os.Stdout},
		NewRollingLog(logDir, "/http.request.log"),
	)

	aLogger := zerolog.New(aWriters).With().Logger()
	pLogger := zerolog.New(pWriters).With().Logger()
	eLogger := zerolog.New(eWriters).With().Logger()
	hLogger := zerolog.New(hWriters).With().Logger()

	return &ZerologLogRepository{
		accessLogger:      &aLogger,
		postLogger:        &pLogger,
		errorLogger:       &eLogger,
		httpRequestLogger: &hLogger,
	}
}

// WriteAccessLog writes the access log to the appropriate logger.
func (r *ZerologLogRepository) WriteAccessLog(log *entity.AccessLog) {
	b, _ := json.Marshal(log)
	r.accessLogger.Info().RawJSON("log", b).Msg("")
}

// WritePostLog writes the post log to the appropriate logger.
func (r *ZerologLogRepository) WritePostLog(log *entity.PostLog) {
	b, _ := json.Marshal(log)
	r.postLogger.Info().RawJSON("log", b).Msg("")
}

// WriteErrorLog writes the error log to the appropriate logger.
func (r *ZerologLogRepository) WriteErrorLog(log *entity.ErrorLog) {
	b, _ := json.Marshal(log)
	r.errorLogger.Info().RawJSON("log", b).Msg("")
}

// WriteHTTPRequestLog writes the http-request log to the appropriate logger.
func (r *ZerologLogRepository) WriteHTTPRequestLog(log *entity.HTTPRequestLog) {
	b, _ := json.Marshal(log)
	r.httpRequestLogger.Info().RawJSON("log", b).Msg("")
}

// NewRollingLog creates a new lumberjack Logger for rolling log files.
func NewRollingLog(logDir string, filename string) io.Writer {
	return &lumberjack.Logger{
		Filename:   path.Join(logDir, filename),
		MaxBackups: 100,
		MaxSize:    100, // megabytes
		MaxAge:     30,  // days
	}
}
