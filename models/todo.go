package models

type Todo struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Todo string `json:"todo"`
}
