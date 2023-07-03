package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cypt",
	Short: "A commander for cypt",
	Long:  `A command-line tool for different situations`,
}

// Execute runs the root command
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "config.example.yaml", "Config file")
}

// loadConfiguration loads the configuration from a config file
func loadConfiguration(cfgFile string) *viper.Viper {
	config := viper.New()
	config.SetConfigFile(cfgFile)
	config.AddConfigPath("./configs/")
	config.AddConfigPath(".")

	if err := config.ReadInConfig(); err != nil {
		fmt.Println("Cannot read config:", err)
	}

	return config
}
