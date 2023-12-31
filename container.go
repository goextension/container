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

func (container *Container) Make(abstract any, parameters map[string]any) any {

	reflector := reflect.TypeOf(abstract)

	if reflector.Kind() == reflect.String {

		alias := pure.GetAbstractName(abstract)

		if container.hasSingleton(alias) {
			return container.getConcrete(alias)
		}
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

func (container *Container) build(reflector reflect.Type, parameters map[string]any) any {

	if reflector.Kind() == reflect.Interface {

		if !container.hasSingleton(reflector.Elem().String()) {
			panic("Target[" + reflector.Elem().String() + "]is not instantiable.")
		}

		return container.getConcrete(reflector.Elem().String())
	}

	if reflector.NumField() == 0 {
		return reflect.New(reflector)
	}

	return ""
}

func (container *Container) resolve() any {

}

func (container *Container) Bind(abstract any, concrete contacts.Callable) {
	container.register(abstract, concrete, false)
}

func (container *Container) Singleton(abstract any, closure contacts.Callable) {
	container.register(abstract, closure, true)
}

func (container *Container) register(abstract any, concrete contacts.Callable, shared bool) {

	abstractName := pure.GetAbstractName(abstract)

	property := expression.Ternary[string](shared, "singletons", "bindings")

	reflector := reflect.ValueOf(container).Elem().FieldByName(property)

	reflector = reflect.NewAt(reflector.Type(), unsafe.Pointer(reflector.UnsafeAddr())).Elem()

	reflector.SetMapIndex(reflect.ValueOf(abstractName), reflect.ValueOf(concrete(container)))
}

func (container *Container) When(concrete []any) contacts.Context {

	var alias = make([]string, len(concrete))

	for _, class := range concrete {
		alias = append(alias, pure.GetAbstractName(class))
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
