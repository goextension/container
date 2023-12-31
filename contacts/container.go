package contacts

type Container interface {
	Make(abstract any, parameters map[string]any) any

	Bind(abstract any, concrete Callable)

	Singleton(abstract any, concrete Callable)

	When(concrete []any) Context

	AddContextualBinding(concrete string, abstract string, implementation any)
}
