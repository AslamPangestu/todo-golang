package routes

import (
	"todo-be/controllers"
	"todo-be/lib"
	"todo-be/middlewares"
	"todo-be/repositories"
	"todo-be/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NoteRoute(api *gin.RouterGroup, db *gorm.DB) {
	jwt := lib.InitializeJWT()

	noteRepository := repositories.NewNoteRepository(db)
	authRepository := repositories.NewAuthRepository(db)

	authService := services.NewAuthService(authRepository)
	noteService := services.NewNoteService(noteRepository)

	noteController := controllers.NewNoteController(noteService)

	api.GET("/notes", middlewares.AuthMiddleware(jwt, authService), noteController.GetNotes)
	api.POST("/notes", middlewares.AuthMiddleware(jwt, authService), noteController.CreateNote)
	api.PATCH("/notes/:id", middlewares.AuthMiddleware(jwt, authService), noteController.UpdateStatusNote)
	api.PUT("/notes/:id", middlewares.AuthMiddleware(jwt, authService), noteController.UpdateNote)
	api.DELETE("/notes/:id", middlewares.AuthMiddleware(jwt, authService), noteController.DeleteNote)
}
