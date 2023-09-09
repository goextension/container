package container

import (
	"github.com/golang-components/container/contacts"
	"sync"
)

var (
	container contacts.Container

	syncOnce sync.Once
)

func New() contacts.Container {

	if container == nil {
		syncOnce.Do(func() {
			container = &Container{
				resolved:   make(map[string]any),
				bindings:   make(map[string]any),
				singletons: make(map[string]any),
				aliases:    make(map[string]any),
				contextual: make(map[string]map[string]any),
			}
		})
	}

	return container
}
