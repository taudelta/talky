package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Addr  string `envconfig:"ADDR" default:":8081"`
	DbDSN string `envconfig:"DB_DSN" default:"postgresql://talky:1@localhost:5432/talky?sslmode=disable"`
}

func InitConfig() *Config {
	cfg := &Config{}
	envconfig.MustProcess("", cfg)
	return cfg
}
