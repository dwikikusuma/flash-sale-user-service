package routes

import (
	"user-management-service/internal/api"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, uh api.UserHandler, jwtSecret string) {

	// Get user by ID
	e.POST("/user", uh.CreateUser)  // Create a new user
	e.POST("/user/login", uh.Login) // User login

	router := e.Group("/secure")
	router.Use(echojwt.JWT(jwtSecret))
	router.GET("/user/:id", uh.GetUserByID)
}
