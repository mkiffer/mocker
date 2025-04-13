package main

import (
	"fmt"
	"os"
	"github/mkiffer/mocker/internal/container"
	"github/mkiffer/mocker/internal/registry"
)

var containerRegistry *registry.Registry

//initialise container registry when program starts
func init(){
	containerRegistry = registry.NewRegistry();
}

func main(){
	fmt.Println("Mocker - A simple container runtime")
	fmt.Println("Creating new registry of containers")

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
	
	// Add the container to the registry before starting it
	containerRegistry.Add(cont)

	// Start the container in a goroutine so it doesn't block
	go func() {
		if err := cont.Run(); err != nil {
			fmt.Printf("Error running container: %v\n", err)
			// Update container status on error
			cont.Status = "error"
		}
	}()
	
	fmt.Printf("Started container with ID: %s\n", cont.ID)
}

func listContainers(){
	fmt.Println("Current containers...")
	containers := containerRegistry.List()
	if len(containers) == 0{
		fmt.Println("No containers running")
		return
	}

	for _, ref := range containers{
		fmt.Printf("Container: %v, status: %v", ref.Name, ref.Status)
	}

}

func stopContainer(args []string){
	if len(args) < 1{
		fmt.Println("Not enough arguments")
		fmt.Println("Usage: mocker stop <container-id>")
		os.Exit(1)
	}

	containerID := args[0]

	cont, err := containerRegistry.Get(containerID)
	if err != nil{
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if err :=cont.Stop(); err != nil{
		fmt.Printf("Error stopping container: %v\n", err)
		os.Exit(1)		
	}

	fmt.Printf("Container %s stopped\n", containerID)
}