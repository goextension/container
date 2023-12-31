package container

import (
	"github.com/golang-components/container/contacts"
	"github.com/golang-components/container/pure"
)

type ContextualBindingBuilder struct {
	container contacts.Container

	concrete []string

	needs string
}

func (context *ContextualBindingBuilder) Needs(abstract any) contacts.Context {

	context.needs = pure.GetAbstractName[any](abstract)

	return context
}

func (context *ContextualBindingBuilder) Give(implementation func(container contacts.Container) any) {
	for _, concrete := range context.concrete {
		context.container.AddContextualBinding(concrete, context.needs, implementation)
	}
}
