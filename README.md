# My-Container (Minimal Container Runtime in Go)
This project demonstrates how Linux namespaces work by creating a mini container runtime in pure Go. It re-executes itself into new namespaces (UTS, PID, NET, MNT), giving you an isolated environment — similar to how Docker works under the hood.

## What This Container Runtime Does
My-Container creates process isolation using Linux namespaces, providing:

Process Isolation: Container processes are completely separated from host processes

Hostname Isolation: Container gets its own hostname independent of the host

Basic Command Execution: Run any command inside the isolated environment

Interactive Shell Support: Get a bash shell inside the container

Mount Namespace Creation: Foundation for filesystem isolation

 Prerequisites
Linux environment (Ubuntu/Debian recommended)

Go 1.19+ installed

Root privileges (required for namespace operations)

Basic understanding of command line

 How to Test It - Step by Step
Test 1: Basic Compilation
bash
## First, make sure it compiles
go build -o my-container

## If successful, you'll see the binary
ls -la my-container
Test 2: Check Isolation is Working
Outside Container (Host):

bash
## Check your real hostname
hostname
## Example: your-laptop-name

## Check processes
ps aux | head -5
## Shows all your host processes (hundreds of them)
Inside Container:

bash
## Run your container with a shell
sudo ./my-container run /bin/bash

## Inside container - check hostname
hostname
## Should show: my-container (different from host!)

## Inside container - check processes
ps aux
## Should show only container processes (isolated!)
Test 3: Process Isolation Verification
In the Container:

bash
ps aux
## Notice PID 1 is your /bin/bash (init process of the container)
echo "Container PID: $$"
From the Host (open another terminal):

bash
ps aux | grep my-container
## Same process appears with a different PID from the host perspective
Test 4: Interactive Commands
bash
## Test various commands inside the container
sudo ./my-container run echo "Hello from container!"
sudo ./my-container run whoami
sudo ./my-container run pwd
sudo ./my-container run ls /proc
 Usage Examples
bash
## Basic command execution
sudo ./my-container run echo "Hello World"

## Interactive shell
sudo ./my-container run /bin/bash

## Run system commands
sudo ./my-container run ps aux
sudo ./my-container run hostname
sudo ./my-container run ls -la
⚡ Quick Start
Clone and build:

bash
git clone <your-repo>
cd my-container
go build -o my-container
Test basic functionality:

bash
sudo ./my-container run echo "It works!"
Get interactive shell:

bash
sudo ./my-container run /bin/bash
## Comparison with Docker

| Feature             | My-Container | Docker |
|---------------------|--------------|--------|
| Process Isolation   | Yes          | Yes    |
| Hostname Isolation  | Yes          | Yes    |
| Command Execution   | Yes          | Yes    |
| Interactive Shell   | Yes          | Yes    |
| Filesystem Isolation| No           | Yes    |
| Image Management    | No           | Yes    |
| Networking          | No           | Yes    |
| Resource Limits     | No           | Yes    |
| Volume Management   | No           | Yes    |
| Container Lifecycle | No           | Yes    |
| Registry Support    | No           | Yes    |
| Multi-platform      | No           | Yes    |
