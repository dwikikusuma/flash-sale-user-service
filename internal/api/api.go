package api

import (
	"fmt"
	"strconv"
	"time"
	"user-management-service/internal/entity"
	"user-management-service/internal/service"

	"github.com/golang-jwt/jwt/v5"
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
	Secret      string
	UserService service.UserService
}

func NewUserHandler(userService service.UserService, secret string) UserHandler {
	return &userHandler{
		UserService: userService,
		Secret:      secret,
	}
}

func (uh *userHandler) GetUserByID(c echo.Context) error {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid user ID"})
	}
	user, err := uh.UserService.GetUserDetails(userID)
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
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid user data"})
	}

	user, err := uh.UserService.CreateUser(&request)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to create user"})
	}

	return c.JSON(201, fmt.Sprintf("User created successfully: %+v", user))
}

func (uh *userHandler) Login(c echo.Context) error {
	var request entity.LoginRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid login data"})
	}

	user, err := uh.UserService.Login(&request)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to login"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"email":    user.Email,
		"exp":      jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
	})

	tokenString, err := token.SignedString([]byte(uh.Secret))
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to generate token"})
	}

	return c.JSON(200, map[string]string{"message": "Login successful", "token": tokenString})
}
