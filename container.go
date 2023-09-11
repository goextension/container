package container

import (
	"github.com/golang-components/container/contacts"
	"github.com/golang-components/container/expression"
	"github.com/golang-components/container/pure"
	"reflect"
	"unsafe"
)

type Container struct {
	resolved map[string]any

	bindings map[string]any

	singletons map[string]any

	aliases map[string]any

	contextual map[string]map[string]any
}

func (container *Container) Make(abstract any, parameters []any) any {

	concrete := pure.GetClass(abstract)

	if container.hasSingleton(concrete) {
		return container.getConcrete(concrete)
	}

	return container.build(concrete, parameters)
}

func (container *Container) hasSingleton(concrete string) bool {

	_, boolean := container.resolved[concrete]

	return boolean
}

func (container *Container) getConcrete(concrete string) any {
	return container.resolved[concrete]
}

func (container *Container) build(abstract string, parameters []any) any {

	return ""
}

func (container *Container) Bind(abstract any, concrete contacts.Callable) {
	container.register(abstract, concrete, false)
}

func (container *Container) Singleton(abstract any, closure contacts.Callable) {
	container.register(abstract, closure, true)
}

func (container *Container) register(abstract any, concrete contacts.Callable, shared bool) {

	abstractName := pure.GetClass(abstract)

	property := expression.Ternary[string](shared, "singletons", "bindings")

	reflector := reflect.ValueOf(container).Elem().FieldByName(property)

	reflector = reflect.NewAt(reflector.Type(), unsafe.Pointer(reflector.UnsafeAddr())).Elem()

	reflector.SetMapIndex(reflect.ValueOf(abstractName), reflect.ValueOf(concrete(container)))
}

func (container *Container) When(concrete []any) contacts.Context {

	var alias = make([]string, len(concrete))

	for _, class := range concrete {
		alias = append(alias, pure.GetClass(class))
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
