// Package adapter provides an implementation of the LogRepository interface
package adapter

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"cypt/internal/logger/entity"
	"cypt/internal/logger/repository"
)

var _ repository.LogRepository = (*ZapLogRepository)(nil)

type ZapLogRepository struct {
	accessLogger      *zap.Logger
	postLogger        *zap.Logger
	errorLogger       *zap.Logger
	httpRequestLogger *zap.Logger
}

func NewLogger(filepath string) *zap.Logger {
	cfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Encoding:          "json",
		OutputPaths:       []string{filepath},
		DisableCaller:     true,
		DisableStacktrace: true,
		EncoderConfig: zapcore.EncoderConfig{
			EncodeTime: zapcore.TimeEncoderOfLayout("2006-01-02T15:04:06"),
		},
	}

	return zap.Must(cfg.Build())
}

func NewZapLogRepository(logDir string) *ZapLogRepository {
	aLogger := NewLogger(logDir + "/access.log")
	defer aLogger.Sync()

	pLogger := NewLogger(logDir + "/post.log")
	defer pLogger.Sync()

	eLogger := NewLogger(logDir + "/error.log")
	defer eLogger.Sync()

	hLogger := NewLogger(logDir + "/http.request.log")
	defer hLogger.Sync()

	return &ZapLogRepository{
		accessLogger:      aLogger,
		postLogger:        pLogger,
		errorLogger:       eLogger,
		httpRequestLogger: hLogger,
	}
}

func (lr *ZapLogRepository) WriteAccessLog(log *entity.AccessLog) {
	lr.accessLogger.Info(
		"access log",
		zap.Time("time", log.At),
		zap.String("method", log.Method),
		zap.String("origin", log.Origin),
		zap.String("version", log.HTTPVersion),
		zap.String("user_agent", log.UserAgent),
		zap.String("xff", log.XFF),
		zap.Int("status_code", log.StatusCode),
		zap.Int("length", log.ContentLength),
		zap.Duration("latency", log.Latency),
		zap.String("host", log.Host),
		zap.String("ip", log.IP),
		zap.Int("domain", log.Domain),
		zap.String("request_id", log.RequestID),
		zap.String("session_id", log.SessionID),
		zap.String("full_path", log.FullPath),
		zap.String("agent", log.Agent),
	)
}

func (lr *ZapLogRepository) WritePostLog(log *entity.PostLog) {
	lr.postLogger.Info(
		"post log",
		zap.Time("time", log.At),
		zap.String("ip", log.IP),
		zap.String("method", log.Method),
		zap.String("origin", log.Origin),
		zap.Int("status_code", log.StatusCode),
		zap.Int("length", log.ContentLength),
		zap.String("host", log.Host),
		zap.Int("domain", log.Domain),
		zap.String("request_id", log.RequestID),
		zap.String("req_body", log.RequestBody),
		zap.String("res_body", log.ResponseBody),
	)
}

func (lr *ZapLogRepository) WriteErrorLog(log *entity.ErrorLog) {
	lr.errorLogger.Info(
		"error log",
		zap.Time("time", log.At),
		zap.String("method", log.Method),
		zap.String("origin", log.Origin),
		zap.Int("domain", log.Domain),
		zap.String("host", log.Host),
		zap.String("request_id", log.RequestID),
		zap.String("ip", log.IP),
		zap.String("req_body", log.RequestBody),
		zap.Object("error", &LogError{
			Code:       log.Error.Code,
			Message:    log.Error.Message,
			StatusCode: log.Error.StatusCode,
			Detail:     log.Error.Detail,
			Previous:   log.Error.Previous,
		}),
	)
}

func (lr *ZapLogRepository) WriteHTTPRequestLog(log *entity.HTTPRequestLog) {
	lr.httpRequestLogger.Info(
		"http-request log",
		zap.Time("time", log.At),
		zap.String("host", log.Host),
		zap.String("method", log.Method),
		zap.String("origin", log.Origin),
		zap.Any("req_headers", log.ReqHeader),
		zap.Any("req_body", log.ReqBody),
		zap.Int("status_code", log.StatusCode),
		zap.Duration("latency", log.Latency),
		zap.Error(log.Error),
		zap.Any("res_headers", log.ResHeader),
		zap.Any("res_body", log.ResBody),
	)
}

type LogError struct {
	Code       string
	Message    string
	StatusCode int
	Detail     string
	Previous   error
}

func (e *LogError) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("code", e.Code)
	enc.AddString("message", e.Message)
	enc.AddInt("status_code", e.StatusCode)
	enc.AddString("detail", e.Detail)
	enc.AddString("error", e.Previous.Error())

	return nil
}
