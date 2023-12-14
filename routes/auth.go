package routes

import (
	"github.com/Ucuping/todo-app/handlers"
	"github.com/Ucuping/todo-app/pkg/middleware"
	"github.com/Ucuping/todo-app/pkg/mysql"
	"github.com/Ucuping/todo-app/repositories"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func AuthRoute(api fiber.Router) {
	authRepository := repositories.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(authRepository)

	api.Post("login", h.Login)

	api.Use(jwtware.New(middleware.JWTMiddlewareConfig))

	api.Get("verify", h.Verify)
	api.Get("logout", h.Logout)
}
