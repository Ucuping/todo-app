package handlers

import (
	"strconv"

	"github.com/Ucuping/todo-app/models"
	"github.com/Ucuping/todo-app/pkg/validator"
	"github.com/Ucuping/todo-app/repositories"
	"github.com/Ucuping/todo-app/request"
	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
)

type handlerRole struct {
	RoleRepository repositories.RoleRepository
}

func HandlerRole(RoleRepository repositories.RoleRepository) *handlerRole {
	return &handlerRole{RoleRepository}
}

func (h *handlerRole) GetAllRole(c *fiber.Ctx) error {
	pg := paginate.New()

	model, roles := h.RoleRepository.GetAllRole()

	return c.JSON(pg.With(model).Request(c.Request()).Response(&roles))
}

func (h *handlerRole) CreateRole(c *fiber.Ctx) error {
	request := new(request.RoleRequest)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Status:  "error",
			Message: "Internal Server Error",
		})
	}

	errors := validator.Validator(request)

	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errorResponse{
			Status: "fail",
			Errors: errors,
		})
	}

	role := models.Role{
		Name: request.Name,
	}

	data, err := h.RoleRepository.CreateRole(role)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Status:  "error",
			Message: "Internal Server Error",
		})
	}

	return c.JSON(successResponse{
		Status:  "success",
		Message: "Successfully created role",
		Data:    data,
	})
}

func (h *handlerRole) GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role, err := h.RoleRepository.GetRole(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(errorResponse{
			Status:  "fail",
			Message: "Role not found",
		})
	}

	return c.JSON(successResponse{
		Status: "success",
		Data:   role,
	})
}

func (h *handlerRole) UpdateRole(c *fiber.Ctx) error {
	request := new(request.RoleRequest)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Status:  "error",
			Message: "Internal Server Error",
		})
	}

	id, _ := strconv.Atoi(c.Params("id"))

	role, err := h.RoleRepository.GetRole(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(errorResponse{
			Status:  "fail",
			Message: "Role not found",
		})
	}

	role.Name = request.Name

	data, err := h.RoleRepository.UpdateRole(role)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Status:  "error",
			Message: "Internal Server Error",
		})
	}

	return c.JSON(successResponse{
		Status: "success",
		Data:   data,
	})
}

func (h *handlerRole) DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role, err := h.RoleRepository.GetRole(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(errorResponse{
			Status:  "fail",
			Message: "Role not found",
		})
	}

	data, err := h.RoleRepository.DeleteRole(id, role)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Status:  "error",
			Message: "Internal Server Error",
		})
	}

	return c.JSON(successResponse{
		Status:  "success",
		Message: "Successfully deleted role",
		Data:    data,
	})
}

func (h *handlerRole) GetAllPermission(c *fiber.Ctx) error {
	permissions, err := h.RoleRepository.GetAllPermission()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Status:  "error",
			Message: "Internal Server Error",
		})
	}

	return c.JSON(successResponse{
		Status: "success",
		Data:   permissions,
	})
}

func (h *handlerRole) ChangePermission(c *fiber.Ctx) error {
	request := new(request.ChangePermissionRequest)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Status:  "error",
			Message: "Internal Server Error",
		})
	}

	id, _ := strconv.Atoi(c.Params("id"))

	role, err := h.RoleRepository.GetRole(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(errorResponse{
			Status:  "fail",
			Message: "Role not found",
		})
	}

	var roleHasPermissions []models.RoleHasPermission

	if len(request.Permissions) > 0 {
		for _, permission := range request.Permissions {
			err := h.RoleRepository.CheckPermission(permission)
			if err != nil {
				return c.Status(fiber.StatusNotFound).JSON(errorResponse{
					Status:  "fail",
					Message: "Permission not found",
				})
			}

			var rhp models.RoleHasPermission
			rhp.PermissionID = permission
			rhp.RoleID = role.ID
			roleHasPermissions = append(roleHasPermissions, rhp)
		}
	}

	data, err := h.RoleRepository.ChangePermission(roleHasPermissions, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Status:  "error",
			Message: "Internal Server Error",
		})
	}

	return c.JSON(successResponse{
		Status:  "success",
		Message: "Successfully changed permission",
		Data:    data,
	})

}
