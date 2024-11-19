package controllers

import (
	"net/http"
	"strconv"
	"todo-be/dtos"
	"todo-be/helper"
	"todo-be/lib"
	"todo-be/services"

	"github.com/gin-gonic/gin"
)

type authController struct {
	service services.AuthInteractor
	jwtLib  lib.JWTInteractor
}

func NewAuthController(service services.AuthInteractor, jwtLib lib.JWTInteractor) *authController {
	return &authController{service, jwtLib}
}

/*
Route: /api/v1/register
Method: POST
*/
func (h *authController) Register(c *gin.Context) {
	var request dtos.RegisterRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": helper.ErrorResponseAdapter(err)}
		errResponse := helper.ResponseAdapter("Register Validation Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	user, err := h.service.AddUser(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseAdapter("Register Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	token, err := h.jwtLib.GenerateToken(strconv.Itoa(user.ID))
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseAdapter("GenerateToken Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	data := dtos.AuthAdapter(user, token)
	res := helper.ResponseAdapter("Register Successful", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}

/*
Route: /api/v1/login
Method: POST
*/
func (h *authController) Login(c *gin.Context) {
	var request dtos.LoginRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": helper.ErrorResponseAdapter(err)}
		errResponse := helper.ResponseAdapter("Login Validation Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	userLogged, err := h.service.GetUserByUserID(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseAdapter("Login Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	token, err := h.jwtLib.GenerateToken(strconv.Itoa(userLogged.ID))
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseAdapter("GenerateToken Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	data := dtos.AuthAdapter(userLogged, token)
	res := helper.ResponseAdapter("Login Successful", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}
