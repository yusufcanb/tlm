package shell

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func GetShell() string {
	if runtime.GOOS == "windows" {
		return "powershell"
	}

	if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
		return "bash"
	}

	return "bash"
}

func IsPowershell() bool {
	return runtime.GOOS == "windows"
}

func Exec(cmd string) *exec.Cmd {
	if GetShell() == "powershell" {
		return exec.Command(GetShell(), "-Command", cmd)
	}

	return exec.Command(GetShell(), "-c", cmd)
}

func Confirm(s string, tries int) bool {
	r := bufio.NewReader(os.Stdin)

	for ; tries > 0; tries-- {
		fmt.Printf("%s [y/n]: ", s)

		res, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Empty input (i.e. "\n")
		if len(res) < 3 {
			continue
		}

		return strings.ToLower(strings.TrimSpace(res))[0] == 'y'
	}

	return false
}

func AskQuestion(question string, defaultValue string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s [%s]: ", question, defaultValue)
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(answer)
	if answer == "" {
		return defaultValue
	}
	return answer
}
