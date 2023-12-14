package request

type TodoRequest struct {
	Todo string `json:"todo" validate:"required"`
}
