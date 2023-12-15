package models

import "github.com/google/uuid"

type RoleHasPermission struct {
	RoleID       uuid.UUID  `gorm:"primaryKey" json:"role_id"`
	PermissionID uuid.UUID  `gorm:"primaryKey;index:,unique" json:"permission_id"`
	Role         Role       `gorm:"foreignKey:RoleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Permission   Permission `gorm:"foreignKey:PermissionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
