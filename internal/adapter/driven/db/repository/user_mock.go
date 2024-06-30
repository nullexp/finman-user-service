package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/nullexp/finman-user-service/internal/domain/model"
)

type MockUserRepository struct {
	users map[string]model.User // Simulated in-memory database
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[string]model.User),
	}
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user model.User) (string, error) {
	id := uuid.New().String() // Generate UUID
	user.Id = id
	m.users[id] = user
	return id, nil
}

func (m *MockUserRepository) GetUserById(id string) (model.User, error) {
	user, ok := m.users[id]
	if !ok {
		return model.User{}, errors.New("user not found")
	}
	return user, nil
}

func (m *MockUserRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	for _, user := range m.users {
		users = append(users, user)
	}
	return users, nil
}

func (m *MockUserRepository) UpdateUser(user model.User) error {
	if _, ok := m.users[user.Id]; !ok {
		return errors.New("user not found")
	}
	m.users[user.Id] = user
	return nil
}

func (m *MockUserRepository) DeleteUser(id string) error {
	if _, ok := m.users[id]; !ok {
		return errors.New("user not found")
	}
	delete(m.users, id)
	return nil
}

func (m *MockUserRepository) GetUserByUsernameAndPassword(username, password string) (model.User, error) {
	for _, user := range m.users {
		if user.Username == username && user.Password == password {
			return user, nil
		}
	}
	return model.User{}, errors.New("user not found")
}

func (m *MockUserRepository) GetUsersWithPagination(offset, limit int) ([]model.User, error) {
	var users []model.User
	count := 0
	for _, user := range m.users {
		if count >= offset && count < offset+limit {
			users = append(users, user)
		}
		count++
	}
	return users, nil
}
