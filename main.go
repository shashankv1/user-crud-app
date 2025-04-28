package main

import (
	"user-crud-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    // Initialize Gin router
    r := gin.Default()

    // Register user-related routes
    routes.RegisterUserRoutes(r)

    // Start the server on port 8080
    r.Run(":8080")
}
