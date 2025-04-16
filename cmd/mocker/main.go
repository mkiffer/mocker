package main

import (
	"fmt"
	"os"
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
	
	go func () {
		if err := cont.Run(); err != nil {
			fmt.Printf("Error running container: %v\n", err)
			// Update container status on error
			cont.Status = "error"
		}
	} ()
	
	
    //Save container to storage
	if err := containerStorage.SaveContainer(cont); err != nil{
		fmt.Printf("Warning: failed to save container. Info: %v\n", err)
	}
	

	
	fmt.Printf("Started container with ID: %s\n", cont.ID)
}

func listContainers(){
	fmt.Println("Current containers...")
	containerList , err := containerStorage.ListContainers() 
	if err != nil{
		fmt.Printf("Error listing containers from storage. Info: %n\n",err)
	}
	if len(containerList) == 0{
		fmt.Println("No containers running")
		return
	}

	for _, ref := range containerList{
		fmt.Printf("Container: %s, status: %s", ref.Name, ref.Status)
	}

}

func stopContainer(args []string){
	if len(args) < 1{
		fmt.Println("Not enough arguments")
		fmt.Println("Usage: mocker stop <container-id>")
		os.Exit(1)
	}

	containerID := args[0]

	cont, err := containerStorage.LoadContainer(containerID)
	if err != nil{
		fmt.Printf("Error loading container with ID: %s. Info %v\n", containerID, err)
		os.Exit(1)
	}

	if err :=cont.Stop(); err != nil{
		fmt.Printf("Error stopping container: %v\n", err)
		os.Exit(1)		
	}

	fmt.Printf("Container %s stopped\n", containerID)
}