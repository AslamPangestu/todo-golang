package middlewares

import (
	"net/http"
	"strconv"
	"strings"
	"todo-be/helper"
	"todo-be/lib"
	"todo-be/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(jwtLib lib.JWTInteractor, authService services.AuthInteractor) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			errResponse := helper.ResponseAdapter("Unauthorize", http.StatusUnauthorized, "failed", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
			return
		}

		var parseToken string
		splitToken := strings.Split(authHeader, " ")
		if len(splitToken) == 2 {
			parseToken = splitToken[1]
		}

		token, err := jwtLib.ValidateToken(parseToken)
		if err != nil {
			errResponse := helper.ResponseAdapter("Unauthorize", http.StatusUnauthorized, "failed", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !(ok || token.Valid) {
			errResponse := helper.ResponseAdapter("Unauthorize", http.StatusUnauthorized, "failed", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
			return
		}

		userID, err := strconv.Atoi(claim["sub"].(string))
		if err != nil {
			errResponse := helper.ResponseAdapter("Unauthorize", http.StatusUnauthorized, "failed", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
			return
		}

		user, err := authService.GetUserByID(userID)
		if err != nil {
			errResponse := helper.ResponseAdapter("Unauthorize", http.StatusUnauthorized, "failed", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
			return
		}
		c.Set("current_user", user)
	}
}
