package controllers

import (
	"net/http"
	"strconv"
	"todo-be/dtos"
	"todo-be/entities"
	"todo-be/helper"
	"todo-be/services"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type Params struct {
	ID string `uri:"id" binding:"required"`
}

type noteController struct {
	service services.NoteInteractor
}

func NewNoteController(service services.NoteInteractor) *noteController {
	return &noteController{service}
}

/*
Route: /api/v1/notes
Method: GET
*/
func (h *noteController) GetNotes(c *gin.Context) {
	currentUser := c.MustGet("current_user").(entities.User)
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	query := helper.Paginate{
		Page:     page,
		PageSize: pageSize,
	}
	notes, err := h.service.GetNotes(currentUser.ID, query)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseAdapter("GetNotes Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	models := []entities.Note{}
	mapstructure.Decode(notes.Data, &models)
	data := dtos.NotesAdapter(models)
	res := helper.ResponseAdapter("GetNotes Successful", http.StatusOK, "success", dtos.NotesResponse{
		Data:       data,
		Pagination: notes.Pagination,
	})
	c.JSON(http.StatusOK, res)
}

/*
Route: /api/v1/notes
Method: POST
*/
func (h *noteController) CreateNote(c *gin.Context) {
	currentUser := c.MustGet("current_user").(entities.User)

	var request []dtos.FormNoteRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": helper.ErrorResponseAdapter(err)}
		errResponse := helper.ResponseAdapter("CreateNote Validation Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	_, err = h.service.AddNotes(currentUser.ID, request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseAdapter("CreateNote Failed Created", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	res := helper.ResponseAdapter("CreateNote Successful Created", http.StatusOK, "success", entities.Note{})
	c.JSON(http.StatusOK, res)
}

/*
Route: /api/v1/notes/:id
Method: PATCH
*/
func (h *noteController) UpdateStatusNote(c *gin.Context) {
	currentUser := c.MustGet("current_user").(entities.User)
	var uri Params

	err := c.ShouldBindUri(&uri)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseAdapter("UpdateStatusNote Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	var request dtos.UpdateStatusNoteRequest
	err = c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": helper.ErrorResponseAdapter(err)}
		errResponse := helper.ResponseAdapter("UpdateStatusNote Validation Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	id, _ := strconv.Atoi(uri.ID)
	_, err = h.service.UpdateStatusNote(id, currentUser.ID, request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseAdapter("UpdateStatusNote Failed Created", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	res := helper.ResponseAdapter("UpdateStatusNote Successful Created", http.StatusOK, "success", entities.Note{})
	c.JSON(http.StatusOK, res)
}

/*
Route: /api/v1/notes/:id
Method: PUT
*/

func (h *noteController) UpdateNote(c *gin.Context) {
	currentUser := c.MustGet("current_user").(entities.User)
	var uri Params

	err := c.ShouldBindUri(&uri)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseAdapter("UpdateNote Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	var request dtos.FormNoteRequest
	err = c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": helper.ErrorResponseAdapter(err)}
		errResponse := helper.ResponseAdapter("UpdateNote Validation Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	id, _ := strconv.Atoi(uri.ID)
	_, err = h.service.UpdateNote(id, currentUser.ID, request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseAdapter("UpdateNote Failed Created", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	res := helper.ResponseAdapter("UpdateNote Successful Created", http.StatusOK, "success", entities.Note{})
	c.JSON(http.StatusOK, res)
}

/*
Route: /api/v1/notes/:id
Method: DELETE
*/

func (h *noteController) DeleteNote(c *gin.Context) {
	currentUser := c.MustGet("current_user").(entities.User)
	var uri Params

	err := c.ShouldBindUri(&uri)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseAdapter("DeleteNote Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	id, _ := strconv.Atoi(uri.ID)
	errAction := h.service.DeleteNote(id, currentUser.ID)
	if errAction != nil {
		errorMessage := gin.H{"errors": errAction.Error()}
		errResponse := helper.ResponseAdapter("DeleteNote Failed Created", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	res := helper.ResponseAdapter("DeleteNote Successful Created", http.StatusOK, "success", entities.Note{})
	c.JSON(http.StatusOK, res)
}
