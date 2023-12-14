package models

type Role struct {
	ID          uint          `gorm:"primaryKey" json:"id"`
	Name        string        `gorm:"type:varchar(255)" json:"name"`
	Permissions []*Permission `gorm:"many2many:role_has_permissions;" json:"permissions"`
	Users       []*User       `gorm:"many2many:user_has_roles;" json:"users"`
}
