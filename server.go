package main

import (
	"github.com/labstack/echo"
	"github.com/rad08d/api/controllers"
	_ "github.com/rad08d/api/docs"
	"github.com/swaggo/echo-swagger"
)

func main() {
	e := echo.New()
	// Bind all routes
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/health", controllers.Health)
	e.POST("/search", controllers.CheckAvailability)

	e.Logger.Fatal(e.Start(":8080"))
}
