package middleware

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UploadMiddleware(next fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		uniqueUuid := uuid.New()
		imageFile, err := c.FormFile("image")
		// file, err := c.FormFile("file")

		imageAllowedTypes := []string{"image/jpeg", "image/png"}
		const MAX_UPLOAD_FILE_SIZE = 1024 * 2048

		// fmt.Println(imageFile)

		if imageFile != nil {
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{
					Status:  "error",
					Message: "Internal Server Error",
				})
			}

			isValid := false
			for _, iat := range imageAllowedTypes {
				if iat == imageFile.Header.Get("Content-Type") {
					isValid = true
				}
			}

			if !isValid {
				return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{
					Status: "fail",
					Errors: []fiber.Map{
						{
							"field":   "image",
							"message": "File type is not supported",
						},
					},
				})
			}

			if imageFile.Size > MAX_UPLOAD_FILE_SIZE {
				return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{
					Status: "fail",
					Errors: []fiber.Map{
						{
							"field":   "image",
							"message": "File size is too large",
						},
					},
				})
			}

			newImageName := fmt.Sprintf("%s-%s", uniqueUuid.String(), imageFile.Filename)
			destination := fmt.Sprintf("./uploads/%s", newImageName)

			err := c.SaveFile(imageFile, destination)

			if err != nil {
				log.Println("Error in saving image :", err.Error())
				return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{
					Status:  "fail",
					Message: "Internal Server Error",
				})
			}

			c.Locals("imageName", newImageName)

			return next(c)

		}

		c.Locals("imageName", "")

		return next(c)
	}
}
