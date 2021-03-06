// Package config contains all configuration info for databases
package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type mongoConfig struct {
	URI string `env:"MONGO_URI" envDefault:"mongodb://localhost:27017"`
}

// GetMongoURI returns URI to connect to mongoDb
func GetMongoURI() (string, error) {
	cfg := mongoConfig{}
	if err := env.Parse(&cfg); err != nil {
		return "", fmt.Errorf("error in parsing mongo env variable %w", err)
	}
	return cfg.URI, nil
}
