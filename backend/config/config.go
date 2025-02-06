package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	config *viper.Viper
)

// Init initializes the config
func Init(env string) {
	config = viper.New()
	config.SetConfigFile("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("/app/config")
	config.AddConfigPath("./config")
	config.AddConfigPath("../config")
	config.AddConfigPath("../../config")
	config.AddConfigPath("/app/backend/config")

	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	config.AutomaticEnv()
	config.BindEnv("AWS_ACCESS_KEY_ID")
	config.BindEnv("AWS_SECRET_ACCESS_KEY")
	config.BindEnv("AWS_REGION")
	config.BindEnv("SENTRY_DSN")
	config.BindEnv("DATABASE_USERNAME")
	config.BindEnv("DATABASE_PASSWORD")
	config.BindEnv("CF_CERT_URL")
	config.BindEnv("CF_AUD")
}

// GetConfig gets the config settings
func GetConfig() *viper.Viper {
	return config
}
