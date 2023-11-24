// routes/routes.go

package routes

import (
	"your-package/controllers"
	"your-package/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes and configures the Gin router
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Auth routes
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", controllers.Register)
		authRoutes.POST("/login", controllers.Login)
	}

	// Event routes
	eventRoutes := router.Group("/events")
	eventRoutes.Use(middleware.AuthMiddleware())
	{
		eventRoutes.POST("/", controllers.CreateEvent)
		eventRoutes.PUT("/:id", controllers.EditEvent)
		eventRoutes.DELETE("/:id", controllers.DeleteEvent)
		eventRoutes.GET("/:id", controllers.GetEvent)
		eventRoutes.GET("/", controllers.GetEvents)
	}

	return router
}
