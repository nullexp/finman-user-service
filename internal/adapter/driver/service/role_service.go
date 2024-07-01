package service

import (
	"context"

	"github.com/nullexp/finman-user-service/internal/domain"
	domainModel "github.com/nullexp/finman-user-service/internal/domain/model"
	"github.com/nullexp/finman-user-service/internal/port/driven/db/repository"
	"github.com/nullexp/finman-user-service/internal/port/model"
)

type RoleService struct {
	roleRepository repository.RoleRepository
	userRepository repository.UserRepository
}

func NewRoleService(roleRepository repository.RoleRepository, userRepository repository.UserRepository) *RoleService {
	return &RoleService{roleRepository: roleRepository, userRepository: userRepository}
}

func (rs RoleService) CreateRole(ctx context.Context, request model.CreateRoleRequest) (*model.CreateRoleResponse, error) {
	if err := request.Validate(ctx); err != nil {
		return nil, err
	}

	id, err := rs.roleRepository.CreateRole(ctx, domainModel.Role{
		Name:        request.Name,
		Permissions: request.Permissions,
	})
	if err != nil {
		return nil, err
	}

	return &model.CreateRoleResponse{Id: id}, nil
}

func (rs RoleService) GetRoleById(ctx context.Context, request model.GetRoleByIdRequest) (*model.GetRoleByIdResponse, error) {
	if err := request.Validate(ctx); err != nil {
		return nil, err
	}

	role, err := rs.roleRepository.GetRoleById(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	if role == nil {
		return nil, domain.ErrRoleNotFound
	}

	return &model.GetRoleByIdResponse{
		Role: castRoleToReadable(role),
	}, nil
}

func (rs RoleService) GetAllRoles(ctx context.Context) (*model.GetAllRolesResponse, error) {
	roles, err := rs.roleRepository.GetAllRoles(ctx)
	if err != nil {
		return nil, err
	}

	return &model.GetAllRolesResponse{
		Roles: castRolesToReadable(roles),
	}, nil
}

func (rs RoleService) UpdateRole(ctx context.Context, request model.UpdateRoleRequest) error {
	if err := request.Validate(ctx); err != nil {
		return err
	}

	err := rs.roleRepository.UpdateRole(ctx, domainModel.Role{
		Id:          request.Id,
		Name:        request.Name,
		Permissions: request.Permissions,
	})
	if err != nil {
		return err
	}

	return nil
}

func (rs RoleService) DeleteRole(ctx context.Context, request model.DeleteRoleRequest) error {
	if err := request.Validate(ctx); err != nil {
		return err
	}

	err := rs.roleRepository.DeleteRole(ctx, request.Id)
	if err != nil {
		return err
	}

	return nil
}

func (rs RoleService) IsUserPermittedToPermission(ctx context.Context, request model.IsUserPermittedToPermissionRequest) (*model.IsUserPermittedToPermissionResponse, error) {
	if err := request.Validate(ctx); err != nil {
		return nil, err
	}

	user, err := rs.userRepository.GetUserById(ctx, request.UserId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, domain.ErrUserNotFound
	}

	role, err := rs.roleRepository.GetRoleById(ctx, user.RoleId)
	if err != nil {
		return nil, err
	}

	if role == nil {
		return nil, domain.ErrRoleNotFound
	}

	for _, v := range role.Permissions {
		if v == request.Permission {
			return &model.IsUserPermittedToPermissionResponse{IsPermitted: true}, nil
		}
	}
	return &model.IsUserPermittedToPermissionResponse{IsPermitted: false}, nil
}

func castRoleToReadable(role *domainModel.Role) model.Role {
	return model.Role{
		Id:          role.Id,
		Name:        role.Name,
		Permissions: role.Permissions,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}
}

func castRolesToReadable(roles []domainModel.Role) []model.Role {
	readableRoles := make([]model.Role, len(roles))
	for i, role := range roles {
		readableRoles[i] = castRoleToReadable(&role)
	}
	return readableRoles
}
