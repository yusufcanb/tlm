package app

import (
	"github.com/yusufcanb/tlama/pkg/api"
	"github.com/yusufcanb/tlama/pkg/config"
	"github.com/yusufcanb/tlama/pkg/doctor"
	prompt2 "github.com/yusufcanb/tlama/pkg/prompt"
	"github.com/yusufcanb/tlama/pkg/setup"

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
		Action:      prompt2.PromptAction,
		Commands: []*cli.Command{
			&cli.Command{
				Name:  "version",
				Usage: "Print tlama version.",
				Action: func(c *cli.Context) error {
					cli.ShowVersion(c)
					return nil
				},
			},
			config.GetCommand(),
			setup.GetCommand(),
			doctor.GetCommand(),
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
