# base-golang-api
Base skeleton to create REST API using Golang

**DISCLAIMER**

```
This repo only provide base skeleton to create a REST API using
Golang, to save your time from building REST API from scratch.
You can modify all of processes based on your needs!
```

---

## Dependencies 

- gorilla/mux       : For routing
- gorilla/handlers  : For http middleware
- alexsasharegan/dotenv : For env variables

---

## Command Line Usages

```
go run --race *.go
```

Optional parameters:
- `-debug` : Enable debug mode
- `-enablePrintStack` : Enable to print stacktraces when panic happened
- `-address` : Set your ip and port
- `-timeout` : Set your write & read
- `-env` : A path to your env file

Examples:

```
go run --race *.go -address=127.0.0.1:7000
```
---

## Routes 

All of your api controllers should be placed at `handler/`, then you need to
register your controller in `routes.go`, example:

```go
routes = append(routes, Route{"GET", "/hello", controller.HelloWorld})
```

And that's it.  For now, i'm not think about sub router or others, but you can 
modify this route's process based on your needs.

---

## Middlewares

Default middlewares:

- Access logger : To log any of request and response access to your api
- Content negotiator : To filter and spesify http header content types
- Recovery : To prevent your api from crash if any panic happened

You can create your own middleware, just make sure that your function signature 
follow gorilla's middleware function:

```go
type MiddlewareFunc func(http.Handler) http.Handler
```