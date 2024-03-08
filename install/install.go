package install

import (
	ollama "github.com/jmorganca/ollama/api"
)

type Install struct {
	api *ollama.Client

	suggestModelfile string
	explainModelfile string
}

func New(api *ollama.Client, suggestModelfile string, explainModelfile string) *Install {
	return &Install{
		api:              api,
		suggestModelfile: suggestModelfile,
		explainModelfile: explainModelfile,
	}
}
