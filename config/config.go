package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

const (
	Local = "local"
	Dev   = "dev"
	Prod  = "prod"
)

type AppConfig struct {
	App struct {
		Name        string `mapstructure:"name"`
		Port        int    `mapstructure:"port"`
		Environment string `mapstructure:"environment"`
	} `mapstructure:"app"`
	Postgres struct {
		Connstr string `mapstructure:"connstr"`
	} `mapstructure:"postgres"`
	Logger struct {
		Encoding string `mapstructure:"encoding"`
	} `mapstructure:"logger"`
}

func GetConfig() (AppConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		return AppConfig{}, fmt.Errorf("error reading config file: %w", err)
	}

	config := AppConfig{}

	err = viper.Unmarshal(&config)
	if err != nil {
		return AppConfig{}, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return config, nil
}
