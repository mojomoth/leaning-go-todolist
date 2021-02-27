package main

import (
	"log"

	"github.com/mojomoth/leaning-go-todolist/server"
)

func main() {
	if err := server.ListenAndServe(server.Config{
		Address:     ":8080",
		DatabaseURL: "postgres://postgres:1q2w3e4r@localhost:5432/postgres?sslmode=disable",
	}); err != nil {
		log.Fatalln(err)
	}
}
