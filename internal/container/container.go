package container

import(
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github/mkiffer/mocker/internal/namespace"
)

// Container represents a running container 
type Container struct {
	ID      string
	Name    string 
	Image   string
	Command []string
	Status  string
	PID		int 
}

//Creates a new container instance (returns pointer to container object)
func NewContainer(image string, command []string) *Container {
	return &Container{
		ID: generateID(), //implement later
		Image: image,
		Command: command,
		Status: "created",
	}
}

//Run starts the container
func (c *Container) Run() error{

	fmt.Printf("Starting container with ID: %s\n", c.ID)
	fmt.Printf("Command: %v\n", c.Command)

	//prepare the command
	cmd := exec.Command(c.Command[0], c.Command[1:]...)

	//set up namespaces for isolation
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: namespace.NamespaceFlags(),
	}

	//Connect standard IO
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	//Start the container process
	if err:= cmd.Start(); err != nil {
		return fmt.Errorf("failed to start container: %v", err)
	}
	
	// Record the PID
	c.PID = cmd.Process.Pid
	c.Status = "running"
	
	fmt.Printf("Container is running with PID: %d\n", c.PID)

	// Wait for the container to finish
	if err := cmd.Wait(); err != nil{
		return fmt.Errorf("container exited with error: %v", err)
	}

	c.Status = "exited"
	fmt.Println("Container has extied")

	return nil

}

//stops the container
func (c *Container) Stop() error {
	if c.Status != "running"{
		return fmt.Errorf("container is not running")
	}
	
	//Send SIGNTERM to the container process
	process, err := os.FindProcess(c.PID)
	if err != nil{
		fmt.Printf("could not find container process: %v", err)
	}

	if err := process.Signal(syscall.SIGTERM); err != nil{
		return fmt.Errorf("failed to send SIGTERM: %v", err)
	}

	c.Status = "stopped"

	return nil
}

//creature a unique ID for the container
func generateID() string {
	return fmt.Sprintf("mocker-%d", os.Getpid())
}
