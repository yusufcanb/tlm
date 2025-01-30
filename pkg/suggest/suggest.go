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
	system  string
	model   string
	style   string
}

func (s *Suggest) Tag() string {
	return s.model
}

func New(api *ollama.Client, version string) *Suggest {
	model := viper.GetString("llm.model")
	style := viper.GetString("llm.suggest")
	return &Suggest{api: api, model: model, system: system, style: style, version: version}
}
