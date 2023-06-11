/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "cypt",
	Short: "A commander for cypt",
	Long:  `A commandline tool for the different situations`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "config.example.yaml", "Config file")
}

func loadConfiguration(cfgFile string) *viper.Viper {
	config := viper.New()
	config.SetConfigFile(cfgFile)
	config.AddConfigPath("./configs/")
	config.AddConfigPath(".")

	if err := config.ReadInConfig(); err != nil {
		fmt.Println("Can not read config:", err)
	}

	return config
}
