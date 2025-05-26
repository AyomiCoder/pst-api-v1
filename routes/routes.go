package routes

import (
	"godb/handlers"
	"godb/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures all the routes for our application
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Public routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	// Protected routes
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// Post routes
		posts := api.Group("/posts")
		{
			posts.POST("", handlers.CreatePost)
			posts.GET("", handlers.GetPosts)
			posts.GET("/:id", handlers.GetPost)
			posts.PUT("/:id", handlers.UpdatePost)
			posts.DELETE("/:id", handlers.DeletePost)
		}
	}

	return r
} 