package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Env string     `mapstructure:"env"`
	Web *WebConfig `mapstructure:"web"`
}

type WebConfig struct {
	Http struct {
		Server struct {
			Port int
		}
	}
}

func setDefault() {
	viper.SetDefault("web.http.server.port", 5001)
	viper.SetDefault("env", "development")

}

func NewConfig() (*Config, error) {
	setDefault()

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
