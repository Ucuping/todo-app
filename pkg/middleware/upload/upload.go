package uploadMiddleware

// func UploadMiddleware(next fiber.Handler) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		imageFile, err := c.FormFile("image")
// 		file, err := c.FormFile("file")

// 		imageAllowedTypes := []string{"image/jpeg", "image/png"}

// 		// fmt.Println(err)

// 		if imageFile != nil {
// 			if err != nil {
// 				return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{
// 					Status:  "error",
// 					Message: err.Error(),
// 				})
// 			}
// 			isValid := false
// 			for _, iat := range imageAllowedTypes {
// 				if iat == imageFile.Header.Get("Content-Type") {
// 					isValid = true
// 				}
// 			}

// 			if !isValid {
// 				return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{
// 					Status:  "fail",
// 					Message: "File type is not supported",
// 				})
// 			}

// 			newImageName := fmt.Sprintf("%s-%s", time.Now(), imageFile.Filename)
// 			destination := fmt.Sprintf("./uploads/%s", newImageName)

// 			err := c.SaveFile(file, destination)

// 			if err != nil {
// 				log.Println("Error in saving image :", err.Error())
// 				return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{
// 					Status:  "fail",
// 					Message: "Internal Server Error",
// 				})
// 			}

// 			c.Set("imageName", newImageName)

// 			return next(c)

// 		}

// 		c.Set("imageName", "")

// 		return next(c)
// 	}
// }

// type Middleware struct {
// 	config Config
// }

// func New(config ...Config) *Middleware {
// 	cfg, err := configDefault(config...)
// 	if err != nil {
// 		panic(fmt.Errorf("Fiber: permission middleware error -> %w", err))
// 	}

// 	return &Middleware{
// 		config: cfg,
// 	}
// }
