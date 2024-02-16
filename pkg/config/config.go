package config

import (
	"log"
	"os"
	"path"

	"github.com/spf13/viper"
	"github.com/yusufcanb/tlama/pkg/shell"
)

func isExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func loadOrCreateConfig() (*TlamaConfig, error) {
	viper.SetConfigName(".tlama")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	configPath := path.Join(homeDir, ".tlama.yaml")
	if !isExists(configPath) {
		viper.Set("shell", shell.GetShell())
		viper.Set("llm.host", defaultLLMHost)
		viper.Set("llm.model", defaultLLMModel)
		viper.Set("llm.parameters.temperature", defaultTemperature)
		viper.Set("llm.parameters.top_p", defaultTopP)

		if err := viper.WriteConfigAs(path.Join(homeDir, ".tlama.yaml")); err != nil {
			return nil, err
		}
	}

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	tlamaConfig := &TlamaConfig{}
	err = tlamaConfig.LoadConfig()
	if err != nil {
		return nil, err
	}

	return tlamaConfig, nil
}

func New() *TlamaConfig {
	cfg, err := loadOrCreateConfig()
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
