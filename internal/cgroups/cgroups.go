package cgroups
import (
	"fmt"
)

// contains resources constraints for a container
type ResourceLimits struct {
	MemoryLimit int64 //in bytes
	CPUShares	int64
}

// applies resource limits to a process
func ApplyLimits(pid int, limits ResourceLimits) error {
	fmt.Printf("Applying limits to PID %d (not yet implemented) \n", pid)
	return nil
}