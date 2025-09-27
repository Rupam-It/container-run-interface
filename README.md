My-Container (Minimal Container in Go)

This project demonstrates how Linux namespaces work by creating a mini container runtime in pure Go.
It re-executes itself into new namespaces (UTS, PID, NET, MNT), giving you an isolated environment — similar to how Docker works under the hood.

How to Test It - Step by Step
Test 1: Basic Compilation
# First, make sure it compiles
go build -o my-container

# If successful, you'll see the binary
ls -la my-container

Test 2: Check Isolation is Working
Outside Container (Host):
# Check your real hostname
hostname
# Example: your-laptop-name

# Check processes
ps aux | head -5
# Shows all your host processes

Inside Container:
# Run your container with a shell
sudo ./my-container run /bin/bash

# Inside container - check hostname
hostname
# Should show: container (different from host!)

# Inside container - check processes
ps aux
# Should show only container processes (isolated!)

Test 3: Process Isolation Verification
In the Container:
ps aux
# Notice PID 1 is your /bin/bash (init process of the container)

From the Host:
ps aux | grep my-container
# Same process appears with a different PID from the host perspective
