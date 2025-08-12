package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"user-management-service/internal/api"
	"user-management-service/internal/infrasturcture"
	"user-management-service/internal/repository"
	"user-management-service/internal/service"
)

func main() {
	infrasturcture.InitLogger()

	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := api.NewUserHandler(userService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userHandler.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
