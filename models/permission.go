package models

type Permission struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `json:"name"`
	DisplayName string  `json:"display_name"`
	Roles       []*Role `gorm:"many2many:role_has_permissions;" json:"roles"`
}
