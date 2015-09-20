package main

import (
	"os"
	"os/exec"
)

func openSublText(projectPath string) (cmd *exec.Cmd) {
	args := []string{"-a", "Sublime Text", projectPath}
	cmd = exec.Command("open", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return
}

func deleteFile(projectPath string) (cmd *exec.Cmd) {
	args := []string{projectPath}
	cmd = exec.Command("rm", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return
}

func copyFile(src, dest string) (cmd *exec.Cmd) {
	args := []string{src, dest}
	cmd = exec.Command("cp", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return
}
