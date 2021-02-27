package server

import (
	"net/http"

	"github.com/mojomoth/leaning-go-todolist/api"
	"github.com/mojomoth/leaning-go-todolist/db"
)

// Config server
type Config struct {
	Address     string
	DatabaseURL string
}

// ListenAndServe starts up the server
func ListenAndServe(cfg Config) error {
	if err := db.Connect(cfg.DatabaseURL); err != nil {
		return err
	}

	return http.ListenAndServe(cfg.Address, loggingMiddleware(api.TodoListAPI()))
}
