package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID          uuid.UUID     `gorm:"primaryKey" json:"id"`
	Name        string        `gorm:"type:varchar(255)" json:"name"`
	Permissions []*Permission `gorm:"many2many:role_has_permissions;" json:"permissions"`
	Users       []*User       `gorm:"many2many:user_has_roles;" json:"users"`
}

func (role *Role) BeforeCreate(scope *gorm.DB) error {
	newUuid := uuid.NewString()
	scope.Statement.SetColumn("ID", newUuid)
	return nil
}
