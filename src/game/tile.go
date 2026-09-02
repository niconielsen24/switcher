package game

type Color string

const (
	Red    Color = "red"
	Blue   Color = "blue"
	Green  Color = "green"
	Yellow Color = "yellow"
)

type Tile struct {
	color Color
}

func NewTile(color Color) *Tile {
	return &Tile{color: color}
}
