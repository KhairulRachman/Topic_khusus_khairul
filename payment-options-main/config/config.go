package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

func LoadConfig() *Config {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("Failed to load env config: %v", err)
	}
	return &cfg
}

type Config struct {
	ListenPort string `envconfig:"LISTEN_PORT" required:"true" default:"8089"`
}
