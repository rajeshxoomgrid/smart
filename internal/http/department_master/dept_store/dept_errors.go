package deptstore

import "errors"

var (
	// Department does not exist or is already deleted.
	ErrNotFound = errors.New("department not found")

	// Department name already exists.
	ErrDuplicate = errors.New("department already exists")
)
