package grpc

import (
	"context"
	"time"

	userv1 "github.com/nullexp/finman-user-service/internal/adapter/driver/grpc/proto/user/v1"
	"github.com/nullexp/finman-user-service/internal/port/driver"
	"github.com/nullexp/finman-user-service/internal/port/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RoleService struct {
	userv1.UnimplementedRoleServiceServer
	service driver.RoleService
}

func NewRoleService(rs driver.RoleService) *RoleService {
	return &RoleService{service: rs}
}

func (rs RoleService) CreateRole(ctx context.Context, request *userv1.CreateRoleRequest) (*userv1.CreateRoleResponse, error) {
	response, err := rs.service.CreateRole(ctx, model.CreateRoleRequest{
		Name:        request.Name,
		Permissions: request.Permissions,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "CreateRole failed: %v", err)
	}

	return &userv1.CreateRoleResponse{Id: response.Id}, nil
}

func (rs RoleService) GetRoleById(ctx context.Context, request *userv1.GetRoleByIdRequest) (*userv1.GetRoleByIdResponse, error) {
	response, err := rs.service.GetRoleById(ctx, model.GetRoleByIdRequest{Id: request.Id})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "GetRoleById failed: %v", err)
	}

	return &userv1.GetRoleByIdResponse{Role: castRoleToGrpc(response.Role)}, nil
}

func (rs RoleService) GetAllRoles(ctx context.Context, request *userv1.GetAllRolesRequest) (*userv1.GetAllRolesResponse, error) {
	response, err := rs.service.GetAllRoles(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "GetAllRoles failed: %v", err)
	}

	roles := castRolesToGrpc(response.Roles)
	return &userv1.GetAllRolesResponse{Roles: roles}, nil
}

func (rs RoleService) UpdateRole(ctx context.Context, request *userv1.UpdateRoleRequest) (*userv1.UpdateRoleResponse, error) {
	err := rs.service.UpdateRole(ctx, model.UpdateRoleRequest{
		Id:          request.Id,
		Name:        request.Name,
		Permissions: request.Permissions,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "UpdateRole failed: %v", err)
	}

	return &userv1.UpdateRoleResponse{}, nil
}

func (rs RoleService) DeleteRole(ctx context.Context, request *userv1.DeleteRoleRequest) (*userv1.DeleteRoleResponse, error) {
	err := rs.service.DeleteRole(ctx, model.DeleteRoleRequest{Id: request.Id})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "DeleteRole failed: %v", err)
	}

	return &userv1.DeleteRoleResponse{}, nil
}

func (rs RoleService) IsUserPermittedToPermission(ctx context.Context, request *userv1.IsUserPermittedToPermissionRequest) (*userv1.IsUserPermittedToPermissionResponse, error) {
	response, err := rs.service.IsUserPermittedToPermission(ctx, model.IsUserPermittedToPermissionRequest{
		UserId:     request.UserId,
		Permission: request.Permission,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "IsUserPermittedToPermission failed: %v", err)
	}

	return &userv1.IsUserPermittedToPermissionResponse{IsPermitted: response.IsPermitted}, nil
}

func castRoleToGrpc(role model.Role) *userv1.Role {
	return &userv1.Role{
		Id:          role.Id,
		Name:        role.Name,
		Permissions: role.Permissions,
		CreatedAt:   role.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   role.UpdatedAt.Format(time.RFC3339),
	}
}

func castRolesToGrpc(roles []model.Role) []*userv1.Role {
	grpcRoles := make([]*userv1.Role, len(roles))
	for i, role := range roles {
		grpcRoles[i] = castRoleToGrpc(role)
	}
	return grpcRoles
}
