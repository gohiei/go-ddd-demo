// Package usecase provides use case implementations for logging.
package usecase

import (
	"encoding/json"
	"fmt"

	"cypt/internal/dddcore"
	"cypt/internal/logger/entity"
	"cypt/internal/logger/entity/events"
	"cypt/internal/logger/repository"
)

// LogAccessUseCaseInput represents the input data for the LogAccessUseCase.
type LogAccessUseCaseInput entity.AccessLog

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
	return events.RequestDoneEventName
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
	log := &entity.AccessLog{
		At:            input.At,
		Method:        input.Method,
		Origin:        input.Origin,
		HTTPVersion:   input.HTTPVersion,
		UserAgent:     input.UserAgent,
		XFF:           input.XFF,
		StatusCode:    input.StatusCode,
		ContentLength: input.ContentLength,
		Latency:       input.Latency,
		Domain:        input.Domain,
		Host:          input.Host,
		RequestID:     input.RequestID,
		FullPath:      input.FullPath,
		SessionID:     input.SessionID,
		Agent:         input.Agent,
	}

	uc.logger.WriteAccessLog(log)

	return LogAccessUseCaseOutput{}, nil
}
