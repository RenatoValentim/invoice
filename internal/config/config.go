package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func setDefaults() {
	// Server Defaults
	viper.SetDefault("server_host", "127.0.0.1")
	viper.SetDefault("server_port", 1323)

	// External Servcer Currency
	viper.SetDefault("currency_host", "http://127.0.0.1")
	viper.SetDefault("currency_port", 3001)

	// DB Defaults
	viper.SetDefault("db_host", "")
	viper.SetDefault("db_port", "")
	viper.SetDefault("db_name", "")
	viper.SetDefault("db_username", "")
	viper.SetDefault("db_password", "")
}

func bindEnvironmentVariables() {
	// Env Prefix
	viper.SetEnvPrefix("INVOICE")

	// Server Variables
	viper.BindEnv("server_host")
	viper.BindEnv("server_port")

	// External Server Currency
	viper.BindEnv("currency_host")
	viper.BindEnv("currency_port")

	// Databse Variables
	viper.BindEnv("db_host")
	viper.BindEnv("db_port")
	viper.BindEnv("db_name")
	viper.BindEnv("db_username")
	viper.BindEnv("db_password")
}

func LoadConfig(configPath string) {
	setDefaults()

	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("err: %v\n", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			bindEnvironmentVariables()
		} else {
			log.Fatalf("Failed to load configuration from configuration file. %s", err.Error())
		}
	}
}
