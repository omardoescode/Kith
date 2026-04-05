package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port              int           `envconfig:"PORT" default:"8000"`
	ConnectionString  string        `envconfig:"CONNECTION_STRING" required:"true"`
	DBMaxOpenConns    int           `envconfig:"DB_MAX_OPEN_CONNS" default:"25"`
	DBMaxIdleConns    int           `envconfig:"DB_MAX_IDLE_CONNS" default:"25"`
	DBConnMaxLifetime time.Duration `envconfig:"DB_CONN_MAX_LIFETIME" default:"5m"`
	DBConnMaxIdleTime time.Duration `envconfig:"DB_CONN_MAX_IDLE_TIME" default:"5m"`
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
