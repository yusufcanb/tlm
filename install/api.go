package install

import (
	"context"
	"errors"
	"fmt"
	ollama "github.com/jmorganca/ollama/api"
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

func (i *Install) installModelfile(modelfile string) error {
	return nil
}

func (i *Install) removeContainer(containerName string) error {
	cmd := exec.Command("docker", "rm", "-f", containerName)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func (i *Install) createContainer(containerName string) error {
	return nil
}

func (i *Install) installAndConfigureOllama(form *InstallForm2) error {
	fmt.Println("Installing and configuring Ollama...")

	// 1. Check if Ollama volume exists
	if !i.isVolumeInstalled(form.ollamaVolume) {
		fmt.Println("Ollama volume not found. Creating a new volume.")
		if err := i.createVolume(form.ollamaVolume); err != nil {
			return fmt.Errorf("error creating volume: %v", err)
		}
	} else {
		fmt.Println("Ollama volume found. Using existing volume.")
	}

	// 2. Check if Ollama container exists
	containerExists, err := i.isContainerRunning(i.defaultContainerName)
	if err != nil {
		return fmt.Errorf("error checking for existing container: %v", err)
	}

	// 3. Remove old container (if it exists) and recreate
	if containerExists {
		fmt.Println("Existing Ollama container found. Removing and recreating.")
		if err := i.removeContainer(i.defaultContainerName); err != nil {
			return fmt.Errorf("error removing existing container: %v", err)
		}
	}

	// 4. Run the Docker command
	cmd := exec.Command("docker", "run", "-d", "--gpus=all",
		"-v", "ollama:/root/.ollama", "-p", "11434:11434", "--name", "ollama", "ollama/ollama")

	stdout, stderr := cmd.CombinedOutput()
	if stderr != nil {
		return errors.New(stderr.Error())
	}
	fmt.Println(string(stdout))

	// 5. Pull CodeLLaMa if not exists
	onProgressResponse := func(res ollama.ProgressResponse) error {
		return nil
	}

	err = i.api.Pull(context.Background(), &ollama.PullRequest{Model: "codellama:7b"}, onProgressResponse)
	if err != nil {
		return err
	}

	// 6. Install the modelfile (Suggest)
	onModelResponse := func(res ollama.ProgressResponse) error {
		return nil
	}

	err = i.api.Create(context.Background(), &ollama.CreateRequest{Model: "suggest:7b", Modelfile: form.suggestModelfile}, onModelResponse)
	if err != nil {
		return err
	}

	// 7. Install the modelfile (Suggest)
	err = i.api.Create(context.Background(), &ollama.CreateRequest{Model: "explain:7b", Modelfile: form.explainModelfile}, onModelResponse)
	if err != nil {
		return err
	}
	return nil
}
