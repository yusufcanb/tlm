package install

import (
	_ "embed"
	ollama "github.com/jmorganca/ollama/api"
)

type Install struct {
	api *ollama.Client

	defaultContainerName string

	suggestModelfile string
	explainModelfile string
}

func New(api *ollama.Client, suggestModelfile string, explainModelfile string) *Install {
	return &Install{
		api:                  api,
		defaultContainerName: "ollama",

		suggestModelfile: suggestModelfile,
		explainModelfile: explainModelfile,
	}
}
