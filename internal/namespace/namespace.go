package namespace

import (
	"fmt"
	"syscall"
)

//configure the namespaces for a container
func SetupNameSpaces() error {
	fmt.Println("Setting up namespaces (not yet implemented)")
	return nil
}

//puts the current process ina  new namespace
func Isolate() error {
	fmt.Println("Isolating process (not yet implemented)")
	return nil
}

func NamespaceFlags() uintptr{
	// Combine namespace flags for isolation:
	// CLONE_NEWUTS: Isolate hostname
	// CLONE_NEWPID: Isolate process IDs
	// CLONE_NEWNS: Isolate mount points
	// CLONE_NEWNET: Isolate network (we'll add this later)
	// CLONE_NEWIPC: Isolate IPC (we'll add this later)
	return syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS
}

