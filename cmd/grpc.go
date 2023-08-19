package cmd

import (
	"context"
	app "cypt/internal"
	"fmt"
	"log"
	"net"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// grpcServerCmd represents the grpc command
var grpcServerCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Run grpc api server",
	Long:  `A grpc server`,
	Run: func(cmd *cobra.Command, args []string) {
		cfgFile, _ := cmd.Parent().PersistentFlags().GetString("config")
		address, _ := cmd.Flags().GetString("address")
		port, _ := cmd.Flags().GetInt("port")

		config := loadConfiguration(cfgFile)

		currentServerSetting := grpcSetting{
			address: address,
			port:    port,
			config:  config,
			app:     app.NewAppGrpcServer,
		}
		runGrpcServer(currentServerSetting)
	},
}

func init() {
	rootCmd.AddCommand(grpcServerCmd)

	grpcServerCmd.Flags().IntP("port", "p", 8080, "Listen port")
	grpcServerCmd.Flags().StringP("address", "a", "127.0.0.1", "Bind address")
}

// serverSetting holds the server configuration
type grpcSetting struct {
	address string
	port    int
	config  *viper.Viper
	app     func(*grpc.Server, *viper.Viper)
}

func runGrpcServer(setting grpcSetting) {
	addr := fmt.Sprintf("%s:%d", setting.address, setting.port)
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	log.Println("gRPC server is running.")

	setting.app(server, setting.config)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Println("Listen on:", addr)

		if err := server.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-ctx.Done()

	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server.GracefulStop()

	log.Println("gRPC Server exiting")
}
