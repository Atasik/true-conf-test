package domain

import "errors"

var (
	ErrUserNotFound   = errors.New("user_not_found")
	ErrNoUpdateValues = errors.New("update_structure_has_no_values")
)
