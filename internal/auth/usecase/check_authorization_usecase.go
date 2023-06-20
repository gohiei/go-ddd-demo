package auth

import (
	entity "cypt/internal/auth/entity"
	"cypt/internal/dddcore"
	"net/http"
)

var _ dddcore.Input = (*CheckAuthorizationUsecaseInput)(nil)
var _ dddcore.Output = (*CheckAuthorizationUsecaseOutput)(nil)
var _ dddcore.UseCase[CheckAuthorizationUsecaseInput, CheckAuthorizationUsecaseOutput] = (*CheckAuthorizationUsecase)(nil)

type CheckAuthorizationUsecaseInput struct {
	Token  string
	Method string
	URL    string
	IP     string
	XFF    string
}

type CheckAuthorizationUsecaseOutput struct {
	Authorized bool
}

type CheckAuthorizationUsecase struct {
	eventBus dddcore.EventBus
}

func NewCheckAuthorizationUsecase(bus dddcore.EventBus) *CheckAuthorizationUsecase {
	return &CheckAuthorizationUsecase{eventBus: bus}
}

func (uc *CheckAuthorizationUsecase) Execute(input *CheckAuthorizationUsecaseInput) (CheckAuthorizationUsecaseOutput, error) {
	if entity.IgnoreRoute(input.Method, input.URL) {
		return CheckAuthorizationUsecaseOutput{Authorized: true}, nil
	}

	jwtToken := entity.NewJwtToken(input.Token, entity.Request{
		Method: input.Method,
		URL:    input.URL,
		IP:     input.IP,
		XFF:    input.XFF,
	})

	isValid := jwtToken.Valid()

	if !isValid {
		uc.eventBus.PostAll(jwtToken)

		return CheckAuthorizationUsecaseOutput{}, dddcore.NewErrorS(
			"00001",
			"authentication failed",
			http.StatusForbidden,
		)
	}

	return CheckAuthorizationUsecaseOutput{Authorized: isValid}, nil
}
