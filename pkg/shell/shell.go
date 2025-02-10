package shell

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/exec"

	"github.com/charmbracelet/lipgloss"
	ollama "github.com/jmorganca/ollama/api"
)

var Version string

func Ok() string {
	style := lipgloss.NewStyle()

	style = style.Bold(true)
	style = style.Foreground(lipgloss.Color("2"))

	return style.Render("(ok)")
}

func SuccessMessage(message string) string {
	style := lipgloss.NewStyle()
	style = style.Foreground(lipgloss.Color("2"))
	return style.Render(message)
}

func WarnMessage(message string) string {
	style := lipgloss.NewStyle()
	style = style.Foreground(lipgloss.Color("202"))
	return style.Render(message)
}

func Err() string {
	style := lipgloss.NewStyle()

	style = style.Bold(true)
	style = style.Foreground(lipgloss.Color("9"))

	return style.Render("(err)")
}

func Warn() string {
	style := lipgloss.NewStyle()

	style = style.Bold(true)
	style = style.Foreground(lipgloss.Color("202"))

	return style.Render("(warn)")
}

func Exec2(command string) (*exec.Cmd, *bytes.Buffer, *bytes.Buffer) {
	var stdout, stderr bytes.Buffer
	var cmd *exec.Cmd

	if GetShell() == "powershell" {
		cmd = exec.Command(GetShell(), "-Command", command)
	} else {
		cmd = exec.Command(GetShell(), "-c", command)
	}

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	return cmd, &stdout, &stderr
}

func CheckOllamaIsUp(api *ollama.Client) error {
	_, err := api.Version(context.Background())
	if err != nil {
		return fmt.Errorf(" Ollama connection failed. Please check your Ollama if it's running or configured correctly. \n%s", err.Error())
	}
	return nil
}

func CheckOllamaIsSet() error {
	host := os.Getenv("OLLAMA_HOST")
	if host == "" {
		return errors.New("OLLAMA_HOST environment variable is not set.")
	}

	// parse url
	u, err := url.Parse(host)
	if err != nil {
		return fmt.Errorf("OLLAMA_HOST url is not valid: %v", err)
	}

	// check if scheme is http or https
	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.New("OLLAMA_HOST must use http or https protocol.")
	}

	return nil
}
