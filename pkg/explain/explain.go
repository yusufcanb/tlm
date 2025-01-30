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
	model   string
	system  string
	mode    string
}

func (e *Explain) Tag() string {
	return e.model
}

func New(api *ollama.Client, version string) *Explain {
	model := viper.GetString("llm.model")
	mode := viper.GetString("llm.explain")
	return &Explain{api: api, model: model, system: system, mode: mode, version: version}
}
