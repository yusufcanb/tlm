package app

import (
	_ "embed"
	"fmt"
	ollama "github.com/jmorganca/ollama/api"
	"github.com/yusufcanb/tlm/config"
	"github.com/yusufcanb/tlm/explain"
	"github.com/yusufcanb/tlm/install"
	"github.com/yusufcanb/tlm/suggest"
	"io/fs"
	"runtime"

	"github.com/urfave/cli/v2"
)

//go:embed Modelfile.explain
var explainModelfile string

//go:embed Modelfile.suggest
var suggestModelfile string

type TlmApp struct {
	writer *fs.File

	explainModelfile string
	suggestModelfile string

	App *cli.App
}

func New(version, buildSha string) *TlmApp {
	con := config.New()
	con.LoadOrCreateConfig()

	o, _ := ollama.ClientFromEnvironment()
	sug := suggest.New(o)
	exp := explain.New(o)
	ins := install.New(o, suggestModelfile, explainModelfile)

	cliApp := &cli.App{
		Name:            "tlm",
		Usage:           "terminal copilot, powered by CodeLLaMa.",
		UsageText:       "tlm explain <command>\ntlm suggest <prompt>",
		Version:         version,
		CommandNotFound: notFound,
		Before:          beforeRun(),
		After:           afterRun(ins, version),
		Action: func(c *cli.Context) error {
			return cli.ShowAppHelp(c)
		},
		Commands: []*cli.Command{
			sug.Command(),
			exp.Command(),
			ins.DeployCommand(),
			con.Command(),
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Prints tlm version.",
				Action: func(c *cli.Context) error {
					fmt.Printf("tlm %s (%s) on %s/%s", version, buildSha, runtime.GOOS, runtime.GOARCH)
					return nil
				},
			},
		},
	}

	app := &TlmApp{
		App: cliApp,
	}

	return app
}
