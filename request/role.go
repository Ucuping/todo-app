package request

import "github.com/google/uuid"

type RoleRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type ChangePermissionRequest struct {
	Permissions []uuid.UUID `json:"permissions" form:"permissions"`
}
