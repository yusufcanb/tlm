package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
)

var defaultLLMModel = "codellama:7b"
var defaultLLMHost = "http://localhost:11434"
var defaultTemperature = 0.2
var defaultTopP = 0.9

type llmConfig struct {
	Host       string
	Model      string
	Parameters llmParametersConfig
}

type llmParametersConfig struct {
	Temperature float64
	TopP        float64
}

type TlamaConfig struct {
	Shell string
	Llm   llmConfig
}

func (t *TlamaConfig) SaveConfig() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	if err := viper.WriteConfigAs(path.Join(homeDir, ".tlama.yaml")); err != nil {
		return err
	}
	return nil
}

func (t *TlamaConfig) LoadConfig() error {
	t.Shell = viper.Get("shell").(string)
	t.Llm = llmConfig{
		Host:  viper.Get("llm.host").(string),
		Model: viper.Get("llm.model").(string),
		Parameters: llmParametersConfig{
			Temperature: viper.Get("llm.parameters.temperature").(float64),
			TopP:        viper.Get("llm.parameters.top_p").(float64),
		},
	}
	return nil
}
