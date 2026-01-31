package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	DB dbConfig
}

var cfg *Config

type dbConfig struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	User     string `env:"DB_USER" envDefault:"my-commerce-admin"`
	Password string `env:"DB_PASSWORD" envDefault:"password"`
	Port     string `env:"DB_PORT" envDefault:"5432"`
}

func Initialize() *Config {
	if cfg == nil {
		cfg = &Config{}
		err := env.Parse(cfg)
		if err != nil {
			log.Fatal("[env] does not load the environment variables: ", err)
		}
	}

	return cfg
}
