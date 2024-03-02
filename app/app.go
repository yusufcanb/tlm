package app

import (
	_ "embed"
	"fmt"
	ollama "github.com/jmorganca/ollama/api"
	"github.com/yusufcanb/tlm/config"
	"github.com/yusufcanb/tlm/explain"
	"github.com/yusufcanb/tlm/install"
	"github.com/yusufcanb/tlm/shell"
	"github.com/yusufcanb/tlm/suggest"
	"os"
	"runtime"

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
		Name:      "tlm",
		Usage:     "terminal copilot, powered by CodeLLaMa.",
		UsageText: "tlm explain <command>\ntlm suggest <prompt>",
		Version:   version,
		CommandNotFound: func(context *cli.Context, s string) {
			fmt.Println(shell.Err() + " command not found.")
			os.Exit(-1)
		},
		Action: func(c *cli.Context) error {
			return cli.ShowAppHelp(c)
		},
		Commands: []*cli.Command{
			sug.Command(),
			exp.Command(),
			ins.Command(),
			con.Command(),
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Prints tlm version.",
				Action: func(c *cli.Context) error {
					fmt.Printf("tlm %s (%s/%s)", version, runtime.GOOS, runtime.GOARCH)
					return nil
				},
			},
		},
	}

	return &TlmApp{
		App: cliApp,
	}
}
