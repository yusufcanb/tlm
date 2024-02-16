package shell

import (
	"os"
	"os/exec"
	"runtime"
)

func GetShell() string {
	if runtime.GOOS == "windows" {
		return "powershell"
	}

	if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
		return "sh"
	}

	return ""
}

func Exec(cmd string) *exec.Cmd {
	if GetShell() == "powershell" {
		return exec.Command(GetShell(), "-Command", cmd)
	}

	return exec.Command(GetShell(), "-c", cmd)
}

func ClearConsole() {
	cmd := Exec("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
