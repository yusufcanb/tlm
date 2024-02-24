package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
	"github.com/yusufcanb/tlm/shell"
)

func (c *Config) Action(_ *cli.Context) error {
	var err error

	form := ConfigForm{
		host:    viper.GetString("llm.host"),
		explain: viper.GetString("llm.explain"),
		suggest: viper.GetString("llm.suggest"),
	}

	err = form.Run()
	if err != nil {
		return err
	}

	viper.Set("shell", form.shell)
	viper.Set("llm.host", form.host)
	viper.Set("llm.explain", form.explain)
	viper.Set("llm.suggest", form.suggest)

	err = viper.WriteConfig()
	if err != nil {
		return err
	}

	fmt.Println(shell.Ok() + " configuration saved")
	return nil
}

func (c *Config) Command() *cli.Command {
	return &cli.Command{
		Name:    "config",
		Aliases: []string{"c"},
		Usage:   "configure preferences.",
		Action:  c.Action,
	}
}
