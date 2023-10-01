package web

import "github.com/go-chi/chi/v5"

func InitializeWebRoutes() *chi.Mux {
	webRouter := chi.NewRouter()
	return webRouter
}
