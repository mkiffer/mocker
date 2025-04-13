package fs

import(
	"fmt"
	"os"
)

// SetupContainerFS prepares a filesystem for a container
// It should:
// 1. Create a temporary directory to serve as the container's rootfs
// 2. Either copy minimal files or extract an image to this directory
// 3. Return the path to the prepared filesystem
func SetupContainerFS(image string) (string, error){
	//Create a temporary directory for the container root fs
	//for now you could use ioutil.TempDir or os.MkdirTemp
	tempdir, err := os.MkdirTemp("", "mocker-fs-") 
	if err != nil{
		err = fmt.Errorf("could not create container filesystem")
	}
	//Here extract the "image" or copy basic files
	// for a simple start, perhaps just create basic directories
	//like /bin, /etc, /proc, /sys that containers need
	
	

	return tempdir, err 
}

// PivotRoot changes the root filesystem to the provided path
// This is a key function for filesystem isolation
// It should:
// 1. Create a "pivot" directory inside the new root
// 2. Call syscall.PivotRoot to switch the root filesystem
// 3. Clean up the old root reference
func PivotRoot(newRoot string) error {
	// Make sure newRoot is an absolute path
	
	// Create a directory for the old root
	
	// Call syscall.PivotRoot(newRoot, putOld)
	
	// Change working directory to the new root
	
	// Unmount the old root and remove the temporary directory
	
	return fmt.Errorf("not yet implemented")
}

//MountProc mounts the proc filesystem in the container
//Essential for process visibility within the container
func MountProc(rootfs string) error {
	//Create /proc directory if it doesn't exist

	//Mount proc filesystem
	//Use syscall.Mount("proc", procPath, "proc", 0 "")

	return fmt.Errorf("not yet implemented")
}

//MountSys mounts the sysfs in the container (optional but useful)
func MountSys(rootfs string) error {
	//similar to MountProc but for /sys

	return fmt.Errorf("not yet implemented")
}

//CleanupFS removes the container filesystem when no longer needed
func CleanupFS(rootfs string) error {
	//Unmount any filesystems mounted within rootfs

	//Remove the rootfs directory

	return fmt.Errorf("not yet implemented")
}

