package suggest

import (
	_ "embed"
	"fmt"
	ollama "github.com/jmorganca/ollama/api"
)

//go:embed Modelfile.suggest
var suggestModelfile string

type Suggest struct {
	api       *ollama.Client
	tag       string
	modelfile string
}

func (s *Suggest) Tag() string {
	return s.tag
}

func (s *Suggest) Modelfile() string {
	return s.modelfile
}

func New(api *ollama.Client, version string) *Suggest {
	tag := fmt.Sprintf("tlm:%s-s", version)
	return &Suggest{api: api, tag: tag, modelfile: suggestModelfile}
}
