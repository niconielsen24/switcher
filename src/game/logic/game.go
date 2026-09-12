package logic

import (
	"switcher/game"
	"uuid"
)

type Game struct {
	Id      uuid.UUID      `json:"id"`
	Players []*game.Player `json:"players"`
	Board   *game.Board    `json:"board"`
}

func NewGame(players []*game.Player) *Game {
	board := game.NewBoard()
	board.Initialize()
	return &Game{Id: uuid.New(), Players: players, Board: board}
}

func (g *Game) ValidateFigure(shape game.Shape, color game.GameColor, x, y int) bool {
	return true
}
