package game

import "math/rand"

const BoardSize = 6

type Board struct {
	Tiles [][]Tile `json:"tiles"`
}

func NewBoard() *Board {
	tiles := make([][]Tile, BoardSize)
	for i := range tiles {
		tiles[i] = make([]Tile, BoardSize)
	}
	return &Board{Tiles: tiles}
}

var colors = []GameColor{Red, Blue, Green, Yellow}

func (b *Board) Initialize() {
	tilesPerColor := (BoardSize * BoardSize) / len(colors)
	pool := make([]GameColor, 0, BoardSize*BoardSize)
	for _, color := range colors {
		for range tilesPerColor {
			pool = append(pool, color)
		}
	}

	rand.Shuffle(len(pool), func(i, j int) {
		pool[i], pool[j] = pool[j], pool[i]
	})

	i := 0
	for y := range BoardSize {
		for x := range BoardSize {
			b.SetTile(x, y, pool[i])
			i++
		}
	}
}

func (b *Board) SetTile(x, y int, color GameColor) {
	tile := NewTile(color)
	tile.SetPosition(x, y)
	b.Tiles[y][x] = *tile
}

func (b *Board) UpdateTileColor(x, y int, color GameColor) {
	b.Tiles[y][x].Color = color
}

func (b *Board) SwapTiles(pos1, pos2 Position) {
	tile1 := b.Tiles[pos1.Y][pos1.X]
	tile2 := b.Tiles[pos2.Y][pos2.X]

	color1 := tile1.Color
	color2 := tile2.Color

	b.UpdateTileColor(pos1.X, pos1.Y, color2)
	b.UpdateTileColor(pos2.X, pos2.Y, color1)
}
