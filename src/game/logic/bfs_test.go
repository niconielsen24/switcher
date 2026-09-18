package logic

import (
	"switcher/game"
	"testing"
)

func newBlueBoard() *game.Board {
	board := game.NewBoard()
	for y := range game.BoardSize {
		for x := range game.BoardSize {
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
		if IsFigure(game.Position{X: 0, Y: game.BoardSize}, game.ShapeSquare, *board) {
			t.Error("expected IsFigure to return false for a position outside the board, got true")
		}
	})

	t.Run("square with extra diagonal neighbor tiles", func(t *testing.T) {
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

		// These tiles are red too, but only touch the square diagonally,
		// so they are not part of its 4-directionally connected component
		// and must not affect detection.
		extraPositions := []game.Position{
			{X: 2, Y: 1},
			{X: 5, Y: 1},
			{X: 2, Y: 4},
			{X: 5, Y: 4},
		}
		for _, pos := range extraPositions {
			board.SetTile(pos.X, pos.Y, game.Red)
		}

		if !IsFigure(game.Position{X: 3, Y: 2}, game.ShapeSquare, *board) {
			t.Error("expected IsFigure to detect the red square despite extra diagonal red tiles, got false")
		}
	})
}

func TestIsFigureDetectsShapes(t *testing.T) {
	anchor := game.Position{X: 1, Y: 1}

	for shape, positions := range game.ShapeMap {
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

var diagonalDirections = []game.Position{
	{X: 1, Y: 1},
	{X: 1, Y: -1},
	{X: -1, Y: 1},
	{X: -1, Y: -1},
}

// paintDiagonalNeighborsRed colors every diagonal neighbor of the shape red,
// skipping any position that is already part of the shape or that is
// orthogonally connected to it. These tiles should be invisible to the BFS
// used by IsFigure, since it only walks up/down/left/right.
func paintDiagonalNeighborsRed(board *game.Board, shapeSet map[game.Position]bool) {
	isOrthogonalToShape := func(pos game.Position) bool {
		for _, dir := range directions {
			if shapeSet[game.Position{X: pos.X + dir.X, Y: pos.Y + dir.Y}] {
				return true
			}
		}
		return false
	}

	for shapePos := range shapeSet {
		for _, dir := range diagonalDirections {
			candidate := game.Position{X: shapePos.X + dir.X, Y: shapePos.Y + dir.Y}
			if !board.InBounds(candidate) {
				continue
			}
			if shapeSet[candidate] || isOrthogonalToShape(candidate) {
				continue
			}
			board.SetTile(candidate.X, candidate.Y, game.Red)
		}
	}
}

func TestIsFigureIgnoresDiagonalNeighborTiles(t *testing.T) {
	anchor := game.Position{X: 1, Y: 1}

	for shape, positions := range game.ShapeMap {
		t.Run(shape, func(t *testing.T) {
			board := newBlueBoard()

			shapeSet := make(map[game.Position]bool)
			for _, pos := range positions {
				abs := game.Position{X: anchor.X + pos.X, Y: anchor.Y + pos.Y}
				shapeSet[abs] = true
				board.SetTile(abs.X, abs.Y, game.Red)
			}

			paintDiagonalNeighborsRed(board, shapeSet)

			start := game.Position{X: anchor.X + positions[0].X, Y: anchor.Y + positions[0].Y}
			if !IsFigure(start, shape, *board) {
				t.Errorf("expected IsFigure to detect the red %s shape despite extra diagonal red tiles, got false", shape)
			}
		})
	}
}

func TestIsFigureDetectsRotatedShapes(t *testing.T) {
	// TPositions is three tiles in a row with one hanging below the middle.
	// Rotated 90 degrees clockwise it becomes three tiles in a column with
	// one poking out to the left of the middle.
	rotatedTPositions := []game.Position{
		{X: 2, Y: 2},
		{X: 2, Y: 3},
		{X: 2, Y: 4},
		{X: 1, Y: 3},
	}

	board := newBlueBoard()
	for _, pos := range rotatedTPositions {
		board.SetTile(pos.X, pos.Y, game.Red)
	}

	if !IsFigure(game.Position{X: 2, Y: 2}, game.ShapeT, *board) {
		t.Error("expected IsFigure to detect the T shape rotated 90 degrees, got false")
	}
}

func TestIsFigureDetectsRedSquarePlusExtraOrthogonalNeighbors(t *testing.T) {
	t.Run("square at the origin", func(t *testing.T) {
		board := newBlueBoard()
		squarePositions := []game.Position{
			{X: 0, Y: 0},
			{X: 1, Y: 0},
			{X: 0, Y: 1},
			{X: 1, Y: 1},
			{X: 2, Y: 0}, // extra neighbor
		}
		for _, pos := range squarePositions {
			board.SetTile(pos.X, pos.Y, game.Red)
		}

		if !IsFigure(game.Position{X: 0, Y: 0}, game.ShapeSquare, *board) {
			t.Error("expected IsFigure to detect the red square, got false")
		}
	})
}
