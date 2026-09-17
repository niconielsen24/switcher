package logic

import "switcher/game"

var directions = []game.Position{
	{X: 0, Y: 1},  // Up
	{X: 0, Y: -1}, // Down
	{X: 1, Y: 0},  // Right
	{X: -1, Y: 0}, // Left
}

const BoardSize = 6

func IsFigure(pos game.Position, shape game.Shape, board game.Board) bool {
	if pos.X < 0 || pos.X >= BoardSize || pos.Y < 0 || pos.Y >= BoardSize {
		return false
	}
	rootTile := board.Tiles[pos.Y][pos.X]

	valid_neighbors := bfs(rootTile, board)

	normalized_positions := normalizePositions(valid_neighbors)

	return game.ComparePositions(normalized_positions, shape)
}

func bfs(startTile game.Tile, board game.Board) []game.Position {
	pos := startTile.Position
	queue := NewQueue()
	queue = queue.Enqueue(startTile)

	visited := make(map[game.Position]bool)
	visited[pos] = true

	for queue.length > 0 {
		currentTile, updatedQueue := queue.Dequeue()
		queue = updatedQueue

		for _, neighbor := range getNeighbors(currentTile, board, visited) {
			pos := neighbor.Position
			if !visited[pos] {
				visited[pos] = true
				queue = queue.Enqueue(neighbor)
			}
		}

	}

	valid_neighbors := []game.Position{}
	for pos := range visited {
		valid_neighbors = append(valid_neighbors, pos)
	}
	return valid_neighbors
}

func getNeighbors(tile game.Tile, board game.Board, visited map[game.Position]bool) []game.Tile {
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
	return pos.X >= 0 && pos.X < BoardSize && pos.Y >= 0 && pos.Y < BoardSize
}

func normalizePositions(positions []game.Position) []game.Position {
	if len(positions) == 0 {
		return nil
	}

	minX := getMinX(positions)
	minY := getMinY(positions)

	normalized := make([]game.Position, len(positions))
	for i, pos := range positions {
		normalized[i] = game.Position{
			X: pos.X - minX,
			Y: pos.Y - minY,
		}
	}

	return normalized
}

func getMinY(positions []game.Position) int {
	if len(positions) == 0 {
		return 0
	}
	minY := positions[0].Y
	for _, pos := range positions {
		if pos.Y < minY {
			minY = pos.Y
		}
	}
	return minY
}

func getMinX(positions []game.Position) int {
	if len(positions) == 0 {
		return 0
	}
	minX := positions[0].X
	for _, pos := range positions {
		if pos.X < minX {
			minX = pos.X
		}
	}
	return minX
}
