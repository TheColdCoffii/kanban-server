package api

import "github.com/go-chi/chi/v5"

func InitializeApiRoutes() *chi.Mux {
	apiRouter := chi.NewRouter()
	users := userRoute()
	auth := AuthRoutes()
	apiRouter.Mount("/users", users)
	apiRouter.Mount("/", auth)
	return apiRouter
}
