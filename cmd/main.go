package main

import (
	"time"
	"user-management-service/config"
	infrasturcture "user-management-service/infrasturcture/log"
	"user-management-service/internal/api"
	"user-management-service/internal/repository"
	"user-management-service/internal/service"
	reqMiddleware "user-management-service/middleware"
	"user-management-service/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	appConfig := config.LoadConfig(
		config.WithConfigFolder([]string{"./files/config"}),
		config.WithConfigFile("config"),
		config.WithConfigType("yaml"),
	)

	infrasturcture.InitLogger()

	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := api.NewUserHandler(userService, appConfig.Secret.JWTSecret)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiterWithConfig(reqMiddleware.GetRateLimiter()))
	e.Use(middleware.ContextTimeout(10 * time.Second))

	routes.SetupRoutes(e, userHandler)

	e.Logger.Fatal(e.Start(":" + appConfig.App.Port))
}
