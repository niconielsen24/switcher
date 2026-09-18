package main

import (
	"os"
	"switcher/api"
	"switcher/middleware"
	"switcher/store"

	"github.com/labstack/echo/v5"
)

func main() {
	game_store := store.NewInmemGameStore()
	user_store := store.NewInmemUserStore()

	e := echo.New()

	e.Use(middleware.CORS)
	e.Use(middleware.NewGameStore(game_store))
	e.Use(middleware.NewUserStore(user_store))

	e.GET("/", api.NewUser)
	e.POST("/games/create", api.CreateGame)
	e.PUT("/games/update", api.UpdateGame)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
