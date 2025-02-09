package app

import (
	_ "embed"
	"fmt"
	"runtime"

	ollama "github.com/jmorganca/ollama/api"
	"github.com/yusufcanb/tlm/pkg/ask"
	"github.com/yusufcanb/tlm/pkg/config"
	"github.com/yusufcanb/tlm/pkg/explain"
	"github.com/yusufcanb/tlm/pkg/suggest"

	"github.com/urfave/cli/v2"
)

var usageText string = `tlm suggest "<prompt>"
tlm s --model=qwen2.5-coder:1.5b --style=stable "<prompt>"

tlm explain "<command>" # explain a command 
tlm e --model=llama3.2:1b --style=balanced "<command>" # explain a command with a overrided model

tlm ask "<prompt>" # ask a question
tlm ask --context . --include *.md "<prompt>" # ask a question with a context`

type TlmApp struct {
	App *cli.App
}

func New(version, buildSha string) *TlmApp {
	o, _ := ollama.ClientFromEnvironment()

	con := config.New(o)
	con.LoadOrCreateConfig()

	sug := suggest.New(o, version)
	exp := explain.New(o, version)
	ask := ask.New(o, version)

	cliApp := &cli.App{
		Name:            "tlm",
		Usage:           "terminal copilot, powered by open-source models.",
		UsageText:       usageText,
		Version:         version,
		CommandNotFound: notFound,
		Before:          beforeRun(o),
		After:           afterRun(version),
		Action: func(c *cli.Context) error {
			return cli.ShowAppHelp(c)
		},
		Commands: []*cli.Command{
			ask.Command(),
			sug.Command(),
			exp.Command(),
			con.Command(),
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Prints tlm version.",
				Action: func(c *cli.Context) error {
					fmt.Printf("tlm %s (%s) on %s/%s\n", version, buildSha, runtime.GOOS, runtime.GOARCH)
					return nil
				},
			},
		},
		Metadata: map[string]interface{}{
			"releaseManager": NewReleaseManager("yusufcanb", "tlm"),
		},
	}

	app := &TlmApp{
		App: cliApp,
	}

	return app
}
