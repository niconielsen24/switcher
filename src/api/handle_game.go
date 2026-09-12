package api

import (
	"switcher/game"
	"switcher/game/logic"

	"github.com/labstack/echo/v5"
)

var DEBUG_PLAYERS = []*game.Player{
	game.NewPlayer("Player 1"),
	game.NewPlayer("Player 2"),
}

func CreateGame(c *echo.Context) error {
	game := logic.NewGame(DEBUG_PLAYERS)

	return c.JSON(200, game)
}
