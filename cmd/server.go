package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	app "cypt/internal"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// restfulServerCmd represents the server command
var restfulServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Run restful api server",
	Long:  `A web server`,
	Run: func(cmd *cobra.Command, args []string) {
		cfgFile, _ := cmd.Parent().PersistentFlags().GetString("config")
		address, _ := cmd.Flags().GetString("address")
		port, _ := cmd.Flags().GetInt("port")

		config := loadConfiguration(cfgFile)

		currentServerSetting := serverSetting{
			address: address,
			port:    port,
			config:  config,
			app:     app.NewAppRestfulServer,
		}
		runServer(currentServerSetting)
	},
}

func init() {
	rootCmd.AddCommand(restfulServerCmd)

	restfulServerCmd.Flags().IntP("port", "p", 8080, "Listen port")
	restfulServerCmd.Flags().StringP("address", "a", "127.0.0.1", "Bind address")
}

// serverSetting holds the server configuration
type serverSetting struct {
	address string
	port    int
	config  *viper.Viper
	app     func(*gin.Engine, *viper.Viper)
}

// runServer starts the HTTP server
func runServer(s serverSetting) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.Default()
	s.app(router, s.config)

	addr := fmt.Sprintf("%s:%d", s.address, s.port)

	srv := &http.Server{
		Addr:              addr,
		Handler:           router,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		log.Println("Listen on:", addr)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("Shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
