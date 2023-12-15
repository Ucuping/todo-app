package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	ID   uuid.UUID `gorm:"primaryKey" json:"id"`
	Todo string    `json:"todo"`
}

func (todo *Todo) BeforeCreate(scope *gorm.DB) error {
	newUuid := uuid.NewString()
	scope.Statement.SetColumn("ID", newUuid)
	return nil
}
