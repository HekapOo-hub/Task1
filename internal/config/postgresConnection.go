package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

// Config contains info about postgres database connection
type Config struct {
	UserName string `env:"POSTGRES_USER" envDefault:"vitalijprokopenya"`
	Host     string `env:"HOST" envDefault:"postgres"`
	Port     string `env:"PORT" envDefault:"5432"`
	Password string `env:"POSTGRES_PASSWORD" envDefault:"1234"`
	DBName   string `env:"DB_NAME" envDefault:"vitalijprokopenya"`
	URL      string `env:"URL" envDefault:"postgres://"`
}

// GetURL returns URL to connect to postgres database
func (c *Config) GetURL() string {
	res := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.UserName, c.Password, c.Host, c.Port, c.DBName)
	return res
}

// NewConfig returns new config of postgresDB parsed from environment variables
func NewConfig() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
