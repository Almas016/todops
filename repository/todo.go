package repository

import (
	"TodoPS/models"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}
func (r *TodoRepository) GetAll() (todo []models.Todo, err error) {
	err = r.db.Model(&models.Todo{}).Find(&todo).Error
	return
}

func (r *TodoRepository) GetOne(id int) (todo *models.Todo, err error) {
	err = r.db.Model(&models.Todo{}).Where("id = ?", id).First(&todo).Error
	return
}

func (r *TodoRepository) Create(model models.Todo) error {
	return r.db.Model(models.Todo{}).Create(&model).Error
}

func (r *TodoRepository) Update(model models.Todo) error {
	return r.db.Model(&model).Where("id = ?", model.Id).Updates(&model).Error
}

func (r *TodoRepository) Delete(model models.Todo) error {
	return r.db.Model(&model).Where("id = ?", model.Id).Delete(&model).Error
}
