package routes

import (
	"event-booking/controllers"
	"event-booking/middleware"
	"event-booking/models"

	"github.com/gin-gonic/gin"
)

func EventRoutes(router *gin.Engine) {
	protected := router.Group("/events")
	protected.Use(middleware.JWTAuth())
	{
		protected.POST("", middleware.BindAndValidate(models.Event{}), controllers.CreateEvent)
		protected.GET("", controllers.GetEvents)
		protected.GET("/:id", controllers.GetEvent)
		protected.PUT("/:id", middleware.BindAndValidate(models.Event{}), controllers.UpdateEvent)
		protected.DELETE("/:id", controllers.DeleteEvent)
	}
}
