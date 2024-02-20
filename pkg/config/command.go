package config

import (
	"fmt"
	"strconv"

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

			val := viper.Get(c.Args().Get(0))
			if val == "" {
				fmt.Println(fmt.Sprintf("ERROR: %s not found", c.Args().Get(0)))
				return nil
			}

			fmt.Println(val)
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
			key := c.Args().Get(0)
			// check key is in the list
			if key != "llm.host" && key != "llm.model" && key != "llm.parameters.temperature" && key != "llm.parameters.top_p" {
				fmt.Println(fmt.Sprintf("%s is not a tlm parameter", key))
				return nil
			}

			if key == "llm.parameters.temperature" || key == "llm.parameters.top_p" {
				value, err := strconv.ParseFloat(c.Args().Get(1), 64)
				viper.Set(key, value)

				if err != nil {
					fmt.Println(fmt.Sprintf("%s is not a valid float value", c.Args().Get(1)))
					return nil
				}
				return nil
			}

			viper.Set(key, c.Args().Get(1))
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
		Name:    "config",
		Aliases: []string{"c"},
		Usage:   "Configure tlama parameters.",
		Action: func(c *cli.Context) error {
			return nil
		},
		Subcommands: []*cli.Command{
			configGetCommand(),
			configSetCommand(),
		},
	}
}
