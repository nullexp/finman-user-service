package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/nullexp/finman-user-service/internal/domain/model"
)

type MockRoleRepository struct {
	roles map[string]model.Role // Simulated in-memory database
}

func NewMockRoleRepository() *MockRoleRepository {
	return &MockRoleRepository{
		roles: make(map[string]model.Role),
	}
}

func (m *MockRoleRepository) CreateRole(ctx context.Context, role model.Role) (string, error) {
	id := uuid.New().String() // Generate UUID
	role.Id = id
	m.roles[id] = role
	return id, nil
}

func (m *MockRoleRepository) GetRoleById(ctx context.Context, id string) (*model.Role, error) {
	role, exists := m.roles[id]
	if !exists {
		return nil, errors.New("role not found")
	}
	return &role, nil
}

func (m *MockRoleRepository) GetAllRoles(ctx context.Context) ([]model.Role, error) {
	var roles []model.Role
	for _, role := range m.roles {
		roles = append(roles, role)
	}
	return roles, nil
}

func (m *MockRoleRepository) UpdateRole(ctx context.Context, role model.Role) error {
	_, exists := m.roles[role.Id]
	if !exists {
		return errors.New("role not found")
	}
	m.roles[role.Id] = role
	return nil
}

func (m *MockRoleRepository) DeleteRole(ctx context.Context, id string) error {
	_, exists := m.roles[id]
	if !exists {
		return errors.New("role not found")
	}
	delete(m.roles, id)
	return nil
}
