package explain

import (
	ollama "github.com/jmorganca/ollama/api"
)

type Explain struct {
	api           *ollama.Client
	modelfileName string
}

func New(api *ollama.Client) *Explain {
	e := &Explain{api: api, modelfileName: "explain:7b"}
	return e
}
