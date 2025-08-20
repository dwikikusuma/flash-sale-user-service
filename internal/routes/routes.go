package routes

import (
	"github.com/labstack/echo/v4"
	"user-management-service/internal/api"
)

func SetupRoutes(e *echo.Echo, uh api.UserHandler) {
	e.GET("/user/:id", uh.GetUserByID) // Get user by ID
	e.POST("/user", uh.CreateUser)     // Create a new user
	e.POST("/user/login", uh.Login)    // User login
}
