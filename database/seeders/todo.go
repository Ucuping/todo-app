package seeders

import (
	"github.com/Ucuping/todo-app/models"
	"github.com/Ucuping/todo-app/pkg/mysql"
)

func TodoSeeder() {
	var todos = []models.Todo{
		{
			Todo: "Test",
		},
		{
			Todo: "Test123",
		},
		{
			Todo: "Test1234",
		},
		{
			Todo: "Test12345",
		},
		{
			Todo: "Test123456",
		},
	}

	err := mysql.DB.Create(&todos).Error

	if err != nil {
		panic(err)
	}
}
