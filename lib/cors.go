package lib

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitializeCORS() gin.HandlerFunc {
	var FRONTEND_URL = os.Getenv("FRONTEND_URL")
	return cors.New(cors.Config{
		AllowOrigins: []string{FRONTEND_URL},
		AllowMethods: []string{"GET", "PATCH", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Authorization", "Content-Type"},
	})
}
