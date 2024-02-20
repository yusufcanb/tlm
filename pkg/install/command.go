package install

import (
	"bytes"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/yusufcanb/tlama/pkg/shell"
)

func GetCommand() *cli.Command {
	return &cli.Command{
		Name:    "install",
		Aliases: []string{"i"},
		Usage:   "Install LLM to your system.",
		Action: func(c *cli.Context) error {
			shell.Confirm("Enable GPU support? (only NVIDIA GPUs are supported)", 3)
			shell.Confirm("\n- Image: ollama:latest\n- Model: codellama:7b\n\nLLaMa will be deployed and model will be pulled for the first time.\nThis process might take a few minutes depending on your network speed.\nProceed?", 3)

			fmt.Printf("Deploying Ollama...")

			var stdout, stderr bytes.Buffer

			////cmd := shell.Exec("docker run -d -v ollama:/root/.ollama -p 11435:11435 --name ollama2 ollama/ollama")
			//cmd.Stdout = &stdout
			//cmd.Stderr = &stderr
			//
			//err := cmd.Run()
			//if err != nil {
			//	fmt.Println(stderr.String())
			//	return err
			//} else {
			//	fmt.Println(stdout.String())
			//}

			cmd := shell.Exec("docker exec -d 1cfeb4b8 /usr/bin/ollama pull codellama:7b")
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			err := cmd.Run()
			if err != nil {
				fmt.Println(stderr.String())
				return err
			} else {
				fmt.Println(stdout.String())
			}

			fmt.Println("Done...")
			return nil
		},
	}
}
