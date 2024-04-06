package install

import (
	ollama "github.com/jmorganca/ollama/api"
	"github.com/yusufcanb/tlm/explain"
	"github.com/yusufcanb/tlm/suggest"
)

var repositoryOwner = "yusufcanb"
var repositoryName = "tlm"

type Install struct {
	api *ollama.Client

	suggest *suggest.Suggest
	explain *explain.Explain

	ReleaseManager *ReleaseManager
}

func New(api *ollama.Client, suggest *suggest.Suggest, explain *explain.Explain) *Install {
	return &Install{
		api:            api,
		suggest:        suggest,
		explain:        explain,
		ReleaseManager: NewReleaseManager(repositoryOwner, repositoryName),
	}
}
