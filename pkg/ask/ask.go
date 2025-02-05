package ask

import (
	ollama "github.com/jmorganca/ollama/api"
	"github.com/spf13/viper"
	"github.com/yusufcanb/tlm/pkg/config"
)

type Ask struct {
	api     *ollama.Client
	version string
	model   string
	style   string
}

func New(o *ollama.Client, version string) *Ask {
	model := viper.GetString("llm.model")
	return &Ask{
		model:   model,
		style:   config.Balanced,
		api:     o,
		version: version,
	}
}
