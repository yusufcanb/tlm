package config

import (
	"errors"
	"fmt"
	"sort"

	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
	"github.com/yusufcanb/tlm/pkg/shell"
)

var (
	shellKey      = "shell"
	llmModelKey   = "llm.model"
	llmExplainKey = "llm.explain"
	llmSuggestKey = "llm.suggest"
)

func (c *Config) before(_ *cli.Context) error {
	return nil
}

func (c *Config) subCommandGet() *cli.Command {
	return &cli.Command{
		Name:      "get",
		Usage:     "get configuration by key",
		UsageText: "tlm config get <key>",
		Action: func(c *cli.Context) error {
			key := c.Args().Get(0)
			value := viper.GetString(key)

			if value == "" {
				fmt.Println(fmt.Sprintf("%s <%s> is not a tlm parameter", shell.Err(), key))
				return nil
			}

			fmt.Println(fmt.Sprintf("%s = %s", key, value))
			return nil
		},
	}
}

func (c *Config) subCommandSet() *cli.Command {
	return &cli.Command{
		Name:  "set",
		Usage: "set configuration with key and value",
		Action: func(c *cli.Context) error {
			key := c.Args().Get(0)

			switch key {
			case llmSuggestKey, llmExplainKey:
				mode := c.Args().Get(1)
				if mode != "stable" && mode != "balanced" && mode != "creative" {
					return errors.New("invalid mode: " + mode)
				}
				viper.Set(key, mode)

			case shellKey:
				s := c.Args().Get(1)
				if s != "bash" && s != "zsh" && s != "auto" && s != "powershell" {
					return errors.New("invalid shell: " + c.Args().Get(1))
				}
				viper.Set(shellKey, s)

			default:
				fmt.Println(fmt.Sprintf("%s <%s> is not a tlm parameter", shell.Err(), key))
				return nil
			}

			viper.Set(key, c.Args().Get(1))
			err := viper.WriteConfig()
			if err != nil {
				return err
			}

			fmt.Println(fmt.Sprintf("%s = %s  %s", key, c.Args().Get(1), shell.Ok()))
			return nil
		},
	}
}

func (c *Config) subCommandLs() *cli.Command {
	return &cli.Command{
		Name:  "ls",
		Usage: "list all configuration",
		Action: func(c *cli.Context) error {
			keys := viper.AllKeys()
			sort.Strings(keys)
			for _, key := range keys {
				fmt.Printf(fmt.Sprintf("%s = %s\n", key, viper.GetString(key)))
			}
			return nil
		},
	}
}

func (c *Config) action(_ *cli.Context) error {
	var err error
	form := ConfigForm{
		model:   viper.GetString(llmModelKey),
		shell:   viper.GetString(shellKey),
		explain: viper.GetString(llmExplainKey),
		suggest: viper.GetString(llmSuggestKey),
	}

	err = form.Run(c.api)
	if err != nil {
		return err
	}

	viper.Set(shellKey, form.shell)
	viper.Set(llmModelKey, form.model)
	viper.Set(llmExplainKey, form.explain)
	viper.Set(llmSuggestKey, form.suggest)

	err = viper.WriteConfig()
	if err != nil {
		return err
	}

	fmt.Println("configuration saved. " + shell.Ok())
	return nil
}

func (c *Config) Command() *cli.Command {
	return &cli.Command{
		Name:    "config",
		Aliases: []string{"c"},
		Usage:   "Configures language model, style and shell",
		Before:  c.before,
		Action:  c.action,
		Subcommands: []*cli.Command{
			c.subCommandGet(),
			c.subCommandSet(),
			c.subCommandLs(),
		},
	}
}
