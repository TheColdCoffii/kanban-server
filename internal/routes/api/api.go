package api

import "github.com/go-chi/chi/v5"

func InitializeApiRoutes() *chi.Mux {
	apiRouter := chi.NewRouter()
	return apiRouter
}
