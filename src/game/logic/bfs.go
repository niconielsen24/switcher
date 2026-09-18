package logic

import (
	"switcher/game"
)

var directions = []game.Position{
	{X: 0, Y: 1},  // Up
	{X: 0, Y: -1}, // Down
	{X: 1, Y: 0},  // Right
	{X: -1, Y: 0}, // Left
}

type connectedMap map[game.Position]bool

func IsFigure(pos game.Position, shape game.Shape, board game.Board) bool {
	if !board.InBounds(pos) {
		return false
	}
	rootTile := board.Tiles[pos.Y][pos.X]

	valid_neighbors := bfs(rootTile, board)

	return containsShape(pos, valid_neighbors, shape)
}

func bfs(startTile game.Tile, board game.Board) []game.Position {
	pos := startTile.Position
	queue := NewQueue()
	queue = queue.Enqueue(startTile)

	visited := make(connectedMap)
	visited[pos] = true

	for queue.length > 0 {
		currentTile, updatedQueue := queue.Dequeue()
		queue = updatedQueue

		for _, neighbor := range getNeighborTiles(currentTile, board, visited) {
			pos := neighbor.Position
			if !visited[pos] {
				visited[pos] = true
				queue = queue.Enqueue(neighbor)
			}
		}

	}

	valid_neighbor_positions := []game.Position{}
	for pos := range visited {
		valid_neighbor_positions = append(valid_neighbor_positions, pos)
	}
	return valid_neighbor_positions
}

func getNeighborTiles(tile game.Tile, board game.Board, visited map[game.Position]bool) []game.Tile {
	neighbors := []game.Tile{}
	for _, dir := range directions {
		pos := tile.Position
		pos.X += dir.X
		pos.Y += dir.Y

		if !board.InBounds(pos) {
			continue
		}
		if visited[pos] {
			continue
		}
		if board.Tiles[pos.Y][pos.X].Color == tile.Color {
			neighbor := board.Tiles[pos.Y][pos.X]
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

// TODO: review shape matching logic, should probably
// check for exact matches instead of ignoring extra tiles.
// For example, if the shape is a square, but the connected tiles are a square with an extra tile attached,
// it should not match. This is because the extra tile would be considered part of the shape,
// and thus the shape would not be a square anymore.
func containsShape(pos game.Position, connected []game.Position, shape game.Shape) bool {
	shapePositions := game.ShapeMap[shape]

	connectedSet := make(connectedMap, len(connected))
	for _, p := range connected {
		connectedSet[p] = true
	}

	for _, orientation := range rotations(shapePositions) {
		if matchesAtSomeTranslation(pos, orientation, connectedSet) {
			return true
		}
	}

	return false
}

func matchesAtSomeTranslation(pos game.Position, shapePositions []game.Position, connected connectedMap) bool {
	for _, anchor := range shapePositions {
		translation := game.Position{X: pos.X - anchor.X, Y: pos.Y - anchor.Y}

		matches := true
		for _, cell := range shapePositions {
			translated := game.Position{X: cell.X + translation.X, Y: cell.Y + translation.Y}
			if !connected[translated] {
				matches = false
				break
			}
		}
		if matches {
			return true
		}
	}
	return false
}

func rotations(positions []game.Position) [][]game.Position {
	variants := make([][]game.Position, 0, 4)
	current := positions
	for range 4 {
		variants = append(variants, current)
		current = rotate90(current)
	}
	return variants
}

func rotate90(positions []game.Position) []game.Position {
	rotated := make([]game.Position, len(positions))
	for i, p := range positions {
		rotated[i] = game.Position{X: -p.Y, Y: p.X}
	}
	return rotated
}
