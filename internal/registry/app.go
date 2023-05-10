package registry

import (
	user "cypt/internal/user/adapter"
)

type AppController struct {
	User interface{ user.UserController }
}
