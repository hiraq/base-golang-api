package handler

import (
	"fmt"
	"net/http"
)

func HelloPost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello world via post")
}
