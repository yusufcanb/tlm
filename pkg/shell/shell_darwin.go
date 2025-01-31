package shell

func GetShell() string {
	return "bash"
}

func GetDeploymentPath() (string, error) {
	return "/usr/local/bin/tlm", nil
}
