package routes

import (
	"time"

	"github.com/E-V-Castillo/kanban-api/internal/routes/api"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Initialize() *chi.Mux {
	api := api.InitializeApiRoutes()
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(time.Second * 60))

	router.Mount("/api", api)

	return router
}
