package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	APIKey   string `mapstructure:"API_KEY"`
	BASEUrl  string `mapstructure:"BASE_URL"`
	HTTPPort string `mapstructure:"HTTP_PORT"`
	GRPCPort string `mapstructure:"GRPC_PORT"`
}

func LoadConfig() (Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
