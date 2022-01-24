package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

// urlExample := "postgres://username:password@localhost:5432/database_name"
type Config struct {
	UserName string `env:"POSTGRES_USER" envDefault:"vitalijprokopenya"`
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     string `env:"PORT" envDefault:"5432"`
	Password string `env:"POSTGRES_PASSWORD" envDefault:"1234"`
	DBName   string `env:"PORT" envDefault:"vitalijprokopenya"`
	URL      string `env:"URL" envDefault:"postgres://"`
}

func (c *Config) GetURL() string {
	res := "postgres://"
	res += fmt.Sprintf("%s:%s@%s:%s/%s", c.UserName, c.Password, c.Host, c.Port, c.DBName)
	return res
}
func NewConfig() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
