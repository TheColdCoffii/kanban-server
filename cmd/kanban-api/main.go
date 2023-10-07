package main

import (
	"log"
	"net/http"

	"github.com/E-V-Castillo/kanban-api/internal/database"
	"github.com/E-V-Castillo/kanban-api/internal/routes"
	"github.com/joho/godotenv"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			log.Println("catched")
			log.Println(r)
		}
	}()

	log.Default().SetFlags(log.Lshortfile)

	// Load env variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Environment variables could not be loaded")
	}

	if err := database.InitializeDatabasePool(); err != nil {
		log.Fatalln(err)
	}
	r := routes.Initialize()
	log.Fatal(http.ListenAndServe(":3000", r))
}
