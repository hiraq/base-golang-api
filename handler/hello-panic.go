package handler

import "net/http"

func HelloPanic(w http.ResponseWriter, r *http.Request) {
	panic("hello panic!")
}
