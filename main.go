package main

import (
	"log"
	"net/http"

	"atsar/config"
	"atsar/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: File .env tidak ditemukan")
	}

	config.ConnectDatabase()

	router := gin.Default()

	// Initialize routes
	routes.SetupUserRoutes(router)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}


