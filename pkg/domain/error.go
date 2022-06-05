package domain

import "errors"

var (
	//ErrInternal Internal error
	ErrInternal = errors.New("internal")
	//ErrNotFound Not found error
	ErrNotFound = errors.New("not found")
	//ErrConflict Conflict error
	ErrConflict = errors.New("conflict")
)
