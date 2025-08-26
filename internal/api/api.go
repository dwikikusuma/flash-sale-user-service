package api

import (
	"fmt"
	"strconv"
	"user-management-service/internal/entity"
	"user-management-service/internal/service"

	_ "github.com/golang-jwt/jwt/v5"
	_ "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4/middleware"
)

type UserHandler interface {
	GetUserByID(c echo.Context) error
	CreateUser(c echo.Context) error
	Login(c echo.Context) error
}

type userHandler struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		UserService: userService,
	}
}

func (uh *userHandler) GetUserByID(c echo.Context) error {
	userIDStr := c.Param("id")
	ctx := c.Request().Context()
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid user ID"})
	}
	user, err := uh.UserService.GetUserDetails(ctx, userID)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to retrieve user details"})
	}

	if user.ID == 0 {
		return c.JSON(404, map[string]string{"error": "User not found"})
	}

	return c.JSON(200, user)
}
func (uh *userHandler) CreateUser(c echo.Context) error {
	var request entity.User
	ctx := c.Request().Context()

	err := c.Bind(&request)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid user data"})
	}

	err = uh.UserService.CreateUser(ctx, &request)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to create user"})
	}

	return c.JSON(201, fmt.Sprintf("User created successfully"))
}

func (uh *userHandler) Login(c echo.Context) error {
	var request entity.LoginRequest
	ctx := c.Request().Context()

	err := c.Bind(&request)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid login data"})
	}

	token, err := uh.UserService.Login(ctx, &request)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to login"})
	}

	return c.JSON(200, map[string]string{"message": "Login successful", "token": token})
}
