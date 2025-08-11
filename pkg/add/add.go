package add

import (
	"fmt"
	"io/ioutil"

	"github.com/urfave/cli/v2"
	"github.com/yusufcanb/tlm/pkg/chroma"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:      "add",
		Usage:     "Adds a document to the ChromaDB collection.",
		ArgsUsage: "<file>",
		Action: func(c *cli.Context) error {
			if c.NArg() == 0 {
				return fmt.Errorf("a file path is required")
			}

			filePath := c.Args().First()
			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("failed to read file: %w", err)
			}

			chromaClient := chroma.NewChromaClient("http://localhost:8000")
			err = chromaClient.Add("tlm-collection", &chroma.AddRequest{
				Documents: []string{string(content)},
				IDs:       []string{filePath},
			})
			if err != nil {
				return fmt.Errorf("failed to add document: %w", err)
			}

			fmt.Printf("Document '%s' added successfully.\n", filePath)
			return nil
		},
	}
}
