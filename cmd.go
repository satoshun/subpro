package subpro

import (
	"os"
	"os/exec"
)

func OpenCommand(projectPath string) (cmd *exec.Cmd) {
	args := []string{"-a", "Sublime Text", projectPath}
	cmd = exec.Command("open", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return
}

func DeleteCommand(projectPath string) (cmd *exec.Cmd) {
	args := []string{projectPath}
	cmd = exec.Command("rm", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return
}

func CopyCommand(src, dest string) (cmd *exec.Cmd) {
	args := []string{src, dest}
	cmd = exec.Command("cp", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return
}
