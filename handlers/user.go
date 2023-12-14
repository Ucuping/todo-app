package handlers

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Ucuping/todo-app/models"
	"github.com/Ucuping/todo-app/pkg/bcrypt"
	"github.com/Ucuping/todo-app/pkg/validator"
	"github.com/Ucuping/todo-app/repositories"
	"github.com/Ucuping/todo-app/request"
	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
)

type handlerUser struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handlerUser {
	return &handlerUser{UserRepository}
}

func (h *handlerUser) GetAllUser(c *fiber.Ctx) error {
	pg := paginate.New()

	model, user := h.UserRepository.GetAllUser()

	return c.JSON(pg.With(model).Request(c.Request()).Response(&user))
}

func (h *handlerUser) CreateUser(c *fiber.Ctx) error {
	request := new(request.UserCreateRequest)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Status: "error", Message: "Internal Server Error"})
	}

	errors := validator.Validator(request)

	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errorResponse{Status: "fail", Errors: errors})
	}

	password, err := bcrypt.EncryptPassword(request.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Status: "error", Message: "Internal Server Error"})
	}

	err = h.UserRepository.CheckRole(request.Role_Id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(errorResponse{
			Status:  "fail",
			Message: "Role not found",
		})
	}

	// isActive, _ := strconv.Atoi(request.Is_Active)

	user := models.User{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
		Password: password,
		IsActive: request.Is_Active,
		Image:    c.Locals("imageName").(string),
	}

	data, err := h.UserRepository.CreateUser(user, request.Role_Id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Status: "error", Message: "Internal Server Error"})
	}

	// err = h.UserRepository.AssignRole(data.ID, request.Role_Id)

	// if err != nil {
	// 	return
	// }

	return c.JSON(successResponse{Status: "success", Message: "Successfully created user", Data: data})
}

func (h *handlerUser) GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user, err := h.UserRepository.GetUser(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(errorResponse{Status: "fail", Message: "User not found"})
	}

	return c.JSON(successResponse{Status: "success", Data: user})
}

func (h *handlerUser) UpdateUser(c *fiber.Ctx) error {
	// _, err := c.FormFile("image")

	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{
	// 		Status:  "error",
	// 		Message: err.Error(),
	// 	})
	// }

	// fmt.Println(err)
	request := new(request.UserUpdateRequest)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Status: "error", Message: "Internal Server Error"})
	}

	id, _ := strconv.Atoi(c.Params("id"))

	user, err := h.UserRepository.GetUser(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(errorResponse{Status: "fail", Message: "User not found"})
	}

	errors := validator.Validator(request)

	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errorResponse{
			Status: "fail",
			Errors: errors,
		})
	}

	err = h.UserRepository.CheckRole(request.Role_Id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(errorResponse{
			Status:  "fail",
			Message: "Role not found",
		})
	}

	// _, err := h.UserRepository.GetUserHasRole(user.ID, request.Role_Id)

	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{
	// 		Status:  "error",
	// 		Message: "Internal Server Error",
	// 	})
	// }

	if user.Image != "" && c.Locals("imageName").(string) != "" {
		if _, err := os.Stat(fmt.Sprintf("./uploads/%s", user.Image)); err == nil {
			err = os.Remove(fmt.Sprintf("./uploads/%s", user.Image))

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{
					Status:  "error",
					Message: "Internal Server Error",
				})
			}
		}
	}

	imageName := ""

	if c.Locals("imageName").(string) != "" {
		imageName = c.Locals("imageName").(string)
	} else {
		imageName = user.Image
	}

	user.Name = request.Name
	user.Username = request.Username
	user.Email = request.Email
	user.IsActive = request.Is_Active
	user.Image = imageName

	data, err := h.UserRepository.UpdateUser(user, request.Role_Id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Status: "fail", Message: "Internal Server Error"})
	}

	return c.JSON(successResponse{Status: "success", Message: "Successfully updated user", Data: data})
}

func (h *handlerUser) DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user, err := h.UserRepository.GetUser(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(errorResponse{Status: "fail", Message: "User not found"})
	}

	data, err := h.UserRepository.DeleteUser(id, user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Status: "fail", Message: "Internal Server Error"})
	}

	return c.JSON(successResponse{Status: "success", Message: "Successfully deleted user", Data: data})
}

func (h *handlerUser) SetActiveUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user, err := h.UserRepository.GetUser(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(errorResponse{
			Status:  "fail",
			Message: "User not found",
		})
	}

	if *user.IsActive == 0 {
		*user.IsActive = 1
	} else {
		*user.IsActive = 0
	}

	data, err := h.UserRepository.SetActiveUser(user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.JSON(successResponse{
		Status:  "success",
		Message: "Successfully updated user",
		Data:    data,
	})
}
