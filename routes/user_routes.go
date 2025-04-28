package routes

import (
	"user-crud-app/controllers" // Adjust this import path based on your project structure

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
    userRoutes := r.Group("/users") // Group all user-related routes under /users
    {
        userRoutes.GET("/", controllers.GetUsers)          // Get all users
        userRoutes.GET("/:id", controllers.GetUserByID)    // Get user by ID
        userRoutes.POST("/", controllers.CreateUser)       // Create a new user
        userRoutes.PUT("/:id", controllers.UpdateUser)     // Update an existing user
        userRoutes.DELETE("/:id", controllers.DeleteUser)  // Delete user by ID
    }
}
