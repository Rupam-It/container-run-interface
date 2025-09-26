package main

import (
	"fmt"
	"log"
	"os"
)

// TODO: Add imports for fmt, log, os
func main() {
	// TODO: Check if we have at least 2 arguments
	// TODO: Print usage message if not enough args
	if len(os.Args) < 2 {
		log.Fatal("Usage: ./my-container run <command>")
	}

	// TODO: Create switch statement for os.Args[1]
	// TODO: Handle "run" case - call runContainer function
	// TODO: Handle "child" case - call runChild function
	// TODO: Handle default case - show error
	switch os.Args[1] {
	case "run":
		if len(os.Args) < 3 {
			log.Fatal("Usage: ./my-container run <command>")
		}
		runContainer(os.Args[2:])
	case "child":
		// This is called internally for the child process
		runChild()
	default:
		log.Fatal("Unknown command:", os.Args[1])
	}

}

func runContainer(args []string) {
	// TODO: Print what we're about to run
	// TODO: Create Container struct with command args
	// TODO: Call container.Run() method
	// TODO: Handle any errors
	fmt.Printf("Running container with command: %v\n", args)

	container := &Container{
		Command: args,
	}

	if err := container.Run(); err != nil {
		log.Fatal(err)
	}
}

func runChild() {
	// TODO: Print "Running inside container!"
	// TODO: Call setupNamespaces() function
	// TODO: Call setupFilesystem() function
	// TODO: Call runCommand() function
	// TODO: Handle all errors with log.Fatal
	fmt.Println("Running inside container!")

	if err := setupNamespaces(); err != nil {
		log.Fatal("Failed to setup namespaces:", err)
	}

	if err := setupFilesystem(); err != nil {
		log.Fatal("Failed to setup filesystem:", err)
	}

	if err := runCommand(); err != nil {
		log.Fatal("Failed to run command:", err)
	}
}
