package install

import (
	"fmt"
	"github.com/charmbracelet/huh"
)

type InstallForm2 struct {
	form *huh.Form

	redeploy     bool
	version      string
	gpuEnabled   bool
	ollamaImage  string
	ollamaVolume string

	suggestModelfile string
	explainModelfile string
}

func (i *InstallForm2) Run() error {

	i.form = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[bool]().
				Title("GPU Support").
				Options(
					huh.NewOption("Disable", false),
					huh.NewOption("Enable", true),
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
			Description(fmt.Sprintf("Ollama (%s) instance is running on 11434, redeploy?", i.version)).
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

	} else {
		return i.form.WithTheme(huh.ThemeBase16()).Run()
	}

}

func NewInstallForm2(version string, suggestModelfile string, explainModelfile string) *InstallForm2 {
	ollamaImage := "ollama:latest"
	ollamaVolume := "ollama"

	var redeploy bool
	if version != "" {
		redeploy = true
	} else {
		redeploy = false
	}

	return &InstallForm2{
		ollamaImage:  ollamaImage,
		ollamaVolume: ollamaVolume,
		redeploy:     redeploy,

		version:          version,
		suggestModelfile: suggestModelfile,
		explainModelfile: explainModelfile,
	}
}
