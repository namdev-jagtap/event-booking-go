package routes

import (
	"event-booking/controllers"
	"event-booking/middleware"
	"event-booking/models"

	"github.com/gin-gonic/gin"
)

func BookingRoutes(router *gin.Engine) {
	protected := router.Group("/bookings")
	protected.Use(middleware.JWTAuth())
	{
		protected.POST("", middleware.BindAndValidate(models.Booking{}), controllers.BookEvent)
		protected.GET("/:id", controllers.GetBookings)
	}
}
