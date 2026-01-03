package command

import (
	"fmt"
	"os"
	"os/exec"
)

func DefaultShellCommand(cmd string, args []string) {
	runCommand(cmd, args)
}
func runCommand(cmd string, args []string) {
	path, ok := searchPath(cmd)
	if !ok {
		fmt.Printf("%s: not found\n", cmd)
		return
	}
	runCmd := exec.Command(cmd, args...)
	runCmd.Path = path
	runCmd.Stdin = os.Stdin
	runCmd.Stdout = os.Stdout
	runCmd.Stderr = os.Stderr
	_ = runCmd.Run()
}
