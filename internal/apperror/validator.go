package apperror

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// ValidateStruct validates the given struct using `validate` tags via the
// go-playground/validator package.
//
// It returns nil if validation passes. If validation fails, it returns the
// first encountered validation error mapped to a specific application-level error.
func ValidateStruct(obj any) error {
	validate := validator.New()
	err := validate.Struct(obj)

	if err == nil {
		return nil
	}

	validationErrors, ok := err.(validator.ValidationErrors)

	if !ok || len(validationErrors) == 0 {
		return err
	}

	ve := validationErrors[0]
	field := ve.StructField()

	switch ve.Tag() {
	case "required":
		return fmt.Errorf("%s: %w", field, ErrRequiredField)
	case "max":
		return fmt.Errorf("%s: %w", field, ErrMaxValueExceeded)
	case "min":
		return fmt.Errorf("%s: %w", field, ErrMinValueNotReached)
	case "email":
		return fmt.Errorf("%s: %w", field, ErrInvalidEmail)
	default:
		return fmt.Errorf("%s: %w", field, ErrInvalidField)
	}
}
