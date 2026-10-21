package partconstant

import "errors"

var (
	ErrNotFound  = errors.New("Part not found")
	ErrDuplicate = errors.New("part number already exists")
)
