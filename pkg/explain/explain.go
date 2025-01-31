package explain

import (
	_ "embed"

	ollama "github.com/jmorganca/ollama/api"
	"github.com/spf13/viper"
)

//go:embed SYSTEM
var system string

type Explain struct {
	api     *ollama.Client
	version string
	system  string
	model   string
	style   string
}

func (e *Explain) Tag() string {
	return e.model
}

func New(api *ollama.Client, version string) *Explain {
	model := viper.GetString("llm.model")
	style := viper.GetString("llm.explain")
	return &Explain{api: api, model: model, system: system, style: style, version: version}
}
