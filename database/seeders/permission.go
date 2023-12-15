package seeders

import (
	"github.com/Ucuping/todo-app/models"
	"github.com/Ucuping/todo-app/pkg/mysql"
)

func PermissionSeeder() {
	var permissions = []models.Permission{
		{
			Name:        "read-users",
			DisplayName: "Read Users",
		},
		{
			Name:        "create-users",
			DisplayName: "Create Users",
		},
		{
			Name:        "update-users",
			DisplayName: "Update Users",
		},
		{
			Name:        "delete-users",
			DisplayName: "Delete Users",
		},
	}

	result := mysql.DB.Create(&permissions)

	if result.Error != nil {
		panic(result.Error)
	}
}
