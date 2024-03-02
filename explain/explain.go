package explain

import (
	ollama "github.com/jmorganca/ollama/api"
)

type Explain struct {
	api           *ollama.Client
	modelfile     string
	modelfileName string
}

func New(api *ollama.Client, modelfile string) *Explain {
	e := &Explain{api: api, modelfile: modelfile, modelfileName: "explain:7b"}
	return e
}
