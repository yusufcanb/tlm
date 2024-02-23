package install

import (
	"github.com/charmbracelet/huh"
)

type InstallForm2 struct {
	form *huh.Form

	redeploy     bool
	gpuEnabled   bool
	ollamaImage  string
	ollamaVolume string
}

func (i *InstallForm2) Run() error {

	i.form = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[bool]().
				Title("GPU Support").
				Options(
					huh.NewOption("Enable", true),
					huh.NewOption("Disable", false),
				).
				Value(&i.gpuEnabled),

			huh.NewInput().
				Title("Ollama Image").
				Value(&i.ollamaImage),

			huh.NewInput().
				Title("Ollama Volume").
				Value(&i.ollamaVolume),
		),
	)

	if i.redeploy {
		var c bool
		err := huh.NewConfirm().
			Title("Redeploy").
			Description("An Ollama instance is running on 11434, redeploy?").
			Affirmative("Proceed").
			Negative("Abort").
			Value(&c).
			WithTheme(huh.ThemeBase16()).
			Run()

		if err != nil {
			return err
		}

		if c {
			return i.form.WithTheme(huh.ThemeBase16()).Run()
		}

		return nil

	}

	return nil
}

func NewInstallForm2() *InstallForm2 {
	ollamaImage := "ollama:latest"
	ollamaVolume := "ollama"

	return &InstallForm2{
		ollamaImage:  ollamaImage,
		ollamaVolume: ollamaVolume,
		redeploy:     true,
	}
}
