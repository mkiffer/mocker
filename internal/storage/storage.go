package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github/mkiffer/mocker/internal/container"
)

var StoragePath = filepath.Join(os.TempDir(), "mocker-containers.json")

//ContainerData represents data we'll store about containers

type ContainerData struct{
	ID 	string `json:"id"`
	Image   string   `json:"image"`
	Command []string `json:"command"`
	Status  string   `json:"status"`
	PID     int      `json:"pid"`	
}
// Storage provides persistent storage for container information
type Storage struct {
	path string
	mutex sync.Mutex
}

func NewStorage(path string) *Storage {
	if path ==  "" {
		path = StoragePath
	}
	return &Storage{
		path: path,
	}
}

func (s *Storage) SaveContainer(c *container.Container) error{
	s.mutex.Lock()
	defer s.mutex.Unlock()

	//Load existing containers
	containers, err := s.loadContainers()
	if err != nil{
		return err
	}

	containers[c.ID] = ContainerData{
		ID: c.ID,
		Image: c.Image,
		Command: c.Command,
		Status: c.Status,
		PID: c.PID,
	}
	return s.saveContainers(containers)
}

func (s *Storage) LoadContainer(id string) (*container.Container,error){
	s.mutex.Lock()
	defer s.mutex.Unlock()
	//Load existing containers
	
	containers, err := s.loadContainers()
	if err != nil{
		return nil, err
	}

	data, exists := containers[id]
	if !exists{
		return nil, fmt.Errorf("container with ID %s not found", id)
	}

	//Create container from instance stored in file
	c := &container.Container{
		ID: data.ID,
		Image: data.Image,
		Command: data.Command,
		Status: data.Status,
		PID: data.PID,
	}

	return c, nil

}

// UpdateContainerStatus updates a container's status in storage
func (s *Storage) UpdateContainerStatus(id string, status string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	containers, err := s.loadContainers()
	if err != nil {
		return err
	}

	data, exists := containers[id]
	if !exists {
		return fmt.Errorf("container with ID %s not found", id)
	}

	//update the status
	data.Status = status
	//update the container
	containers[id] = data

	return s.saveContainers(containers)
}

// ListContainers returns all containers from storage
func (s *Storage) ListContainers() ([]*container.Container, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	containerMap, err := s.loadContainers()
	if err != nil {
		return nil, err
	}

	// Convert map to slice
	containers := make([]*container.Container, 0, len(containerMap))
	for _, data := range containerMap {
		containers = append(containers, &container.Container{
			ID:      data.ID,
			Image:   data.Image,
			Command: data.Command,
			Status:  data.Status,
			PID:     data.PID,
		})
	}

	return containers, nil
}

// RemoveContainer removes a container from storage
func (s *Storage) RemoveContainer(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	containers, err := s.loadContainers()
	if err != nil {
		return err
	}

	if _, exists := containers[id]; !exists {
		return fmt.Errorf("container with ID %s not found", id)
	}

	delete(containers, id)
	return s.saveContainers(containers)
}

func (s *Storage) loadContainers() (map[string]ContainerData, error){
	containers := make(map[string]ContainerData)

	if _,err := os.Stat(s.path); os.IsNotExist(err){
		//The file doesnt exist return an empty map
		return containers, nil
	}

	data, err := os.ReadFile(s.path)
	if err != nil{
		return nil, fmt.Errorf("failed to read container storage file: %v", err)
	}

	if len(data)==0{
		return containers, nil
	}

	//Unmarshal JSON
	if err := json.Unmarshal(data, &containers); err != nil{
		return nil, fmt.Errorf("failed to parse container storage file: %v", err)
	}

	return containers, nil
}

// saveContainers saves all containers to the storage file
func (s *Storage) saveContainers(containers map[string]ContainerData) error {
	// Create the directory if it doesn't exist
	dir := filepath.Dir(s.path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create storage directory: %v", err)
	}

	// Marshal the containers to JSON
	data, err := json.MarshalIndent(containers, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode containers: %v", err)
	}

	// Write to the file
	if err := os.WriteFile(s.path, data, 0644); err != nil {
		return fmt.Errorf("failed to write containers file: %v", err)
	}

	return nil
}