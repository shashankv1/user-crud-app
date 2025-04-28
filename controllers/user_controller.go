package controllers

import (
	"net/http"
	"strconv"
	"user-crud-app/models"
	"user-crud-app/services"

	"github.com/gin-gonic/gin"
)

// Get all users
func GetUsers(c *gin.Context) {
    users := services.GetUsers()  // Use the service function
    c.JSON(http.StatusOK, users)
}

// Get user by ID
func GetUserByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    user, found := services.GetUserByID(id)
    if found {
        c.JSON(http.StatusOK, user)
    } else {
        c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
    }
}

// Create a new user
func CreateUser(c *gin.Context) {
    var newUser models.User
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }
    createdUser := services.CreateUser(newUser)  // Use the service function
    c.JSON(http.StatusCreated, createdUser)
}

// Update an existing user
func UpdateUser(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var updatedUser models.User
    if err := c.ShouldBindJSON(&updatedUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }
    updated, found := services.UpdateUser(id, updatedUser)  // Use the service function
    if found {
        c.JSON(http.StatusOK, updated)
    } else {
        c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
    }
}

// Delete a user
func DeleteUser(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    deleted := services.DeleteUser(id)  // Use the service function
    if deleted {
        c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
    } else {
        c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
    }
}
