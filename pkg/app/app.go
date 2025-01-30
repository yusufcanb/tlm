package app

import (
	_ "embed"
	"fmt"
	"runtime"

	ollama "github.com/jmorganca/ollama/api"
	"github.com/yusufcanb/tlm/pkg/config"
	"github.com/yusufcanb/tlm/pkg/explain"
	"github.com/yusufcanb/tlm/pkg/suggest"

	"github.com/urfave/cli/v2"
)

type TlmApp struct {
	App *cli.App
}

func New(version, buildSha string) *TlmApp {
	o, _ := ollama.ClientFromEnvironment()

	con := config.New(o)
	con.LoadOrCreateConfig()

	sug := suggest.New(o, version)
	exp := explain.New(o, version)

	cliApp := &cli.App{
		Name:            "tlm",
		Usage:           "terminal copilot, powered by open-source models.",
		UsageText:       "tlm explain '<command>'\ntlm suggest '<prompt>'",
		Version:         version,
		CommandNotFound: notFound,
		Before:          beforeRun(o),
		After:           afterRun(version),
		Action: func(c *cli.Context) error {
			return cli.ShowAppHelp(c)
		},
		Commands: []*cli.Command{
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
