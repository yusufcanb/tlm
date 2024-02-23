package suggest

import (
	ollama "github.com/jmorganca/ollama/api"
)

type Suggest struct {
	api       *ollama.Client
	modelfile string
}

func New(api *ollama.Client, modelfile string) *Suggest {
	return &Suggest{api: api, modelfile: modelfile}
}
