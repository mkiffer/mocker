package registry

import(
	"fmt"
	"sync"

	"github/mkiffer/mocker/internal/container"
)

//Registry manages container instances
type Registry struct{
	containers 	map[string]*container.Container
	mutex 		sync.Mutex
}

//NewRegistry creates container registry
func NewRegistry() *Registry{
	return &Registry{
		containers: make(map[string]*container.Container),
	}
}

//Add stores a container in the registry
func (r *Registry) Add(container *container.Container){
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.containers[container.ID] = container
}

//Get retrieves a container by ID
func (r *Registry) Get(id string) (*container.Container, error){
	r.mutex.Lock()
	defer r.mutex.Unlock()

	container, exists := r.containers[id]
	if !exists{
		return nil, fmt.Errorf("container with ID %s not found", id)
	}

	return container, nil

}

func (r *Registry) Remove(id string) error{
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.containers[id]; !exists{
		return fmt.Errorf("container with ID %s not found", id)
	}

	delete(r.containers,id)
	return nil
}

//List returns all containers in the registry
func (r *Registry) List() []*container.Container{
	r.mutex.Lock()
	defer r.mutex.Unlock()

	containers := make([]*container.Container,0, len(r.containers))
	for _, container := range r.containers {
		containers = append(containers, container)
	}

	return containers
}