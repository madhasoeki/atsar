package routes

import (
	"atsar/controllers"
	"atsar/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		// Protecting POST /api/users with SuperAdminMiddleware
		api.POST("/users", middlewares.SuperAdminMiddleware(), controllers.CreateUser)
	}
}
