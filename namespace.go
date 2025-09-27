package main

import (
	"os"
	"syscall"
)

func setupNamespaces() error {
	if err := syscall.Sethostname([]byte("container")); err != nil {
		return err
	}
	if err := os.Chdir("/"); err != nil {
		return err
	}

	return nil

}
