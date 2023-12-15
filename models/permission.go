package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	ID          uuid.UUID `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	DisplayName string    `json:"display_name"`
	Roles       []*Role   `gorm:"many2many:role_has_permissions;" json:"roles"`
}

func (permission *Permission) BeforeCreate(scope *gorm.DB) error {
	newUuid := uuid.NewString()
	scope.Statement.SetColumn("ID", newUuid)
	return nil
}
