package service

import (
	"context"
	"testing"
	"time"

	"github.com/nullexp/finman-user-service/internal/adapter/driven/db/repository"
	"github.com/nullexp/finman-user-service/internal/domain/model"
	portModels "github.com/nullexp/finman-user-service/internal/port/model"
	"github.com/stretchr/testify/assert"
)

func TestRoleServiceCreateRole(t *testing.T) {
	mockRoleRepository := repository.NewMockRoleRepository()
	userMockRepository := repository.NewMockUserRepository()

	rs := NewRoleService(mockRoleRepository, userMockRepository)

	ctx := context.Background()
	request := portModels.CreateRoleRequest{
		Name:        "testrole",
		Permissions: []string{"permission1", "permission2"},
	}

	response, err := rs.CreateRole(ctx, request)
	assert.NoError(t, err)
	assert.NotEmpty(t, response.Id)

	createdRole, err := mockRoleRepository.GetRoleById(ctx, response.Id)
	assert.NoError(t, err)
	assert.Equal(t, "testrole", createdRole.Name)
	assert.Equal(t, []string{"permission1", "permission2"}, createdRole.Permissions)
}

func TestRoleServiceGetRoleById(t *testing.T) {
	mockRoleRepository := repository.NewMockRoleRepository()
	userMockRepository := repository.NewMockUserRepository()

	rs := NewRoleService(mockRoleRepository, userMockRepository)

	ctx := context.Background()
	mockRole := model.Role{
		Name:        "testrole",
		Permissions: []string{"permission1", "permission2"},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	id, err := mockRoleRepository.CreateRole(ctx, mockRole)
	assert.NoError(t, err)

	request := portModels.GetRoleByIdRequest{Id: id}
	response, err := rs.GetRoleById(ctx, request)
	assert.NoError(t, err)
	assert.Equal(t, "testrole", response.Role.Name)
	assert.Equal(t, []string{"permission1", "permission2"}, response.Role.Permissions)
}

func TestRoleServiceGetAllRoles(t *testing.T) {
	mockRoleRepository := repository.NewMockRoleRepository()
	userMockRepository := repository.NewMockUserRepository()

	rs := NewRoleService(mockRoleRepository, userMockRepository)

	ctx := context.Background()
	mockRole1 := model.Role{
		Name:        "testrole1",
		Permissions: []string{"permission1", "permission2"},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	mockRole2 := model.Role{
		Name:        "testrole2",
		Permissions: []string{"permission3", "permission4"},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err := mockRoleRepository.CreateRole(ctx, mockRole1)
	assert.NoError(t, err)
	_, err = mockRoleRepository.CreateRole(ctx, mockRole2)
	assert.NoError(t, err)

	response, err := rs.GetAllRoles(ctx)
	assert.NoError(t, err)
	assert.Len(t, response.Roles, 2)
}

func TestRoleServiceUpdateRole(t *testing.T) {
	mockRoleRepository := repository.NewMockRoleRepository()
	userMockRepository := repository.NewMockUserRepository()

	rs := NewRoleService(mockRoleRepository, userMockRepository)

	ctx := context.Background()
	mockRole := model.Role{
		Name:        "testrole",
		Permissions: []string{"permission1", "permission2"},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	id, err := mockRoleRepository.CreateRole(ctx, mockRole)
	assert.NoError(t, err)

	updateRequest := portModels.UpdateRoleRequest{
		Id:          id,
		Name:        "updatedrole",
		Permissions: []string{"permission3", "permission4"},
	}
	err = rs.UpdateRole(ctx, updateRequest)
	assert.NoError(t, err)

	updatedRole, err := mockRoleRepository.GetRoleById(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, "updatedrole", updatedRole.Name)
	assert.Equal(t, []string{"permission3", "permission4"}, updatedRole.Permissions)
}

func TestRoleServiceDeleteRole(t *testing.T) {
	mockRoleRepository := repository.NewMockRoleRepository()
	userMockRepository := repository.NewMockUserRepository()

	rs := NewRoleService(mockRoleRepository, userMockRepository)

	ctx := context.Background()
	mockRole := model.Role{
		Name:        "testrole",
		Permissions: []string{"permission1", "permission2"},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	id, err := mockRoleRepository.CreateRole(ctx, mockRole)
	assert.NoError(t, err)

	deleteRequest := portModels.DeleteRoleRequest{Id: id}
	err = rs.DeleteRole(ctx, deleteRequest)
	assert.NoError(t, err)

	deletedRole, err := mockRoleRepository.GetRoleById(ctx, id)
	assert.Error(t, err)
	assert.Nil(t, deletedRole)
}
