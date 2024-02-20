package config

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	ollama "github.com/jmorganca/ollama/api"
	"io"
	"net/http"
	"os"
)

type Ollama struct {
	cfg *TlamaConfig
	Api *ollama.Client
}

func (o *Ollama) IsInstalled() bool {
	fmt.Print("- Checking LLaMa is installed...")

	resp, err := http.Get(o.cfg.Llm.Host)
	if err != nil {
		fmt.Println("\t\tno")
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("\t\tno")
			return false
		}
		bodyString := string(bodyBytes)

		if bodyString == "Ollama is running" {
			fmt.Println("\t\tok")
			return true
		}
	}

	fmt.Println("\t\tno")
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
