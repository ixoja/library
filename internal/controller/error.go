package controller

import "github.com/pkg/errors"

var (
	ErrNotFound = errors.New("book not found")
	ErrBadArgument = errors.New("bad argument")
	ErrInternal = errors.New("internal")
	ErrConflict = errors.New("book update conflict")
)
