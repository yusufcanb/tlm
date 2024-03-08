package suggest

import (
	ollama "github.com/jmorganca/ollama/api"
)

type Suggest struct {
	api           *ollama.Client
	modelfileName string
}

func New(api *ollama.Client) *Suggest {
	return &Suggest{api: api, modelfileName: "suggest:7b"}
}
