package main

import (
	"os"
	"os/exec"
	"syscall"
)

type Container struct {
	Command []string
}

func (c *Container) Run() error {

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, c.Command...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNET |
			syscall.CLONE_NEWUTS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	return cmd.Run()
}

func runCommand() error {
	args := os.Args[2:]
	if len(args) == 0 {
		args = []string{"/bin/sh"}
	}

	// Execute the command
	return syscall.Exec(args[0], args, os.Environ())

}
