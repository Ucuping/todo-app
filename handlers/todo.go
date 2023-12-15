package handlers

import (
	"github.com/Ucuping/todo-app/models"
	"github.com/Ucuping/todo-app/pkg/validator"
	"github.com/Ucuping/todo-app/repositories"
	"github.com/Ucuping/todo-app/request"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/morkid/paginate"
)

type handlerTodo struct {
	TodoRepository repositories.TodoRepository
}

func HandlerTodo(TodoRepository repositories.TodoRepository) *handlerTodo {
	return &handlerTodo{TodoRepository}
}

func (h *handlerTodo) GetAllTodo(c *fiber.Ctx) error {
	pg := paginate.New()

	model, todos := h.TodoRepository.GetAllTodo()

	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Status: "fail", Message: "Internal Server Error"})
	// }

	return c.JSON(pg.With(model).Request(c.Request()).Response(&todos))
}

func (h *handlerTodo) CreateTodo(c *fiber.Ctx) error {
	request := new(request.TodoRequest)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Status: "fail", Message: "Internal Server Error"})
	}

	errors := validator.Validator(request)

	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errorResponse{Status: "fail", Errors: errors})
	}

	todo := models.Todo{
		Todo: request.Todo,
	}

	data, err := h.TodoRepository.CreateTodo(todo)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Status: "fail", Message: "Internal Server Error"})
	}

	return c.JSON(successResponse{Status: "success", Message: "Successfuly created todo", Data: data})
}

func (h *handlerTodo) GetTodo(c *fiber.Ctx) error {
	id := uuid.MustParse(c.Params("id"))

	todo, err := h.TodoRepository.GetTodo(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(errorResponse{Status: "fail", Message: "Todo not found"})
	}

	return c.JSON(successResponse{Status: "success", Data: todo})
}

func (h *handlerTodo) UpdateTodo(c *fiber.Ctx) error {
	request := new(request.TodoRequest)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Status: "fail", Message: "Internal Server Error"})
	}

	id := uuid.MustParse(c.Params("id"))

	todo, err := h.TodoRepository.GetTodo(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(errorResponse{Status: "fail", Message: "Todo not found"})
	}

	todo.Todo = request.Todo

	data, err := h.TodoRepository.UpdateTodo(todo)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Status: "fail", Message: "Internal Server Error"})
	}

	return c.JSON(successResponse{Status: "success", Message: "Successfully updated todo", Data: data})
}

func (h *handlerTodo) DeleteTodo(c *fiber.Ctx) error {
	id := uuid.MustParse(c.Params("id"))

	todo, err := h.TodoRepository.GetTodo(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(errorResponse{Status: "fail", Message: "Todo not found"})
	}

	data, err := h.TodoRepository.DeleteTodo(id, todo)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Status: "fail", Message: "Internal Server Error"})
	}

	return c.JSON(successResponse{Status: "success", Message: "Successfully deleted todo", Data: data})
}
