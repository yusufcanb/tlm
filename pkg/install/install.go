package install

import (
	_ "embed"
	ollama "github.com/jmorganca/ollama/api"
)

type Install struct {
	api *ollama.Client
}

func New(api *ollama.Client) *Install {
	return &Install{
		api: api,
	}
}
