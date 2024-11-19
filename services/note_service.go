package services

import (
	"fmt"
	"todo-be/dtos"
	"todo-be/entities"
	"todo-be/helper"
	"todo-be/repositories"
)

type NoteInteractor interface {
	GetNotes(userID int, query helper.Paginate) (helper.PaginationResult, error)
	AddNotes(userId int, payload []dtos.FormNoteRequest) (int64, error)
	UpdateNote(id int, userID int, payload dtos.FormNoteRequest) (entities.Note, error)
	UpdateStatusNote(id int, userID int, payload dtos.UpdateStatusNoteRequest) (entities.Note, error)
	DeleteNote(id int, userID int) error
}

type noteService struct {
	repository repositories.NoteInteractor
}

func NewNoteService(repository repositories.NoteInteractor) *noteService {
	return &noteService{repository}
}

func (s *noteService) GetNotes(userID int, query helper.Paginate) (helper.PaginationResult, error) {
	result, err := s.repository.FindAll(query, userID)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *noteService) AddNotes(userId int, payload []dtos.FormNoteRequest) (int64, error) {

	count, err := s.repository.Count()
	if err != nil {
		return 0, err
	}

	var models []entities.Note

	for i, note := range payload {
		counter := count + int64(i+1)
		models = append(models, entities.Note{
			Title:       note.Title,
			Description: note.Description,
			Status:      0,
			UserID:      userId,
			ActivityNo:  fmt.Sprintf("AC-%04d", counter),
		})
	}

	rowsAffected, err := s.repository.BatchCreate(models)
	if err != nil {
		return rowsAffected, err
	}

	return rowsAffected, nil
}

func (s *noteService) UpdateNote(id int, userID int, payload dtos.FormNoteRequest) (entities.Note, error) {
	result, err := s.repository.FindOne(id, userID)
	if err != nil {
		return result, err
	}

	result.Title = payload.Title
	result.Description = payload.Description

	updated, err := s.repository.Update(result)
	if err != nil {
		return updated, err
	}

	return updated, nil
}

func (s *noteService) UpdateStatusNote(id int, userID int, payload dtos.UpdateStatusNoteRequest) (entities.Note, error) {
	result, err := s.repository.FindOne(id, userID)
	if err != nil {
		return result, err
	}

	result.Status = payload.Status

	updated, err := s.repository.Update(result)
	if err != nil {
		return updated, err
	}

	return updated, nil
}

func (s *noteService) DeleteNote(id int, userID int) error {
	err := s.repository.Delete(id, userID)
	if err != nil {
		return err
	}

	return nil
}
