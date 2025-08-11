package main

import (
	"github.com/labstack/echo/v4"
	"user-management-service/internal/api"
	"user-management-service/internal/repository"
	"user-management-service/internal/service"
)

func main() {
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := api.NewUserHandler(userService)

	e := echo.New()
	userHandler.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
