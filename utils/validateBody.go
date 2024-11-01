// utils/validated_data.go
package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// RetrieveValidatedData retrieves and converts the validated data from the context
func RetrieveValidatedData[T any](c *gin.Context) (*T, error) {
	validatedBody, exists := c.Get("validatedBody")
	if !exists {
		return nil, errors.New("unable to retrieve validated data")
	}

	// Type assertion
	data, ok := validatedBody.(*T)
	if !ok {
		return nil, errors.New("invalid data type")
	}
	return data, nil
}
