// middleware/validation.go
package middleware

import (
	"net/http"
	"reflect"

	"event-booking/utils"

	"github.com/gin-gonic/gin"
)

// BindAndValidate is a middleware that binds and validates any struct type passed as protoType
func BindAndValidate(protoType interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a new instance of the provided type
		reqBody := reflect.New(reflect.TypeOf(protoType)).Interface()

		// Bind the request body to the instance
		if err := c.ShouldBindJSON(reqBody); err != nil {
			// Output the error for debugging
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
			c.Abort()
			return
		}

		// Validate the bound instance
		if validationErrors := utils.ValidateStruct(reqBody); validationErrors != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
			c.Abort()
			return
		}

		// Pass the validated instance to the request context
		c.Set("validatedBody", reqBody)
		c.Next()
	}
}
