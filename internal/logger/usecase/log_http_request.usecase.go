package usecase

import (
	"encoding/json"

	"cypt/internal/dddcore"
	"cypt/internal/logger/entity"
	"cypt/internal/logger/repository"
)

type LogHTTPRequestUseCaseInput entity.HTTPRequestLog

type LogHTTPRequestUseCaseOutput struct{ Result bool }

type LogHTTPRequestUseCase struct {
	logger repository.LogRepository
}

var _ dddcore.UseCase[LogHTTPRequestUseCaseInput, LogHTTPRequestUseCaseOutput] = (*LogHTTPRequestUseCase)(nil)

func NewLogHTTPRequestUseCase(logger repository.LogRepository, eventBus dddcore.EventBus) *LogHTTPRequestUseCase {
	uc := &LogHTTPRequestUseCase{logger: logger}
	_ = eventBus.Register(uc)

	return uc
}

func (uc *LogHTTPRequestUseCase) Name() string {
	return "logger.http_request"
}

func (uc *LogHTTPRequestUseCase) EventName() string {
	return "http_request.done"
}

func (uc *LogHTTPRequestUseCase) When(eventName string, message []byte) error {
	var input LogHTTPRequestUseCaseInput

	if err := json.Unmarshal(message, &input); err != nil {
		return err
	}

	_, err := uc.Execute(&input)
	return err
}

func (uc *LogHTTPRequestUseCase) Execute(input *LogHTTPRequestUseCaseInput) (*LogHTTPRequestUseCaseOutput, error) {
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

	return &LogHTTPRequestUseCaseOutput{Result: true}, nil
}
