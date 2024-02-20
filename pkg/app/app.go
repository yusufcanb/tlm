package app

import (
	"github.com/yusufcanb/tlama/pkg/api"
	"github.com/yusufcanb/tlama/pkg/config"
	"github.com/yusufcanb/tlama/pkg/explain"
	"github.com/yusufcanb/tlama/pkg/install"
	"github.com/yusufcanb/tlama/pkg/suggest"

	"github.com/urfave/cli/v2"
)

type TlamaApp struct {
	App    *cli.App
	Config *config.TlamaConfig
}

func New(version string) *TlamaApp {

	cliApp := &cli.App{
		Name:        "tlama",
		Usage:       "Terminal Intelligence /w Local LLM",
		Description: "tlama is a command line tool to provide terminal intelligence locally with LLaMa.",
		Version:     version,
		Action: func(c *cli.Context) error {
			return cli.ShowAppHelp(c)
		},
		Commands: []*cli.Command{
			&cli.Command{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Print tlama version.",
				Action: func(c *cli.Context) error {
					cli.ShowVersion(c)
					return nil
				},
			},
			suggest.GetCommand(),
			explain.GetCommand(),
			install.GetCommand(),
			config.GetCommand(),
		},
	}

	cliApp.HideHelpCommand = true
	cliApp.Metadata = make(map[string]interface{})

	cliApp.Metadata["config"] = config.New()
	cliApp.Metadata["api"] = api.New(cliApp.Metadata["config"].(*config.TlamaConfig))

	return &TlamaApp{
		App: cliApp,
	}
}
