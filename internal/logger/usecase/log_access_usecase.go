package logger

import (
	"encoding/json"
	"fmt"
	"time"

	dddcore "cypt/internal/dddcore"
	entity "cypt/internal/logger/entity"
	repository "cypt/internal/logger/repository"
)

// LogAccessUseCaseInput represents the input data for the LogAccessUseCase.
type LogAccessUseCaseInput struct {
	At            time.Time `json:"at"`
	UserAgent     string    `json:"user_agent"`
	XFF           string    `json:"x_forwarded_for"`
	RequestId     string    `json:"request_id"`
	Host          string    `json:"host"`
	Domain        string    `json:"domain"`
	StatusCode    int       `json:"status_code"`
	ContentLength int       `json:"content_length"`
	Latency       int64     `json:"latency"`
	IP            string    `json:"ip"`
	Method        string    `json:"method"`
	Origin        string    `json:"origin"`
	HttpVersion   string    `json:"http_version"`
}

// LogAccessUseCaseOutput represents the output data for the LogAccessUseCase.
type LogAccessUseCaseOutput struct{}

// LogAccessUseCase handles logging of access events.
type LogAccessUseCase struct {
	logger repository.LogRepository
}

// NewLogAccessUseCase creates a new instance of LogAccessUseCase.
func NewLogAccessUseCase(logger repository.LogRepository, eventBus dddcore.EventBus) *LogAccessUseCase {
	uc := &LogAccessUseCase{logger: logger}
	eventBus.Register(uc)

	return uc
}

// Name returns the name of the LogAccessUseCase.
func (uc *LogAccessUseCase) Name() string {
	return "logger.access"
}

// EventName returns the name of the event handled by the LogAccessUseCase.
func (uc *LogAccessUseCase) EventName() string {
	return "request.done"
}

// When handles the incoming event and executes the use case.
func (uc *LogAccessUseCase) When(eventName string, message []byte) {
	var input LogAccessUseCaseInput

	if err := json.Unmarshal(message, &input); err != nil {
		fmt.Println("err ", err)
		return
	}

	uc.Execute(&input)
}

// Execute performs the logging of access events based on the provided input.
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
