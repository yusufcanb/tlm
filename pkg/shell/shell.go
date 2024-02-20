package shell

import (
	"bytes"
	"os/exec"
	"runtime"
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
