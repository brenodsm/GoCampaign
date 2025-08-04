package apperror

import "errors"

var (
	ErrInternal error = errors.New("internal server error")
)
