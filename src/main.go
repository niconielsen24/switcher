package main

import (
	"fmt"
	"switcher/game"
)

func main() {
	player1 := game.NewPlayer("Alice")
	player2 := game.NewPlayer("Bob")
	players := []*game.Player{player1, player2}
	game := game.NewGame(players)
	fmt.Printf("New game: %f\n", game)
}
