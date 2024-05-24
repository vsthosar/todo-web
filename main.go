package main

import (
	"os"
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	// "todo-web/todo"
)

func main () {

	file, err := os.ReadFile(".todos.json")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}


	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/list", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(file))
	})

	http.ListenAndServe(":3000", r)
}