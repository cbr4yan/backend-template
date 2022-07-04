package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Debug      bool
	HttpServer Server
}

type Server struct {
	Addr string
}

func Setup(prefix string) (*Config, error) {
	c := &Config{}
	err := envconfig.Process(prefix, c)
	return c, err
}
