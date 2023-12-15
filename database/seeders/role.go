package seeders

import (
	"github.com/Ucuping/todo-app/models"
	"github.com/Ucuping/todo-app/pkg/mysql"
)

func RoleSeeder() {
	var roles = []models.Role{
		{
			Name: "Developer",
		},
	}

	result := mysql.DB.Create(&roles)

	if result.Error != nil {
		panic(result.Error)
	}
}

func RoleHasPermissionSeeder() {
	var permissions []models.Permission
	err := mysql.DB.Find(&permissions).Error

	if err != nil {
		panic(err)
	}

	var roles []models.Role
	err = mysql.DB.Find(&roles).Error

	if err != nil {
		panic(err)
	}

	var roleHasPermissions []*models.RoleHasPermission

	for _, permission := range permissions {
		for _, role := range roles {
			var rhp models.RoleHasPermission
			rhp.RoleID = role.ID
			rhp.PermissionID = permission.ID
			roleHasPermissions = append(roleHasPermissions, &rhp)
		}
	}

	err = mysql.DB.Create(&roleHasPermissions).Error

	if err != nil {
		panic(err)
	}
}
