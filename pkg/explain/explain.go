package explain

import (
	ollama "github.com/jmorganca/ollama/api"
)

type Explain struct {
	api       *ollama.Client
	modelfile string
	system    string
}

func New(api *ollama.Client, modelfile string) *Explain {
	e := &Explain{api: api, modelfile: modelfile}
	e.system = `You are software program specifically for Command Line Interface usage.
User will ask you some thing that can be convertible to a UNIX or Windows command.
You won't provide information or explanations and your output will be just an executable shell command inside three backticks.
`
	return e
}
