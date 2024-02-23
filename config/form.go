package config

import "github.com/charmbracelet/huh"

type ConfigForm struct {
	form *huh.Form

	host    string
	shell   string
	explain string
	suggest string
}

func (c *ConfigForm) Run() error {
	c.form = huh.NewForm(
		huh.NewGroup(

			huh.NewInput().
				Title("Ollama").
				Value(&c.host),

			huh.NewSelect[string]().
				Title("Default Shell (Windows)").
				Options(
					huh.NewOption("Windows Powershell", "powershell"),
					huh.NewOption("Windows Command Prompt", "cmd"),
				).
				Value(&c.shell),

			huh.NewSelect[string]().
				Title("Suggestion Preference").
				Description("This sets how suggestions should be in placed").
				Options(
					huh.NewOption("Stable", "stable"),
					huh.NewOption("Balanced", "balanced"),
					huh.NewOption("Creative", "creative"),
				).
				Value(&c.explain),

			huh.NewSelect[string]().
				Title("Explain Preference").
				Description("This configuration sets explain responses").
				Options(
					huh.NewOption("Stable", "stable"),
					huh.NewOption("Balanced", "balanced"),
					huh.NewOption("Creative", "creative"),
				).
				Value(&c.suggest),
		),
	)

	return c.form.WithTheme(huh.ThemeBase16()).Run()
}
