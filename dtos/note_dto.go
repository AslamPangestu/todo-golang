package dtos

import (
	"todo-be/entities"
	"todo-be/helper"
)

type noteResponse struct {
	ID          int    `json:"id"`
	ActivityNo  string `json:"activity_no"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type NotesResponse struct {
	Pagination helper.Pagination `json:"pagination"`
	Data       []noteResponse    `json:"data"`
}

func noteAdapter(note entities.Note) noteResponse {
	return noteResponse{
		ID:          note.ID,
		Title:       note.Title,
		Status:      note.Status,
		ActivityNo:  note.ActivityNo,
		Description: note.Description,
	}
}

func NotesAdapter(notes []entities.Note) []noteResponse {
	results := []noteResponse{}
	for _, campaign := range notes {
		campaignAdapter := noteAdapter(campaign)
		results = append(results, campaignAdapter)
	}
	return results
}
