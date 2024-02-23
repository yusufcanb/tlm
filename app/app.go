package app

import (
	_ "embed"
	ollama "github.com/jmorganca/ollama/api"
	"github.com/yusufcanb/tlm/config"
	"github.com/yusufcanb/tlm/explain"
	"github.com/yusufcanb/tlm/install"
	"github.com/yusufcanb/tlm/suggest"

	"github.com/urfave/cli/v2"
)

//go:embed Modelfile.explain
var explainModelfile string

//go:embed Modelfile.suggest
var suggestModelfile string

type TlmApp struct {
	App *cli.App

	explainModelfile string
	suggestModelfile string
}

func New(version string) *TlmApp {
	con := config.New()
	con.LoadOrCreateConfig()

	o, _ := ollama.ClientFromEnvironment()
	sug := suggest.New(o, suggestModelfile)
	exp := explain.New(o, explainModelfile)
	ins := install.New(o, suggestModelfile, explainModelfile)

	cliApp := &cli.App{
		Name:            "tlm",
		Usage:           "local terminal companion powered by CodeLLaMa.",
		Version:         version,
		HideHelpCommand: true,
		Action: func(c *cli.Context) error {
			return cli.ShowAppHelp(c)
		},
		Commands: []*cli.Command{
			sug.Command(),
			exp.Command(),
			ins.Command(),
			con.Command(),
			&cli.Command{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "print version.",
				Action: func(c *cli.Context) error {
					cli.ShowVersion(c)
					return nil
				},
			},
		},
	}

	return &TlmApp{
		App: cliApp,
	}
}
