package contacts

type Context interface {
	Needs(abstract any) Context

	Give(implementation Callable)
}
