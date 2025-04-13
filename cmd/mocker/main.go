package main

import (
	"fmt"
	"os"
	"github/mkiffer/mocker/internal/container"
)

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
		fmt.Println("Running container...")
		//implement calling run functionality
		if len(os.Args) < 3{
			fmt.Println("not enough arguments")
			fmt.Println("Uasge: mocker run <command>")
			os.Exit(1)
		}
		//placeholder image and the provided command
		image := "alpine"
		command := os.Args[2:]

		cont := container.NewContainer(image,command)
		if err := cont.Run(); err != nil {
			fmt.Printf("Error running container: %v\n", err)
			os.Exit(1)
		}

		
	case "lc":
		fmt.Println("Current containers...")
		//implement calling container list
	case "stop":
		fmt.Println("Stopping container...")
		//implement stopping container
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}