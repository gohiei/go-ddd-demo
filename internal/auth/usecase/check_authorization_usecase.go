// Package usecase contains the logic for checking authorization in the authentication system.
package usecase

import (
	"net/http"

	"cypt/internal/auth/entity"
	"cypt/internal/dddcore"
)

// CheckAuthorizationUsecaseInput represents the input for the CheckAuthorizationUsecase.
type CheckAuthorizationUsecaseInput struct {
	Token  string
	Method string
	URL    string
	IP     string
	XFF    string
}

// CheckAuthorizationUsecaseOutput represents the output of the CheckAuthorizationUsecase.
type CheckAuthorizationUsecaseOutput struct {
	Authorized bool
}

// CheckAuthorizationUsecase is responsible for checking authorization.
type CheckAuthorizationUsecase struct {
	eventBus dddcore.EventBus
}

// NewCheckAuthorizationUsecase creates a new CheckAuthorizationUsecase instance.
func NewCheckAuthorizationUsecase(bus dddcore.EventBus) *CheckAuthorizationUsecase {
	return &CheckAuthorizationUsecase{eventBus: bus}
}

// Execute performs the authorization check.
func (uc *CheckAuthorizationUsecase) Execute(input *CheckAuthorizationUsecaseInput) (CheckAuthorizationUsecaseOutput, error) {
	if entity.IgnoreRoute(input.Method, input.URL) {
		return CheckAuthorizationUsecaseOutput{Authorized: true}, nil
	}

	jwtToken, err := entity.NewJwtToken(input.Token, entity.Request{
		Method: input.Method,
		URL:    input.URL,
		IP:     input.IP,
		XFF:    input.XFF,
	})

	if err != nil {
		return CheckAuthorizationUsecaseOutput{Authorized: false}, err
	}

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
