package explain

import (
	"context"
	"fmt"
	"github.com/jmorganca/ollama/api"
	"github.com/urfave/cli/v2"
	"github.com/yusufcanb/tlama/pkg/config"
)

func explainAction(c *cli.Context) error {
	cfg := c.App.Metadata["config"].(*config.TlamaConfig)
	ollama := cfg.GetOllamaApi()

	myResponseFunc := func(resp api.GenerateResponse) error {
		// Process the response here (e.g., print it)
		fmt.Print(resp.Response)
		return nil // Or return an error if processing fails
	}

	// Call the Generate function
	err := ollama.Api.Generate(context.Background(), &api.GenerateRequest{
		Model:  cfg.Llm.Model,
		Prompt: "Explain the command briefly: " + c.Args().First(),
		Options: map[string]interface{}{
			"num_predict": 128,
			"temperature": 0.1,
			"top_p":       0.25,
		},
	}, myResponseFunc)
	if err != nil {
		fmt.Println("Error during generation:", err)
	}
	return nil
}
