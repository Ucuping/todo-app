package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var CorsConfig = cors.Config{
	AllowOriginsFunc: func(origin string) bool {
		return os.Getenv("ENV") == "development"
	},
	AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
	// AllowOrigins:     "http://localhost:3000",
	AllowCredentials: true,
	AllowMethods: strings.Join([]string{
		fiber.MethodGet,
		fiber.MethodPost,
		fiber.MethodHead,
		fiber.MethodDelete,
	}, ","),
}
