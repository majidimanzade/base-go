package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/majidimanzade/base-go/internal/api"
	"github.com/majidimanzade/base-go/internal/database"
	"github.com/majidimanzade/base-go/internal/services"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}
	database.InitializeDB()
	servicesContainer := services.NewServices()

	r := gin.Default()
	api.InitializeHandlers(r, servicesContainer)

	serverAddress := fmt.Sprintf(":%s", port)
	log.Printf("Server running on %s", serverAddress)
	if err := r.Run(serverAddress); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
