package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

var CacheConfig = cache.Config{
	Next: func(c *fiber.Ctx) bool {
		return c.Query("noCache") == "true"
	},
	Expiration:   5 * time.Minute,
	CacheControl: true,
}
