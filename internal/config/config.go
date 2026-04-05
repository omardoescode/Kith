package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port int `default:"8000"`
}

func (c *Config) Load() error {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("load .env: %w", err)
	}

	if err := envconfig.Process("", c); err != nil {
		return fmt.Errorf("parse env config: %w", err)
	}

	return nil
}
