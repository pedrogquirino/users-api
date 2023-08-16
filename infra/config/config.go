package config

import (
	"log"

	"github.com/Netflix/go-env"
)

type Config struct {
	Http  Http
	SqlDB SqlDB
}

type Http struct {
	Port string `env:"HTTP_PORT",default:"8080"`
}

type SqlDB struct {
	HOST     string `env:"DB_HOST",default=localhost`
	DB       string `env:"DB_NAME",default=postgres`
	SSLMode  string `env:"SSL_MODE",default=disable`
	User     string `env:"DB_USER",default=api`
	Password string `env:"DB_PASSWORD",default=api`
}

func New() *Config {
	var config Config

	_, err := env.UnmarshalFromEnviron(&config)
	if err != nil {
		log.Panicf("error parsing config: %+v", err)
	}
	return &config
}
