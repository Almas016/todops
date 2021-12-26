package repository

import "gorm.io/gorm"

type Repository struct {
	TodoRepository *TodoRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		TodoRepository: NewTodoRepository(db),
	}
}
