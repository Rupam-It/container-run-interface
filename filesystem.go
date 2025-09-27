package main

import (
	"os"
	"syscall"
)

func setupFilesystem() error {
	dirs := []string{"/proc", "/sys", "/dev", "/tmp"}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			continue
		}
	}

	if err := syscall.Mount("proc", "/proc", "proc", 0, ""); err != nil {
	}

	// Mount sysfs
	if err := syscall.Mount("sysfs", "/sys", "sysfs", 0, ""); err != nil {
	}

	return nil
}
