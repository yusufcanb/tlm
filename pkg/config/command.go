package config

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

func printAllConfig() {
	for _, k := range viper.AllKeys() {
		fmt.Println(fmt.Sprintf("%s = %s", k, viper.GetString(k)))
	}
}

func configGetCommand() *cli.Command {
	return &cli.Command{
		Name:      "get",
		Aliases:   []string{"g"},
		Args:      true,
		ArgsUsage: "<key> get from config file",
		HelpName:  "config get",
		Action: func(c *cli.Context) error {

			arg := c.Args().Get(0)
			if arg == "" {
				printAllConfig()
				return nil
			}

			asd := viper.Get(c.Args().Get(0)).(string)
			if asd == "" {
				fmt.Println(fmt.Sprintf("ERROR: %s not found", c.Args().Get(0)))
				return nil
			}

			fmt.Println(asd)
			return nil
		},
	}
}

func configSetCommand() *cli.Command {
	return &cli.Command{
		Name:      "set",
		Args:      true,
		ArgsUsage: "<key> - <value> set to config file",
		Action: func(c *cli.Context) error {
			viper.Set(c.Args().Get(0), c.Args().Get(1))
			err := viper.WriteConfig()
			if err != nil {
				return err
			}

			printAllConfig()
			return nil
		},
	}
}

func GetCommand() *cli.Command {
	return &cli.Command{
		Name:  "config",
		Usage: "Configure tlama parameters.",
		Action: func(c *cli.Context) error {
			return nil
		},
		Subcommands: []*cli.Command{
			configGetCommand(),
			configSetCommand(),
		},
	}
}
