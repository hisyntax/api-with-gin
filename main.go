package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	//init gin router
	router := gin.Default()

	//its great to version ypur API's
	v1 := router.Group("/api/v1")

	//handle error when a route is not defined
	router.NoRoute(func(c *gin.Context) {
		//return a json response
		c.JSON(404, gin.H{"message": "Not founf"})
	})

	//init our server
	router.Run(":" + port)
}
