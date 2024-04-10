package vl

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func Validate(request interface{}) *validator.ValidationErrors {
	if err := validator.New().Struct(request); err != nil {
		var validatorErrs validator.ValidationErrors
		errors.As(err, &validatorErrs)

		return &validatorErrs
	}

	return nil
}
