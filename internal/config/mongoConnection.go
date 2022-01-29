package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

type mongoConfig struct {
	URI string `env:"DB_HOST" envDefault:"mongodb://localhost:27017"`
}

func GetMongoURI() (string, error) {
	cfg := mongoConfig{}
	if err := env.Parse(&cfg); err != nil {
		log.WithField("error", err.Error()).Warn("error in parsing mongo env variable")
		return "", fmt.Errorf("error in parsing mongo env variable %w", err)
	}
	return cfg.URI, nil
}
