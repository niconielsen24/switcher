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
	if pos.X < 0 || pos.X >= game.BoardSize || pos.Y < 0 || pos.Y >= game.BoardSize {
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

		if !inBounds(pos) {
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

func inBounds(pos game.Position) bool {
	return pos.X >= 0 && pos.X < game.BoardSize && pos.Y >= 0 && pos.Y < game.BoardSize
}

// containsShape reports whether shape occurs, at some translation, inside
// connected, with pos aligned to one of the shape's own cells. Extra
// connected tiles beyond the shape's cells are allowed.
func containsShape(pos game.Position, connected []game.Position, shape game.Shape) bool {
	shapePositions := game.Shapes[shape]

	connectedSet := make(connectedMap, len(connected))
	for _, p := range connected {
		connectedSet[p] = true
	}

	for _, anchor := range shapePositions {
		translation := game.Position{X: pos.X - anchor.X, Y: pos.Y - anchor.Y}

		matches := true
		for _, cell := range shapePositions {
			translated := game.Position{X: cell.X + translation.X, Y: cell.Y + translation.Y}
			if !connectedSet[translated] {
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
