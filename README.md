# mocker

A simple container runtime implementation in Go, designed for learning about containers, 
Linux namespaces, cgroups, and Go systems programming.

## Features (Planned)
- Process isolation using namespaces
- Resource constraints with cgroups
- Container filesystem management
- Basic networking
- Simple CLI

## Usage
```bash
mocker run <image> <command>  # Run a command in a container
mocker lc                     # List running containers
mocker stop <container-id>    # Stop a container
