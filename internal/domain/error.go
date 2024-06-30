package domain

import "errors"

var ErrAdminCantBeRemoved = errors.New("ADMIN_CANT_BE_REMOVED: Admins cannot be removed")
var ErrUserNotFound = errors.New("USER_NOT_FOUND: User not found")
