package permission

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Middleware struct {
	config Config
}

func New(config ...Config) *Middleware {
	cfg, err := configDefault(config...)
	if err != nil {
		panic(fmt.Errorf("Fiber: permission middleware error -> %w", err))
	}

	return &Middleware{
		config: cfg,
	}
}

func (m *Middleware) RequiresPermissions(permission string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if permission == "" {
			return c.Next()
		}

		token := m.config.Lookup(c)
		if token == nil {
			return m.config.Unauthorized(c)
		}

		isAccess := false

		for _, data := range token["permissions"].([]interface{}) {
			if data.(map[string]interface{})["name"] == permission {
				isAccess = true
			}
		}

		if !isAccess {
			return m.config.Forbidden(c)
		}

		return c.Next()

	}
}
