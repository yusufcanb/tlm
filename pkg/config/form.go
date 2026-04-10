package config

import (
	"context"
	"fmt"
	"sort"

	"github.com/charmbracelet/huh"
	ollama "github.com/jmorganca/ollama/api"
)

type ConfigForm struct {
	form *huh.Form

	model   string
	shell   string
	explain string
	suggest string
}

func normalizeConfigFormValues(modelNames []string, model, shellValue, suggest, explain string) (string, string, string, string) {
	if len(modelNames) > 0 && !contains(modelNames, model) {
		model = modelNames[0]
	}

	validShells := []string{"auto", "powershell", "bash", "zsh"}
	if !contains(validShells, shellValue) {
		shellValue = defaultShell
	}

	validStyles := []string{Stable, Balanced, Creative}
	if !contains(validStyles, suggest) {
		suggest = defaultSuggestionPolicy
	}
	if !contains(validStyles, explain) {
		explain = defaultExplainPolicy
	}

	return model, shellValue, suggest, explain
}

func contains(values []string, value string) bool {
	for _, candidate := range values {
		if candidate == value {
			return true
		}
	}
	return false
}

func (c *ConfigForm) Run(api *ollama.Client) error {

	// get available models from Ollama
	models, err := api.List(context.Background())
	if err != nil {
		fmt.Printf("Error fetching models: %v\n", err)
		return err
	}

	if len(models.Models) == 0 {
		return fmt.Errorf("no Ollama models found. Run `ollama pull <model_name>` first")
	}

	sort.Slice(models.Models, func(i, j int) bool {
		return models.Models[i].Name < models.Models[j].Name
	})

	modelNames := make([]string, 0, len(models.Models))
	for _, model := range models.Models {
		modelNames = append(modelNames, model.Name)
	}

	c.model, c.shell, c.suggest, c.explain = normalizeConfigFormValues(
		modelNames,
		c.model,
		c.shell,
		c.suggest,
		c.explain,
	)

	// create model options from available Ollama models
	modelOptions := make([]huh.Option[string], 0, len(models.Models))
	for _, model := range models.Models {
		modelOptions = append(modelOptions, huh.NewOption(
			fmt.Sprintf("%s (%.2f GB)", model.Name, float64(model.Size)/(1024*1024*1024)),
			model.Name,
		))
	}

	c.form = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Model Selection").
				Description("Sets a default model from the list of all available models.\nUse `ollama pull <model_name>` to download new models.\n").
				Options(
					modelOptions...,
				).
				Value(&c.model),

			huh.NewSelect[string]().
				Title("Shell").
				Description("Overrides platform's shell for suggestions.\n").
				Options(
					huh.NewOption("Automatic", "auto"),
					huh.NewOption("Powershell (Windows)", "powershell"),
					huh.NewOption("Bash (Linux)", "bash"),
					huh.NewOption("Zsh (macOS)", "zsh"),
				).
				Value(&c.shell),

			huh.NewSelect[string]().
				Title("Suggestion Style").
				Description("Sets style for command suggestions. \n").
				Options(
					huh.NewOption("Precise", Stable),
					huh.NewOption("Balanced", Balanced),
					huh.NewOption("Creative", Creative),
				).
				Value(&c.suggest),

			huh.NewSelect[string]().
				Title("Explain Style").
				Description("Sets style for command explanations. \n").
				Options(
					huh.NewOption("Precise", Stable),
					huh.NewOption("Balanced", Balanced),
					huh.NewOption("Creative", Creative),
				).
				Value(&c.explain),
		),
	)

	return c.form.WithTheme(huh.ThemeBase16()).Run()
}
