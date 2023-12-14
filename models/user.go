package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Email     string    `gorm:"size:30;unique" json:"email"`
	Username  string    `gorm:"size:50;not null;unique" json:"username"`
	Password  string    `gorm:"size:255;not null" json:"password"`
	Image     string    `gorm:"null" json:"image"`
	IsActive  *int      `gorm:"type:tinyint(1);default:1;column:is_active" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Roles     []*Role   `gorm:"many2many:user_has_roles;" json:"roles"`
}
