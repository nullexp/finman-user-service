package driver

import (
	"context"

	"github.com/nullexp/finman-user-service/internal/port/model"
)

type RoleService interface {
	CreateRole(context.Context, model.CreateRoleRequest) (*model.CreateRoleResponse, error)
	GetRoleById(context.Context, model.GetRoleByIdRequest) (*model.GetRoleByIdResponse, error)
	GetAllRoles(context.Context) (*model.GetAllRolesResponse, error)
	UpdateRole(context.Context, model.UpdateRoleRequest) error
	DeleteRole(context.Context, model.DeleteRoleRequest) error
	IsUserPermittedToPermission(context.Context, model.IsUserPermittedToPermissionRequest) (*model.IsUserPermittedToPermissionResponse, error)
}
