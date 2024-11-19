package repositories

import (
	"todo-be/entities"
	"todo-be/helper"

	"gorm.io/gorm"
)

type NoteInteractor interface {
	Count() (int64, error)
	FindAll(query helper.Paginate, userID int) (helper.PaginationResult, error)
	FindOne(id int, userID int) (entities.Note, error)
	BatchCreate(payload []entities.Note) (int64, error)
	Update(payload entities.Note) (entities.Note, error)
	Delete(id int, userID int) error
}

type noteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) *noteRepository {
	return &noteRepository{db}
}

func (r *noteRepository) Count() (int64, error) {
	var total int64
	err := r.db.Table("notes").Count(&total).Error
	if err != nil {
		return total, err
	}

	return total, err
}

func (r *noteRepository) FindAll(query helper.Paginate, userID int) (helper.PaginationResult, error) {
	var models []entities.Note
	var pagination helper.PaginationResult

	err := r.db.Scopes(helper.PaginationScope(query.Page, query.PageSize)).Find(&models).Where("user_id = ?", userID).Error
	if err != nil {
		return pagination, err
	}

	total, errTotal := r.Count()
	if errTotal != nil {
		return pagination, errTotal
	}

	pagination = helper.PaginationResultAdapter(query.Page, query.PageSize, int(total), models)
	return pagination, nil
}

func (r *noteRepository) FindOne(id int, userID int) (entities.Note, error) {
	var model entities.Note
	err := r.db.First(&model, id).Where("user_id = ?", userID).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *noteRepository) BatchCreate(payload []entities.Note) (int64, error) {
	result := r.db.Create(&payload)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (r *noteRepository) Update(payload entities.Note) (entities.Note, error) {
	err := r.db.Save(&payload).Error
	if err != nil {
		return payload, err
	}
	return payload, nil
}

func (r *noteRepository) Delete(id int, userID int) error {
	err := r.db.Delete(&entities.Note{}, id).Where("user_id = ?", userID).Error
	if err != nil {
		return err
	}
	return nil
}
