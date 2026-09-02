package game

const BoardSize = 6

type Board struct {
	Tiles [][]Tile
}

func NewBoard() *Board {
	tiles := make([][]Tile, BoardSize)
	for i := range tiles {
		tiles[i] = make([]Tile, BoardSize)
	}
	return &Board{Tiles: tiles}
}

func (b *Board) Initialize() {
	for i := range BoardSize {
		for j := range BoardSize {
			b.Tiles[i][j] = Tile{color: "red"}
		}
	}
}
