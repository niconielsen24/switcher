package main

import (
	"os"
	"switcher/api"
	"switcher/middleware"

	"github.com/labstack/echo/v5"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS)
	e.GET("/", api.NewUser)
	e.POST("/games/create", api.CreateGame)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
