package routes

import (
	"event-booking/controllers"
	"event-booking/middleware"
	"event-booking/models"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/register", middleware.BindAndValidate(models.User{}), controllers.RegisterUser)
	router.POST("/login", middleware.BindAndValidate(models.User{}), controllers.LoginUser)
}
