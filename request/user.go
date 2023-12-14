package request

type UserCreateRequest struct {
	Name      string `json:"name" form:"name" validate:"required"`
	Username  string `json:"username" form:"username" validate:"required"`
	Email     string `json:"email" form:"email" validate:"required"`
	Password  string `json:"password" form:"password" validate:"required"`
	Is_Active *int   `json:"is_active" form:"is_active" validate:"required"`
	Role_Id   uint   `json:"role_id" form:"role_id" validate:"required"`
}

type UserUpdateRequest struct {
	Name      string `json:"name" form:"name" validate:"required"`
	Username  string `json:"username" form:"username" validate:"required"`
	Email     string `json:"email" form:"email" validate:"required,email"`
	Is_Active *int   `json:"is_active" form:"is_active" validate:"required"`
	Role_Id   uint   `json:"role_id" form:"role_id" validate:"required"`
}
