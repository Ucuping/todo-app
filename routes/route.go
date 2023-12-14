package routes

import (
	"github.com/Ucuping/todo-app/pkg/middleware"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func Route(app *fiber.App) {
	app.Use(
		helmet.New(middleware.HelmetConfig),
		compress.New(middleware.CompressConfig),
		// cache.New(middleware.CacheConfig),
		// csrf.New(middleware.CSRFConfig),
	)

	app.Route("/api/v1/", func(router fiber.Router) {
		AuthRoute(router)
		router.Use(
			jwtware.New(middleware.JWTMiddlewareConfig),
		)
		RoleRoute(router)
		TodoRoute(router)
		UserRoute(router)
	})
}
