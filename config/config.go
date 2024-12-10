package config

import (
	"github.com/alexfalkowski/go-service/config"
)

// Config for the client.
type Config struct {
	*config.Config `yaml:",inline" json:",inline" toml:",inline"`
}

// Valid or error.
func (c Config) Valid() error {
	if c.Config == nil {
		return config.ErrInvalidConfig
	}

	return c.Config.Valid()
}

func decorateConfig(cfg *Config) *config.Config {
	return cfg.Config
}
