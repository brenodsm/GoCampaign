package apperror

import "errors"

var (
	ErrInternal error = errors.New("internal server error")

	ErrRequiredField      error = errors.New("field is required")
	ErrMaxValueExceeded   error = errors.New("value exceeds max limit")
	ErrMinValueNotReached error = errors.New("value below min limit")
	ErrInvalidEmail       error = errors.New("invalid email")
	ErrInvalidField       error = errors.New("invalid field")
	ErrCampaignNotFound   error = errors.New("campaign not found")
)
