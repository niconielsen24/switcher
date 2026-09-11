package main

import (
	"os"
	"switcher/api"

	"github.com/labstack/echo/v5"
)

func main() {
	e := echo.New()
	e.GET("/", api.NewUser)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
