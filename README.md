# Go-container

#### install
```shell
$ go get vendor/container
```

#### bind
```go
app.Bind((*Notifier)(nil), func(app *container.Container) any {
	return &RegisterNotifier{message: "bind successful."}
})
```