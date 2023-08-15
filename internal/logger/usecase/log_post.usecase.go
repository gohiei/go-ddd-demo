package usecase

import (
	"encoding/json"

	"cypt/internal/dddcore"
	"cypt/internal/logger/entity"
	"cypt/internal/logger/entity/events"
	"cypt/internal/logger/repository"
)

// LogPostUseCaseInput represents the input data for the LogPostUseCase.
type LogPostUseCaseInput entity.PostLog

// LogPostUseCaseOutput represents the output data for the LogPostUseCase.
type LogPostUseCaseOutput struct{}

// LogPostUseCase handles logging of post request events.
type LogPostUseCase struct {
	logger repository.LogRepository
}

// NewLogPostUseCase creates a new instance of LogPostUseCase.
func NewLogPostUseCase(logger repository.LogRepository, eventBus dddcore.EventBus) *LogPostUseCase {
	uc := &LogPostUseCase{logger: logger}
	_ = eventBus.Register(uc)

	return uc
}

// it's an usecase
// and it's an event handler too.
var _ dddcore.UseCase[LogPostUseCaseInput, LogPostUseCaseOutput] = (*LogPostUseCase)(nil)
var _ dddcore.EventHandler = (*LogPostUseCase)(nil)

// Name returns the name of the LogPostUseCase.
func (uc *LogPostUseCase) Name() string {
	return "logger.post"
}

// EventName returns the name of the event handled by the LogPostUseCase.
func (uc *LogPostUseCase) EventName() string {
	return events.RequestDoneEventName
}

// When handles the incoming event and executes the use case.
func (uc *LogPostUseCase) When(eventName string, message []byte) error {
	var input LogPostUseCaseInput

	if err := json.Unmarshal(message, &input); err != nil {
		return err
	}

	_, err := uc.Execute(&input)
	return err
}

// Execute performs the logging of post request events based on the provided input.
func (uc *LogPostUseCase) Execute(input *LogPostUseCaseInput) (*LogPostUseCaseOutput, error) {
	log := &entity.PostLog{
		At:            input.At,
		IP:            input.IP,
		Method:        input.Method,
		Origin:        input.Origin,
		StatusCode:    input.StatusCode,
		ContentLength: input.ContentLength,
		Domain:        input.Domain,
		Host:          input.Host,
		RequestID:     input.RequestID,
		RequestBody:   input.RequestBody,
		ResponseBody:  input.ResponseBody,
	}

	uc.logger.WritePostLog(log)

	return &LogPostUseCaseOutput{}, nil
}
