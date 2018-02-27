package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

var (
	contentTypes = []string{
		"application/json",
		"application/vnd.api+json",
	}
)

func AccessLogMiddleware(next http.Handler) http.Handler {
	loggedMiddleware := handlers.LoggingHandler(os.Stdout, next)
	return loggedMiddleware
}

func ContentNegotiatorMiddleware(next http.Handler) http.Handler {
	contentTypeNegotiatorMiddleware := handlers.ContentTypeHandler(next, contentTypes...)
	return contentTypeNegotiatorMiddleware
}
