package main

import (
	"log"
	"net/http"
	"vigour/infrastructure"
	"vigour/router"
)

// @title Swagger UI for vigour API
// @version 1.0
// @description API lists for vigour API

// @host localhost:1900
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	// Go run main.go
	log.Println("Database name: ", infrastructure.GetDBName())
	log.Fatal(http.ListenAndServe(":"+infrastructure.GetAppPort(), router.Router()))

	
}

