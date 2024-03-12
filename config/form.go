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
				Title("Shell").
				Description("Overrides platform's shell for suggestions").
				Options(
					huh.NewOption("Automatic", "auto"),
					huh.NewOption("Powershell (Windows)", "powershell"),
					huh.NewOption("Bash (Linux)", "bash"),
					huh.NewOption("Zsh (macOS)", "zsh"),
				).
				Value(&c.shell),

			huh.NewSelect[string]().
				Title("Suggestion Preference").
				Description("Sets preference for command suggestions").
				Options(
					huh.NewOption("Precise", "stable"),
					huh.NewOption("Balanced", "balanced"),
					huh.NewOption("Creative", "creative"),
				).
				Value(&c.suggest),

			huh.NewSelect[string]().
				Title("Explain Preference").
				Description("Sets preference for command explanations").
				Options(
					huh.NewOption("Precise", "stable"),
					huh.NewOption("Balanced", "balanced"),
					huh.NewOption("Creative", "creative"),
				).
				Value(&c.explain),
		),
	)

	return c.form.WithTheme(huh.ThemeBase16()).Run()
}
