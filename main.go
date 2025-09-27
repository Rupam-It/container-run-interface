package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: ./my-container run <command>")
	}

	switch os.Args[1] {
	case "run":
		if len(os.Args) < 3 {
			log.Fatal("Usage: ./my-container run <command>")
		}
		runContainer(os.Args[2:])
	case "child":
		runChild()
	default:
		log.Fatal("Unknown command:", os.Args[1])
	}

}

func runContainer(args []string) {
	fmt.Printf("Running container with command: %v\n", args)

	container := &Container{
		Command: args,
	}

	if err := container.Run(); err != nil {
		log.Fatal(err)
	}
}

func runChild() {
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
