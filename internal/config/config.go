package config

import (
	"fmt"

	"github.com/vrischmann/envconfig"
)

var Cfg *Config

type Config struct {
	CacheSize int `envconfig:"default=5"`
}

func InitConfig(prefix string) (*Config, error) {
	conf := &Config{}
	if err := envconfig.InitWithPrefix(conf, prefix); err != nil {
		return nil, fmt.Errorf("init config error: %w", err)
	}

	return conf, nil
}
