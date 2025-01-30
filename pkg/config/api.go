package config

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/spf13/viper"
	"github.com/yusufcanb/tlm/pkg/shell"
)

var (
	defaultSuggestionPolicy = "stable"
	defaultExplainPolicy    = "creative"
	defaultShell            = "auto"
)

func isExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (c *Config) LoadOrCreateConfig() {
	viper.SetConfigName(".tlm")
	viper.SetConfigType("yml")
	viper.AddConfigPath("$HOME")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	configPath := path.Join(homeDir, ".tlm.yml")
	if !isExists(configPath) {
		viper.Set("shell", defaultShell)
		viper.Set("llm.suggest", defaultSuggestionPolicy)
		viper.Set("llm.explain", defaultExplainPolicy)

		if err := viper.WriteConfigAs(path.Join(homeDir, ".tlm.yml")); err != nil {
			fmt.Printf(shell.Err()+" error writing config file, %s", err)
		}
	}

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

}
