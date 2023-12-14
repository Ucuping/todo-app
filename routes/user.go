package routes

import (
	"github.com/Ucuping/todo-app/handlers"
	"github.com/Ucuping/todo-app/pkg/middleware"
	"github.com/Ucuping/todo-app/pkg/middleware/permission"
	"github.com/Ucuping/todo-app/pkg/mysql"
	"github.com/Ucuping/todo-app/repositories"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(api fiber.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	authz := permission.New(permission.Config{
		DB: mysql.DB,
	})

	api.Get("users", authz.RequiresPermissions("read-users"), h.GetAllUser)
	api.Post("users", authz.RequiresPermissions("create-users"), middleware.UploadMiddleware(h.CreateUser))
	api.Get("users/:id", authz.RequiresPermissions("update-users"), h.GetUser)
	api.Post("users/:id", authz.RequiresPermissions("update-users"), middleware.UploadMiddleware(h.UpdateUser))
	api.Delete("users/:id", authz.RequiresPermissions("delete-users"), h.DeleteUser)
	api.Get("users/:id/set-active", authz.RequiresPermissions("update-users"), h.SetActiveUser)
}
