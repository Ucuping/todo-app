package uploadMiddleware

// type Config struct {
// 	FileName string
// 	AllowedType         []string
// 	Lookup              func(c *fiber.Ctx) (*multipart.FileHeader, error)
// 	Unauthorized        fiber.Handler
// 	InternalServerError fiber.Handler
// }

// type errorResponse struct {
// 	Status  string `json:"status"`
// 	Message string `json:"message"`
// }

// var ConfigDefault = Config{
// 	FileName: "",
// 	AllowedType: []string{"image/jpeg", "image/png"},
// 	Lookup: func(c *fiber.Ctx) (*multipart.FileHeader, error) {
// 		file, err := c.FormFile(c.Get(""))

// 		return file, err
// 	},
// 	InternalServerError: func(c *fiber.Ctx) error {
// 		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{
// 			Status:  "error",
// 			Message: "Internal Server Error",
// 		})
// 	},
// }

// func configDefault(config ...Config) (Config, error) {
// 	if len(config) < 1 {
// 		return ConfigDefault, nil
// 	}

// 	cfg := config[0]

// 	if cfg.AllowedType == nil {
// 		cfg.AllowedType = ConfigDefault.AllowedType
// 	}

// 	if cfg.Lookup == nil {
// 		cfg.Lookup = ConfigDefault.Lookup
// 	}

// 	if cfg.InternalServerError == nil {
// 		cfg.InternalServerError = ConfigDefault.InternalServerError
// 	}

// 	return cfg, nil
// }
