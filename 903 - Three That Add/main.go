package main

import (
	"fmt"
	"github.com/camilovietnam/threethatadd/handlers"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Post("/add-item/{item}", handlers.AddItem())
	r.Get("/get-items", handlers.GetAllItems())
	r.Delete("/delete-all", handlers.DeleteAllItems())
	r.Get("/find-three/{sum}", handlers.FindThree())

	fmt.Println("running server in port 8000 ...")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
