package install

import (
	"bytes"
	"fmt"
	"github.com/yusufcanb/tlama/pkg/shell"
)

func installOllama(gpuSupport bool) error {
	var stdout, stderr bytes.Buffer
	var cmdStr string

	if gpuSupport {
		cmdStr = "docker run -d --gpus=all -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama"
	} else {
		cmdStr = "docker run -d -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama"
	}

	cmd := shell.Exec(cmdStr)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(stderr.String())
		return err
	} else {
		fmt.Println(stdout.String())
	}

	return nil
}
