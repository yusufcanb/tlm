package config

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	ollama "github.com/jmorganca/ollama/api"
	"github.com/yusufcanb/tlama/pkg/shell"
	"io"
	"net/http"
	"os"
	"strings"
)

type Ollama struct {
	cfg *TlamaConfig
	Api *ollama.Client
}

func (o *Ollama) IsInstalled() bool {
	resp, err := http.Get(o.cfg.Llm.Host)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return false
		}
		bodyString := string(bodyBytes)

		if bodyString == "Ollama is running" {
			return true
		}
	}

	return false
}

func (o *Ollama) IsModelExists() bool {
	fmt.Printf("- Checking %s model is installed...", o.cfg.Llm.Model)

	payload := map[string]string{"name": o.cfg.Llm.Model}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return false
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", o.cfg.Llm.Host+"/api/show", bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("\t\tno")
		return false
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("\t\tno")
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("\t\tok")
	}

	return resp.StatusCode == http.StatusOK
}

func (o *Ollama) IsVolumeInstalled() bool {
	_, stdout, _ := shell.Exec2("docker volume ls -q -f name=ollama")

	for _, volume := range strings.Split(stdout.String(), "\n") {
		if volume == "ollama" {
			return true
		}
	}

	return false
}

func (o *Ollama) InstallModel() error {
	payload := map[string]string{"name": "codellama:34b"}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Construct the request
	req, err := http.NewRequest("POST", o.cfg.Llm.Host+"/api/pull", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Handle streaming response
	decoder := json.NewDecoder(resp.Body)
	for {
		var responseChunk map[string]interface{} // Adjust the type if your API sends structured data
		if err := decoder.Decode(&responseChunk); err != nil {
			if err == io.EOF {
				break // End of stream
			}
			return err // Error during decoding
		}
		// Log the response chunk
		fmt.Println(responseChunk["status"]) // Replace with your preferred logging
	}

	return nil // Installation successful (adjust if needed)
}

func (o *Ollama) List() error {
	ctx := context.Background()
	list, err := o.Api.List(ctx)
	if err != nil {
		return err
	}

	for _, model := range list.Models {
		fmt.Println(model.Name)
	}
	return nil
}

func (o *Ollama) Pull() error {
	ctx := context.Background()
	err := o.Api.Pull(ctx, &ollama.PullRequest{
		Model: o.cfg.Llm.Model,
	}, func(response ollama.ProgressResponse) error {
		fmt.Println(response.Status)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (o *Ollama) Install() error {
	ok := o.IsVolumeInstalled()
	if ok {
		fmt.Println("Ollama volume found. Using existing volume.")
	}

	_, stdout, stderr := shell.Exec2("docker run -d --gpus=all -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama\n")

	if stderr.String() != "" {
		fmt.Println(stderr.String())
		return errors.New(stderr.String())
	}

	fmt.Println(stdout.String())

	return nil
}

func NewOllama(cfg *TlamaConfig) *Ollama {
	os.Setenv("OLLAMA_HOST", cfg.Llm.Host)
	api, err := ollama.ClientFromEnvironment()
	if err != nil {
		panic(err)
	}

	return &Ollama{
		cfg: cfg,
		Api: api,
	}
}
