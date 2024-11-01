// routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	UserRoutes(router)
	EventRoutes(router)
	BookingRoutes(router)
}
