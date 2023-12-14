package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
)

var CSRFConfig = csrf.Config{
	KeyLookup:      "header:X-Csrf-Token",
	CookieName:     "csrf_token",
	CookieSameSite: "Lax",
	Expiration:     10 * time.Minute,
	KeyGenerator:   utils.UUIDv4,
	CookieHTTPOnly: true,
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.Status(fiber.StatusForbidden).JSON(errorResponse{Status: "error", Message: "Forbidden"})
	},
}
