# Go-container

#### install
```shell
$ go get gopher/container
```

#### bind
```go
app.Bind((*Notifier)(nil), func(app *container.Container) any {
	return &RegisterNotifier{message: "bind successful."}
})
```