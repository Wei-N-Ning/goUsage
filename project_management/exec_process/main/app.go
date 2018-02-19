package main

import (
	"os/exec"
	"os"
	"syscall"
)

func main() {
	execPath, err := exec.LookPath("ls")
	if err != nil {
		panic(err)
	}
	cmdLine := []string{"ls", "-a", "-l", "-h"}
	env := os.Environ()
	if err := syscall.Exec(execPath, cmdLine, env); err != nil {
		panic(err)
	}
}

