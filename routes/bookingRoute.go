package routes

import (
	"event-booking/controllers"
	"event-booking/middleware"

	"github.com/gin-gonic/gin"
)

func BookingRoutes(router *gin.Engine) {
	protected := router.Group("/bookings")
	protected.Use(middleware.JWTAuth())
	{
		protected.POST("", controllers.BookEvent)
		protected.GET("/:id", controllers.GetBookings)
	}
}
