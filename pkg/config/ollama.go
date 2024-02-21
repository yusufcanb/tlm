package config

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/briandowns/spinner"
	ollama "github.com/jmorganca/ollama/api"
	"github.com/yusufcanb/tlama/pkg/shell"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Ollama struct {
	cfg *TlamaConfig
	Api *ollama.Client
}

func (o *Ollama) createVolume() error {
	cmd := exec.Command("docker", "volume", "create", "ollama")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func (o *Ollama) isContainerRunning(name string) (bool, error) {
	cmd := exec.Command("docker", "ps", "-aqf", "name=ollama")
	out, err := cmd.Output()
	if err != nil {
		return false, err
	}
	return strings.TrimSpace(string(out)) != "", nil
}

func (o *Ollama) removeContainer(name string) error {
	cmd := exec.Command("docker", "rm", "-f", name)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
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
	cmd := exec.Command("docker", "volume", "inspect", "ollama")
	err := cmd.Run()
	return err == nil // Returns true if the volume exists, false otherwise
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
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (o *Ollama) Install() error {
	var s *spinner.Spinner

	// 1. Check ollama volume exists
	if !o.IsVolumeInstalled() {
		fmt.Println("- Ollama volume not found. Creating a new volume.")
		if err := o.createVolume(); err != nil {
			return fmt.Errorf("error creating Ollama volume: %v", err)
		}
	} else {
		fmt.Println("- Ollama volume found. Using existing volume.")
	}

	// 2. Check ollama container exists
	containerExists, err := o.isContainerRunning("ollama")
	if err != nil {
		return fmt.Errorf("error checking for existing container: %v", err)
	}

	if containerExists {
		if err := o.removeContainer("ollama"); err != nil {
			return fmt.Errorf("error removing existing container: %v", err)
		}
	}

	// 3. Run the Docker command
	s = nil
	s = spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = " Creating Ollama container. (might take a few minutes)"
	s.Start()
	cmd, _, stderr := shell.Exec2("docker run -d --gpus=all -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama")

	err = cmd.Run()
	if err != nil {
		return errors.New(stderr.String())
	}
	s.Stop()
	fmt.Println("- Creating Ollama container. done")

	// 4. Pull CodeLlama model
	s = nil
	s = spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = " Downloading CodeLLaMa. (might take a few minutes)"
	s.Start()
	err = o.Pull()
	if err != nil {
		return err
	}
	s.Stop()
	fmt.Println("- Downloading CodeLLaMa. done")

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
