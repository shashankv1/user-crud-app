package services

import (
	"user-crud-app/models" // Import the models
)

// In-memory slice for users (for demo purposes)
var users = []models.User{
    {ID: 1, Name: "John Doe", Email: "john@example.com"},
    {ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
}

// Get all users
func GetUsers() []models.User {
    return users
}

// Get a user by ID
func GetUserByID(id int) (models.User, bool) {
    for _, user := range users {
        if user.ID == id {
            return user, true
        }
    }
    return models.User{}, false
}

// Create a new user
func CreateUser(newUser models.User) models.User {
    newUser.ID = len(users) + 1 // Auto-increment ID
    users = append(users, newUser)
    return newUser
}

// Update an existing user
func UpdateUser(id int, updatedUser models.User) (models.User, bool) {
    for i, user := range users {
        if user.ID == id {
            users[i] = updatedUser
            updatedUser.ID = id // Keep the original ID
            return updatedUser, true
        }
    }
    return models.User{}, false
}

// Delete a user
func DeleteUser(id int) bool {
    for i, user := range users {
        if user.ID == id {
            users = append(users[:i], users[i+1:]...)
            return true
        }
    }
    return false
}
