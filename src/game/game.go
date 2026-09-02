package game

type Game struct {
	Players []*Player
	Board   *Board
}

func NewGame(players []*Player) *Game {
	board := NewBoard()
	board.Initialize()
	return &Game{Players: players, Board: board}
}
