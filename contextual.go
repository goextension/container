package container

import "github.com/xgbnl/container/contacts"

type ContextualBindingBuilder struct {
	container contacts.Container

	concrete []string

	needs string
}

func (context *ContextualBindingBuilder) Needs(abstract any) contacts.Context {

	context.needs = context.container.GetStructName(abstract)

	return context
}

func (context *ContextualBindingBuilder) Give(implementation func(container contacts.Container) any) {
	for _, concrete := range context.concrete {
		context.container.AddContextualBinding(concrete, context.needs, implementation)
	}
}
