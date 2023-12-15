package seeders

import (
	"github.com/Ucuping/todo-app/models"
	"github.com/Ucuping/todo-app/pkg/bcrypt"
	"github.com/Ucuping/todo-app/pkg/mysql"
)

func UserSeeder() {
	userPassword, _ := bcrypt.EncryptPassword("root")
	var isActive *int
	isActive = new(int)
	*isActive = 1

	var users = []models.User{
		{
			Name:     "Developer",
			Email:    "dev@example.com",
			Username: "root",
			Password: userPassword,
			IsActive: isActive,
		},
	}

	result := mysql.DB.Create(&users)

	if result.Error != nil {
		panic(result.Error)
	}
}

func UserHasRoleSeeder() {
	var roles []models.Role
	err := mysql.DB.Find(&roles).Error

	if err != nil {
		panic(err)
	}

	var devUser models.User

	err = mysql.DB.First(&devUser, "name = ?", "Developer").Error

	if err != nil {
		panic(err)
	}

	var userHasRoles []*models.UserHasRole

	for _, role := range roles {
		var uhr models.UserHasRole
		uhr.RoleID = role.ID
		uhr.UserID = devUser.ID
		userHasRoles = append(userHasRoles, &uhr)
	}

	err = mysql.DB.Create(&userHasRoles).Error

	if err != nil {
		panic(err)
	}
}
