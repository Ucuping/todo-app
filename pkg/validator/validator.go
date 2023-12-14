package validator

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func validationMessage(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "This filed is type email"
	case "number":
		return "This field is type number"
	case "boolean":
		return "This field is type boolean"
	case "uuid4":
		return "This field is type uuid"
	}
	return ""
}

type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var validate *validator.Validate

func Validator(field interface{}) []*ErrorResponse {
	validate = validator.New()
	var errors []*ErrorResponse

	err := validate.Struct(field)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = strings.ToLower(err.Field())
			element.Message = validationMessage(err.Tag())
			errors = append(errors, &element)
		}
	}
	return errors
}
