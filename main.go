package main

import (
	"os"

	"github.com/Ucuping/todo-app/database"
	"github.com/Ucuping/todo-app/pkg/mysql"
	"github.com/Ucuping/todo-app/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env file")
	}

	mysql.DatabaseInit()
	database.MigrateTable()
}

func main() {
	app := fiber.New()

	app.Static("/uploads", "./uploads")

	routes.Route(app)

	// password := "root"
	// fmt.Println(bcrypt.EncryptPassword(password))

	var port string

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	} else {
		port = "8080"
	}

	app.Listen(":" + port)
}
