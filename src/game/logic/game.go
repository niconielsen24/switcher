package logic

import (
	"switcher/game"
	"sync"
	"time"
	"uuid"
)

const MAX_PLAYERS = 4
const TurnDuration = 2 * time.Minute

type Game struct {
	Id      uuid.UUID      `json:"id"`
	Players []*game.Player `json:"players"`
	Board   *game.Board    `json:"board"`
	Turn    int            `json:"turn"`
	Host    *game.Player   `json:"host"`

	closed    bool
	turnTimer *time.Timer
	sync.Mutex
}

func NewGame(players []*game.Player) *Game {
	board := game.NewBoard()
	board.Initialize()

	g := &Game{
		Id:      uuid.New(),
		Players: players,
		Board:   board,
		Turn:    0,
		Host:    players[0],
		closed:  false,
	}
	g.turnTimer = time.AfterFunc(TurnDuration, g.advanceTurn)
	return g
}

func (g *Game) Close() {
	g.Lock()
	defer g.Unlock()
	if !g.closed {
		g.turnTimer.Stop()
		g.closed = true
	}
}

func (g *Game) advanceTurn() {
	g.Lock()
	defer g.Unlock()
	g.nextTurn()
}

func (g *Game) EndTurn(playerId string) {
	if !g.isTurn(playerId) {
		return
	}

	g.Lock()
	defer g.Unlock()
	g.turnTimer.Stop()
	g.nextTurn()
}

func (g *Game) nextTurn() {
	if !g.closed {
		g.Turn = (g.Turn + 1) % len(g.Players)
		g.turnTimer = time.AfterFunc(TurnDuration, g.advanceTurn)
	}
}

func (g *Game) isTurn(playerId string) bool {
	return g.Players[g.Turn].Id.String() == playerId
}

func (g *Game) MakeMove(playerId string, cardId string, start, end game.Position) {
	g.Lock()
	defer g.Unlock()

	if !g.isTurn(playerId) {
		return
	}

	p := g.Players[g.Turn]
	card, ok := p.MovementDeck[cardId]
	if !ok {
		return
	}

	if !g.Board.InBounds(start) || !g.Board.InBounds(end) {
		return
	}

	if !card.ValidSwap(start, end) {
		return
	}

	g.Board.SwapTiles(start, end)
}

func (g *Game) ValidateFigure(shape game.Shape, color game.GameColor, x, y int) bool {
	return true
}
