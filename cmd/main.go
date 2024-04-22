package main

import (
	"log"
	"net/http"

	"github.com/GabriellGds/petshop-dti/internal/router"
	"github.com/go-chi/chi/v5"
)

func main() {
	mux := chi.NewRouter()

	router.InitRoutes(mux)

	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal("error to initialize: ", err)
	}
}