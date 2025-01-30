package app

import (
	_ "embed"
	"fmt"
	"io/fs"
	"runtime"

	ollama "github.com/jmorganca/ollama/api"
	"github.com/yusufcanb/tlm/pkg/config"
	"github.com/yusufcanb/tlm/pkg/explain"
	"github.com/yusufcanb/tlm/pkg/install"
	"github.com/yusufcanb/tlm/pkg/suggest"

	"github.com/urfave/cli/v2"
)

type TlmApp struct {
	writer *fs.File

	App *cli.App
}

func New(version, buildSha string) *TlmApp {
	con := config.New()
	con.LoadOrCreateConfig()

	o, _ := ollama.ClientFromEnvironment()
	sug := suggest.New(o, version)
	exp := explain.New(o, version)
	ins := install.New(o, sug, exp)

	cliApp := &cli.App{
		Name:            "tlm",
		Usage:           "terminal copilot, powered by CodeLLaMa.",
		UsageText:       "tlm explain '<command>'\ntlm suggest '<prompt>'",
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
