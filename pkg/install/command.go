package install

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/yusufcanb/tlama/pkg/config"
	"github.com/yusufcanb/tlama/pkg/shell"
)

func GetCommand() *cli.Command {
	return &cli.Command{
		Name:    "install",
		Aliases: []string{"i"},
		Usage:   "Install LLM to your system.",
		Action: func(c *cli.Context) error {

			cfg := c.App.Metadata["config"].(*config.TlamaConfig)

			ollama := cfg.GetOllamaApi()
			if ollama.IsInstalled() {
				confirm := shell.Confirm("Ollama is already deployed and running, proceed?", 3)
				if !confirm {
					return nil
				}
			}

			gpuSupport := shell.Confirm("Enable GPU support? (only NVIDIA GPUs are supported)", 3)
			//proceed := shell.Confirm("\n- Image: ollama:latest\n- Model: codellama:7b\n\nLLaMa will be deployed and model will be pulled for the first time.\nThis process might take a few minutes depending on your network speed.\nProceed?", 3)
			proceed := shell.Confirm(`
- Image: ollama:latest
- Model: codellama:7b
- Volume: ollama

LLaMa will be deployed and model will be pulled for the first time.
This process might take a few minutes depending on your network speed.

Continue?`, 3)
			if !proceed {
				return nil
			}

			fmt.Printf("Deploying Ollama...")
			err := installOllama(gpuSupport)
			if err != nil {
				return err
			}

			fmt.Println("Done...")
			return nil
		},
	}
}
