package config

import (
	ollama "github.com/jmorganca/ollama/api"
)

type Config struct {
	api *ollama.Client
}

func New(o *ollama.Client) *Config {
	return &Config{
		api: o,
	}
}
