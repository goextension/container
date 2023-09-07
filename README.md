# Go-container

#### install
```shell
$ go get koala/container
```

#### bind
```go
app.Bind((*Notifier)(nil), func(app *container.Container) any {
	return &RegisterNotifier{message: "bind successful."}
})
```