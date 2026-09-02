package main

import (
	"switcher/api"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", api.HandleGetGame)

	e.Logger.Fatal(e.Start(":8080"))
}
