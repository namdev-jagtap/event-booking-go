package routes

import (
	"event-booking/controllers"
	"event-booking/middleware"

	"github.com/gin-gonic/gin"
)

func EventRoutes(router *gin.Engine) {
	protected := router.Group("/events")
	protected.Use(middleware.JWTAuth())
	{
		protected.POST("", controllers.CreateEvent)
		protected.GET("", controllers.GetEvents)
		protected.GET("/:id", controllers.GetEvent)
		protected.PUT("/:id", controllers.UpdateEvent)
		protected.DELETE("/:id", controllers.DeleteEvent)
	}
}
