package app

import (
	"tlama/pkg/api"
	"tlama/pkg/config"
	"tlama/pkg/doctor"
	prompt2 "tlama/pkg/prompt"
	"tlama/pkg/setup"

	"github.com/urfave/cli/v2"
)

type TlamaApp struct {
	App    *cli.App
	Config *config.TlamaConfig
}

func New() *TlamaApp {
	var prompt string

	cliApp := &cli.App{
		Name:        "tlama",
		Usage:       "tllama -p \"List all go files in the current directory\"",
		Description: "tlama is a command line tool to provide terminal intelligence locally with LLaMa",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "prompt",
				Aliases:     []string{"p"},
				Usage:       "Prompt for a task",
				Destination: &prompt,
			},
		},
		Action: prompt2.PromptAction, // Default action is prompt
		Commands: []*cli.Command{
			config.GetCommand(),
			setup.GetCommand(),
			doctor.GetCommand(),
		},
	}

	cliApp.Metadata = make(map[string]interface{})
	cliApp.Metadata["config"] = config.New()
	cliApp.Metadata["api"] = api.New(cliApp.Metadata["config"].(*config.TlamaConfig))

	return &TlamaApp{
		App: cliApp,
	}
}
