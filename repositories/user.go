package repositories

import (
	"github.com/Ucuping/todo-app/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser() (*gorm.DB, []models.User)
	CreateUser(user models.User, RoleID uint) (models.User, error)
	GetUser(ID int) (models.User, error)
	UpdateUser(user models.User, RoleID uint) (models.User, error)
	DeleteUser(ID int, user models.User) (models.User, error)
	SetActiveUser(user models.User) (models.User, error)
	CheckRole(RoleID uint) error
	GetUserHasRole(ID uint, RoleID uint) (models.UserHasRole, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllUser() (*gorm.DB, []models.User) {
	// var users []models.User
	model := r.db.Preload("Roles").Joins("JOIN user_has_roles on user_has_roles.user_id = users.id JOIN roles on user_has_roles.role_id = roles.id AND roles.name != ?", "Dev").Model(&models.User{})

	return model, []models.User{}
}

func (r *repository) CreateUser(user models.User, RoleID uint) (models.User, error) {
	var err error
	// var userHasRole models.UserHasRole
	r.db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&user).Error; err != nil {
			return err
		}

		if err = tx.Create(&models.UserHasRole{RoleID: RoleID, UserID: user.ID}).Error; err != nil {
			return err
		}

		return nil
	})
	// err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.Preload("Roles").First(&user, ID).Error
	return user, err
}

func (r *repository) UpdateUser(user models.User, RoleID uint) (models.User, error) {
	var err error
	var userHasRole models.UserHasRole

	r.db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Save(&user).Error; err != nil {
			return err
		}

		if err = tx.Model(&userHasRole).Where("user_id = ?", user.ID).Update("role_id", RoleID).Error; err != nil {
			return err
		}

		return nil
	})
	// err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) DeleteUser(ID int, user models.User) (models.User, error) {
	err := r.db.Delete(&user, ID).Scan(&user).Error

	return user, err
}

func (r *repository) SetActiveUser(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) CheckRole(RoleID uint) error {
	var role models.Role
	err := r.db.First(&role, RoleID).Error

	return err
}

func (r *repository) GetUserHasRole(UserID uint, RoleID uint) (models.UserHasRole, error) {
	var userHasRole models.UserHasRole
	err := r.db.First(&userHasRole, UserID, RoleID).Error

	return userHasRole, err
}
