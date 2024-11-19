package routes

import (
	"todo-be/controllers"
	"todo-be/lib"
	"todo-be/repositories"
	"todo-be/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoute(api *gin.RouterGroup, db *gorm.DB) {
	jwt := lib.InitializeJWT()

	authRepository := repositories.NewAuthRepository(db)

	authService := services.NewAuthService(authRepository)

	userHandler := controllers.NewAuthController(authService, jwt)

	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)
}
