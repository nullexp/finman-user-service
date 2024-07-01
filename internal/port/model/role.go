package model

import (
	"context"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

type Role struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Permissions []string  `json:"permissions"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CreateRoleRequest struct {
	Name        string   `json:"name" validate:"required"`
	Permissions []string `json:"permissions" validate:"required"`
}

var validPermissions = map[string]bool{
	"ManageUsers":        true,
	"ManageTransactions": true,
	"ManageRoles":        true,
}

func (dto CreateRoleRequest) Validate(ctx context.Context) error {
	validate := validator.New()
	if err := validate.StructCtx(ctx, dto); err != nil {
		return err
	}

	// Remove duplicates and validate each permission
	uniquePermissions := make(map[string]bool)
	var validPermissionsList []string
	for _, perm := range dto.Permissions {
		if validPermissions[perm] {
			if !uniquePermissions[perm] {
				uniquePermissions[perm] = true
				validPermissionsList = append(validPermissionsList, perm)
			}
		} else {
			return errors.New("invalid permission: " + perm)
		}
	}

	dto.Permissions = validPermissionsList
	return nil
}

type CreateRoleResponse struct {
	Id string `json:"id"`
}

type GetRoleByIdRequest struct {
	Id string `json:"id" validate:"required,uuid"`
}

func (dto GetRoleByIdRequest) Validate(ctx context.Context) error {
	validate := validator.New()
	return validate.StructCtx(ctx, dto)
}

type GetRoleByIdResponse struct {
	Role Role `json:"role"`
}

type GetAllRolesResponse struct {
	Roles []Role `json:"roles"`
}

type UpdateRoleRequest struct {
	Id          string   `json:"id" validate:"required,uuid"`
	Name        string   `json:"name" validate:"required"`
	Permissions []string `json:"permissions" validate:"required"`
}

func (dto UpdateRoleRequest) Validate(ctx context.Context) error {
	validate := validator.New()
	err := validate.StructCtx(ctx, dto)
	if err != nil {
		return err
	}

	// Remove duplicates and validate each permission
	uniquePermissions := make(map[string]bool)
	var validPermissionsList []string
	for _, perm := range dto.Permissions {
		if validPermissions[perm] {
			if !uniquePermissions[perm] {
				uniquePermissions[perm] = true
				validPermissionsList = append(validPermissionsList, perm)
			}
		} else {
			return errors.New("invalid permission: " + perm)
		}
	}

	dto.Permissions = validPermissionsList
	return nil
}

type DeleteRoleRequest struct {
	Id string `json:"id" validate:"required,uuid"`
}

func (dto DeleteRoleRequest) Validate(ctx context.Context) error {
	validate := validator.New()
	return validate.StructCtx(ctx, dto)
}

type IsUserPermittedToPermissionRequest struct {
	UserId     string `json:"userId" validate:"required,uuid"`
	Permission string `json:"permission" validate:"required"`
}

func (dto IsUserPermittedToPermissionRequest) Validate(ctx context.Context) error {
	validate := validator.New()
	return validate.StructCtx(ctx, dto)
}

type IsUserPermittedToPermissionResponse struct {
	IsPermitted bool `json:"isPermitted"`
}
