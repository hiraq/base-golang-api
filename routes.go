package main

import (
	controller "base-golang-api/handler"
	"net/http"

	"github.com/gorilla/mux"
)

// AppHandler used as http endpoint controller
type AppHandler func(w http.ResponseWriter, r *http.Request)

// Route used as a dictionary provide an information about
// endpoint path and endpoint handler (AppHandler)
type Route struct {
	Method  string
	Path    string
	Handler AppHandler
}

var routes []Route

func init() {
	routes = append(routes, Route{"GET", "/hello", controller.HelloWorld})
	routes = append(routes, Route{"GET", "/hello-panic", controller.HelloPanic})
	routes = append(routes, Route{"POST", "/hello-post", controller.HelloPost})
}

// RegisterRoutes used to activate all configured routes register them to
// gorilla/mux routers.
func RegisterRoutes(r *mux.Router) *mux.Router {
	if len(routes) >= 1 {
		for _, route := range routes {
			r.HandleFunc(route.Path, route.Handler).Methods(route.Method)
		}
	}

	return r
}
