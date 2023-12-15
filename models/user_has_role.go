package models

import "github.com/google/uuid"

type UserHasRole struct {
	UserID uuid.UUID `gorm:"primaryKey" json:"user_id"`
	RoleID uuid.UUID `gorm:"primaryKey;index:,unique" json:"role_id"`
	User   User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Role   Role      `gorm:"foreignKey:RoleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
