package main

import (
	"log"
	"os"
	"todo-be/lib"
	"todo-be/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	var APP_ENV = os.Getenv("APP_ENV")
	if APP_ENV != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("ENV Failure, %v\n", err.Error())
		}
	}

	database := lib.InitializeDatabase()
	router := gin.Default()
	router.Use(lib.InitializeCORS())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/api/v1")
	routes.AuthRoute(v1, database)
	routes.NoteRoute(v1, database)

	router.Run()
}
