package request

type RoleRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type ChangePermissionRequest struct {
	Permissions []uint `json:"permissions" form:"permissions"`
}
