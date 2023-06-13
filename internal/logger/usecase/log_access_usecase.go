package logger

import (
	"encoding/json"
	"fmt"
	"time"

	dddcore "cypt/internal/dddcore"
	entity "cypt/internal/logger/entity"
	repository "cypt/internal/logger/repository"
)

type LogAccessUseCaseInput struct {
	At            time.Time `json:"at"`
	UserAgent     string    `json:"user_agent"`
	XFF           string    `json:"x_forwarded_for"`
	RequestId     string    `json:"request_id"`
	Host          string    `json:"host"`
	Domain        string    `json:"domain"`
	StatusCode    int       `json:"status_code"`
	ContentLength int       `json:"content_log"`
	Latency       int64     `json:"latency"`
	IP            string    `json:"ip"`
	Method        string    `json:"method"`
	Origin        string    `json:"origin"`
	HttpVersion   string    `json:"http_version"`
}

type LogAccessUseCaseOutput struct{}

type LogAccessUseCase struct {
	logger repository.LogRepository
}

func NewLogAccessUseCase(logger repository.LogRepository, eventBus dddcore.EventBus) *LogAccessUseCase {
	uc := &LogAccessUseCase{logger: logger}
	eventBus.Register(uc)

	return uc
}

// it's an usecase
// and it's an event handler too.
var _ dddcore.UseCase[LogAccessUseCaseInput, LogAccessUseCaseOutput] = (*LogAccessUseCase)(nil)
var _ dddcore.EventHandler = (*LogAccessUseCase)(nil)

func (uc *LogAccessUseCase) Name() string {
	return "logger.access"
}

func (uc *LogAccessUseCase) EventName() string {
	return "request.done"
}

func (uc *LogAccessUseCase) When(eventName string, message []byte) {
	var input LogAccessUseCaseInput

	if err := json.Unmarshal(message, &input); err != nil {
		fmt.Println("err ", err)
		return
	}

	uc.Execute(&input)
}

func (uc *LogAccessUseCase) Execute(input *LogAccessUseCaseInput) (LogAccessUseCaseOutput, error) {
	log := entity.AccessLog{
		At:            input.At,
		Method:        input.Method,
		Origin:        input.Origin,
		HttpVersion:   input.HttpVersion,
		UserAgent:     input.UserAgent,
		XFF:           input.XFF,
		StatusCode:    input.StatusCode,
		ContentLength: input.ContentLength,
		Latency:       input.Latency,
		Domain:        input.Domain,
		Host:          input.Host,
		RequestId:     input.RequestId,
	}

	uc.logger.WriteAccessLog(log)

	return LogAccessUseCaseOutput{}, nil
}
