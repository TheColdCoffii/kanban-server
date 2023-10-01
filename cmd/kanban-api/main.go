package main

import (
	"net/http"

	"github.com/E-V-Castillo/kanban-api/internal/routes"
)

func main() {
	router := routes.InitializeRoutes()
	http.ListenAndServe(":3000", router)
}
