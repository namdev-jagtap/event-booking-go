// utils/validator.go
package utils

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidationErrors is a map that holds field validation errors
type ValidationErrors map[string]string

// ValidateStruct validates any struct and returns a map of errors if validation fails
func ValidateStruct(s interface{}) ValidationErrors {
	err := validate.Struct(s)
	if err != nil {
		errors := make(ValidationErrors)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = err.Tag() // Customize the error message here if needed
		}
		return errors
	}
	return nil
}
