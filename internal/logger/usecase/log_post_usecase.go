package logger

import (
	"encoding/json"
	"time"

	dddcore "cypt/internal/dddcore"
	entity "cypt/internal/logger/entity"
	repository "cypt/internal/logger/repository"
)

type LogPostUseCaseInput struct {
	At            time.Time `json:"at"`
	UserAgent     string    `json:"user_agent"`
	XFF           string    `json:"x_forwarded_for"`
	RequestId     string    `json:"request_id"`
	Host          string    `json:"host"`
	Domain        string    `json:"domain"`
	StatusCode    int       `json:"status_code"`
	ContentLength int       `json:"content_log"`
	Latency       int       `json:"latency"`
	IP            string    `json:"ip"`
	Method        string    `json:"method"`
	Origin        string    `json:"origin"`
	HttpVersion   string    `json:"http_version"`
	RequestBody   string    `json:"request_body"`
	ResponseData  string    `json:"response_data"`
}

type LogPostUseCaseOutput struct{}

type LogPostUseCase struct {
	logger repository.LogRepository
}

func NewLogPostUseCase(logger repository.LogRepository, eventBus dddcore.EventBus) *LogPostUseCase {
	uc := &LogPostUseCase{logger: logger}
	eventBus.Register(uc)

	return uc
}

// it's an usecase
// and it's an event handler too.
var _ dddcore.UseCase[LogPostUseCaseInput, LogPostUseCaseOutput] = (*LogPostUseCase)(nil)
var _ dddcore.EventHandler = (*LogPostUseCase)(nil)

func (uc *LogPostUseCase) Name() string {
	return "logger.post"
}

func (uc *LogPostUseCase) EventName() string {
	return "request.done"
}

func (uc *LogPostUseCase) When(eventName string, message []byte) {
	var input LogPostUseCaseInput

	if err := json.Unmarshal(message, &input); err != nil {
		return
	}

	uc.Execute(&input)
}

func (uc *LogPostUseCase) Execute(input *LogPostUseCaseInput) (LogPostUseCaseOutput, error) {
	log := entity.PostLog{
		At:            input.At,
		IP:            input.IP,
		Method:        input.Method,
		Origin:        input.Origin,
		StatusCode:    input.StatusCode,
		ContentLength: input.ContentLength,
		Domain:        input.Domain,
		Host:          input.Host,
		RequestId:     input.RequestId,
		RequestBody:   input.RequestBody,
		ResponseData:  input.ResponseData,
	}

	uc.logger.WritePostLog(log)

	return LogPostUseCaseOutput{}, nil
}
