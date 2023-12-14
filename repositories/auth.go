package repositories

import (
	"github.com/Ucuping/todo-app/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Login(username string) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Login(username string) (models.User, error) {
	var user models.User
	err := r.db.Preload("Roles.Permissions").First(&user, "username = ?", username).Error

	return user, err
}
