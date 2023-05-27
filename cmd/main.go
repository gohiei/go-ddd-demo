package main

import (
	dddcore "cypt/internal/dddcore/adapter"
	user "cypt/internal/user/adapter/restful"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("configs/.env")

	r := gin.Default()
	NewAppController(r)

	r.Run(":8080")
}

func NewAppController(router *gin.Engine) {
	eventBus := dddcore.NewWatermillEventBus()

	user.NewUserRestful(router, &eventBus)
}
