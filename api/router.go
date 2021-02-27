package api

import (
	"fmt"
	"net/http"
)

// TodoListAPI todo
func TodoListAPI() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})
}
