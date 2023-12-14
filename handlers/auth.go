package handlers

import (
	"os"
	"strconv"
	"time"

	"github.com/Ucuping/todo-app/pkg/bcrypt"
	jwtToken "github.com/Ucuping/todo-app/pkg/jwt"
	"github.com/Ucuping/todo-app/pkg/validator"
	"github.com/Ucuping/todo-app/repositories"
	"github.com/Ucuping/todo-app/request"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Login(c *fiber.Ctx) error {
	request := new(request.AuthRequest)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Status: "fail", Message: err.Error()})
	}

	errors := validator.Validator(request)

	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errorResponse{Status: "fail", Errors: errors})
	}

	user, err := h.AuthRepository.Login(request.Username)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errorResponse{Status: "fail", Message: "Wrong username or password"})
	}

	isValid := bcrypt.CheckEncryptedPassword(request.Password, user.Password)

	if !isValid {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errorResponse{Status: "fail", Message: "Wrong username or password"})
	}

	type permissionsStruct struct {
		Name string `json:"name"`
	}

	var permissions []*permissionsStruct

	for _, permision := range user.Roles[0].Permissions {
		var p permissionsStruct
		p.Name = permision.Name
		permissions = append(permissions, &p)
	}

	jwtExp, _ := strconv.Atoi(os.Getenv("JWT_EXPIRY"))

	claims := jwt.MapClaims{
		"iss":         user.ID,
		"exp":         time.Now().Add(time.Minute * time.Duration(jwtExp)).Unix(),
		"permissions": permissions,
	}

	token, err := jwtToken.GenerateToken(&claims)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Status: "fail", Message: "Internal Server Error"})
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "Authorization"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Minute * time.Duration(jwtExp)).UTC()
	cookie.HTTPOnly = true
	cookie.SameSite = fiber.CookieSameSiteLaxMode
	cookie.Secure = false

	c.Cookie(cookie)

	return c.JSON(successResponse{Status: "success", Message: "Login success", Data: fiber.Map{
		"user": user,
		// "token": token,
	}})
}

func (h *handlerAuth) Verify(c *fiber.Ctx) error {
	return c.JSON(successResponse{Status: "success", Message: "Login verified"})
}

func (h *handlerAuth) Logout(c *fiber.Ctx) error {
	cookie := new(fiber.Cookie)
	cookie.Name = "Authorization"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-time.Hour).UTC()
	cookie.HTTPOnly = true
	cookie.SameSite = fiber.CookieSameSiteLaxMode
	cookie.Secure = false

	c.Cookie(cookie)

	return c.JSON(successResponse{
		Status:  "success",
		Message: "Logged out",
	})
}
