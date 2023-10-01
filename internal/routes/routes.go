package routes

import (
	"github.com/E-V-Castillo/kanban-api/internal/routes/api"
	"github.com/E-V-Castillo/kanban-api/internal/routes/web"
	"github.com/go-chi/chi/v5"
)

func InitializeRoutes() *chi.Mux {

	api := api.InitializeApiRoutes()
	web := web.InitializeWebRoutes()
	router := chi.NewRouter()

	router.Mount("/api", api)
	router.Mount("/", web)
	return router
}
