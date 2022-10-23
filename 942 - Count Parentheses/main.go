package main

import (
	"fmt"
	"github.com/dailyproblem/942-count-parentheses/handlers"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Get("/min_remove/{string}", handlers.MinRemove())

	fmt.Println("Server running in port 8000...")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal("Could not start server", err)
	}
}
