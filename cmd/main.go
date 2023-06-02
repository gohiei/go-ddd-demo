package main

import (
	"context"
	dddcore "cypt/internal/dddcore"
	adapter "cypt/internal/dddcore/adapter"
	logger "cypt/internal/logger/adapter/restful"
	user "cypt/internal/user/adapter/restful"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	godotenv.Load("configs/.env")

	eventBus := adapter.NewWatermillEventBus()

	router := gin.Default()
	NewAppController(router, &eventBus)

	srv := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("Shutting down gracefullly, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server existing")
}

func NewAppController(router *gin.Engine, eventBus dddcore.EventBus) {
	logger.NewLoggerRestful(router, eventBus)
	user.NewUserRestful(router, eventBus)
}
