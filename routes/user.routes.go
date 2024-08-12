package routes

import (
	"restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/users", controllers.GetAllUsers)
	router.GET("/users/:userId", controllers.GetUser)
	router.POST("/users/register", controllers.CreateUser)
	router.POST("/users/login", controllers.Login)	
}