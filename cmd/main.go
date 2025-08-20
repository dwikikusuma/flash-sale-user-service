package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
	infrasturcture2 "user-management-service/infrasturcture/log"
	"user-management-service/internal/api"
	"user-management-service/internal/repository"
	"user-management-service/internal/service"
	middleware2 "user-management-service/middleware"
	"user-management-service/routes"
)

func main() {
	infrasturcture2.InitLogger()

	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := api.NewUserHandler(userService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiterWithConfig(middleware2.GetRateLimiter()))
	e.Use(middleware.ContextTimeout(10 * time.Second))

	routes.SetupRoutes(e, userHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
