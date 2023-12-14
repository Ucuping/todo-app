package database

import (
	"fmt"

	"github.com/Ucuping/todo-app/models"
	"github.com/Ucuping/todo-app/pkg/mysql"
)

func MigrateTable() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		// &models.RoleHasPermission{},
		&models.Permission{},
		&models.Role{},
		&models.Todo{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migrate table failed")
	}

	fmt.Println("Migrate table success")
}
