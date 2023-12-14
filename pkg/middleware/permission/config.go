package permission

import (
	jwtToken "github.com/Ucuping/todo-app/pkg/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type Config struct {
	DB           *gorm.DB
	Lookup       func(c *fiber.Ctx) jwt.MapClaims
	Unauthorized fiber.Handler
	Forbidden    fiber.Handler
}

type errorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var ConfigDefault = Config{
	DB: &gorm.DB{},
	Lookup: func(c *fiber.Ctx) jwt.MapClaims {
		reqToken := c.Cookies("Authorization")
		// token := strings.Split(reqToken, "Bearer ")[1]
		decodeToken, err := jwtToken.DecodeToken(reqToken)
		if err != nil {
			return nil
		}

		return decodeToken
	},
	Unauthorized: func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusUnauthorized).JSON(errorResponse{Status: "fail", Message: "Unauthorized"})
	},
	Forbidden: func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusForbidden).JSON(errorResponse{Status: "fail", Message: "Forbidden"})
	},
}

func configDefault(config ...Config) (Config, error) {
	if len(config) < 1 {
		return ConfigDefault, nil
	}

	cfg := config[0]

	// if cfg.DB == nil {
	// 	cfg.DB = ConfigDefault.DB
	// }

	if cfg.Lookup == nil {
		cfg.Lookup = ConfigDefault.Lookup
	}

	if cfg.Unauthorized == nil {
		cfg.Unauthorized = ConfigDefault.Unauthorized
	}

	if cfg.Forbidden == nil {
		cfg.Forbidden = ConfigDefault.Forbidden
	}

	return cfg, nil
}
