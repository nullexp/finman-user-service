package repository

import (
	"context"

	"github.com/nullexp/finman-user-service/internal/domain/model"
)

type RoleRepository interface {
	CreateRole(context.Context, model.Role) (string, error)
	GetRoleById(context.Context, string) (*model.Role, error)
	GetAllRoles(context.Context) ([]model.Role, error)
	UpdateRole(context.Context, model.Role) error
	DeleteRole(context.Context, string) error
}
