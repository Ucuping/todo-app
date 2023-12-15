package repositories

import (
	"github.com/Ucuping/todo-app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoRepository interface {
	GetAllTodo() (*gorm.DB, []models.Todo)
	CreateTodo(todo models.Todo) (models.Todo, error)
	GetTodo(ID uuid.UUID) (models.Todo, error)
	UpdateTodo(todo models.Todo) (models.Todo, error)
	DeleteTodo(ID uuid.UUID, todo models.Todo) (models.Todo, error)
}

func RepositoryTodo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllTodo() (*gorm.DB, []models.Todo) {
	// var todos []models.Todo
	model := r.db.Model(&models.Todo{})

	return model, []models.Todo{}
}

func (r *repository) CreateTodo(todo models.Todo) (models.Todo, error) {
	err := r.db.Create(&todo).Error
	return todo, err
}

func (r *repository) GetTodo(ID uuid.UUID) (models.Todo, error) {
	var todo models.Todo
	err := r.db.First(&todo, ID).Error
	return todo, err
}

func (r *repository) UpdateTodo(todo models.Todo) (models.Todo, error) {
	err := r.db.Save(&todo).Error

	return todo, err
}

func (r *repository) DeleteTodo(ID uuid.UUID, todo models.Todo) (models.Todo, error) {
	err := r.db.Delete(&todo, ID).Scan(&todo).Error

	return todo, err
}
