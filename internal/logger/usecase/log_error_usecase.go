package logger

import (
	"encoding/json"
	"fmt"
	"time"

	dddcore "cypt/internal/dddcore"
	entity "cypt/internal/logger/entity"
	repository "cypt/internal/logger/repository"
)

// LogErrorUseCaseInput represents the input data for the LogErrorUseCase.
type LogErrorUseCaseInput struct {
	At          time.Time     `json:"at"`
	RequestId   string        `json:"request_id"`
	Host        string        `json:"host"`
	Domain      string        `json:"domain"`
	IP          string        `json:"ip"`
	Method      string        `json:"method"`
	Origin      string        `json:"origin"`
	RequestBody string        `json:"request_body"`
	Error       dddcore.Error `json:"error"`
}

// LogErrorUseCaseOutput represents the output data for the LogErrorUseCase.
type LogErrorUseCaseOutput struct{}

// LogErrorUseCase handles logging of error events.
type LogErrorUseCase struct {
	logger repository.LogRepository
}

// NewLogErrorUseCase creates a new instance of LogErrorUseCase.
func NewLogErrorUseCase(logger repository.LogRepository, eventBus dddcore.EventBus) *LogErrorUseCase {
	uc := &LogErrorUseCase{logger: logger}
	eventBus.Register(uc)

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
	return "error.raised"
}

// When handles the incoming event and executes the use case.
func (uc *LogErrorUseCase) When(eventName string, message []byte) {
	var input LogErrorUseCaseInput

	if err := json.Unmarshal(message, &input); err != nil {
		fmt.Println("err ", err)
		return
	}

	uc.Execute(&input)
}

// Execute performs the logging of error events based on the provided input.
func (uc *LogErrorUseCase) Execute(input *LogErrorUseCaseInput) (LogErrorUseCaseOutput, error) {
	log := entity.ErrorLog{
		At:          input.At,
		IP:          input.IP,
		Method:      input.Method,
		Origin:      input.Origin,
		Domain:      input.Domain,
		Host:        input.Host,
		RequestId:   input.RequestId,
		RequestBody: input.RequestBody,
		Error:       input.Error,
	}

	uc.logger.WriteErrorLog(log)

	return LogErrorUseCaseOutput{}, nil
}
