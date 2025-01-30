package suggest

import (
	_ "embed"

	ollama "github.com/jmorganca/ollama/api"
	"github.com/spf13/viper"
)

//go:embed SYSTEM
var system string

type Suggest struct {
	api     *ollama.Client
	version string
	model   string
	system  string
	mode    string
}

func (s *Suggest) Tag() string {
	return s.model
}

func New(api *ollama.Client, version string) *Suggest {
	model := viper.GetString("llm.model")
	mode := viper.GetString("llm.suggest")
	return &Suggest{api: api, model: model, system: system, mode: mode, version: version}
}
