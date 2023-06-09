package logger

import (
	"encoding/json"
	"fmt"

	dddcore "cypt/internal/dddcore"
	entity "cypt/internal/logger/entity"
	repository "cypt/internal/logger/repository"
)

type LogHTTPRequestUseCaseInput entity.HTTPRequestLog

type LogHTTPRequestUseCaseOutput bool

type LogHTTPRequestUseCase struct {
	logger repository.LogRepository
}

func NewLogHTTPRequestUseCase(logger repository.LogRepository, eventBus dddcore.EventBus) *LogHTTPRequestUseCase {
	uc := &LogHTTPRequestUseCase{logger: logger}
	eventBus.Register(uc)

	return uc
}

func (uc *LogHTTPRequestUseCase) Name() string {
	return "logger.http_request"
}

func (uc *LogHTTPRequestUseCase) EventName() string {
	return "http_request.done"
}

func (uc *LogHTTPRequestUseCase) When(eventName string, message []byte) {
	var input LogHTTPRequestUseCaseInput

	if err := json.Unmarshal(message, &input); err != nil {
		fmt.Println("err ", err)
		return
	}

	uc.Execute(&input)
}

func (uc *LogHTTPRequestUseCase) Execute(input *LogHTTPRequestUseCaseInput) (LogHTTPRequestUseCaseOutput, error) {
	log := &entity.HTTPRequestLog{
		At:         input.At,
		Method:     input.Method,
		Origin:     input.Origin,
		Host:       input.Host,
		ReqHeader:  input.ReqHeader,
		ReqBody:    input.ReqBody,
		StatusCode: input.StatusCode,
		Latency:    input.Latency,
		Error:      input.Error,
		ResHeader:  input.ResHeader,
		ResBody:    input.ResBody,
	}
	uc.logger.WriteHTTPRequestLog(log)

	return true, nil
}
