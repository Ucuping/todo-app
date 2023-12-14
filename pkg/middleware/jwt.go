package middleware

import (
	jwtToken "github.com/Ucuping/todo-app/pkg/jwt"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

var JWTMiddlewareConfig = jwtware.Config{
	SigningKey:  jwtware.SigningKey{Key: []byte(jwtToken.SecretKey)},
	TokenLookup: "cookie:Authorization",
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.Status(fiber.StatusUnauthorized).JSON(errorResponse{
			Status:  "fail",
			Message: "Your not authorized",
		})
	},
}
