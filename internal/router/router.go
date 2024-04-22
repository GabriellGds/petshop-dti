package router

import (
	"github.com/GabriellGds/petshop-dti/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRoutes(mux *chi.Mux) {
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Use(handlers.CORS)
	mux.Get("/api/petshops", handlers.ListPetshops)
}