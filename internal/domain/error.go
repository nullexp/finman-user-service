package domain

import "errors"

var (
	ErrAdminCantBeRemoved = errors.New("ADMIN_CANT_BE_REMOVED: Admins cannot be removed")
	ErrUserNotFound       = errors.New("USER_NOT_FOUND: User not found")
)
