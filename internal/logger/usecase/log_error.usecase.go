package usecase

import (
	"encoding/json"

	"cypt/internal/dddcore"
	"cypt/internal/logger/entity"
	"cypt/internal/logger/entity/events"
	"cypt/internal/logger/repository"
)

// LogErrorUseCaseInput represents the input data for the LogErrorUseCase.
type LogErrorUseCaseInput entity.ErrorLog

// LogErrorUseCaseOutput represents the output data for the LogErrorUseCase.
type LogErrorUseCaseOutput struct{}

// LogErrorUseCase handles logging of error events.
type LogErrorUseCase struct {
	logger repository.LogRepository
}

// NewLogErrorUseCase creates a new instance of LogErrorUseCase.
func NewLogErrorUseCase(logger repository.LogRepository, eventBus dddcore.EventBus) *LogErrorUseCase {
	uc := &LogErrorUseCase{logger: logger}
	_ = eventBus.Register(uc)

	return uc
}

// it's an usecase
// and it's an event handler too.
var _ dddcore.UseCase[LogErrorUseCaseInput, LogErrorUseCaseOutput] = (*LogErrorUseCase)(nil)
var _ dddcore.EventHandler = (*LogErrorUseCase)(nil)

// Name returns the name of the LogErrorUseCase.
func (uc *LogErrorUseCase) Name() string {
	return "logger.error"
}

// EventName returns the name of the event handled by the LogErrorUseCase.
func (uc *LogErrorUseCase) EventName() string {
	return events.UnexpectedErrorRaisedEventName
}

// When handles the incoming event and executes the use case.
func (uc *LogErrorUseCase) When(eventName string, message []byte) error {
	var input LogErrorUseCaseInput

	if err := json.Unmarshal(message, &input); err != nil {
		return err
	}

	_, err := uc.Execute(&input)
	return err
}

// Execute performs the logging of error events based on the provided input.
func (uc *LogErrorUseCase) Execute(input *LogErrorUseCaseInput) (*LogErrorUseCaseOutput, error) {
	log := &entity.ErrorLog{
		At:          input.At,
		IP:          input.IP,
		Method:      input.Method,
		Origin:      input.Origin,
		Domain:      input.Domain,
		Host:        input.Host,
		RequestID:   input.RequestID,
		RequestBody: input.RequestBody,
		Error:       input.Error,
	}

	uc.logger.WriteErrorLog(log)

	return &LogErrorUseCaseOutput{}, nil
}
