package app

import (
	_ "embed"
	ollama "github.com/jmorganca/ollama/api"
	"github.com/yusufcanb/tlama/pkg/config"
	"github.com/yusufcanb/tlama/pkg/explain"
	"github.com/yusufcanb/tlama/pkg/install"
	"github.com/yusufcanb/tlama/pkg/suggest"

	"github.com/urfave/cli/v2"
)

type TlmApp struct {
	App *cli.App
}

func New(version string, suggestModelfile string, explainModelfile string) *TlmApp {
	con := config.New()
	con.LoadOrCreateConfig()

	o, _ := ollama.ClientFromEnvironment()
	sug := suggest.New(o, suggestModelfile)
	exp := explain.New(o, explainModelfile)
	ins := install.New(o)

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
