package repositories

import (
	"github.com/Ucuping/todo-app/models"
	"gorm.io/gorm"
)

type RoleRepository interface {
	GetAllRole() (*gorm.DB, []models.Role)
	CreateRole(role models.Role) (models.Role, error)
	GetRole(ID int) (models.Role, error)
	UpdateRole(role models.Role) (models.Role, error)
	DeleteRole(ID int, role models.Role) (models.Role, error)
	GetAllPermission() ([]models.Permission, error)
	ChangePermission(roleHasPermissions []models.RoleHasPermission, RoleID int) (models.Role, error)
	CheckPermission(ID uint) error
}

func RepositoryRole(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllRole() (*gorm.DB, []models.Role) {
	// var role models.Role

	model := r.db.Preload("Permissions").Model(&models.Role{}).Where("name != ?", "Dev")

	return model, []models.Role{}
}

func (r *repository) CreateRole(role models.Role) (models.Role, error) {
	err := r.db.Create(&role).Error
	return role, err
}

func (r *repository) GetRole(ID int) (models.Role, error) {
	var role models.Role
	err := r.db.First(&role, ID).Error
	return role, err
}

func (r *repository) UpdateRole(role models.Role) (models.Role, error) {
	err := r.db.Save(&role).Error
	return role, err
}

func (r *repository) DeleteRole(ID int, role models.Role) (models.Role, error) {
	err := r.db.Delete(&role, ID).Scan(&role).Error
	return role, err
}

func (r *repository) GetAllPermission() ([]models.Permission, error) {
	var permissions []models.Permission

	err := r.db.Find(&permissions).Error

	return permissions, err
}

func (r *repository) ChangePermission(roleHasPermissions []models.RoleHasPermission, RoleID int) (models.Role, error) {
	var err error
	var rhp models.RoleHasPermission
	var role models.Role

	r.db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Delete(&rhp, "role_id = ?", RoleID).Error; err != nil {
			return err
		}

		if err = tx.Create(&roleHasPermissions).Error; err != nil {
			return err
		}
		return nil
	})
	err = r.db.Preload("Permissions").First(&role, "id = ?", RoleID).Error

	return role, err
}

func (r *repository) CheckPermission(ID uint) error {
	var permission models.Permission
	err := r.db.First(&permission, ID).Error

	return err
}
