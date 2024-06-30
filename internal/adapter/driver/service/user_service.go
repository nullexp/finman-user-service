package service

import (
	"context"

	"github.com/nullexp/finman-user-service/internal/domain"
	domainModel "github.com/nullexp/finman-user-service/internal/domain/model"
	"github.com/nullexp/finman-user-service/internal/port/driven"
	"github.com/nullexp/finman-user-service/internal/port/driven/db/repository"
	"github.com/nullexp/finman-user-service/internal/port/model"
)

type UserService struct {
	userRepository  repository.UserRepository
	passwordService driven.PasswordService
}

func NewUserService(userRepository repository.UserRepository, passwordService driven.PasswordService) *UserService {
	return &UserService{userRepository: userRepository, passwordService: passwordService}
}

func (us UserService) CreateUser(ctx context.Context, request model.CreateUserRequest) (*model.CreateUserResponse, error) {
	if err := request.Validate(ctx); err != nil {
		return nil, err
	}

	ps, err := us.passwordService.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	id, err := us.userRepository.CreateUser(ctx, domainModel.User{
		Username: request.Username,
		RoleId:   request.RoleId,
		IsAdmin:  false,
		Password: ps,
	})

	if err != nil {
		return nil, err
	}

	return &model.CreateUserResponse{Id: id}, nil
}

func castUserToReadable(user *domainModel.User) model.UserReadable {
	return model.UserReadable{
		Id:        user.Id,
		Username:  user.Username,
		RoleId:    user.RoleId,
		IsAdmin:   user.IsAdmin,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (us UserService) GetUserById(ctx context.Context, request model.GetUserByIdRequest) (*model.GetUserByIdResponse, error) {
	if err := request.Validate(ctx); err != nil {
		return nil, err
	}

	user, err := us.userRepository.GetUserById(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, domain.ErrUserNotFound
	}

	return &model.GetUserByIdResponse{
		User: castUserToReadable(user),
	}, nil
}
func (us UserService) GetAllUsers(ctx context.Context) (*model.GetAllUsersResponse, error) {

	users, err := us.userRepository.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	out := []model.UserReadable{}

	for _, user := range users {
		out = append(out, castUserToReadable(&user))
	}
	return &model.GetAllUsersResponse{
		Users: out,
	}, nil

}
func (us UserService) UpdateUser(ctx context.Context, request model.UpdateUserRequest) error {

	// Please note that this operation should be transactional but it is not
	if err := request.Validate(ctx); err != nil {
		return err
	}

	user, err := us.userRepository.GetUserById(ctx, request.Id)
	if err != nil {
		return err
	}

	if user == nil {
		return domain.ErrUserNotFound
	}

	user.RoleId = request.RoleId

	ps, err := us.passwordService.HashPassword(request.Password)
	if err != nil {
		return err
	}

	user.Password = ps

	return us.userRepository.UpdateUser(ctx, *user)
}
func (us UserService) DeleteUser(ctx context.Context, request model.DeleteUserRequest) error {
	if err := request.Validate(ctx); err != nil {
		return err
	}

	user, err := us.userRepository.GetUserById(ctx, request.Id)
	if err != nil {
		return err
	}

	if user == nil {
		return domain.ErrUserNotFound
	}

	if user.IsAdmin {
		return domain.ErrAdminCantBeRemoved
	}

	return us.userRepository.DeleteUser(ctx, request.Id)
}
func (us UserService) GetUserByUsernameAndPassword(ctx context.Context, request model.GetUserByUsernameAndPasswordRequest) (*model.GetUserByUsernameAndPasswordResponse, error) {
	if err := request.Validate(ctx); err != nil {
		return nil, err
	}

	ps, err := us.passwordService.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	user, err := us.userRepository.GetUserByUsernameAndPassword(ctx, request.Username, ps)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, domain.ErrUserNotFound
	}

	return &model.GetUserByUsernameAndPasswordResponse{User: castUserToReadable(user)}, nil
}
func (us UserService) GetUsersWithPagination(ctx context.Context, request model.GetUsersWithPaginationRequest) (*model.GetUsersWithPaginationResponse, error) {
	if err := request.Validate(ctx); err != nil {
		return nil, err
	}

	users, err := us.userRepository.GetUsersWithPagination(ctx, request.Offset, request.Limit)
	if err != nil {
		return nil, err
	}

	out := []model.UserReadable{}

	for _, user := range users {
		out = append(out, castUserToReadable(&user))
	}
	return &model.GetUsersWithPaginationResponse{
		Users: out,
	}, nil
}
