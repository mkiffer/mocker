package main

import (
	"fmt"
	"os"
	"syscall"
	"github/mkiffer/mocker/internal/container"
	//"github/mkiffer/mocker/internal/registry"
	"github/mkiffer/mocker/internal/storage"
)

var containerStorage *storage.Storage

//initialise container registry when program starts
func init(){
	
	containerStorage = storage.NewStorage("")

}

func main(){
	fmt.Println("Mocker - A simple container runtime")

	if len(os.Args) < 2 {
		fmt.Println("Usage: mocker <command> [arguments]")
		fmt.Println("Available commands:")
		fmt.Println("	run 	Run a container")
		fmt.Println("	lc 		List containers")
		fmt.Println("	stop 	Stop a container")
		os.Exit(1)
	}

	command :=os.Args[1]

	switch command {
	case "run":
		runContainer(os.Args[2:])
	case "lc":
		listContainers()
		//implement calling container list
	case "stop":
		stopContainer(os.Args[2:])
		//implement stopping container
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}

func runContainer(args []string) {
	if len(args) < 1 {
		fmt.Println("Not enough arguments")
		fmt.Println("Usage: mocker run <command>")
		os.Exit(1)
	}

	// Placeholder image and the provided command
	image := "alpine"
	command := args

	cont := container.NewContainer(image, command)
	    //Save container to storage
	if err := containerStorage.SaveContainer(cont); err != nil{
		fmt.Printf("Warning: failed to save container. Info: %v\n", err)
	}
		
	go func () {
		if err := cont.Run(); err != nil {
			fmt.Printf("Error running container: %v\n", err)
			// Update container status on error
			cont.Status = "error"
			containerStorage.UpdateContainerStatus(cont.ID, "error")
		}
		//update container status upon completion
		containerStorage.UpdateContainerStatus(cont.ID, cont.Status)

	} ()
	
	fmt.Printf("Started container with ID: %s\n", cont.ID)
}

func listContainers(){
	fmt.Println("Current containers...")
	containers , err := containerStorage.ListContainers() 
	if err != nil{
		fmt.Printf("Error listing containers from storage. Info: %n\n",err)
	}
	if len(containers) == 0{
		fmt.Println("No containers running")
		return
	}
	fmt.Println("ID\t\tIMAGE\tSTATUS\tPID\tCOMMAND")
	for _, cont := range containers {
		// Check if the process is still running
		process, err := os.FindProcess(cont.PID)
		if err != nil || process.Signal(syscall.Signal(0)) != nil {
			// Process doesn't exist anymore or can't be signaled
			if cont.Status == "running" {
				cont.Status = "exited"
				containerStorage.UpdateContainerStatus(cont.ID, "exited")
			}
		}

		fmt.Printf("%s\t%s\t%s\t%d\t%v\n", 
			cont.ID, 
			cont.Image, 
			cont.Status, 
			cont.PID, 
			cont.Command)
	}

}

func stopContainer(args []string) {
	if len(args) < 1 {
		fmt.Println("Not enough arguments")
		fmt.Println("Usage: mocker stop <container-id>")
		os.Exit(1)
	}

	containerID := args[0]

	// Load container from storage
	cont, err := containerStorage.LoadContainer(containerID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if err := cont.Stop(); err != nil {
		fmt.Printf("Error stopping container: %v\n", err)
		os.Exit(1)
	}
	
	// Update container status in storage
	if err := containerStorage.UpdateContainerStatus(containerID, "stopped"); err != nil {
		fmt.Printf("Warning: Failed to update container status: %v\n", err)
	}

	fmt.Printf("Container %s stopped\n", containerID)
}