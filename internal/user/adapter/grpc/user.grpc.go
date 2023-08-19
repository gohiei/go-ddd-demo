// Package grpc provides the gRPC server implementation for the user service.
package grpc

import (
	context "context"

	"cypt/internal/dddcore"
	"cypt/internal/user/adapter/grpc/protobuffer"
	"cypt/internal/user/usecase"
)

// UserServer is responsible for handling user-related gRPC requests.
type UserServer struct {
	RegisterUserUsecase dddcore.UseCase[usecase.RegisterUserUseCaseInput, usecase.RegisterUserUseCaseOutput]
}

// RegisterUser handles the RegisterUser gRPC request.
func (c *UserServer) RegisterUser(ctx context.Context, in *protobuffer.RegisterUserInput) (*protobuffer.RegisterUserOutput, error) {
	input := usecase.RegisterUserUseCaseInput{
		Username: in.GetUsername(),
		Password: in.GetPassword(),
	}

	out, err := c.RegisterUserUsecase.Execute(&input)

	if err != nil {
		return nil, err
	}

	return &protobuffer.RegisterUserOutput{
		Result: "ok",
		Ret: &protobuffer.RegisterUserUseCaseOutput{
			Id:       out.ID,
			Username: out.Username,
			UserID:   out.UserID,
		},
	}, nil
}
