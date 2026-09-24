package deptservice

import "errors"

var (
	ErrUserIDRequired = errors.New(
		"user id is required",
	)

	ErrDepartmentRequired = errors.New(
		"department is required",
	)

	ErrInvalidID = errors.New(
		"invalid department id",
	)

	ErrInvalidName = errors.New(
		"invalid department name",
	)
)
