package registry

import (
	dddcore "cypt/internal/dddcore"
	"cypt/internal/infra"
	adapter "cypt/internal/user/adapter"
)

func (r *registry) NewUserController(eventBus dddcore.EventBus) adapter.UserController {
	db, err := infra.NewUserDB()

	if err != nil {
		panic(err)
	}

	return adapter.NewController(
		adapter.NewMySqlUserRepository(db),
		eventBus,
	)
}
