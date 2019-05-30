package controller

import "errors"

var (
	ErrNotFound = errors.New("book not found")
	ErrBadArgument = errors.New("bad argument")
)
