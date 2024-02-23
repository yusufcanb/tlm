package install

import (
	"context"
	"errors"
	"fmt"
	"github.com/charmbracelet/huh/spinner"
	ollama "github.com/jmorganca/ollama/api"
	"github.com/yusufcanb/tlm/shell"
	"os"
	"os/exec"
	"strings"
)

func (i *Install) isVolumeInstalled(volumeName string) bool {
	cmd := exec.Command("docker", "volume", "inspect", volumeName)
	err := cmd.Run()
	return err == nil
}

func (i *Install) isContainerRunning(containerName string) (bool, error) {
	cmd := exec.Command("docker", "ps", "-aqf", fmt.Sprintf("name=%s", containerName))
	out, err := cmd.Output()
	if err != nil {
		return false, err
	}
	return strings.TrimSpace(string(out)) != "", nil
}

func (i *Install) createVolume(volumeName string) error {
	if !i.isVolumeInstalled(volumeName) {
		cmd := exec.Command("docker", "volume", "create", volumeName)
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}

func (i *Install) installModelfile(name, modelfile string) error {
	var err error
	err = i.api.Create(context.Background(), &ollama.CreateRequest{Model: name, Modelfile: modelfile}, func(res ollama.ProgressResponse) error {
		return nil
	})
	return err
}

func (i *Install) removeContainer(containerName string) error {
	cmd := exec.Command("docker", "rm", "-f", containerName)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func (i *Install) createContainer(containerName string, gpuEnabled bool) (string, error) {
	var cmd *exec.Cmd

	withCPU := []string{"run", "-d", "--gpus=all",
		"-v", "ollama:/root/.ollama", "-p", "11434:11434", "--name", containerName, "ollama/ollama"}

	withGPU := []string{"run", "-d", "--gpus=all",
		"-v", "ollama:/root/.ollama", "-p", "11434:11434", "--name", containerName, "ollama/ollama"}

	if gpuEnabled {
		cmd = exec.Command("docker", withGPU...)
	} else {
		cmd = exec.Command("docker", withCPU...)

	}

	cmd = exec.Command("docker", "run", "-d", "--gpus=all",
		"-v", "ollama:/root/.ollama", "-p", "11434:11434", "--name", containerName, "ollama/ollama")

	stdout, stderr := cmd.CombinedOutput()
	if stderr != nil {
		return "", errors.New(stderr.Error())
	}

	return string(stdout), nil
}

func (i *Install) installAndConfigureOllama(form *InstallForm2) error {

	// 1. Check if Ollama volume exists
	i.checkOllamaVolumeExists(form)

	// 2. Check if Ollama container exists
	containerExists, err := i.isContainerRunning(i.defaultContainerName)
	if err != nil {
		return fmt.Errorf("- Checking for existing container: %s", shell.Err())
	}

	// 3. Remove old container (if it exists) and recreate
	if containerExists {
		fmt.Println("- Existing Ollama container found, removing and re-creating. " + shell.Ok())
		if err := i.removeContainer(i.defaultContainerName); err != nil {
			return fmt.Errorf("error removing existing container: %v", err)
		}
	}

	// 4. Run the Docker command
	_ = spinner.New().Type(spinner.Line).Title(" Creating Ollama container").Action(func() {
		_, err = i.createContainer(i.defaultContainerName, form.gpuEnabled)
		if err != nil {
			fmt.Printf("- Creating Ollama container. %s", shell.Err())
			fmt.Println(err.Error())
			os.Exit(-1)
		}
	}).Run()
	fmt.Println("- Creating Ollama container. " + shell.Ok())

	// 5. Pull CodeLLaMa if not exists
	i.installCodeLLaMa()

	// 6. Install the modelfile (Suggest)
	_ = spinner.New().Type(spinner.Line).Title(" Creating Modelfile for suggestions").Action(func() {
		err = i.installModelfile("suggest:7b", form.suggestModelfile)
		if err != nil {
			fmt.Println("- Creating Modelfile for suggestions. " + shell.Err())
			os.Exit(-1)
		}
	}).Run()
	fmt.Println("- Creating Modelfile for suggestions. " + shell.Ok())

	// 7. Install the modelfile (Suggest)
	_ = spinner.New().Type(spinner.Line).Title(" Creating Modelfile for explanations").Action(func() {
		err = i.installModelfile("suggest:7b", form.suggestModelfile)
		if err != nil {
			fmt.Println("- Creating Modelfile for explanations. " + shell.Err())
			os.Exit(-1)
		}
	}).Run()
	fmt.Println("- Creating Modelfile for explanations. " + shell.Ok())

	return nil
}

func (i *Install) installCodeLLaMa() {
	var err error
	_ = spinner.New().Type(spinner.Line).Title(" Getting latest CodeLLaMa").Action(func() {
		err = i.api.Pull(context.Background(), &ollama.PullRequest{Model: "codellama:7b"}, func(res ollama.ProgressResponse) error {
			return nil
		})
		if err != nil {
			fmt.Println("- Installing CodeLLaMa. " + shell.Err())
			os.Exit(-1)
		}
	}).Run()

	fmt.Println("- Getting latest CodeLLaMa. " + shell.Ok())
}

func (i *Install) checkOllamaVolumeExists(form *InstallForm2) {
	if !i.isVolumeInstalled(form.ollamaVolume) {
		fmt.Println("- Ollama volume not found, creating a new volume. " + shell.Ok())
		if err := i.createVolume(form.ollamaVolume); err != nil {
			fmt.Printf("- Error creating volume: %v", err)
			os.Exit(-1)
		}
	} else {
		fmt.Println("- Ollama volume found, using existing volume. " + shell.Ok())
	}
}
