package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Web *WebConfig `mapstructure:"web"`
}

type WebConfig struct {
	Http struct {
		Server struct {
			Port string
		}
	}
}

func setDefault() {
	viper.SetDefault("web.http.server.port", "5000")
}

func NewConfig() (*Config, error) {
	setDefault()

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
