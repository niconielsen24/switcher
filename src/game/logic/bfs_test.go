package logic

import (
	"switcher/game"
	"testing"
)

func newBlueBoard() *game.Board {
	board := game.NewBoard()
	for y := range BoardSize {
		for x := range BoardSize {
			board.SetTile(x, y, game.Blue)
		}
	}
	return board
}

func TestIsFigureDetectsRedSquare(t *testing.T) {
	t.Run("square at the origin", func(t *testing.T) {
		board := newBlueBoard()
		squarePositions := []game.Position{
			{X: 0, Y: 0},
			{X: 1, Y: 0},
			{X: 0, Y: 1},
			{X: 1, Y: 1},
		}
		for _, pos := range squarePositions {
			board.SetTile(pos.X, pos.Y, game.Red)
		}

		if !IsFigure(game.Position{X: 0, Y: 0}, game.ShapeSquare, *board) {
			t.Error("expected IsFigure to detect the red square, got false")
		}
	})

	t.Run("square away from the origin", func(t *testing.T) {
		board := newBlueBoard()
		squarePositions := []game.Position{
			{X: 3, Y: 2},
			{X: 4, Y: 2},
			{X: 3, Y: 3},
			{X: 4, Y: 3},
		}
		for _, pos := range squarePositions {
			board.SetTile(pos.X, pos.Y, game.Red)
		}

		if !IsFigure(game.Position{X: 3, Y: 2}, game.ShapeSquare, *board) {
			t.Error("expected IsFigure to detect the red square away from the origin, got false")
		}
	})

	t.Run("invalid position", func(t *testing.T) {
		board := newBlueBoard()

		if IsFigure(game.Position{X: -1, Y: 0}, game.ShapeSquare, *board) {
			t.Error("expected IsFigure to return false for a negative position, got true")
		}
		if IsFigure(game.Position{X: 0, Y: BoardSize}, game.ShapeSquare, *board) {
			t.Error("expected IsFigure to return false for a position outside the board, got true")
		}
	})
}

func TestIsFigureDetectsShapes(t *testing.T) {
	anchor := game.Position{X: 1, Y: 1}

	for shape, positions := range game.Shapes {
		t.Run(shape, func(t *testing.T) {
			board := newBlueBoard()
			for _, pos := range positions {
				board.SetTile(anchor.X+pos.X, anchor.Y+pos.Y, game.Red)
			}

			start := game.Position{X: anchor.X + positions[0].X, Y: anchor.Y + positions[0].Y}
			if !IsFigure(start, shape, *board) {
				t.Errorf("expected IsFigure to detect the red %s shape, got false", shape)
			}
		})
	}
}
