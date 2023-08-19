package grpc

import (
	"cypt/internal/dddcore"
	"cypt/internal/user/adapter"
	"cypt/internal/user/adapter/grpc/protobuffer"
	"cypt/internal/user/infra"
	"cypt/internal/user/usecase"

	"google.golang.org/grpc"
)

type UserGrpcConfig struct {
	UserWriteDatabaseDSN string
	UserReadDatabaseDSN  string
	IDRedisDSN           string
}

func NewUserGrpc(server *grpc.Server, eventBus dddcore.EventBus, config UserGrpcConfig) *UserServer {
	db, _ := infra.NewUserDB(
		config.UserWriteDatabaseDSN,
		config.UserReadDatabaseDSN,
	)
	userRepo := adapter.NewMySQLUserRepository(db)

	redisConn, _ := infra.NewIDRedis(config.IDRedisDSN)
	idRepo := adapter.NewRedisIDRepository(redisConn)

	userServer := &UserServer{
		RegisterUserUsecase: usecase.NewRegisterUserUseCase(userRepo, idRepo, eventBus),
	}
	protobuffer.RegisterUserServer(server, userServer)

	return userServer
}
