package routes

import (
	"github.com/Ucuping/todo-app/handlers"
	"github.com/Ucuping/todo-app/pkg/mysql"
	"github.com/Ucuping/todo-app/repositories"
	"github.com/gofiber/fiber/v2"
)

func RoleRoute(api fiber.Router) {
	roleRepository := repositories.RepositoryRole(mysql.DB)
	h := handlers.HandlerRole(roleRepository)

	api.Get("roles", h.GetAllRole)
	api.Post("roles", h.CreateRole)
	api.Get("roles/get-permissions", h.GetAllPermission)
	api.Get("roles/:id", h.GetRole)
	api.Post("roles/:id", h.UpdateRole)
	api.Delete("roles/:id", h.DeleteRole)
	api.Post("roles/:id/change-permission", h.ChangePermission)
}
