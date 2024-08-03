package lendoErrors

import (
	"errors"
)

var ErrVertexAlreadyExists = errors.New("vertex already exists")
var ErrVertexDoesNotExist = errors.New("vertex does not exist")
