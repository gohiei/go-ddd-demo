package registry

import (
	"cypt/internal/infra"
	adapter "cypt/internal/user/adapter"
)

func (r *registry) NewUserController() adapter.UserController {
	db, err := infra.NewUserDB()

	if err != nil {
		panic(err)
	}

	return adapter.NewController(adapter.NewMySqlUserRepository(db))
}
