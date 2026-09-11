package logic

import (
	"switcher/game"
	"uuid"
)

type Game struct {
	Id      uuid.UUID
	Players []*game.Player
	Board   *game.Board
}

func NewGame(players []*game.Player) *Game {
	board := game.NewBoard()
	board.Initialize()
	return &Game{Id: uuid.New(), Players: players, Board: board}
}

func (g *Game) ValidateFigure(shape game.Shape, color game.GameColor, x, y int) bool {
	return true
}
