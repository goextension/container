package container

import (
	`koala/container/contacts`
	"reflect"
	"sync"
)

type Container struct {
	resolved map[string]any

	bindings map[string]any

	singletons map[string]any

	aliases map[string]func()

	contextual map[string]map[string]any

	instance *Container
}

func (container *Container) Make(abstract any, parameters []any) any {

	reflector := container.reflectionStruct(abstract)

	structName := reflector.String()

	if container.hasSingleton(structName) {
		return container.getConcrete(structName)
	}

	return container.build(reflector, parameters)
}

func (container *Container) hasSingleton(concrete string) bool {

	_, boolean := container.resolved[concrete]

	return boolean
}

func (container *Container) getConcrete(concrete string) any {
	return container.resolved[concrete]
}

func (container *Container) build(reflector reflect.Type, parameters []any) any {

	return ""
}

func (container *Container) Bind(abstract any, concrete contacts.Callable) {
	container.register(abstract, concrete, false)
}

func (container *Container) Singleton(abstract any, closure contacts.Callable) {
	container.register(abstract, closure, true)
}

func (container *Container) register(abstract any, concrete contacts.Callable, shared bool) {

	abstractStruct := container.guessAbstractName(abstract)

	if shared {
		container.singletons[abstractStruct] = concrete(container)
	} else {
		container.bindings[abstractStruct] = concrete(container)
	}
}

func (container *Container) When(concrete []any) contacts.Context {

	var alias = make([]string, len(concrete))

	for _, class := range concrete {
		alias = append(alias, container.GetStructName(class))
	}

	return &ContextualBindingBuilder{container: container, concrete: alias}
}

func (container *Container) AddContextualBinding(concrete string, abstract string, implementation any) {
	container.contextual[concrete][abstract] = implementation
}

func (container *Container) getAlias(abstract string) any {

	alias, boolean := container.aliases[abstract]

	if boolean {
		return container.aliases[abstract]
	}

	return alias
}

func (container *Container) guessAbstractName(abstract any) string {
	return container.reflectionStruct(abstract).Elem().String()
}

func (container *Container) GetStructName(abstract any) string {
	return container.reflectionStruct(abstract).String()
}

func (container *Container) reflectionStruct(structName any) reflect.Type {
	return reflect.TypeOf(structName)
}

func (container *Container) Instance() contacts.Container {

	var once sync.Once

	if container.instance == nil {
		once.Do(func() {
			container.instance = &Container{
				resolved:   make(map[string]any),
				bindings:   make(map[string]any),
				singletons: make(map[string]any),
				aliases:    make(map[string]func()),
				contextual: make(map[string]map[string]any),
			}
		})

	}

	return container.instance
}
