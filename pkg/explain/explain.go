package explain

import (
	_ "embed"
	"fmt"
	ollama "github.com/jmorganca/ollama/api"
)

//go:embed Modelfile.explain
var modelFile string

type Explain struct {
	api       *ollama.Client
	version   string
	tag       string
	modelfile string
}

func (e *Explain) Tag() string {
	return e.tag
}

func (e *Explain) Modelfile() string {
	return e.modelfile
}

func New(api *ollama.Client, version string) *Explain {
	modelfileName := fmt.Sprintf("tlm:%s-e", version)
	return &Explain{api: api, tag: modelfileName, modelfile: modelFile, version: version}
}
