package user

import (
	dddcore "cypt/internal/dddcore"
	"cypt/internal/infra"
	adapter "cypt/internal/user/adapter"
	repo "cypt/internal/user/repository"
	usecase "cypt/internal/user/usecase"

	"github.com/gin-gonic/gin"
)

type _UserController interface {
	RegisterUser(ctx *gin.Context)
	Rename(ctx *gin.Context)
}

func NewController(r repo.UserRepository, eventBus dddcore.EventBus) _UserController {
	return &UserService{
		RegisterUserUseCase:  usecase.NewRegisterUserUseCase(r, eventBus),
		RenameUseCase:        usecase.NewRenameUseCase(r, eventBus),
		NotifyManagerHandler: usecase.NewNotifyManagerHandler(eventBus),
	}
}

func NewUserController(r *gin.Engine, eventBus dddcore.EventBus) {
	db, err := infra.NewUserDB()

	if err != nil {
		panic(err)
	}

	controller := NewController(
		adapter.NewMySqlUserRepository(db),
		eventBus,
	)

	r.POST("/user", controller.RegisterUser)
	r.PUT("/user/:id", controller.Rename)
}
