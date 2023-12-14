package routes

import (
	"github.com/Ucuping/todo-app/handlers"
	"github.com/Ucuping/todo-app/pkg/middleware/permission"
	"github.com/Ucuping/todo-app/pkg/mysql"
	"github.com/Ucuping/todo-app/repositories"
	"github.com/gofiber/fiber/v2"
)

func TodoRoute(api fiber.Router) {
	todoRepository := repositories.RepositoryTodo(mysql.DB)
	h := handlers.HandlerTodo(todoRepository)

	authz := permission.New(permission.Config{
		DB: mysql.DB,
	})

	api.Get("todos", authz.RequiresPermissions("read-users"), h.GetAllTodo)
	api.Post("todos", h.CreateTodo)
	api.Get("todos/:id", h.GetTodo)
	api.Post("todos/:id", h.UpdateTodo)
	api.Delete("todos/:id", h.DeleteTodo)
}
