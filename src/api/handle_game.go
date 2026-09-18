package api

import (
	"switcher/game"
	"switcher/game/logic"
	"switcher/store"

	"github.com/labstack/echo/v5"
)

var DEBUG_PLAYERS = []*game.Player{
	game.NewPlayer("Player 1"),
	game.NewPlayer("Player 2"),
}

func CreateGame(c *echo.Context) error {
	if s, ok := c.Get("gameStore").(store.GameStore); ok {
		game := logic.NewGame(DEBUG_PLAYERS)
		s.Set(game.Id.String(), game)
		return c.JSON(200, game)
	}
	return c.JSON(500, map[string]string{"error": "store not found"})
}

type UpdateGameRequest struct {
	GameId    string          `json:"gameId"`
	Positions []game.Position `json:"positions"`
}

func UpdateGame(c *echo.Context) error {
	s, ok := c.Get("gameStore").(store.GameStore)
	if !ok {
		return c.JSON(500, map[string]string{"error": "store not found"})
	}

	r := new(UpdateGameRequest)
	c.Bind(r)

	game, err := s.Get(r.GameId)
	if err != nil {
		return c.JSON(404, map[string]string{"error": "game not found"})
	}

	game.Board.SwapTiles(r.Positions[0], r.Positions[1])

	return c.JSON(200, game)
}
